package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:lintauunited26@tcp(127.0.0.1:3306)/fga_golang_assignment"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Item{})
	database.AutoMigrate(&Order{})

	DB = database
}
