package models

import (
	"gologin/inits"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique"`
	Password string `json:"password"`
	Post     []Post `json:"post"`
}

func (this *User) Create() error {
	result := inits.DB.Create(this)

	return result.Error
}

func (this *User) FindByEmail() (*User, error) {
	var user User
	result := inits.DB.Where("email = ?", this.Email).First(&user)

	return &user, result.Error
}

func (this *User) FindByName() (*User, error) {
	var user User
	result := inits.DB.Where("name = ?", this.Name).First(&user)

	return &user, result.Error
}
