package main

import (
	"fmt"
	"sync"
	"time"
)

var data = [][]int{
	{
		1, 2, 3, 4, 5,
	},
	{
		1, 2, 4, 5,
	},
	{
		1, 2, 3, 4,
	},
}

func main() {
	msgChannel := make(chan string, 3)

	wg := sync.WaitGroup{}

	wg.Add(3)
	go func() {
		printSums(data, msgChannel)
		wg.Done()
	}()
	go func() {
		printSums(data, msgChannel)
		wg.Done()
	}()
	go func() {
		printSums(data, msgChannel)
		wg.Done()
	}()

	go func() {
		for msg := range msgChannel {
			fmt.Println(msg)
		}
	}()

	wg.Wait()
}

func printSums(streams [][]int, msgChannel chan<- string) {
	for _, list := range streams {
		msgChannel <- fmt.Sprintf("Sum(%v) = %d", list, sum(list))
		time.Sleep(time.Millisecond * 10)
	}
}

type Age int

type User struct {
	Name string
	Age  Age
}

func NewUser(name string, age int) (*User, error) {
	if age < 18 {
		return nil, fmt.Errorf("our user is not 18 yo; got age = %d", age)
	}

	return &User{
		Name: name,
		Age:  Age(age),
	}, nil
}

func (u *User) GetAge() Age {
	return u.Age
}

func (u *User) HappyBirthday() {
	u.Age++
	fmt.Printf("Happy %d Birthday, %s!\n", u.Age, u.Name)
}

func (u *User) String() string {
	return fmt.Sprintf("User[name=%s age=%d]", u.Name, u.Age)
}

type Set[T comparable] map[T]bool

func (s Set[T]) Add(v T) {
	s[v] = true
}

func (s Set[T]) Has(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Delete(v T) {
	delete(s, v)
}

func (s Set[T]) String() string {
	res := "Set["

	for k := range s {
		res += fmt.Sprint(k) + " "
	}
	return res + "]"
}

type Numbers interface {
	int | float32 | byte | int32
}

func sum[T Numbers](values []T) T {
	res := values[0]
	for i := 1; i < len(values); i++ {
		res = res + values[i]
	}
	return res
}

func reverse[T any](ints []T) []T {
	var res = make([]T, len(ints))
	for i := range ints {
		res[i] = ints[len(ints)-1-i]
	}
	return res
}

// func testPanic() string {
// 	fmt.Println("inside function; recover():", recover())

// 	defer func() {
// 		if v := recover(); v != nil {
// 			fmt.Println("we were in panic mode! but now everything is fine; recover():", v)
// 		}
// 	}()

// 	panic("Something is entirely wrong!")

// 	return "test"
// }
