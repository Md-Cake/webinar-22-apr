package main_test

import (
	main "github.com/md-cake/simple-app"

	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSimpleApp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SimpleApp Suite")
}

var _ = Describe("new user register", func() {
	username := "admin"
	email := "admin@gmail.com"

	var user *main.User
	var err error

	BeforeEach(func() {
		user, err = main.NewUserSource().RegisterUser(username, email)
	})

	It("Should not return error", func() {
		Expect(err).To(BeNil())
	})

	It("Should return correct email", func() {
		Expect(user.Email).To(Equal(email))
	})
})
