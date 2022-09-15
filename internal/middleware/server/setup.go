// server 包定义路由及其处理函数。
package server

import (
	"fmt"
	frontend "k12-math-paper-generator/frontend/dist"
	"k12-math-paper-generator/internal/models"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Setup() {
	r := gin.Default() // 创建 Gin 实例

	store := cookie.NewStore([]byte("secret")) // 使用 cookie 存储 session

	// r.StaticFS(path.Join(""), http.FS(frontend.FS)) // 服务静态文件
	r.Use(frontendHandler) // 先经过 handler 判断是否为前端请求
	r.Use(cors.Default())  // 允许跨域

	api := r.Group("/api")                         // 设置 API 路由组
	api.Use(sessions.Sessions("mysession", store)) // 设置 Gin 中间件

	api.GET("/ping", func(c *gin.Context) { // 测试 API
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 用户相关路由
	api.POST("/user", createUser)             // 创建用户
	api.POST("/user/pswd", createPswd)        // 创建密码
	api.GET("/user/session", login)           // 登录
	api.DELETE("/user/session", auth, logout) // 登出
	api.PATCH("/user", auth, changePswd)      // 修改密码

	// 试卷相关路由
	api.GET("/paper", auth, getPaper) // 获取试卷
	api.POST("/paper", auth, judge)   // 判卷

	r.Run(":" + models.ServerConf.Port) // 启动服务器

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
}

func frontendHandler(c *gin.Context) {
	if strings.HasPrefix(c.Request.URL.Path, "/api") { // API 请求，不处理
		c.Next()
		return
	}
	// 返回 frontend.FS 中的文件
	c.FileFromFS(c.Request.URL.Path, http.FS(frontend.FS))
	c.Abort()
}
