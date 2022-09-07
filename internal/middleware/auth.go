package middleware

import "k12-math-paper-generator/internal/models"

func Auth(username, password string) bool {
	user := models.GetUserByName(username)
	if user.Password == password {
		models.CurrentUser = user
		return true
	}
	return false
}
