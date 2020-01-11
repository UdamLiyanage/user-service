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
	user.Password = hashPassword(c, user.Password)
	user.CreatedAt = time.Now()
	insertResult, err := DB.Collection.InsertOne(context.TODO(), user)
	checkError(c, err)
	if oid, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
		c.JSON(201, map[string]interface{}{
			"id":             oid.String(),
			"first_name":     user.FirstName,
			"last_name":      user.LastName,
			"email":          user.Email,
			"contact_number": user.ContactNumber,
		})
	} else {
		c.AbortWithStatus(500)
	}
}

func hashPassword(c *gin.Context, p string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p), 10)
	checkError(c, err)
	return string(bytes)
}
