package models

import (
	"errors"
)

// 用户模型，用于存储用户信息
type User struct {
	AccountType int    // 账号类型
	Username    string // 用户名
	Password    string // 密码
	Note        string // 备注
}

const (
	UserTypePrimarySchool = iota // 小学
	UserTypeJuniorHigh    = iota // 初中
	UserTypeSeniorHigh    = iota // 高中
)

var AccountTypeText = []string{"小学", "初中", "高中"}

var CurrentUser *User

// 存储附表 1 中的账户密码表。
var predefinedUsers = []User{
	{UserTypePrimarySchool, "张三1", "123", ""},
	{UserTypePrimarySchool, "张三2", "123", ""},
	{UserTypePrimarySchool, "张三3", "123", ""},
	{UserTypeJuniorHigh, "李四1", "123", ""},
	{UserTypeJuniorHigh, "李四2", "123", ""},
	{UserTypeJuniorHigh, "李四3", "123", ""},
	{UserTypeSeniorHigh, "王五1", "123", ""},
	{UserTypeSeniorHigh, "王五2", "123", ""},
	{UserTypeSeniorHigh, "王五3", "123", ""},
}

// 使用用户名查找用户，保留后期接入数据库的可扩展性。
func GetUserByName(name string) (User, error) {
	for _, usr := range predefinedUsers {
		if usr.Username == name {
			return usr, nil
		}
	}
	return User{}, errors.New("user not found")
}

func ChangeUserType(newType string) {
	switch newType {
	case "小学":
		CurrentUser.AccountType = UserTypePrimarySchool
	case "初中":
		CurrentUser.AccountType = UserTypeJuniorHigh
	case "高中":
		CurrentUser.AccountType = UserTypeSeniorHigh
	}
}

func ClearUser() {
	CurrentUser = nil
}
