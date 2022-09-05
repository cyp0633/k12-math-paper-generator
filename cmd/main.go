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
	getProblems()
}

func auth() {
	var username, password string
	fmt.Printf("请输入用户名和密码\n")
	fmt.Scanf("%s %s\n", &username, &password)
	if middleware.Auth(username, password) {
		fmt.Printf("当前选择为%v出题\n", models.AccountTypeText[models.CurrentUser.AccountType]) // 根据当前用户类型选择教育等级名称
	} else {
		fmt.Printf("请输入正确的用户名、密码\n")
		auth() // 使用递归重新验证。Go 的递归深度是无限的，所以没有问题
	}
}

func getProblems() {
	fmt.Printf("准备生成%v数学题目，请输入生成题目数量（输入-1将退出当前用户，重新登录）：\n", models.AccountTypeText[models.CurrentUser.AccountType])
	var num int
	var tmp string
	fmt.Scanf("%s\n", &tmp)
	matchChangeType, err := regexp.MatchString(`切换为[a-zA-Z]+`, tmp) // 使用正则表达式判断是否切换用户类型
	if err != nil {
		log.Printf("正则表达式匹配错误：%v", err)
	}
	if matchChangeType {
		models.ChangeUserType(tmp[:])
		models.ClearUser()
		getProblems()
	}
	fmt.Sscanf(tmp, "%d", &num)
	if num == -1 {
		models.ClearUser()
		auth()        // 重新登录
		getProblems() // 重新获取题目
	} else {
		return // TODO: 生成题目
	}
}
