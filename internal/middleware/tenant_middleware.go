package middleware

import (
	"cw/internal/services"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetTenantContext(service services.TenantService) gin.HandlerFunc {
	return func(c *gin.Context) {
		
		slug := c.Param("tenant_slug")
		if slug == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Tenant slug inválido na URL"})
			c.Abort() 
			return
		}

		tenant, err := service.GetBySlug(slug)
		if err != nil {
			
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Pizzaria não encontrada"})
				c.Abort()
				return
			}
			
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao validar pizzaria"})
			c.Abort()
			return
		}
	
		c.Set("tenantID", tenant.ID)

		c.Next()
	}
}