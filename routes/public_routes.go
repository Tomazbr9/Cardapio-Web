package routes

import (
	"cw/internal/handlers"
	"cw/internal/middleware"
	"cw/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)


func RegisterPublicRoutes(
	router *gin.Engine,
	tenantService services.TenantService,
	categoryHandler *handlers.CategoryHandler,
) {

	menuGroup := router.Group("/t/:tenant_slug")

	menuGroup.Use(middleware.SetTenantContext(tenantService))

	{
		menuGroup.GET("/menu", func(c *gin.Context) {
			tenantId := c.MustGet("tenantID")

			c.JSON(http.StatusOK, gin.H{
				"message": "Você acessou o cardápio",
				"tenant_slug": c.Param("tenant_slug"),
				"injected_tenant_id": tenantId,
			})
		})

		menuGroup.GET("/categories", categoryHandler.ListTenantCategories)
	}
}