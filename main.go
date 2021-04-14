package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"quickstart/helpers"
	"quickstart/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

var database = helpers.ConnectDB()
var collection = database.Collection("users")

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []models.User

	// Get all the users
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helpers.GetError(err, w)
		return
	}

	// Close the cursor once finished
	defer cur.Close(context.TODO())

	// Iterate through each
	for cur.Next(context.TODO()) {
		var user models.User
		// Decode the user
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}

		// Add item to the array
		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Encode the data to the database
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	// Decode the user
	_ = json.NewDecoder(r.Body).Decode(&user)

	// Insert the user into the collection
	result, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		helpers.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/users", createUser).Methods("POST")

	config := helpers.GetConfiguration()
	log.Fatal(http.ListenAndServe(config.Port, r))
}
