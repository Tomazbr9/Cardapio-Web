package main

import (
	"cw/internal/controllers"
	"cw/internal/database"
	"cw/internal/repositories"
	"cw/internal/services"
	"cw/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


func main() {
	
	database.InitDB()

	db := database.DB
	tenantRepository := repositories.NewTenantRepository(db)
	tenantService := services.NewTenantService(*tenantRepository)
	tenantHandler := controllers.NewTenantHandler(*tenantService)

	router := gin.Default()

	router.GET("/health", func (c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H{
			"status":    "online",
			"framework": "gin",
			"database":  "connected",
		})
	})

	routes.RegisterAdminRoutes(router, tenantHandler)

	log.Println("Iniciando API na porta 8080...")
	
	log.Fatal(router.Run(":8080"))
}