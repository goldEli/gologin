package router

import (
	"gologin/controllers"

	"github.com/gin-gonic/gin"
)

func User(r *gin.Engine) {
	user := r.Group("/users")
	user.POST("/validate", controllers.Validate)
	user.POST("/", controllers.Signup)
	user.POST("/login", controllers.Login)
}
