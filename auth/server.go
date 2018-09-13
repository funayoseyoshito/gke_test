package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello auth")
	})

	r.GET("/introspect", func(c *gin.Context) {
		c.String(http.StatusOK, "introspect response(from auth)")
	})

	r.GET("/consent", func(c *gin.Context) {
		c.String(http.StatusOK, "consent response(from auth)")
	})

	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "auth OK")
	})
	r.Run(":7001")
}
