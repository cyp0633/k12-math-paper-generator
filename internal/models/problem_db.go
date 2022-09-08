package models

import (
	"log"

	"gorm.io/gorm"
)

type Problem struct {
	gorm.Model
	User       string  // 用户名
	Expression string  // 题目表达式
	Answer     float64 // 答案
	FilePath   string  // 文件路径
}

// CheckDupProblem 检查是否有重复的题目
func CheckDupProblem(problem string) bool {
	result := DB.Where(&Problem{Expression: problem, User: CurrentUser.Username}).Find(&Problem{})
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		log.Printf("查询数据库失败: %v", result.Error)
	}
	if result.RowsAffected == 0 { // 没有找到重复题目
		return false
	}
	return true
}

func WriteProblemToDb(problem string, ans float64, usr string) {
	var p = Problem{
		User:       usr,
		Expression: problem,
		Answer:     ans,
	}
	result := DB.Create(&p)
	if result.Error != nil {
		log.Printf("写入数据库失败: %v", result.Error)
	}
}

func ClearProblems() {
	result := DB.Where("1=1").Unscoped().Delete(&Problem{})
	if result.Error != nil {
		log.Printf("清空数据库失败: %v", result.Error)
	}
}
