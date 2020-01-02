package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func createUser(c *gin.Context) {
	var user User
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	checkError(c, err)
	user.Password = hashPassword(user.Password)
	user.CreatedAt = time.Now()
	insertResult, err := db.Collection.InsertOne(context.TODO(), user)
	checkError(c, err)
	user.ID = insertResult.InsertedID.(primitive.ObjectID)
	c.JSON(201, user)
}

func hashPassword(p string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p), 10)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
