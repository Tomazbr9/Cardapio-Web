package routes

import (
	"cw/internal/handlers"
	"cw/internal/middleware"
	"cw/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterTenantAdminRoutes(
	router *gin.Engine,
	tenantService services.TenantService,
	categoryHandler *handlers.CategoryHandler,
	productHandler *handlers.ProductHandler, 
	pizzaFlavorHandler *handlers.PizzaFlavorHandler,
	pizzaSizeHandler *handlers.PizzaSizeHandler,
) {

	adminGroup := router.Group("/api/t/:tenant_slug/admin")
	
	adminGroup.Use(middleware.SetTenantContext(tenantService))

	// adminGroup.Use(middleware.RequireTenantLogin())
	
	{
		categories := adminGroup.Group("/categories")
		{
			categories.POST("/", categoryHandler.CreateCategory)
			
		}

		products := adminGroup.Group("/products")
		{
			products.GET("/:id", productHandler.GetProduct)
			products.POST("/", productHandler.CreateProduct)
		}

		pizzaFlavors := adminGroup.Group("/pizza-flavors")
		{
			pizzaFlavors.GET("/", pizzaFlavorHandler.ListFlavors)
			pizzaFlavors.GET("/:id", pizzaSizeHandler.GetPizzaFlavor)
			pizzaFlavors.POST("/", pizzaFlavorHandler.CreatePizzaFlavor)
		}

		pizzaSizes := adminGroup.Group("/pizza-sizes")
		{
			pizzaSizes.GET("/", pizzaSizeHandler.ListPizzaSizes)
			pizzaSizes.GET("/:id", pizzaSizeHandler.GetPizzaSize)
			pizzaSizes.POST("/", pizzaSizeHandler.CreatePizzaSize)
		}
	}
}