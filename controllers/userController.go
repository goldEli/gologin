package controllers

import (
	"gologin/inits"
	"gologin/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(ctx *gin.Context) {
	var body struct {
		Name     string
		Email    string
		Password string
	}

	if ctx.BindJSON(&body) != nil {
		ctx.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err})
		return
	}

	user := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: string(hash),
	}

	result := inits.DB.Create(&user)

	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}
	ctx.JSON(200, gin.H{"data": user})
}

func Login(ctx *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if ctx.BindJSON(&body) != nil {
		ctx.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	var user models.User

	result := inits.DB.Where("email = ?", body.Email).First(&user)

	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": "user not found"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		ctx.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	// generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		ctx.JSON(500, gin.H{"error": "error signing token"})
		return
	}

	// Set the same-site mode to the default mode
	ctx.SetSameSite(http.SameSiteDefaultMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24*30, "", "localhost", false, true)
	ctx.JSON(200, gin.H{"data": "login success"})
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
