package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func attachDeviceDefinition(c *gin.Context) {
	var attach AttachDeviceDefinition
	err := json.NewDecoder(c.Request.Body).Decode(&attach)
	checkError(c, err)
	objectID, err := primitive.ObjectIDFromHex(attach.UserID)
	checkError(c, err)
	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$addToSet": bson.M{
			"device-definitions": attach.DefinitionID,
		},
	}
	res, err := DB.Collection.UpdateOne(context.TODO(), filter, update)
	checkError(c, err)
	c.JSON(200, res)
}

func removeDeviceDefinition(c *gin.Context) {
	var remove map[string]string
	err := json.NewDecoder(c.Request.Body).Decode(&remove)
	checkError(c, err)
	objectID, err := primitive.ObjectIDFromHex(remove["user_id"])
	checkError(c, err)
	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$pull": bson.M{
			"device-definitions": bson.M{
				"$in": bson.A{remove["definition_id"]},
			},
		},
	}
	res, err := DB.Collection.UpdateOne(context.TODO(), filter, update)
	checkError(c, err)
	c.JSON(200, res)
}
