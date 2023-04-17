// migrationとは
// auto migration とは
package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //use postgreSQL in gorm
	// "gorm.io/driver/postgres"

	"go-backend-tuto/entity"
)

var (
	db *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {
	db, err = gorm.Open(
		"postgres", 
		"host=localhost port=5432 user=gorm dbname=go-backend-tuto_db_1 password=gorm sslmode=disable",
	)
	if err != nil {
		panic(err)
	}

	autoMigration()
}

//GetDB is called in module
func GetDB() *gorm.DB {
	return db
}

// Close is closing db
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration() {
	db.AutoMigrate(&entity.User{})
}