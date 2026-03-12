package main

import (
	"log"

	"github.com/AlvinSanudharma/books-api/controller"
	"github.com/AlvinSanudharma/books-api/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.InitDb()

	app.Post("/api/books", controller.CreateBookController)
	app.Get("/api/books/:id", controller.ShowBookController)
	app.Delete("/api/books/:id", controller.DeleteBookController)

	log.Fatal(app.Listen(":3000"))
}
