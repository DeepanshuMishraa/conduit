package main

import (
	"log"

	"github.com/DeepanshuMishraa/conduit.git/config"
	"github.com/DeepanshuMishraa/conduit.git/db"
	"github.com/DeepanshuMishraa/conduit.git/routes"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	cfg, err := config.Load()
	reg := prometheus.NewRegistry()

	reg.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)

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
	router.GET("/metrics", gin.WrapH(promhttp.HandlerFor(reg, promhttp.HandlerOpts{})))

	router.Run(":" + cfg.PORT)
}
