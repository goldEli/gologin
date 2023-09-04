package controllers

import (
	"gologin/dto"
	"gologin/response"
	"gologin/service"
	"gologin/utils"
	"gologin/vo"
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
		ctx.JSON(http.StatusBadRequest, response.ResponseCodeInternalServerError)
	}

	ctx.JSON(http.StatusOK, response.ResponseNoData(response.ResponseCodeOk))
}

func Login(ctx *gin.Context) {

	var user vo.LoginVo

	ctx.ShouldBindJSON(&user)

	if err := utils.Validate.Struct(user); err != nil {
		response.FailWithMessage(utils.Translate(err), ctx)
		return
	}

	tokenString, code := service.Login(user.Email, user.Password)

	if code != response.ResponseCodeOk {
		response.FailWithCode(code, ctx)
		return
	}

	dto := dto.LoginDto{
		Jwt:   tokenString,
		Email: user.Email,
	}

	response.OkWithData(dto, ctx)
}

func Logout(ctx *gin.Context) {
	// Set the same-site mode to lax
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", "", -1, "", "localhost", false, true)
	ctx.JSON(200, gin.H{"data": "You are logged out!"})
}
