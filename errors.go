package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func checkError(c *gin.Context, err error) {
	if err == nil {
		return
	}
	if err == mongo.ErrNoDocuments {
		c.AbortWithStatusJSON(404, err)
	} else {
		c.AbortWithStatusJSON(500, err)
	}
}
