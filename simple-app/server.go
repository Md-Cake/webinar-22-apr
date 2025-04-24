package main

import (
	"fmt"
	"sync"
)

type Server struct {
	lock *sync.Mutex

	userSource UserSource
	postSource PostSource

	followingUsers map[int]Set[int]
}

func NewServer(userSource UserSource, postSource PostSource) *Server {
	return &Server{
		lock: &sync.Mutex{},

		userSource: userSource,
		postSource: postSource,

		followingUsers: make(map[int]Set[int]),
	}
}

func (s *Server) RegisterUser(username, email string) (*User, error) {
	u, err := s.userSource.RegisterUser(username, email)
	if err != nil {
		return nil, err
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	s.followingUsers[u.Id] = make(Set[int])
	return u, nil
}

func (s *Server) MakePost(author *User, text string) (*Post, error) {
	if !s.userSource.HasUser(author.Id) {
		return nil, fmt.Errorf("cannot create post from non-existing user %d", author.Id)
	}

	return s.postSource.NewPost(author, text)
}

func (s *Server) FollowUser(current *User, toFollow *User) error {
	if !s.userSource.HasUser(current.Id) {
		return fmt.Errorf("current user doesn't exist: %d", current.Id)
	}
	if !s.userSource.HasUser(toFollow.Id) {
		return fmt.Errorf("user to follow doesn't exist: %d", toFollow.Id)
	}

	s.followingUsers[current.Id].Add(toFollow.Id)
	return nil
}

func (s *Server) GetFeed(current *User) ([]*Post, error) {
	if !s.userSource.HasUser(current.Id) {
		return nil, fmt.Errorf("current user doesn't exist: %d", current.Id)
	}

	var ids []int
	for id := range s.followingUsers[current.Id] {
		ids = append(ids, id)
	}

	feed := s.postSource.GetPostsByAuthors(ids)
	return feed, nil
}
