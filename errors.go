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
		c.AbortWithStatus(404)
	} else {
		c.AbortWithStatus(500)
	}
}
