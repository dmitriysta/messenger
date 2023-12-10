package repository

import (
	"github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnect(logger *logrus.Logger) *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=messenger port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.WithFields(logrus.Fields{
			"module": "repository",
			"func":   "DatabaseConnect",
			"error":  err.Error(),
		}).Fatalf("failed to connect database: %v", err)

		logger.Fatalf("failed to connect database")
	}

	logger.Info("Database connected")
	return db
}
