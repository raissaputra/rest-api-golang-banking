package handlers

import (
	"banking-management-api/internal/dto"
	"banking-management-api/internal/models"
	"banking-management-api/internal/validators"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

var customers = []models.Customer{
	{
		ID:      1,
		Name:    "John Doe",
		Email:   "john.doe@example.com",
		Phone:   "123-456-7890",
		Address: "123 Main St, City, State 12345",
	},
	{
		ID:      2,
		Name:    "Jane Smith",
		Email:   "jane.smith@example.com",
		Phone:   "098-765-4321",
		Address: "456 Oak Ave, Town, State 67890",
	},
}

// GET - Mendapatkan semua customer
func GetCustomers(c fiber.Ctx) error {
	return c.JSON(customers)
}

// GET - Mendapatkan customer berdasarkan ID
func GetCustomerByID(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid customer ID",
		})
	}

	for _, customer := range customers {
		if customer.ID == id {
			return c.JSON(customer)
		}
	}

	return c.Status(404).JSON(fiber.Map{
		"error": "Customer not found",
	})
}

// POST - Mmembuat customer baru
func CreateCustomer(c fiber.Ctx) error {
	var newCustomer dto.CreateCustomerRequest

	if err := c.Bind().Body(&newCustomer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"errors":  err.Error(),
		})
	}

	if err := validators.Validate.Struct(newCustomer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  err.Error(),
		})
	}

	customer := models.Customer{
		ID:      len(customers) + 1, // Generate ID based on the length of the customers slice
		Name:    newCustomer.Name,
		Email:   newCustomer.Email,
		Phone:   newCustomer.Phone,
		Address: newCustomer.Address,
	}

	customers = append(customers, customer)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":  "Customer created successfully",
		"data": customer,
	})
}

// PUT - Memperbarui customer berdasarkan ID
func UpdateCustomer(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid customer ID",
		})
	}

	var input models.Customer

	if err := c.Bind().Body(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	for i, customer := range customers {
		if customer.ID == id {
			input.ID = customer.ID
			customers[i] = input
			return c.JSON(input)
		}
	}

	return c.Status(404).JSON(fiber.Map{
		"message": "Customer not found",
	})

}

// DELETE - Menghapus customer berdasarkan ID
func DeleteCustomer(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid customer ID",
		})
	}

	for i, customer := range customers {
		if customer.ID == id {
			customers = append(customers[:i], customers[i+1:]...)
			return c.JSON(fiber.Map{
				"message": "Customer deleted successfully",
			})
		}
	}

	return c.Status(404).JSON(fiber.Map{
		"message": "Customer not found",
	})
}
