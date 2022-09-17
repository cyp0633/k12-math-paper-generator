package middleware

import (
	"k12-math-paper-generator/internal/models"
	"testing"
)

func TestAuth(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"1888888", "123456"}, true},
		{"", args{"1999999", "123456"}, true},
		{"", args{"王五3", "123456"}, false},
		{"", args{"刘其", "123"}, false},
	}
	user := &models.User{
		Username: "1888888",
		Password: "123456",
	}
	user2 := &models.User{
		Username: "1999999",
		Password: "123456",
	}
	models.DB.Create(user)
	models.DB.Create(user2)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Auth(tt.args.username, tt.args.password); got != tt.want {
				t.Errorf("Auth() = %v, want %v", got, tt.want)
			}
		})
	}
	models.DB.Delete(user)
	models.DB.Delete(user2)
}
