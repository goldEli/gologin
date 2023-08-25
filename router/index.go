package router

import (
	"gologin/middlewares"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()

	// 设置跨域
	r.Use(middlewares.Cors())

	Post(r)
	User(r)

	r.Run()
}
