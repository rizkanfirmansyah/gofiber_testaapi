package main

import (
	"gofiber/controllers/bookcontroller"
	"gofiber/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	models.ConnectDatabase()

	app := fiber.New()

	apiBook := app.Group("/api/books")

	apiBook.Get("/", bookcontroller.Index)
	apiBook.Get("/:id", bookcontroller.Show)
	apiBook.Post("/", bookcontroller.Create)
	apiBook.Put("/:id", bookcontroller.Update)
	apiBook.Delete("/:id", bookcontroller.Delete)

	app.Listen(":8090")
}
