package main

import "github.com/gin-gonic/gin"

func checkError(c *gin.Context, err error) {
	if err != nil {
		c.AbortWithStatusJSON(500, err)
	}
	return
}
