package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Ping endpoint
// @Description Health check endpoint that returns "pong"
// @Produce json
// @Success 200 {string} string "pong"
// @Router /ping [get]
func handlePing(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
