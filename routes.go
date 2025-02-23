package main

import "github.com/gin-gonic/gin"

func initialiseRoutes(s *server) *gin.Engine {
	r := gin.Default()

	private := r.Group("/")
	private.Use(authMiddleware())

	r.GET("/ping", handlePing)
	r.GET("/logs", s.handlerGetLogsAggregates)

	private.POST("/v2/email/outbound-emails", s.handleSendEmail)

	return r
}
