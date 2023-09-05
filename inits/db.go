package inits

import (
	"gologin/config"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInit() {
	logrus.Info(config.Env.MysqlUrl)
	db, err := gorm.Open(mysql.Open(config.Env.MysqlUrl), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}
