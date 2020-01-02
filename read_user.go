package main

import "github.com/gin-gonic/gin"

func readUser(c *gin.Context) {
	c.String(200, "Read User")
}
