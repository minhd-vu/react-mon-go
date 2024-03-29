package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/minhd-vu/react-mon-go/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB is used to connect to MongoDB
func ConnectDB() *mongo.Database {
	// Load in the environment variables from .env
	config := util.GetConfiguration()

	// Set client options
	clientOptions := options.Client().ApplyURI(config.ConnectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Define what database you are using
	return client.Database(config.DatabaseName)
}

// ErrorResponse model
type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

// GetError is used to prepare the error model
func GetError(err error, w http.ResponseWriter) {
	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}

// Export the database
var Database = ConnectDB()
