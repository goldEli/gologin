package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title  string `json:"title"`
	Body   string
	Likes  string
	Draft  string
	Author string
}
