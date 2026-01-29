package docs

import (
	"net/http"

	"RESTAURANT-MANAGEMENT/internal/config"

	"github.com/gin-gonic/gin"
)

// RegisterDocs registers routes to serve static API documentation when enabled.
func RegisterDocs(router *gin.Engine) {
	// Serve files in the docs folder under the /docs path
	router.Static("/docs", config.DOCS_PATH)

	// Convenience redirect: /swagger -> /docs/swagger.html
	router.GET("/swagger", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/docs/swagger.html")
	})
}
