package db

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Use PostgreSQL in gorm

	"github.com/daitasu/test-module/entity"
)

var (
	db  *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {
	user := os.Getenv("USER_NAME")
	dbname := os.Getenv("DBNAME")
	pass := os.Getenv("PASSWORD")

	db, err = gorm.Open("postgres", "host=0.0.0.0 port=5432 user="+user+" dbname="+dbname+" password="+pass+" sslmode=disable")
	if err != nil {
		panic(err)
	}

	autoMigration()
}

// GetDB is called in models
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
