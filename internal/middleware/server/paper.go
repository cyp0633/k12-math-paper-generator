package server

import (
	"k12-math-paper-generator/services/paper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getPaper(c *gin.Context) {
	var s paper.GetPaperService
	err := c.BindJSON(&s)
	if err != nil { // 未绑定上 JSON，请求格式错误
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "请求格式错误",
		})
		c.Abort()
		return
	}
	s.GetPaper(c)
}
