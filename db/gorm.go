package db

import (
	"log"
	"os"
	dbschema "rfid_payment/db/db_schema"

	"github.com/joho/godotenv" // Import the godotenv package
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectGorm() *gorm.DB {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURI := os.Getenv("SQLALCHEMY_DATABASE_URI")

	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	// Migrate Product Database
	errors := db.AutoMigrate(dbschema.UserTracker{})
	if errors != nil {
		log.Println(errors.Error())
	}

	return db
}
