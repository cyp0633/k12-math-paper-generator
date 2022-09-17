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
	if s.Level < 0 || s.Level > 2 || s.Num < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "请求参数错误",
		})
	}
	user, _ := c.Get("user")
	ok := models.ClearProblems(user.(string))
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  "数据库错误",
		})
	}
	for i := 0; i < s.Num; {
		var l string
		var o [4]float64
		problem := models.GenerateProblem(models.GenNum(1, 6))
		for problem.Level != s.Level || math.IsNaN(problem.Value) || math.IsInf(problem.Value, 0) {
			problem = models.GenerateProblem(models.GenNum(1, 6))
		}
		str := models.GenerateProblemStr(problem)
		if models.CheckDupProblem(str) {
			continue
		}
		l = str
		models.WriteProblemToDb(user.(string), problem.Value, str) // TODO: 如何实现快速查询答案？
		for j := 0; j < 4; j++ {
			o[j] = float64(models.GenNum(1, 100000) / 1000.0) // 随机生成选项数字
		}
		o[models.GenNum(0, 3)] = problem.Value // 随机将答案放入选项中
		var p = problemPost{
			ProblemStr: l,
			Options:    o,
		}
		pSet = append(pSet, p)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": pSet,
	})
}
