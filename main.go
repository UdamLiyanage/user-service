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

	r.GET("/users/:id", readUser)

	r.POST("/users", createUser)

	r.PUT("/users/:id", updateUser)

	r.DELETE("/users/:id", deleteUser)
	return r
}

func main() {
	r := setupRouter()
	log.Fatal(r.Run())
}
