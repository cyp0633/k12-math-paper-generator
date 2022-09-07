package models

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Problem struct {
	User       string  // 用户名
	Expression string  // 题目表达式
	Answer     float64 // 答案
}

// CheckDupProblem 检查是否有重复的题目
func CheckDupProblem(problem string) bool {
	// TODO: 检查是否有重复的题目
	return false
}

func WriteProblemToDb(problem string, ans float64) {
	var _ = Problem{
		User:       CurrentUser.Username,
		Expression: problem,
		Answer:     ans,
	}
}

// WriteProblemsToFile 将题目写入文件
func WriteProblemsToFile(problem []string) {
	t := time.Now()                                                                                                                       // 获取当前时间
	path := fmt.Sprintf("%v/%v-%v-%v-%v-%v-%v.txt", CurrentUser.Username, t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()) // 用户名/年-月-日-时-分-秒
	err := os.MkdirAll(CurrentUser.Username, os.ModePerm)
	if err != nil {
		log.Printf("创建文件夹失败: %v", err)
	}
	file, err := os.ReadFile(path)
	if err != nil && !os.IsNotExist(err) {
		log.Printf("打开文件失败: %v", err)
	}
	for i, p := range problem {
		problemStr := fmt.Sprintf("%v. %v\n\n", i+1, p) // 题号，内容，换行
		file = append(file, []byte(problemStr)...)
	}
	err = os.WriteFile(path, file, os.ModePerm)
	if err != nil {
		log.Printf("写入文件失败: %v", err)
	}
}
