package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func readUser(c *gin.Context) {
	var user map[string]interface{}
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	checkError(c, err)
	filter := bson.M{"_id": objID}
	err = DB.Collection.FindOne(context.TODO(), filter).Decode(&user)
	checkError(c, err)
	delete(user, "password")
	c.JSON(200, user)
}
