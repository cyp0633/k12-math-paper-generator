package models

import (
	"fmt"
	"testing"
)

func TestGetUserByName(t *testing.T) {
	tests := []struct {
		name     string
		username string
	}{
		{"", "张三1"},
		{"", "李四3"},
		{"", "王五2"},
		{"", "Tom"},
	}
	u := &User{
		Username: "Tom",
	}
	DB.Create(u)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUserByName(tt.username); got.Username != tt.username {
				t.Errorf("GetUserByName() = %v, want %v", got.Username, tt.username)
			}
		})
	}
	DB.Delete(u)
}

func TestClearUser(t *testing.T) {
	ClearUser()
	if CurrentUser != nil {
		t.Errorf("CurrentUser want nil")
	}
}

func TestCreateUser(t *testing.T) {
	type args struct {
		phone string
		code  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"123456", 123456}, true},
		{"", args{"1234567", 123456}, true},
	}
	for _, tt := range tests {
		user := &User{
			Username: tt.args.phone,
			Note:     fmt.Sprint(tt.args.code),
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateUser(tt.args.phone, tt.args.code); got != tt.want {
				t.Errorf("CreateUser() = %v, want %v", got, tt.want)
			}
		})
		DB.Delete(user)
	}
}

func TestUser_ChangePswd(t *testing.T) {
	type fields struct {
		Username string
		Password string
	}
	tests := []struct {
		name   string
		fields fields
		pswd   string
	}{
		{"", fields{"101010", "111111"}, "123456"},
		{"", fields{"121212", "111111"}, "123456"},
		{"", fields{"151515", "111111"}, "123456"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &User{
				Username: tt.fields.Username,
				Password: tt.fields.Password,
			}
			DB.Create(a)
			a.ChangePswd(tt.pswd)
			DB.Where(&User{Username: tt.fields.Username}).First(a)
			if a.Password != tt.pswd {
				t.Errorf("ChangePswd() = %v, want %v", a.Password, tt.pswd)
			}
			DB.Delete(a)
		})
	}
}
