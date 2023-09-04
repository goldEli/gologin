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

// @Summary 注册
// @Description 账户注册
// @Tags user
// @Accept  json
// @Produce json
// @Param body body vo.RegisterVo true "用户注册"
// @Success 200  {object}  object{code=number,message=string}
// @Router /users/register [post]
func Register(ctx *gin.Context) {

	var data vo.RegisterVo

	ctx.ShouldBindJSON(&data)

	if errMsg := utils.GetErrorMessage(data); errMsg != "" {
		response.FailWithMessage(errMsg, ctx)
		return
	}

	err := service.Register(data.Name, data.Email, data.Password)
	if err != nil {
		response.FailServer(ctx)
		return
	}

	response.Ok(ctx)
}

// @Summary 登录
// @Description 登录账户, 获取 token
// @Tags user
// @Accept  json
// @Produce  json
// @Param body body vo.LoginVo true "User login""
// @Success 200  {object}  object{data=dto.LoginDto,code=number,message=string}
// @Router /users/login [post]
func Login(ctx *gin.Context) {
	var user vo.LoginVo

	ctx.ShouldBindJSON(&user)

	if errMsg := utils.GetErrorMessage(user); errMsg != "" {
		response.FailWithMessage(errMsg, ctx)
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
