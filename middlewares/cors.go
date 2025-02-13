package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CORSConfig handles Cross-Origin Resource Sharing (CORS)
func CORSConfig() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // Change this in production
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})
}
