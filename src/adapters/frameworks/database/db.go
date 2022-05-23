package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	PAGE_SIZE = 25
)

type DbConnector struct {
	db *gorm.DB
}

func NewDbAdapter(connectionString string) (*DbConnector, error) {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Unable to open connection during init:", err)
	}

	connSettings, err := db.DB()
	if err != nil {
		log.Fatal("Error during set max conn:", err)
	}

	connSettings.SetMaxIdleConns(20)
	connSettings.SetMaxOpenConns(200)

	return &DbConnector{db: db}, nil
}
