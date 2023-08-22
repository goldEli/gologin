package main

import (
	"gologin/inits"

	"github.com/gin-gonic/gin"
)

// this init function will be called automatically and cannot have any parameters
func init() {
	inits.LoadEnv()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	r.Run()
}
