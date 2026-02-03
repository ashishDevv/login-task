package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(dbUrl string) (*gorm.DB, error) {

	// log.Printf("DNS is : %v", dbUrl)

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
