package util

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

// Configuration model
type Configuration struct {
	Port             string
	ConnectionString string
	DatabaseName     string
}

const projectDirName = "react-mon-go"

// GetConfiguration populates dotenv configuration information into a configuration model
func GetConfiguration() Configuration {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config := Configuration{
		os.Getenv("PORT"),
		os.Getenv("CONNECTION_STRING"),
		os.Getenv("DATABASE_NAME"),
	}

	return config
}
