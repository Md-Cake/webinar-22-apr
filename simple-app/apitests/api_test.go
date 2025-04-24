package apitests_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Hello(t *testing.T) {
	resp, err := http.Get("https://playground.learnqa.ru/api/hello")
	if err != nil {
		t.Fatal("Unexpected error")
	}
	defer resp.Body.Close()

	answer := struct {
		Answer string `json:"answer"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&answer)
	if err != nil {
		t.Fatal("Unexpected error decoding answer")
	}

	assert.Equal(t, "Hello, someone", answer.Answer)
}

func Test_Post(t *testing.T) {
	request := struct {
		Name string `json:"name"`
	}{
		Name: "Alex",
	}

	body, _ := json.Marshal(request)
	resp, err := http.Post(
		"https://playground.learnqa.ru/api/show_json",
		"application/json",
		bytes.NewReader(body))
	if err != nil {
		t.Fatal("Unexpected error")
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Unexpected error")
	}

	fmt.Println(string(body))
}

func Test_Put(t *testing.T) {
	client := http.Client{}
	request, err := http.NewRequest(
		http.MethodPut,
		"https://playground.learnqa.ru/api/show_json",
		bytes.NewReader([]byte{}))

	// add cookies in your request
	cookie := &http.Cookie{
		Name:    "my_cookie",
		Value:   "some_value",
		Path:    "/",                       // path the cookie is valid for
		Domain:  "example.com",             // domain the cookie is valid for
		Expires: time.Now().Add(time.Hour), // set expiration time if needed
	}
	request.AddCookie(cookie)

	if err != nil {
	}

	request.Header.Add("", "")

	resp, err := client.Do(request)

	for range resp.Header {
	}
}
