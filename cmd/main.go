package main

import (
	"log"

	"github.com/DeepanshuMishraa/conduit.git/config"
	"github.com/DeepanshuMishraa/conduit.git/db"
	"github.com/DeepanshuMishraa/conduit.git/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		log.Fatal(err)
	}

	db, err := db.ConnectDB(cfg.DATABASE_URL)

	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.Use(RequestLogger(cfg.INSTANCE))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Println("[API] RUNNING ON PORT", cfg.PORT)
	router.GET("/slow", routes.SlowRoute())
	router.GET("/fast", routes.FastRoute())
	router.POST("/db", routes.DBRoute(db))

	router.Run(":" + cfg.PORT)
}
