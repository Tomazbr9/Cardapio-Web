package main

import (
	"cw/internal/database"
	"cw/internal/handlers"
	"cw/internal/repositories"
	"cw/internal/services"
	"cw/internal/routes"
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

	productRepository := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepository)
	productHandler := handlers.NewProductHandler(productService)

	pizzaFlavorRepository := repositories.NewPizzaFlavorRepository(db)
	pizzaFlavorService := services.NewPizzaFlavorService(pizzaFlavorRepository)
	pizzaFlavorHandler := handlers.NewPizzaFlavorHandler(pizzaFlavorService)

	pizzaSizeRepository := repositories.NewPizzaSizeRepository(db)
	pizzaSizeService := services.NewPizzaSizeService(pizzaSizeRepository)
	pizzaSizeHandler := handlers.NewPizzaSizeHandler(pizzaSizeService)


	router := gin.Default()

	router.GET("/health", func (c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H{
			"status":    "online",
			"framework": "gin",
			"database":  "connected",
		})
	})

	routes.RegisterSuperAdminRoutes(router, tenantHandler)
	routes.RegisterTenantAdminRoutes(router, *tenantService, categoryHandler, productHandler, pizzaFlavorHandler, pizzaSizeHandler)
	routes.RegisterPublicRoutes(router, *tenantService, categoryHandler, productHandler)

	log.Println("Iniciando API na porta 8080...")
	
	log.Fatal(router.Run(":8080"))
}