package main

import (
	"fmt"
	"k12-math-paper-generator/internal/middleware/server"
	"k12-math-paper-generator/internal/models"
)

func main() {
	fmt.Printf("Hello, world.")
	models.InitConf() // 读取配置文件
	server.Setup()    // 初始化路由并启动服务器
}
