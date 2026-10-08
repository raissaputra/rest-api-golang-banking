package routes

import (
	"banking-management-api/internal/handlers"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api/v1")

	api.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Banking API siap di develop oleh rs",
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
			"status": "successBanget",
		})
	})

	api.Get("/customers", handlers.GetCustomers)
	api.Get("/customers/:id", handlers.GetCustomerByID)
	api.Post("/customers", handlers.CreateCustomer)
	api.Put("/customers/:id", handlers.UpdateCustomer)
	api.Delete("/customers/:id", handlers.DeleteCustomer)
}
