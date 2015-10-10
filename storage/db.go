package storage

import (
	"os"

	"github.com/hnakamur/api2go-gorm-gin-crud-example/model"
	"github.com/jinzhu/gorm"
)

// InitDB creates and migrates the database
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(os.Getenv("DB_DIALECT"), os.Getenv("DB_PARAMS"))
	if err != nil {
		return nil, err
	}

	//db.LogMode(true)
	db.AutoMigrate(&model.User{}, &model.Chocolate{})

	return &db, nil
}
