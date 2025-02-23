package main

import (
	_ "github.com/arunvm123/demtech/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// gin-swagger middleware
// swagger embed files

func initialiseRoutes(s *server) *gin.Engine {
	r := gin.Default()

	private := r.Group("/")
	private.Use(authMiddleware())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", handlePing)
	r.GET("/logs", s.handlerGetLogsAggregates)

	private.POST("/v2/email/outbound-emails", s.handleSendEmail)

	return r
}
