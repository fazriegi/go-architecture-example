package config

import (
	"fmt"
	"log"
	"os"

	"github.com/fazriegi/go-architecture-example/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase() *gorm.DB {
	host := os.Getenv("DBHOST")
	username := os.Getenv("DBUSER")
	password := os.Getenv("DBPASS")
	name := os.Getenv("DBNAME")
	port := os.Getenv("DBPORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host,
		username,
		password,
		name,
		port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	migrate(db)

	return db
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
	)
}
