package database

import (
	"fmt"
	"hology/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

// Init is called in server.go at root project
func Init() {
	configuration := config.GetConfig()
	connect_string := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", configuration.DB_USERNAME, configuration.DB_PASSWORD, configuration.DB_NAME)
	db, err = gorm.Open("mysql", connect_string)
	// defer db.Close()
	if err != nil {
		panic("DB Connection Error")
	}
}

// CreateConn function to get db connection at repository
func CreateConn() *gorm.DB {
	return db
}
