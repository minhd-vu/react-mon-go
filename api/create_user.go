package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minhd-vu/react-mon-go/db"
	"github.com/minhd-vu/react-mon-go/model"
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
	result, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}

	c.JSON(http.StatusOK, result)
}
