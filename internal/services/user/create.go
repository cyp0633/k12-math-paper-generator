package user

import (
	"k12-math-paper-generator/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CreateUserService 用于注册时传入 JSON 的绑定。
type CreateUserService struct {
	Username string `json:"username" binding:"required"`
}

// Create 使用电话号码创建用户。
func (s *CreateUserService) Create(c *gin.Context) {
	u := models.GetUserByName(s.Username)
	if u != nil {
		c.JSON(http.StatusConflict, gin.H{
			"code": -1,
			"msg":  "用户已存在",
		})
		return
	}
	code := models.GenNum(1000, 9999)
	var ok bool
	if _, err := strconv.Atoi(s.Username); err == nil { // 无错误时均为数字，代表为手机号
		ok = models.SendSms(s.Username, code)
	} else { // 否则视为邮箱
		ok = models.SendMail(s.Username, code)
	}
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  "发送验证码失败",
		})
		c.Abort()
		return
	}
	ok = models.CreateUser(s.Username, code)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  "创建用户失败",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "发送验证码成功",
	})
}

// CreatePswdService 用于设置密码时传入 JSON 的绑定。
type CreatePswdService struct {
	Phone    string `json:"username" binding:"required"`
	Verify   string `json:"verify" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (s *CreatePswdService) Create(c *gin.Context) {
	u := models.GetUserByName(s.Phone)
	if u == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "用户不存在",
		})
		return
	}
	if u.Note != s.Verify { // 备注（预设验证码）与提供的不相同
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "验证码错误",
		})
		return
	}
	ok := u.ChangePswd(s.Password)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  "创建密码失败",
		})
		return
	}
	session := sessions.Default(c)
	session.Set("user", u.Username)
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "创建密码成功",
	})
}
