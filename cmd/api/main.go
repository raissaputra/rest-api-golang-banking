package main

import (
	"log"
	"github.com/gofiber/fiber/v3"
	"banking-management-api/internal/routes"
)

func main () {
	app := fiber.New()
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3300"))

}