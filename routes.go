package main

import "github.com/gin-gonic/gin"

func initialiseRoutes(s *server) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", handlePing)

	r.POST("/v2/email/outbound-emails", s.handleSendEmail)

	return r
}
