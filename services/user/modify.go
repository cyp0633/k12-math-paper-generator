package user

import (
	"k12-math-paper-generator/internal/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type ChangePasswordService struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

func (s *ChangePasswordService) ChangePassword(c *gin.Context) {
	session := sessions.Default(c)
	user := models.GetUserByName(session.Get("user").(string))
	if user.Password != s.OldPassword {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": -1,
			"msg":  "用户名或密码错误",
		})
		c.Abort()
		return
	}
	ok := user.ChangePswd(s.NewPassword)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  "修改密码失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "修改密码成功",
	})
}
