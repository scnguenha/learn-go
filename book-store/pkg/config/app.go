package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func Connect() {

	dbUrl := fmt.Sprintf("host=localhost user=postgres dbname=bookstore sslmode=disable password=postgres")

	d, err := gorm.Open("postgres", dbUrl)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
