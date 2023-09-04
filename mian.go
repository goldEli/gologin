package main

import (
	"gologin/config"
	"gologin/inits"
	"gologin/migrations"
	"gologin/router"
)

// this init function will be called automatically and cannot have any parameters
func init() {

	config.GetConfig()
	inits.InitLogrus()
	inits.LoadEnv()
	inits.RedisInit()
	inits.DBInit()
}

func main() {

	migrations.CreateTable()
	router.Router()
}
