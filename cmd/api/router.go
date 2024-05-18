package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	bootstrap "unicomer-test/cmd/bootstrap"
	"unicomer-test/internal/endpoint"
)

type Key string

const BootstrapKey Key = "bootstrap"

func SetupRouter(basePath string, components *bootstrap.Bootstrap) *gin.Engine {
	r := gin.Default()

	r.GET("/health", HealthCheck)
	r.GET(basePath, BootstrapMiddleware(components))

	r.GET("test", func(c *gin.Context) {
		endpoint.NewEndpoint(components.Holidays).Invoke(c)
	})

	return r
}

func BootstrapMiddleware(components *bootstrap.Bootstrap) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), BootstrapKey, components)
		endpoint.NewEndpoint(components.Holidays).Invoke(c)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
