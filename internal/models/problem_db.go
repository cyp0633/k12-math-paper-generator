package models

import (
	"log"

	"gorm.io/gorm"
)

// Problem 描述一个题目在数据库中的存储形式。
type Problem struct {
	gorm.Model
	User       string  // 用户名
	Expression string  // 题目表达式
	Answer     float64 // 答案
	FilePath   string  // 文件路径，暂时没卵用
}

// CheckDupProblem 检查是否有重复的题目
func CheckDupProblem(problem string) bool {
	result := DB.Where(&Problem{Expression: problem, User: CurrentUser.Username}).Find(&Problem{}) // 用 Find 而非 First 以避免不必要的日志输出；只搜索本用户的
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		log.Printf("查询数据库失败: %v", result.Error)
	}
	if result.RowsAffected == 0 { // 没有找到重复题目
		return false
	}
	return true
}

// WriteProblemToDb 将题目信息写入数据库
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

// ClearProblems 将题目数据库清空，防止两次打开之间文件变化
func ClearProblems(username string) bool {
	result := DB.Where(&Problem{User: username}).Unscoped().Delete(&Problem{})
	if result.Error != nil {
		log.Printf("清空数据库失败: %v", result.Error)
		return false
	}
	return true
}
