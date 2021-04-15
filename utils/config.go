package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Configuration model
type Configuration struct {
	Port             string
	ConnectionString string
	DatabaseName     string
}

// GetConfiguration populates dotenv configuration information into a configuration model
func GetConfiguration() Configuration {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	configuration := Configuration{
		os.Getenv("PORT"),
		os.Getenv("CONNECTION_STRING"),
		os.Getenv("DATABASE_NAME"),
	}

	return configuration
}
