package service

import (
	"gologin/config"
	"gologin/inits"
	"gologin/models"
	"gologin/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

// 注册
func Register(name, email, password string) error {
	md5, err := utils.MD5(password)
	if err != nil {
		return err
	}
	user := models.User{
		Name:     name,
		Email:    email,
		Password: md5,
	}

	return inits.DB.Create(&user).Error
}

// 登陆
func Login(user *models.User) (string, error) {
	current, err := user.FindByEmail()
	logrus.Error("login")

	if err != nil {
		return "", err
	}
	// 生成 jwt
	// generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  current.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	logrus.Info("token", token)
	tokenString, err := token.SignedString([]byte(config.Env.SECRET))

	return tokenString, err
}
