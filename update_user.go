package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func updateUser(c *gin.Context) {
	var body map[string]interface{}
	var res *mongo.UpdateResult
	err := json.NewDecoder(c.Request.Body).Decode(&body)
	checkError(c, err)
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	checkError(c, err)
	filter := bson.M{"_id": objID}
	err = nil
	if _, exists := body["first_name"]; exists {
		update := bson.M{
			"$set": bson.M{
				"first_name": body["first_name"],
			},
		}
		res, err = DB.Collection.UpdateOne(context.TODO(), filter, update)
	} else if _, exists := body["last_name"]; exists {
		update := bson.M{
			"$set": bson.M{
				"first_name": body["first_name"],
			},
		}
		res, err = DB.Collection.UpdateOne(context.TODO(), filter, update)
	} else {
		update := bson.M{
			"$set": bson.M{
				"first_name": body["first_name"],
				"last_name":  body["last_name"],
			},
		}
		res, err = DB.Collection.UpdateOne(context.TODO(), filter, update)
	}
	checkError(c, err)
	c.JSON(200, res)
}
