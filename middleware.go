package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userName := c.Request.Header.Get("UserName")
		if userName == "" {
			c.JSON(http.StatusUnauthorized, "Provide a user name")
			c.Abort()
			return
		}

		c.Keys = make(map[string]interface{})
		c.Keys["user_name"] = userName
		c.Next()
	}
}
