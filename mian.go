package main

import (
	"gologin/inits"
	"gologin/router"
)

// this init function will be called automatically and cannot have any parameters
func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {

	router.Router()

}
