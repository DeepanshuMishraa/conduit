package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func RequestLogger(instance string) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("[%s] %s %s",
			instance,
			c.Request.Method,
			c.Request.URL.Path,
		)

		c.Next()
	}
}
