package response

import "github.com/gofiber/fiber/v3"

func Success(c fiber.Ctx, status int, message string, data any) error {
	return c.Status(status).JSON(struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Data    any    `json:"data"`
	}{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(c fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{
		Success: false,
		Message: message,
	})
}