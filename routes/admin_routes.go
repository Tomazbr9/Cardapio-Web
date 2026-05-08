package routes

import (
	"cw/internal/controllers"

	"github.com/gin-gonic/gin"
)


func RegisterAdminRoutes(router *gin.Engine, tenantHandler *controllers.TenantHandler) {

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