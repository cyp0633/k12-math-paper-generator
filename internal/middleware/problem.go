package middleware

import (
	"k12-math-paper-generator/internal/models"
)

// GetProblems 用于生成题目
// num: 题目数量
func GetProblems(num int) {
	var str []string
	var ans []float64
	for i := 0; i < num; {
		problem := models.GenerateProblem(models.GenNum(1, 6))
		for problem.Level != models.CurrentUser.AccountType {
			problem = models.GenerateProblem(models.GenNum(1, 6)) // 生成一个 1~5 个操作数的题目
		}
		problemStr := models.GenerateProblemStr(problem) // 生成题目字符串
		if models.CheckDupProblem(problemStr) {          // 检查是否重复
			continue
		}
		str = append(str, problemStr)
		ans = append(ans, problem.Value)
		// fmt.Printf("%v\n", problemStr)
		i++
	}
	models.WriteProblemsToFile(str)
}
