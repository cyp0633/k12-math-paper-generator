package paper

import (
	"k12-math-paper-generator/internal/models"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPaperService 描述了前端请求生成试卷的 JSON 格式。
type GetPaperService struct {
	Level int `json:"level" form:"level" binding:"required"` // 难度等级
	Num   int `json:"num" form:"num" binding:"required"`     // 题目数量
}

// problemPost 是特定某道题题面的格式。
type problemPost struct {
	ProblemStr string     `json:"problem_str"` // 题干
	Options    [4]float64 `json:"options"`     // 选项
}

// GetPaper 生成试卷并返回给前端。
func (s *GetPaperService) GetPaper(c *gin.Context) {
	var pSet []problemPost
	s.Level-- // 前端传入的参数 1 代表小学，以此类推；0 会触发绑定错误，暂不清楚原因
	if s.Level < 0 || s.Level > 2 || s.Num < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "请求参数错误",
		})
	}
	user := c.GetString("user")
	ok := models.ClearProblems(user)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  "数据库错误",
		})
		c.Abort()
		return
	}
	for i := 0; i < s.Num; {
		var l string
		var o [4]float64
		problem := models.GenerateProblem(models.GenNum(1, 6), s.Level)
		for problem.Level != s.Level || math.IsNaN(problem.Value) || math.IsInf(problem.Value, 0) {
			problem = models.GenerateProblem(models.GenNum(1, 6), s.Level)
		}
		str := models.GenerateProblemStr(problem)
		if models.CheckDupProblem(str, user) {
			continue
		}
		l = str
		models.WriteProblemToDb(str, problem.Value, user)
		for j := 0; j < 4; j++ {
			o[j] = float64(models.GenNum(1, 100000) / 1000.0) // 随机生成选项数字
		}
		o[models.GenNum(0, 3)] = problem.Value // 随机将答案放入选项中
		var p = problemPost{
			ProblemStr: l,
			Options:    o,
		}
		pSet = append(pSet, p)
		i++
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": pSet,
	})
}
