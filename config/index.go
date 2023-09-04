package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type Mysql struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

type Redis struct {
	Host string
	Port int
}

type Config struct {
	PORT     string
	SECRET   string
	Mysql    Mysql
	MysqlUrl string
	Redis    Redis
}

var Env *Config

func GetConfig() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	name := fmt.Sprintf("config.%s", env)
	fmt.Println("configName", name)

	// 初始化配置
	viper.SetConfigName(name)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}
	// 将配置文件中的值解析到 Config 结构体中
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}

	Env = &Config{
		PORT:     config.PORT,
		SECRET:   config.SECRET,
		Mysql:    config.Mysql,
		MysqlUrl: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Mysql.Username, config.Mysql.Password, config.Mysql.Host, strconv.Itoa(config.Mysql.Port), config.Mysql.Database),
		Redis:    config.Redis,
	}
}
