package main

import (
	"fmt"
	"k12-math-paper-generator/internal/middleware"
	"k12-math-paper-generator/internal/models"
	"log"
	"regexp"
)

func main() {
	auth()
	inputHandler()
}

// auth 用于用户登录
func auth() {
	var username, password string
	fmt.Printf("请输入用户名和密码\n")
	for {
		fmt.Scanf("%s %s\n", &username, &password)
		if middleware.Auth(username, password) {
			fmt.Printf("当前选择为%v出题\n", models.CurrentUser.GetUserTypeText()) // 根据当前用户类型选择教育等级名称
			break
		} else {
			fmt.Printf("请输入正确的用户名、密码\n")
		}
	}
}

// inputHandler 用于处理输入，并调用相应的函数
func inputHandler() {
	middleware.ReadAllProblems() // 读取所有题目，准备查重
	regex, err := regexp.Compile(`切换为..`)
	if err != nil {
		log.Printf("正则表达式匹配错误：%v", err)
	}
	for {
		fmt.Printf("准备生成%v数学题目，请输入生成题目数量（输入-1将退出当前用户，重新登录）：\n", models.CurrentUser.GetUserTypeText())
		var num int
		var tmp string
		fmt.Scanf("%s\n", &tmp)
		matchChangeType := regex.MatchString(tmp) // 使用正则表达式判断是否切换用户类型
		if matchChangeType {                      // 切换用户类型
			ok := models.ChangeUserType(tmp)
			for !ok {
				fmt.Printf("请输入小学、初中和高中三个选项中的一个\n")
				fmt.Scanf("%s\n", &tmp)
				ok = models.ChangeUserType(tmp)
			}
			fmt.Printf("准备生成%v数学题目，请输入生成题目数量\n", models.CurrentUser.GetUserTypeText())
			fmt.Scanf("%s\n", &tmp)
		}
		fmt.Sscanf(tmp, "%d", &num)
		if num == -1 {
			models.ClearUser() // 清除当前用户
			auth()             // 重新登录
			continue
		}
		if num > 30 || num < 10 {
			fmt.Printf("题目数量不合法\n")
			continue
		}
		middleware.GetProblems(num)
	}
}
