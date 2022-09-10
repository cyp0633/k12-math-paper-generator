package models

// 用户模型，用于存储用户信息。
type User struct {
	AccountType int    // 账号类型
	Username    string // 用户名
	Password    string // 密码
	Note        string // 备注
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
func GetUserByName(name string) *User {
	for _, usr := range predefinedUsers {
		if usr.Username == name {
			return &usr
		}
	}
	return nil
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
