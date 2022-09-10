package middleware

import "k12-math-paper-generator/internal/models"

// Auth 通过用户名和密码验证用户，并设置当前登录用户
func Auth(username, password string) bool {
	user := models.GetUserByName(username)
	if user != nil && user.Password == password {
		models.CurrentUser = user
		return true
	}
	return false
}
