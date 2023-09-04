package service

import (
	"gologin/config"
	"gologin/inits"
	"gologin/models"
	"gologin/response"
	"gologin/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

// 注册
func Register(name, email, password string) error {
	hashedPassword, err := utils.GetPwd(password)
	if err != nil {
		return err
	}
	user := models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	return inits.DB.Create(&user).Error
}

// 登陆
func Login(email, password string) (string, int) {

	var user models.User

	result := inits.DB.Where("email = ?", email).First(&user)
	logrus.Info(user)

	if result.Error != nil {
		logrus.Error(result.Error)
		return "", response.ResponseCodeUserPasswordError
	}

	// compare password
	// 密码比对
	isEqual := utils.ComparePwd(user.Password, password)
	if !isEqual {
		return "", response.ResponseCodeUserPasswordError
	}

	// 生成 jwt
	// generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(config.Env.SECRET))
	if err != nil {
		return "", response.ResponseCodeInternalServerError
	}

	return tokenString, response.ResponseCodeOk
}
