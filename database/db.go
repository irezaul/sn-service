package database

import (
	"ginorm/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Connection() {
	DB, err = gorm.Open(sqlite.Open("ginorm.db"), &gorm.Config{})
	if err != nil {
		panic("Connection failed")
	}
	DB.AutoMigrate(&model.User{}, &model.Login{}, &model.Service{})
}
