package main_test

import (
	"fmt"

	main "github.com/md-cake/simple-app"

	"testing"
)

// func Test_userSource_RegisterUser(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		username string
// 		email    string
// 		wantErr  bool
// 	}{
// 		{
// 			name:     "with correct data",
// 			username: "admin",
// 			email:    "admin@gmail.com",
// 			wantErr:  false,
// 		},
// 	}

// 	us := main.NewUserSource()

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := us.RegisterUser(tt.username, tt.email)
// 			if tt.wantErr {
// 				require.Error(t, err, "Unexpeceted error")
// 			} else {
// 				require.NoError(t, err, "Unexpeceted error")
// 			}
// 			assert.Equal(t, tt.email, got.Email, "wrong email")
// 		})
// 	}
// }

func Benchmark_UserRegister(b *testing.B) {
	us := main.NewUserSource()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			username := fmt.Sprintf("user_%d_%d", i, j)
			email := fmt.Sprintf("email%d_%d@gmail.com", i, j)

			_, err := us.RegisterUser(username, email)
			if err != nil {
				b.Fatal("Unexpected error")
			}
		}
	}
}
