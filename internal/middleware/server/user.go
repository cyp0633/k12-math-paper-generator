package server

import (
	"k12-math-paper-generator/internal/services/user"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func createUser(c *gin.Context) {
	var s user.CreateUserService
	err := c.BindJSON(&s)
	if err != nil { // 未绑定上 JSON，请求格式错误
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "请求格式错误",
		})
		c.Abort()
		return
	}
	s.Create(c)
}

// login 使用用户名和密码登录，并将用户信息存入 Session。
func login(c *gin.Context) {
	var s user.LoginService
	err := c.BindJSON(&s)
	if err != nil { // 未绑定上 JSON，请求格式错误
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "请求格式错误",
		})
		c.Abort()
		return
	}
	s.Login(c)
}

// logout 清除 Session 中的用户信息。
func logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("user")
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "登出成功",
	})
}

// createPswd 创建密码，用于刚刚注册的阶段，同时校验手机验证码。
func createPswd(c *gin.Context) {
	var s user.CreatePswdService
	err := c.BindJSON(&s)
	if err != nil { // 未绑定上 JSON，请求格式错误
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "请求格式错误",
		})
		c.Abort()
		return
	}
	s.Create(c)
}

// changePswd 修改密码。
func changePswd(c *gin.Context) {
	var s user.ChangePasswordService
	err := c.BindJSON(&s)
	if err != nil { // 未绑定上 JSON，请求格式错误
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "请求格式错误",
		})
		c.Abort()
		return
	}
	s.ChangePassword(c)
}

// auth 验证用户是否登录。
func auth(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": -1,
			"msg":  "未登录",
		})
		c.Abort()
		return
	}
	c.Set("user", user)
}

// getStats 获取用户做题统计数据。
func getStats(c *gin.Context) {
	user.GetProblemStats(c)
}
