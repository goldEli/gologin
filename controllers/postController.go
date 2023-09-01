package controllers

import (
	"fmt"
	"gologin/inits"
	"gologin/models"
	"gologin/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreatePost
func CreatePost(ctx *gin.Context) {
	var body struct {
		UserID uint `json:"user_id"`
		Title  string
		Body   string
		Likes  int
		Draft  bool
		Author string
	}
	ctx.BindJSON(&body)

	// Get the user from the context
	user, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusBadRequest, response.ResponseUserPasswordError())
		return
	}
	body.UserID = user.(models.User).ID

	post := models.Post{
		Title:  body.Title,
		Body:   body.Body,
		Likes:  body.Likes,
		Draft:  body.Draft,
		Author: body.Author,
		UserID: body.UserID,
	}

	fmt.Println(post)
	result := inits.DB.Create(&post)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, response.ResponseServerError())
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseOkWIthData(post))
}

// UpdatePost
func UpdatePost(ctx *gin.Context) {
	var body struct {
		Title  string
		Body   string
		Likes  int
		Draft  bool
		Author string
	}

	ctx.BindJSON(&body)

	var post models.Post

	result := inits.DB.First(&post, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, response.ResponseServerError())
		return
	}
	inits.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body, Likes: body.Likes, Draft: body.Draft, Author: body.Author})

	ctx.JSON(http.StatusOK, response.ResponseOkWIthData(post))
}

// DeletePost
func DeletePost(ctx *gin.Context) {
	id := ctx.Param("id")
	inits.DB.Delete(&models.Post{}, id)

	ctx.JSON(http.StatusOK, response.ResponseOk())
}

// GetPost
func GetPost(ctx *gin.Context) {
	fmt.Println("GetPost")
	var post models.Post
	result := inits.DB.First(&post, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, response.ResponseServerError())
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseOkWIthData(post))
}

// GetPosts
func GetPosts(ctx *gin.Context) {
	var posts []models.Post
	result := inits.DB.Find(&posts)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, response.ResponseServerError())
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseOkWIthData(posts))
}
