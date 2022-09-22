package user

import (
	"k12-math-paper-generator/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProblemStats(c *gin.Context) {
	usr := c.GetString("user")
	u := models.GetUserByName(usr)
	total, correct, ok := (*u).ProblemRecord()
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  "获取做题记录失败",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"total":   total,
		"correct": correct,
	})
}
