package main

import (
	"cw/internal/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


func main() {
	database.InitDB()

	router := gin.Default()

	router.GET("/health", func (c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H{
			"status":    "online",
			"framework": "gin",
			"database":  "connected",
		})
	})

	log.Println("Iniciando API na porta 8080...")
	
	log.Fatal(router.Run(":8080"))
}