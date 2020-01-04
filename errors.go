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
		notFoundResponse(c)
	}
}

func notFoundResponse(c *gin.Context) {
	_ = c.AbortWithError(404, nil)
}
