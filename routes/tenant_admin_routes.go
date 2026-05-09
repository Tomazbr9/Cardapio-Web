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
) {

	adminGroup := router.Group("/api/t/:tenant_slug/admin")
	
	adminGroup.Use(middleware.SetTenantContext(tenantService))

	// adminGroup.Use(middleware.RequireTenantLogin())
	
	{
		categories := adminGroup.Group("/categories")
		{
			categories.POST("/", categoryHandler.CreateCategory)
			
		}
	}
}