package main

import "github.com/gin-gonic/gin"

func getAttachedOrganisations(c *gin.Context) {
	c.String(200, "Get attached organisation")
}
