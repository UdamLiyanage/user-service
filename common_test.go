package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)
	DB.Collection = connect()
	// Run the other tests
	os.Exit(m.Run())
}

func newRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/users/:id", readUser)
	r.POST("/users", createUser)
	r.POST("/users/attach/device", attachDevice)
	r.DELETE("/users/:id", deleteUser)
	r.PUT("/users/:id", updateUser)
	return r
}
