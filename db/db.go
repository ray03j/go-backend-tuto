package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //use postgreSQL in gorm
)

var (
	db *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {
	db, err = gorm.Open("postgres", "host=0.0.0.0 port=5432 user=gorm dbname=gorm password=gorm sslmode=disable")
	if err != nil {
		panic(err)
	}
}

//GetDB is called in module
func GetDB() *gorm.DB {
	return db
}

// Close is closing db
func Close() {
	if err := db.Close; err != nil {
		panic(err)
	}
}