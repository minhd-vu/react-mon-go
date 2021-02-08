package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/faygun/go-rest-api/helper"
	"github.com/gorilla/mux"
	"github.com/minhd-vu/go-project/helpers"
	"github.com/minhd-vu/go-project/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Connection mongoDB with helper class
var collection = helpers.ConnectDB()

func getPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// we created Player array
	var players []models.Player

	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	// Close the cursor once finished
	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var player models.Player
		// & character returns the memory address of the following variable.
		err := cur.Decode(&player) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		players = append(players, player)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(players) // encode similar to serialize process.
}

func getPlayer(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	var player models.Player
	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&player)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(player)
}

func createPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var player models.Player

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&player)

	// insert our player model.
	result, err := collection.InsertOne(context.TODO(), player)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func updatePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	//Get id from parameters
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var player models.Player

	// Create filter
	filter := bson.M{"_id": id}

	// Read update model from body request
	_ = json.NewDecoder(r.Body).Decode(&player)

	// prepare update model.
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "name", Value: player.Name},
		}},
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&player)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	player.ID = id

	json.NewEncoder(w).Encode(player)
}

func deletePlayer(w http.ResponseWriter, r *http.Request) {
	// Set header
	w.Header().Set("Content-Type", "application/json")

	// get params
	var params = mux.Vars(r)

	// string to primitve.ObjectID
	id, err := primitive.ObjectIDFromHex(params["id"])

	// prepare filter.
	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/players", getPlayers).Methods("GET")
	r.HandleFunc("/api/players/{id}", getPlayer).Methods("GET")
	r.HandleFunc("/api/players", createPlayer).Methods("POST")
	r.HandleFunc("/api/players/{id}", updatePlayer).Methods("PUT")
	r.HandleFunc("/api/players/{id}", deletePlayer).Methods("DELETE")

	config := helper.GetConfiguration()
	log.Fatal(http.ListenAndServe(config.Port, r))

}
