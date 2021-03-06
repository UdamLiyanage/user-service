package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func deleteUser(c *gin.Context) {
	opts := options.Delete().SetCollation(&options.Collation{
		Locale:    "en_US",
		Strength:  1,
		CaseLevel: false,
	})
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	checkError(c, err)
	filter := bson.D{bson.E{Key: "_id", Value: objID}}
	_, err = DB.Collection.DeleteOne(context.TODO(), filter, opts)
	checkError(c, err)
	c.JSON(404, nil)
}
