package main

import (
	"log"

	"github.com/AlvinSanudharma/books-api/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.InitDb()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
