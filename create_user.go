package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func createUser(c *gin.Context) {
	var user User
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		panic(err)
	}
	user.CreatedAt = time.Now()
	insertResult, err := db.Collection.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}
	user.ID = insertResult.InsertedID.(primitive.ObjectID)
	c.JSON(200, user)
}
