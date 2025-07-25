package database

import (
	"errors"
	"rodrigoorlandini/urlshortener/analytics/config"
	"rodrigoorlandini/urlshortener/analytics/internal/infrastructure/database/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Database struct {
	Connection *gorm.DB
}

func ResetConnection() {
	db = nil
}

func GetConnection() (*Database, error) {
	if db == nil {
		var err error
		db, err = createConnection()
		if err != nil {
			return nil, err
		}

		err = db.AutoMigrate(&models.URL{})
		if err != nil {
			return nil, errors.New("failed to migrate database")
		}
	}

	return &Database{
		Connection: db,
	}, nil
}

func createConnection() (*gorm.DB, error) {
	dsn := config.NewEnvironment().DatabaseURL
	if dsn == "" {
		return nil, errors.New("DATABASE_URL environment variable is not set")
	}

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed to connect database: " + err.Error())
	}

	return connection, nil
}
