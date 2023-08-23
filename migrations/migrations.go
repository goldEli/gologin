package migrations

import (
	"gologin/inits"
	"gologin/models"
	"log"
)

func CreateTable() {
	err := inits.DB.AutoMigrate(&models.Post{}, &models.User{})
	if err != nil {
		log.Fatal(err)
	}
}
