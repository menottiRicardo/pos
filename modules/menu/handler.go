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

func GetMenuItem(c *fiber.Ctx) error {
	id := c.Params("id")
	item, err := GetItem(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Item not found"})
	}

	return c.JSON(item)
}

func UpdateMenuItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var updateData MenuItem
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	updatedItem, err := UpdateItem(id, updateData) // Database interaction
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(updatedItem)
}

func DeleteMenuItem(c *fiber.Ctx) error {
	id := c.Params("id")
	itemId, err := DeleteItem(id)
	if err != nil { // Database interaction
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"id": itemId})
}

func CreateCategoryHandler(c *fiber.Ctx) error {
	item, err := CreateCategory(c)
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

func UpdateCategoryHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	var updateData MenuCategory
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	updatedItem, err := UpdateCategory(id, updateData) // Database interaction
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(updatedItem)
}

func GetCategoryHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	item, err := GetCategory(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Item not found"})
	}

	return c.JSON(item)
}

func CreateMenuHandler(c *fiber.Ctx) error {
	item, err := CreateMenu(c)
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

func UpdateMenuHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	var updateData Menu
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	updatedItem, err := UpdateMenu(id, updateData) // Database interaction
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(updatedItem)
}

func GetMenuHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	item, err := GetCompleteMenu(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Menu not found"})
	}

	return c.JSON(item)
}
