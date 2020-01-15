package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func attachOrganisation(c *gin.Context) {
	var attach AttachOrganisation
	err := json.NewDecoder(c.Request.Body).Decode(&attach)
	checkError(c, err)
	objectID, err := primitive.ObjectIDFromHex(attach.UserID)
	checkError(c, err)
	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$push": bson.M{
			"organisations": attach.OrganisationID,
		},
	}
	res, err := DB.Collection.UpdateOne(context.TODO(), filter, update)
	checkError(c, err)
	c.JSON(201, res)
}

func getAttachedOrganisations(c *gin.Context) {
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	checkError(c, err)
	values, err := DB.Collection.Distinct(
		context.TODO(),
		"organisations",
		bson.M{"_id": objID},
	)
	c.JSON(200, values)
}
