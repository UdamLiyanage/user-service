package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Database struct {
	Collection *mongo.Collection
}

var DB = Database{}

func init() {
	DB.Collection = connect()
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/users/:id", readUser)
	r.GET("/users/:id/devices", getAttachedDevices)
	r.GET("/users/:id/organisations", getAttachedOrganisations)

	r.POST("/users", createUser)
	r.POST("/login", login)
	r.POST("/users/attach/device", attachDevice)
	r.POST("/users/attach/organisation", attachOrganisation)

	r.PUT("/users/:id", updateUser)

	r.DELETE("/users/:id", deleteUser)
	return r
}

func main() {
	r := setupRouter()
	log.Fatal(r.Run())
}
