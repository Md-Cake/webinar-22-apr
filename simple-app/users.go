package main

import (
	"fmt"
	"regexp"
	"sync"
)

var (
	validUsername = regexp.MustCompile(`^[a-zA-Z]\w+$`)     // admin
	validEmail    = regexp.MustCompile(`^\w+@(\w+\.)+\w+$`) // admin@gmail.com
)

type UserSource interface {
	RegisterUser(username, email string) (*User, error)
	HasUser(id int) bool
}

type userSource struct {
	sync.Mutex

	users []*User
}

func NewUserSource() UserSource {
	return &userSource{}
}

func (us *userSource) RegisterUser(username, email string) (*User, error) {
	us.Lock()
	defer us.Unlock()

	newId := len(us.users)
	u, err := NewUser(newId, username, email)
	if err != nil {
		return nil, err
	}

	us.users = append(us.users, u)
	return u, nil
}

func (us *userSource) HasUser(id int) bool {
	return id >= 0 && id < len(us.users)
}

type User struct {
	Id       int
	Username string
	Email    string
}

func NewUser(id int, username, email string) (*User, error) {
	if !validEmail.MatchString(email) {
		return nil, fmt.Errorf("email is not valid %s", email)
	}
	if !validUsername.MatchString(username) {
		return nil, fmt.Errorf("username is not valid %s", username)
	}

	return &User{
		Id:       id,
		Username: username,
		Email:    email,
	}, nil
}

func (u *User) String() string {
	return fmt.Sprintf("User[id=%d username=%s email=%s]", u.Id, u.Username, u.Email)
}
