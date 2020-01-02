package main

import "github.com/gin-gonic/gin"

func deleteUser(c *gin.Context) {
	c.String(200, "Delete User")
}
