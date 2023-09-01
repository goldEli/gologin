package controllers

import (
	"gologin/dto"
	"gologin/models"
	"gologin/response"
	"gologin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var body struct {
		Name     string
		Email    string
		Password string
	}

	if ctx.BindJSON(&body) != nil {
		ctx.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	err := service.Register(body.Name, body.Email, body.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResponseError())
	}

	ctx.JSON(http.StatusOK, response.ResponseOk())
}

func Login(ctx *gin.Context) {

	var user struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if ctx.BindJSON(&user) != nil {
		ctx.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	tokenString, err := service.Login(&models.User{Email: user.Email, Password: user.Password})

	if err != nil {
		ctx.JSON(500, gin.H{"error": "error signing token"})
		return
	}

	dto := dto.LoginDto{
		Jwt:   tokenString,
		Email: user.Email,
	}

	ctx.JSON(http.StatusOK, response.ResponseOkWIthData(dto))

}

func Logout(ctx *gin.Context) {
	// Set the same-site mode to lax
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", "", -1, "", "localhost", false, true)
	ctx.JSON(200, gin.H{"data": "You are logged out!"})
}

func Validate(ctx *gin.Context) {
	user, exists := ctx.Get("user")

	if !exists {
		ctx.JSON(500, gin.H{"error": "error"})
		return
	}
	ctx.JSON(200, gin.H{"data": "You are logged in!", "user": user})
}
