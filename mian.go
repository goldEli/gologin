package main

import (
	"gologin/controllers"
	"gologin/inits"
	"gologin/middlewares"

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

	r.POST("user", middlewares.RequireAuth, controllers.Validate)
	r.POST("/users", controllers.Signup)
	r.POST("/users/login", controllers.Login)

	r.Run()
}
