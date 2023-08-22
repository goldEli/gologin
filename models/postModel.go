package models

import (
	"time"

	"gorm.io/gorm"
)

// 自定义一个匿名结构体，用于自定义字段的标签
type CustomModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleteAt"`
}

type Post struct {
	CustomModel
	UserID uint   `gorm:"user_id" json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Likes  int    `json:"likes"`
	Draft  bool   `json:"draft"`
	Author string `json:"author"`
}
