package main

import (
	"cw/internal/database"
	"cw/internal/handlers"
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
	tenantHandler := handlers.NewTenantHandler(*tenantService)

	categoryRepository := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryRepository(categoryRepository)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	router := gin.Default()

	router.GET("/health", func (c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H{
			"status":    "online",
			"framework": "gin",
			"database":  "connected",
		})
	})

	routes.RegisterSuperAdminRoutes(router, tenantHandler)
	routes.RegisterTenantAdminRoutes(router, *tenantService, categoryHandler)
	routes.RegisterPublicRoutes(router, *tenantService, categoryHandler)

	log.Println("Iniciando API na porta 8080...")
	
	log.Fatal(router.Run(":8080"))
}