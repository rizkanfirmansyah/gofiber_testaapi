package bookcontroller

import (
	"gofiber/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {
	var books []models.Book
	models.DB.Find(&books)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": books,
	})
}

func Show(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book

	if err := models.DB.First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Data no found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Data no found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": book,
	})
}

func Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if models.DB.Where("id = ?", id).Updates(&book).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data no found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data has been updated",
		"data":    book,
	})
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data has been created",
		"data":    book,
	})
}
func Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	var book models.Book

	if models.DB.Delete(&book, id).Updates(&book).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data no found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data has been deleted",
		"data":    book,
	})
}
