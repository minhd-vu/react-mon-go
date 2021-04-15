package routes

import (
	"context"
	"net/http"

	"github.com/minhd-vu/go-project/db"
	"github.com/minhd-vu/go-project/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User

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
