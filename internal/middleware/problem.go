package middleware

import (
	"k12-math-paper-generator/internal/models"
	"log"
	"math"
	"os"
	"regexp"
)

// GetProblems 用于生成题目
// num: 题目数量
func GetProblems(num int) {
	var str []string
	// var ans []float64
	for i := 0; i < num; {
		problem := models.GenerateProblem(models.GenNum(1, 6))
		for problem.Level != models.CurrentUser.AccountType || math.IsNaN(problem.Value) || math.IsInf(problem.Value, 0) { // 难度不符合，或答案 NaN（出现不合法），重新生成
			problem = models.GenerateProblem(models.GenNum(1, 6)) // 生成一个 1~5 个操作数的题目
		}
		problemStr := models.GenerateProblemStr(problem) // 生成题目字符串
		if models.CheckDupProblem(problemStr) {          // 检查是否重复
			continue
		}
		str = append(str, problemStr)
		// ans = append(ans, problem.Value)
		// fmt.Printf("%v\n", problemStr)
		models.WriteProblemToDb(problemStr, problem.Value, models.CurrentUser.Username) // 将题目写入数据库
		i++
	}
	models.WriteProblemsToFile(str)
}

// ReadAllProblems 将子文件夹内的所有题目读取到数据库中
func ReadAllProblems() {
	ok := models.ClearProblems(models.CurrentUser.Username) // 清空数据库；TODO 更好的查重方案
	if !ok {
		log.Println("清空数据库失败")
	}
	dirs, err := os.ReadDir("./")
	if err != nil {
		log.Printf("读取文件夹失败: %v", err)
	}
	pattern, err := regexp.Compile(`/\d...-\d.-\d.-\d.-\d.-\d.\.txt`) // 匹配日期的文件名规则
	if err != nil {
		log.Printf("正则表达式编译失败: %v", err)
	}
	for _, dir := range dirs { // 遍历子文件/文件夹
		if !dir.IsDir() || models.GetUserByName(dir.Name()) == nil { // 不是目录，或该子目录并不是用户的，跳过
			continue
		}
		files, err := os.ReadDir(dir.Name())
		if err != nil {
			log.Printf("读取文件夹失败: %v", err)
		}
		for _, file := range files { // 遍历子文件/文件夹
			if file.IsDir() || !pattern.MatchString(file.Name()) { // 不是文件，或文件名不符合规范，跳过
				continue
			}
			models.ReadProblemsIntoDb(dir.Name()+"/"+file.Name(), dir.Name()) // 对某个特定的文件进行读取
		}
	}
}
