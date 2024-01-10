package repository

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func DatabaseConnect(logger *logrus.Logger) *sql.DB {

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"module": "repository",
			"func":   "DatabaseConnect",
			"error":  err.Error(),
		}).Fatalf("failed to connect database: %v", err)

		logger.Info("failed to connect database")
	}

	err = db.Ping()
	if err != nil {
		logger.WithFields(logrus.Fields{
			"module": "repository",
			"func":   "DatabaseConnect",
			"error":  err.Error(),
		}).Fatalf("failed to ping database: %v", err)

		logger.Fatalf("failed to ping database")
	}

	logger.Info("Database connected")
	return db
}
