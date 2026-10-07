package routes

import (

	"github.com/gofiber/fiber/v3"
	"banking-management-api/internal/handlers"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api/v1")

	api.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Banking API is running",
		})
	})

	api.Get("/about", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "This is a banking API",
			"version": "1.0.0",
		})
	})	
	
	api.Get("/contact", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "success",
		})
	})

	api.Get("/customers", handlers.GetCustomers)	
	api.Get("/customers/:id", handlers.GetCustomerByID)
	api.Post("/customers", handlers.CreateCustomer)
	api.Put("/customers/:id", handlers.UpdateCustomer)
	api.Delete("/customers/:id", handlers.DeleteCustomer)
}