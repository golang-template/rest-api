package main

import (
	"REST-API/middlewares"
	"REST-API/routes"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover()) // Graceful recovery from panics
	e.Use(middlewares.CORSConfig())
	e.Use(middlewares.Logger())
	e.Validator = middlewares.NewValidator() // Request Validation

	// Routes
	routes.RegisterRoutes(e)

	// Start Server
	log.Println(" Server is running on port :8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatal(" Server Error: ", err)
	}
}
