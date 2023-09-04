package router

import (
	"gologin/controllers"
	"gologin/middlewares"

	"github.com/gin-gonic/gin"
)

func Post(r *gin.Engine) {

	post := r.Group("/post", middlewares.RequireAuth())
	// post.POST("/", controllers.CreatePost)
	post.GET("/", controllers.GetPosts)
	// post.GET("/:id", controllers.GetPost)
	// post.PUT("/:id", controllers.UpdatePost)
	// post.DELETE("/:id", controllers.DeletePost)
}
