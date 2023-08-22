package main

import (
	"gologin/controllers"
	"gologin/inits"

	"github.com/gin-gonic/gin"
)

// this init function will be called automatically and cannot have any parameters
func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	r := gin.Default()

	r.POST("/", controllers.CreatePost)
	r.GET("/", controllers.GetPosts)
	r.GET("/:id", controllers.GetPost)
	r.PUT("/:id", controllers.UpdatePost)
	r.DELETE("/:id", controllers.DeletePost)

	r.Run()
}
