package helpers

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnv(path string) {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetAsString(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

func GetAsInt(name string, defaultValue int) int {
	valueStr := GetAsString(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
