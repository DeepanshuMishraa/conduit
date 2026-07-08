package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SlowRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		slow := time.Duration(time.Second * 30)
		time.Sleep(slow)

		c.JSON(http.StatusOK, gin.H{
			"message": "Its a Slow Route",
		})
	}
}

func FastRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Its a Fast Route",
		})
	}
}
