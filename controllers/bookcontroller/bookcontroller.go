package bookcontroller

import (
	"gofiber/models"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	var books []models.Book
	models.DB.Find(&books)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": books,
	})
}

func Show(c *fiber.Ctx) error {
	return nil
}
func Update(c *fiber.Ctx) error {
	return nil
}
func Create(c *fiber.Ctx) error {
	var book models.Book

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := models.DB.Create(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return nil
}
func Delete(c *fiber.Ctx) error {
	return nil
}
