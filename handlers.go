package main

import (
	"github.com/gin-gonic/gin"
)

// DataHandler provives access to GET / end-point
func DataHandler(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok", "query": "query", "results": ""})
}
