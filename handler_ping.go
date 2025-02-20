package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handlePing(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
