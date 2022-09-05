package middleware

import "k12-math-paper-generator/internal/models"

func Auth(username, password string) bool {
	user, err := models.GetUserByName(username)
	if err != nil {
		return false
	}
	if user.Password == password {
		models.CurrentUser = &user
		return true
	}
	return false
}
