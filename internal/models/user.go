package models

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

// 用户模型，用于存储用户信息。
type User struct {
	gorm.Model
	AccountType int    // 账号类型
	Username    string // 用户名，结对项目中是手机号
	Password    string // 密码，此处经过 SHA256 加密
	Note        string // 备注，服务端中用于手机验证码
}

// 用户类型枚举。
const (
	UserTypePrimarySchool = iota // 小学
	UserTypeJuniorHigh    = iota // 初中
	UserTypeSeniorHigh    = iota // 高中
)

// 用户类型对应名称。
var AccountTypeText = []string{"小学", "初中", "高中"}

// 当前登录用户，仅限于个人项目 CLI。
var CurrentUser *User

// 存储附表 1 中的账户密码表。
var predefinedUsers = []User{
	{
		AccountType: UserTypePrimarySchool,
		Username:    "张三1",
		Password:    "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3",
	},
	{
		AccountType: UserTypePrimarySchool,
		Username:    "张三2",
		Password:    "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3",
	},
	{
		AccountType: UserTypePrimarySchool,
		Username:    "张三3",
		Password:    "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3",
	},
	{
		AccountType: UserTypeJuniorHigh,
		Username:    "李四1",
		Password:    "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3",
	},
	{
		AccountType: UserTypeJuniorHigh,
		Username:    "李四2",
		Password:    "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3",
	},
	{
		AccountType: UserTypeJuniorHigh,
		Username:    "李四3",
		Password:    "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3",
	},
	{
		AccountType: UserTypeSeniorHigh,
		Username:    "王五1",
		Password:    "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3",
	},
	{
		AccountType: UserTypeSeniorHigh,
		Username:    "王五2",
		Password:    "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3",
	},
	{
		AccountType: UserTypeSeniorHigh,
		Username:    "王五3",
		Password:    "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3",
	},
}

// GetUserByName 使用用户名查找用户。
func GetUserByName(name string) *User {
	var user *User
	result := DB.Where(&User{Username: name}).First(&user)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		log.Printf("GetUserByName error:%v", result.Error)
	}
	return user
}

// ChangeUserType 修改用户类型。
func ChangeUserType(newType string) bool {
	switch newType {
	case "切换为小学":
		CurrentUser.AccountType = UserTypePrimarySchool
	case "切换为初中":
		CurrentUser.AccountType = UserTypeJuniorHigh
	case "切换为高中":
		CurrentUser.AccountType = UserTypeSeniorHigh
	default:
		return false // 不合法的学习等级
	}
	return true // 修改成功
}

// ClearUser 清空当前登录用户，仍然仅在 CLI 下有效。
func ClearUser() {
	CurrentUser = nil
}

// User.GetUserTypeText 获取用户类型对应的文本。
func (a User) GetUserTypeText() string {
	return AccountTypeText[a.AccountType]
}

func CreateUser(phone string, code int) bool {
	var user = User{
		Username: phone,
		Note:     fmt.Sprint(code),
	}
	result := DB.Create(&user)
	if result.Error != nil {
		log.Printf("CreateUser error:%v", result.Error)
		return false
	}
	log.Printf("CreateUser success:%v", user)
	return true
}
