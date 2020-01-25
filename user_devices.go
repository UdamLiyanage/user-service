package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func attachDevice(c *gin.Context) {
	var attach AttachDevice
	err := json.NewDecoder(c.Request.Body).Decode(&attach)
	checkError(c, err)
	objectID, err := primitive.ObjectIDFromHex(attach.UserID)
	checkError(c, err)
	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$push": bson.M{
			"devices": attach.DeviceID,
		},
	}
	res, err := DB.Collection.UpdateOne(context.TODO(), filter, update)
	checkError(c, err)
	c.JSON(201, res)
}

func removeDevice(c *gin.Context) {
	var remove map[string]string
	err := json.NewDecoder(c.Request.Body).Decode(&remove)
	checkError(c, err)
	objectID, err := primitive.ObjectIDFromHex(remove["user_id"])
	checkError(c, err)
	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$pull": bson.M{
			"devices": bson.M{
				"$in": bson.A{remove["device_id"]},
			},
		},
	}
	res, err := DB.Collection.UpdateOne(context.TODO(), filter, update)
	checkError(c, err)
	c.JSON(404, res)
}

func getAttachedDevices(c *gin.Context) {
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	checkError(c, err)
	values, err := DB.Collection.Distinct(
		context.TODO(),
		"devices",
		bson.M{"_id": objID},
	)
	if err != nil {
		panic(err)
	}
	c.JSON(200, values)
}
