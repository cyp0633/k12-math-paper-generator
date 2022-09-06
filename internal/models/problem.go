package models

import (
	"fmt"
	"math"
)

// Expression 描述一个数学表达式节点。
type Expression struct {
	Left           *Expression // 左子表达式
	Right          *Expression // 右子表达式
	Operator       int         // 运算符
	Value          float64     // 值
	HasJuniorLevel bool        // 是否有初中难度运算
	HasSeniorLevel bool        // 是否有高中难度运算
}

const (
	operatorAdd     = iota // 加法
	operatorSub     = iota // 减法
	operatorMul     = iota // 乘法
	operatorDiv     = iota // 除法
	operatorSqrt    = iota // 开方
	operationSquare = iota // 平方
	operationSin    = iota // 正弦
	operationCos    = iota // 余弦
	operationTan    = iota // 正切
	operationNone   = -1   // 无运算符
)

// GenerateProblem 生成一个表达式树。
// op 操作数
// 返回值：生成的表达式
func GenerateProblem(op int) *Expression {
	var exp Expression
	var leftOpNum int
	if op == 0 { // 表达式不含运算符，生成一个数字
		exp.Left, exp.Right = nil, nil
		exp.Operator = -1
		exp.HasJuniorLevel, exp.HasSeniorLevel = false, false
		exp.Value = float64(GenNum(0, 101)) // 生成一个 0 到 100 的随机数
		return &exp
	}
	// 确定该表达式的学习等级
	if exp.Left != nil {
		exp.HasJuniorLevel = exp.Left.HasJuniorLevel
		exp.HasSeniorLevel = exp.Left.HasSeniorLevel
	}
	if exp.Right != nil {
		exp.HasJuniorLevel = exp.HasJuniorLevel || exp.Right.HasJuniorLevel
		exp.HasSeniorLevel = exp.HasSeniorLevel || exp.Right.HasSeniorLevel
	}
	exp.Operator = genOp()
	switch exp.Operator { // 根据运算符的不同，生成不同的表达式
	case operatorAdd, operatorSub, operatorMul, operatorDiv: // 二元运算符
		leftOpNum = GenNum(0, op)                       // 分配 0~op-1 个运算符给左子表达式
		exp.Left = GenerateProblem(leftOpNum)           // 使用递归生成左子表达式
		exp.Right = GenerateProblem(op - leftOpNum - 1) // 使用递归生成右子表达式
	case operatorSqrt, operationSquare: // 一元运算符，初中难度
		exp.Left = nil
		exp.Right = GenerateProblem(op - 1) // 使用递归生成右子表达式
		exp.HasJuniorLevel = true
	case operationSin, operationCos, operationTan: // 一元运算符，高中难度
		exp.Left = nil
		exp.Right = GenerateProblem(op - 1) // 使用递归生成右子表达式
		exp.HasSeniorLevel = true
	}
	switch exp.Operator { // 计算表达式的值
	case operatorAdd:
		exp.Value = exp.Left.Value + exp.Right.Value
	case operatorSub:
		exp.Value = exp.Left.Value - exp.Right.Value
	case operatorMul:
		exp.Value = exp.Left.Value * exp.Right.Value
	case operatorDiv:
		exp.Value = exp.Left.Value / exp.Right.Value
	case operatorSqrt:
		exp.Value = math.Sqrt(exp.Right.Value)
	case operationSquare:
		exp.Value = exp.Right.Value * exp.Right.Value
	case operationSin:
		exp.Value = math.Sin(exp.Right.Value)
	case operationCos:
		exp.Value = math.Cos(exp.Right.Value)
	case operationTan:
		exp.Value = math.Tan(exp.Right.Value)
	}
	return &exp
}

// GenerateProblemStr 对表达式树进行中序遍历，生成表达式字符串。
func GenerateProblemStr(root *Expression) string {
	var left, right string
	if root.Left != nil {
		left = GenerateProblemStr(root.Left)
	}
	if root.Right != nil {
		right = GenerateProblemStr(root.Right)
	}
	switch root.Operator {
	case operatorAdd:
		return "(" + left + " + " + right + ")"
	case operatorSub:
		return "(" + left + " - " + right + ")"
	case operatorMul:
		return "(" + left + " * " + right + ")"
	case operatorDiv:
		return "(" + left + " / " + right + ")"
	case operatorSqrt:
		return "sqrt(" + right + ")"
	case operationSquare:
		return "(" + right + " ^ 2)"
	case operationSin:
		return "sin(" + right + ")"
	case operationCos:
		return "cos(" + right + ")"
	case operationTan:
		return "tan(" + right + ")"
	case operationNone:
		return fmt.Sprintf("%v", root.Value)
	}
	return ""
}
