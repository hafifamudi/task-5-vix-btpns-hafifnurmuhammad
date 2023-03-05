package database

import (
	"fmt"
	"log"

	"github.com/hafifamudi/task-5-vix-btpns-hafifnurmuhammad/helpers"
	"github.com/hafifamudi/task-5-vix-btpns-hafifnurmuhammad/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func MigrateDB() {
	//load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbURI := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		helpers.GetAsString("DB_USER", "postgres"),
		helpers.GetAsString("DB_PASSWORD", "mysecretpassword"),
		helpers.GetAsString("DB_HOST", "postgres"),
		helpers.GetAsInt("DB_PORT", 5432),
		helpers.GetAsString("DB_NAME", "postgres"),
	)

	db, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	db.Debug().AutoMigrate(models.User{}, models.Photo{})
}

func GetDB() *gorm.DB {
	return db
}
