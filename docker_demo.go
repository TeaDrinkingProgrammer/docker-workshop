package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

const (
	port = ":8080"
)

var calls = 0

func HelloWorld(c *gin.Context) {
	calls++
	c.String(200, "Hello, world! You have called me %d times.\n", calls)
}

func main() {
	router := gin.Default()
	router.GET("/", HelloWorld)
	
	fmt.Printf("Started server at http://localhost%v.\n", port)
	if err := router.Run(port); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}