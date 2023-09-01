package models

import (
	"time"

	"gorm.io/gorm"
)

type CustomModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleteAt"`
	DeletedOn uint32         `json:"deleted_on"`
	IsDel     uint8          `json:"is_del"`
}
