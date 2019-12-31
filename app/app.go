package app

import (
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
