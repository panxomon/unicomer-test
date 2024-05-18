package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"

	bootstrap "unicomer-test/cmd/bootstrap"
)

type Key string

const BootstrapKey Key = "bootstrap"

func SetUpRoutes(basePath string, components *bootstrap.Bootstrap) {

}

func SetupRouter(components *bootstrap.Bootstrap) *gin.Engine {
	r := gin.Default()

	r.GET("/health", HealthCheck)
	r.GET("/holidays")

	return r
}

// BootstrapMiddleware inyecta el componente Bootstrap en el contexto de Gin.
func BootstrapMiddleware(bootstrap *bootstrap.Bootstrap) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), BootstrapKey, bootstrap)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
