package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// docs is generated by Swag CLI, you have to import it.
	// _ "tzh.com/web/docs"
	_ "gologin/docs" // docs is generated by Swag CLI, you have to import it.
)

// @title 博客系统
// @version 1.0
// @description test to
// @termsOfService https://github.com/go-programming-tour-book

// @contact.name miaoyu
// @contact.url http://coolcat.io/support
// @contact.email miaoyu200@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:5000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func Router() {

	r := gin.Default()

	// 设置跨域
	// r.Use(middlewares.Cors())
	r.Use(cors.Default())

	Post(r)
	User(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
