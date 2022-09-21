// paper 包用于处理试卷相关的请求，如出题与判分等。
package paper

import (
	"k12-math-paper-generator/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JudgeService 前端传入的判分请求格式。
type JudgeService struct {
	UserAnswers []float64 `json:"answers" binding:"required"`
}

// Judge 对照下层传入的答案，判分
func (s *JudgeService) Judge(c *gin.Context) {
	usr := c.GetString("user")
	ans := models.Answers(usr)
	if len(ans) != len(s.UserAnswers) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "答案数量不匹配",
		})
		return
	}
	var score float64
	for i := range ans {
		if ans[i] == s.UserAnswers[i] {
			score += 1.0
		}
	}
	score = setPrecision(score/float64(len(ans))*100.0, 2)
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"score": score,
	})
}
