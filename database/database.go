package database

import (
	"fmt"

	"github.com/Just-Goo/Student-Portal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DSN = "root:@tcp(localhost:3306)/studentportaldb?parseTime=true&loc=Africa%2FLagos"

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("error connecting to database")
	}

	// auto migrate the database
	db.AutoMigrate(&model.Student{})

	return db, nil
}
