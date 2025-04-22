package config

import (
	//	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/moazan-waqar-hpe/gin-gorm-rest/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgres"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.User{})
	DB = db

}
