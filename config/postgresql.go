package config

import (
	"fmt"
	"log"
	"os"

	"golang-assignment/entity"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	host := os.Getenv("PGHOST")
	user := os.Getenv("PGUSERNAME")
	password := os.Getenv("PGPASSWORD")
	dbname := os.Getenv("PGNAME")
	port := os.Getenv("PGPORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	fmt.Println("dsn logging : ", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Set up auto-migration
	log.Println("INFO: Running database migrations...")

	if err := db.AutoMigrate(&entity.User{}); err != nil {
		log.Fatalf("ERROR: Failed to run migrations: %v", err)
	}

	log.Println("INFO: Database connection and migration successful")
	return db
}
