package models

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
)

// ReadProblemsIntoDb 将给定路径的单个文件中的题目读入数据库，同时记录用户信息
func ReadProblemsIntoDb(path string, usr string) {
	file, err := os.ReadFile(path)
	problemPattern := regexp.MustCompile(`\d\d?\..+\n\n`) // 匹配题目
	if err != nil {
		log.Printf("打开文件失败: %v", err)
	}
	fileStr := string(file)
	problems := problemPattern.FindAllString(fileStr, -1) // 找到所有题目
	for _, p := range problems {
		WriteProblemToDb(p[3:len(p)-2], 0, usr) // 去掉换行符和行首的题号，将题目写入数据库
	}
}

// WriteProblemsToFile 将题目写入文件
func WriteProblemsToFile(problem []string) {
	t := time.Now()                                                                         // 获取当前时间
	path := fmt.Sprintf("%v/%v.txt", CurrentUser.Username, t.Format("2006-01-02-15-04-05")) // 用户名/年-月-日-时-分-秒
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
