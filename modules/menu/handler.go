package menu

import (
	"pos/pkg/common"

	"github.com/gofiber/fiber/v2"
)

// CreateMenuItem handles POST requests to create new menu items.
func CreateMenuItem(c *fiber.Ctx) error {
	item, err := CreateItem(c)
	if err != nil {
		// Check if the error is a validation error
		if ve, ok := err.(*fiber.Error); ok && ve.Code == fiber.StatusBadRequest {
			// Parse the validation error message
			errMap := common.ParseValidationError(ve.Message)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": "Validation errors",
				"errors":  errMap,
			})
		}
		// Handle other types of errors
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(item)
}
