package main

import (
	"log"

	"github.com/DeepanshuMishraa/conduit.git/config"
	"github.com/DeepanshuMishraa/conduit.git/routes"
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

	if cfg.PORT == ":8081" {
		log.Println("Response Reached Server 1")
	} else if cfg.PORT == ":8082" {
		log.Println("Response Reached Server 2")
	} else if cfg.PORT == ":8083" {
		log.Println("Response Reached Server 3")
	}

	log.Println("[API] RUNNING ON PORT", cfg.PORT)
	router.GET("/slow", routes.SlowRoute())
	router.GET("/fast", routes.FastRoute())

	router.Run(":" + cfg.PORT)
}
