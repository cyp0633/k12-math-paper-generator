package models

// rand.go 用于生成随机数

import (
	"crypto/rand"
	"log"
	"math/big"
)

// genNum 在 [0, max) 内生成一个随机数
func genNum(max int) int {
	val, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		log.Printf("生成随机数失败：%v", err)
	}
	return int(val.Int64())
}

// genOp 生成一个符合等级的运算符
func genOp() int {
	var maxLevel int
	switch CurrentUser.AccountType {
	case UserTypePrimarySchool:
		maxLevel = 4 // 加减乘除
	case UserTypeJuniorHigh:
		maxLevel = 6 // 平方开方
	case UserTypeSeniorHigh:
		maxLevel = 9 // 三角函数
	default:
		return -1 // 不合法的学习等级
	}
	return genNum(maxLevel)
}
