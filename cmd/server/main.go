package main

import (
	"fmt"
	frontend "k12-math-paper-generator/frontend/dist"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Printf("Hello, world.")
	r := gin.Default()
	r.StaticFS(path.Join(""), http.FS(frontend.FS))
	r.Run(":8000")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
}
