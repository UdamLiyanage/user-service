package main

import "github.com/gin-gonic/gin"

func updateUser(c *gin.Context) {
	c.String(200, "Update User")
}
