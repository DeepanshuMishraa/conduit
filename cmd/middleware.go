package main

import (
	"log"
	"strconv"
	"time"

	"github.com/DeepanshuMishraa/conduit.git/metrics"
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

func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		metrics.InflightRequests.Inc()

		start := time.Now()

		c.Next()

		duration := time.Since(start).Seconds()

		metrics.InflightRequests.Dec()

		metrics.HttpRequestsTotal.
			WithLabelValues(
				c.Request.Method,
				c.FullPath(),
				strconv.Itoa(c.Writer.Status()),
			).
			Inc()

		metrics.HttpRequestsDuration.
			WithLabelValues(
				c.Request.Method,
				c.FullPath(),
			).
			Observe(duration)
	}
}
