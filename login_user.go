package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func login(c *gin.Context) {
	var credentials map[string]interface{}
	var user User
	err := json.NewDecoder(c.Request.Body).Decode(&credentials)
	checkError(c, err)
	filter := bson.M{"email": credentials["email"]}
	err = DB.Collection.FindOne(context.TODO(), filter).Decode(&user)
	checkError(c, err)
	if !checkPassword(credentials["password"].(string), user.Password) {
		c.AbortWithStatusJSON(401, nil)
		return
	}
	c.JSON(200, user)
}

func checkPassword(p, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(p))
	return err == nil
}
