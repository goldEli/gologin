package router

import (
	"gologin/controllers"
	"gologin/middlewares"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()

	// 设置跨域
	r.Use(middlewares.Cors())

	r.POST("/", controllers.CreatePost)
	r.GET("/", controllers.GetPosts)
	r.GET("/:id", controllers.GetPost)
	r.PUT("/:id", controllers.UpdatePost)
	r.DELETE("/:id", controllers.DeletePost)

	User(r)

	r.Run()
}
