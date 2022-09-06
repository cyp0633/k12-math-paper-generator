package middleware

import (
	"fmt"
	"k12-math-paper-generator/internal/models"
)

// GetProblems 用于生成题目
// num: 题目数量
func GetProblems(num int) {
	for i := 0; i < num; {
		problem := models.GenerateProblem(models.GenNum(1, 6)) // 生成一个 1~5 个操作数的题目
		problemStr := models.GenerateProblemStr(problem)       // 生成题目字符串
		fmt.Printf("%v\n", problemStr)
		i++
	}
}
