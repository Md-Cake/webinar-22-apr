package main_test

import (
	"fmt"
	"testing"

	main "github.com/md-cake/simple-app"
	"go.uber.org/mock/gomock"
)

func TestServer_RegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	userSource := NewMockUserSource(ctrl)

	userSource.EXPECT().
		RegisterUser("admin", "admin@gmail.com").
		DoAndReturn(func(username, email string) (*main.User, error) {
			return main.NewUser(0, "user", email)
		}).AnyTimes()

	s := main.NewServer(userSource, main.NewPostSource())

	u, err := s.RegisterUser("admin", "admin@gmail.com")
	if err != nil {
		t.Fatal("Unexpected error")
	}

	fmt.Println(u)
}
