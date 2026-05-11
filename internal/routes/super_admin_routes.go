package routes

import (
	"cw/internal/handlers"

	"github.com/gin-gonic/gin"
)


func RegisterSuperAdminRoutes(router *gin.Engine, tenantHandler *handlers.TenantHandler) {

	adminGroup := router.Group("/api/admin") 

	{
		tenants := adminGroup.Group("/tenants")

		{
			tenants.POST("/", tenantHandler.CreateTenant)
			tenants.GET("/:id", tenantHandler.GetTenant)
			tenants.PUT("/:id", tenantHandler.UpdateTenant)
		}
	}
}