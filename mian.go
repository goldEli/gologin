package main

import (
	"gologin/config"
	"gologin/inits"
	"gologin/migrations"
	"gologin/router"
)

// this init function will be called automatically and cannot have any parameters
func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	config.GetConfig()

	migrations.CreateTable()
	router.Router()
}
