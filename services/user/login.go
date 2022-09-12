package user

import (
	"k12-math-paper-generator/internal/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginService struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (s *LoginService) Login(c *gin.Context) {
	user := models.GetUserByName(s.Username)
	if user.Password != s.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": -1,
			"msg":  "用户名或密码错误",
		})
		c.Abort()
		return
	}
	session := sessions.Default(c)
	session.Set("user", user.Username)
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "登录成功",
	})
}
