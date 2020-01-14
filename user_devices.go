package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func attachDevice(c *gin.Context) {
	var attach AttachDevice
	err := json.NewDecoder(c.Request.Body).Decode(&attach)
	checkError(c, err)
	objectID, err := primitive.ObjectIDFromHex(attach.UserID)
	checkError(c, err)
	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$addToSet": bson.M{
			"devices": bson.M{
				"device_id": attach.DeviceID,
			},
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
				"device_id": remove["device_id"],
			},
		},
	}
	res, err := DB.Collection.UpdateOne(context.TODO(), filter, update)
	checkError(c, err)
	c.JSON(404, res)
}

func getAttachedDevices(c *gin.Context) {
	var res []map[string]interface{}
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	checkError(c, err)
	cursor, err := DB.Collection.Find(
		context.Background(),
		bson.M{"_id": objID},
		options.Find().SetProjection(bson.D{
			bson.E{Key: "devices", Value: 1},
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
