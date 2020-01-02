package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Database struct {
	Collection *mongo.Collection
}

var db = Database{Collection: connect()}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/users/:id")

	r.POST("/users")

	r.PUT("/users/:id")

	r.DELETE("/users/:id")
	return r
}

func main() {
	r := setupRouter()
	log.Fatal(r.Run())
}
