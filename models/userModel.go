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
	result := inits.DB.Find(this, "email = ?", this.Email)

	return &user, result.Error
}

func (this *User) FindByEmailAndPassword() (*User, error) {
	var user User
	result := inits.DB.Where("email = ? AND password = ?", this.Email, this.Password).First(&user)

	return &user, result.Error
}

func (this *User) Find() (*User, error) {
	var user User
	result := inits.DB.Find(this)

	return &user, result.Error
}
