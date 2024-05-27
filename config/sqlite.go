package config

import (
	"os"

	"github.com/thiagoCalazans-dev/go-api-opportunities.git/schema"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	//check if database exist
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("Database file not found, creating ...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	//Create DB and Connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("sqlite start error: %v", err)
		return nil, err
	}

	//Migrate Schema
	err = db.AutoMigrate(&schema.Opportunity{})

	if err != nil {
		logger.Errorf("Migration start error: %v", err)
		return nil, err
	}

	return db, nil
}
