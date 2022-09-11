package server

import (
	"k12-math-paper-generator/internal/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func createUser(c *gin.Context) {}

// login 使用用户名和密码登录，并将用户信息存入 Session。
func login(c *gin.Context) {
	user := models.GetUserByName(c.PostForm("username"))
	if user.Password != c.PostForm("password") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": -1,
			"msg":  "用户名或密码错误",
		})
		c.Abort()
	}
	session := sessions.Default(c)
	session.Set("user", user)
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "登录成功",
	})
}

// logout 清除 Session 中的用户信息。
func logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": -1,
			"msg":  "未登录",
		})
		c.Abort()
	}
	session.Delete("user")
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "登出成功",
	})
}
