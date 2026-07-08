package main

import (
	"log"

	"github.com/DeepanshuMishraa/conduit.git/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Println("[API] RUNNING ON PORT", cfg.PORT)

	router.Run(":" + cfg.PORT)
}
