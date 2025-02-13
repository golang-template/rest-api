package routes

import (
	"REST-API/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/items", handlers.GetItems)
	e.POST("/items", handlers.CreateItem)
	e.PUT("/items/:id", handlers.UpdateItem)
	e.DELETE("/items/:id", handlers.DeleteItem)
}
