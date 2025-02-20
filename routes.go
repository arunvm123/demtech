package main

import "github.com/gin-gonic/gin"

func initialiseRoutes(s *server) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", handlePing)

	return r
}
