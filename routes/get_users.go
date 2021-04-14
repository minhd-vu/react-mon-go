package routes

import (
	"context"
	"log"
	"net/http"

	"github.com/minhd-vu/go-project/helpers"
	"github.com/minhd-vu/go-project/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUsers(c *gin.Context) {
	var users []*models.User

	var database = helpers.ConnectDB()
	var collection = database.Collection("users")

	// Get all the users
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err})
		return
	}

	// Close the cursor once finished
	defer cur.Close(context.TODO())

	// Get all the documents
	err = cur.All(context.TODO(), &users)

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Response with users
	c.JSON(http.StatusOK, users)
}
