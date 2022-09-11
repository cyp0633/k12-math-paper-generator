package main

import (
	"fmt"
	"k12-math-paper-generator/internal/middleware/server"
	_ "k12-math-paper-generator/internal/models"
)

func main() {
	fmt.Printf("Hello, world.")
	server.Setup()
}
