package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getAttachedOrganisations(c *gin.Context) {
	var res []map[string]interface{}
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	checkError(c, err)
	cursor, err := db.Collection.Find(
		context.Background(),
		bson.M{"_id": objID},
		options.Find().SetProjection(bson.D{
			{"organisations", 1},
		}),
	)
	checkError(c, err)
	for cursor.Next(context.TODO()) {
		var user map[string]interface{}
		err := cursor.Decode(&user)
		checkError(c, err)
		res = append(res, user)
	}
	c.JSON(200, res)
}
