package storage

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/hnakamur/api2go-gorm-gin-crud-example/model"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(os.Getenv("DB_DIALECT"), os.Getenv("DB_PARAMS"))
	if err != nil {
		return nil, err
	}

	db.LogMode(true)
	db.AutoMigrate(&model.User{}, &model.Chocolate{})

	return &db, nil
}
