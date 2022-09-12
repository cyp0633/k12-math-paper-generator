package user

import (
	"k12-math-paper-generator/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserService struct {
	Phone string `json:"username" binding:"required"`
}

func (s *CreateUserService) Create(c *gin.Context) {
	u := models.GetUserByName(s.Phone)
	if u != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "用户已存在",
		})
		return
	}
	code := models.GenNum(1000, 9999)
	ok := models.SendSms(s.Phone, code)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  "发送验证码失败",
		})
		return
	}
	ok = models.CreateUser(s.Phone, code)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  "创建用户失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "发送验证码成功",
	})
}
