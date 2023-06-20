package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")
	router := gin.Default()

	// Router tanımlamaları buraya gelecek

	router.Run(":8080")
}
