package api

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minhd-vu/react-mon-go/db"
	"github.com/minhd-vu/react-mon-go/model"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(c *gin.Context) {
	var user model.User

	// Decode the user
	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	collection := db.Database.Collection("users")

	// Insert the user into the collection
	result, err := collection.InsertOne(c, user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetUsers(c *gin.Context) {
	var users []*model.User

	collection := db.Database.Collection("users")

	// Get all the users
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err})
		return
	}

	// Close the cursor once finished
	defer cur.Close(c)

	// Get all the documents
	err = cur.All(c, &users)

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Response with users
	c.JSON(http.StatusOK, users)
}
