package main

import (
	"fmt"
	"sync"
)

func main() {
	// RegisterUser(username, email string)
	// MakePost(current *User, text string)
	// FollowUser(current *User, toFollow *User)
	// GetFeed(current *User)

	s := NewServer(
		NewUserSource(),
		NewPostSource(),
	)

	wg := sync.WaitGroup{}
	wg.Add(10)

	// for i := 0 ; i < 10; i ++ {}
	for batch := range 10 {
		go func() {
			for i := batch; i < 100; i += 10 {
				u, err := s.RegisterUser(
					fmt.Sprintf("user%d", i),
					fmt.Sprintf("user%d@gmail.com", i),
				)
				if err != nil {
					fmt.Println("Error creating user", err)
				} else {
					fmt.Println("New user:", u)
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
