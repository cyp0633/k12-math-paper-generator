package main

import (
	"fmt"
	"k12-math-paper-generator/frontend"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Printf("Hello, world.")
	dir, err := frontend.FS.ReadDir("dist")
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range dir {
		fmt.Println(v.Name())
	}
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Run(":8080")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
}
