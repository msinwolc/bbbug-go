package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDB() *gorm.DB {
	viper.SetConfigName("database")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}

	driver := "mysql"
	host := viper.Get("mysql.hostname")
	port := viper.Get("mysql.port")
	username := viper.Get("mysql.username")
	password := viper.Get("mysql.password")
	database := viper.Get("mysql.database")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		username,
		password,
		host,
		port,
		database)
	fmt.Println(args)
	db, err := gorm.Open(driver, args)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	return db
}
