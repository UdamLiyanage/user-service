package main

import "github.com/gin-gonic/gin"

func createUser(c *gin.Context) {
	c.String(200, "Create Function")
}
