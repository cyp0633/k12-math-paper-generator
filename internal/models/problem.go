package models

import (
	"fmt"
	"math"
)

// Expression 描述数学表达式树的某一个节点。
type Expression struct {
	Left     *Expression // 左子表达式
	Right    *Expression // 右子表达式
	Operator int         // 运算符
	Value    float64     // 值
	Level    int         // 难度等级，与账户类型对应
}

// 运算符。一元运算符统一只有右子表达式。
const (
	operatorAdd    = iota // 加法
	operatorSub    = iota // 减法
	operatorMul    = iota // 乘法
	operatorDiv    = iota // 除法
	operatorSqrt   = iota // 开方
	operatorSquare = iota // 平方
	operatorSin    = iota // 正弦
	operatorCos    = iota // 余弦
	operatorTan    = iota // 正切
	operatorNone   = -1   // 无运算符
)

// opPriority 运算符优先级。
var opPriority = []int{1, 1, 2, 2, 3, 3, 4, 4, 4}

// GenerateProblem 生成一个表达式树。
// op 操作数
// 返回值：生成的表达式
func GenerateProblem(op int) *Expression {
	var exp Expression
	var leftOpNum int
	if op == 0 { // 表达式不含运算符，生成一个数字
		exp.Left, exp.Right = nil, nil
		exp.Operator = -1
		// exp.HasJuniorLevel, exp.HasSeniorLevel = false, false
		exp.Level = UserTypePrimarySchool
		exp.Value = float64(GenNum(1, 101)) // 生成一个 1 到 100 的随机数
		return &exp
	}
	// 确定该表达式的学习等级
	if exp.Left != nil {
		exp.Level = exp.Left.Level
	}
	if exp.Right != nil {
		if exp.Level < exp.Right.Level {
			exp.Level = exp.Right.Level
		}
	}
	exp.Operator = genOp()
	switch exp.Operator { // 根据运算符的不同，生成不同的表达式
	case operatorAdd, operatorSub, operatorMul, operatorDiv: // 二元运算符
		leftOpNum = GenNum(0, op)                       // 分配 0~op-1 个运算符给左子表达式
		exp.Left = GenerateProblem(leftOpNum)           // 使用递归生成左子表达式
		exp.Right = GenerateProblem(op - leftOpNum - 1) // 使用递归生成右子表达式
	case operatorSqrt, operatorSquare: // 一元运算符，初中难度
		exp.Left = nil
		exp.Right = GenerateProblem(op - 1) // 使用递归生成右子表达式
		if exp.Level < UserTypeJuniorHigh {
			exp.Level = UserTypeJuniorHigh
		}
	case operatorSin, operatorCos, operatorTan: // 一元运算符，高中难度
		exp.Left = nil
		exp.Right = GenerateProblem(op - 1) // 使用递归生成右子表达式
		if exp.Level < UserTypeSeniorHigh {
			exp.Level = UserTypeSeniorHigh
		}
	}
	switch exp.Operator { // 计算表达式的值，此处暂时忽略不合法运算，留到后面处理 NaN
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
	case operatorSquare:
		exp.Value = exp.Right.Value * exp.Right.Value
	case operatorSin:
		exp.Value = math.Sin(exp.Right.Value) // 三角函数都是弧度制
	case operatorCos:
		exp.Value = math.Cos(exp.Right.Value)
	case operatorTan:
		exp.Value = math.Tan(exp.Right.Value)
	}
	return &exp
}

// GenerateProblemStr 对表达式树进行中序遍历，生成表达式字符串。
func GenerateProblemStr(root *Expression) string {
	left, right := "", ""
	// 视情况为子表达式加括号（不再为自己加括号！）
	if root.Left != nil {
		left = GenerateProblemStr(root.Left)
		if root.Left.Operator != operatorNone && opPriority[root.Left.Operator] < opPriority[root.Operator] {
			left = "(" + left + ")"
		}
	}
	if root.Right != nil {
		right = GenerateProblemStr(root.Right)
		if (root.Right.Operator != operatorNone && root.Operator != operatorSqrt && opPriority[root.Right.Operator] < opPriority[root.Operator]) || // 开方特判，因为 LaTeX 格式不能像平常一样加括号
			(root.Right.Operator != operatorNone && opPriority[root.Operator] == 4) { // 三角函数特判，除非里面是个常数，否则都要加括号。三角函数是单目运算符，所以只需要处理右侧
			right = "(" + right + ")"
		}
	}
	// 连接左右表达式字符串与运算符
	switch root.Operator {
	case operatorAdd:
		return left + " + " + right
	case operatorSub:
		return left + " - " + right
	case operatorMul:
		return left + " * " + right
	case operatorDiv:
		return left + " / " + right
	case operatorSqrt:
		return "\\sqrt{" + right + "}" // LaTeX 格式
	case operatorSquare:
		return right + " ^ 2"
	case operatorSin:
		return "sin" + right
	case operatorCos:
		return "cos" + right
	case operatorTan:
		return "tan" + right
	case operatorNone:
		return fmt.Sprintf("%v", root.Value) // 仅起到转换为字符串的作用
	}
	return ""
}
