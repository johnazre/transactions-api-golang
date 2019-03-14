package databaseConfig

import (
	"github.com/jinzhu/gorm"
)

func DbStart() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 dbname=transactions-api-go sslmode=disable")

	if err != nil {
		panic(err.Error())
	}

	database := db.DB()

	err = database.Ping()

	if err != nil {
		panic(err.Error())
	}

	return db, err
}
