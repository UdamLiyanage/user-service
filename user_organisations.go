package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func attachOrganisation(c *gin.Context) {
	var attach AttachOrganisation
	err := json.NewDecoder(c.Request.Body).Decode(&attach)
	checkError(c, err)
	objectID, err := primitive.ObjectIDFromHex(attach.UserID)
	checkError(c, err)
	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$addToSet": bson.M{
			"organisations": bson.M{
				"organisation_id":   attach.OrganisationID,
				"organisation_name": attach.OrganisationName,
			},
		},
	}
	res, err := DB.Collection.UpdateOne(context.TODO(), filter, update)
	checkError(c, err)
	c.JSON(201, res)
}

func getAttachedOrganisations(c *gin.Context) {
	var res []map[string]interface{}
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	checkError(c, err)
	cursor, err := DB.Collection.Find(
		context.Background(),
		bson.M{"_id": objID},
		options.Find().SetProjection(bson.D{
			bson.E{Key: "organisations", Value: 1},
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