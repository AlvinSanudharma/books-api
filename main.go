package main

import (
	"log"
	"net/http"

	"github.com/AlvinSanudharma/books-api/database"
	"github.com/gofiber/fiber/v2"
)

type CreateBookRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ISBN        string `json:"isbn"`
	Author      string `json:"author"`
	Genre       string `json:"genre"`
	Stock       int    `json:"stock"`
	PublishDate string `json:"publish_date"`
}

func main() {
	app := fiber.New()

	database.InitDb()

	app.Post("/api/books", func(c *fiber.Ctx) error {
		var request CreateBookRequest

		err := c.BodyParser(&request)
		if err != nil {
			return err
		}

		row := database.DB.QueryRow("INSERT INTO books (title, description, author, genre, isbn, stock, publish_date) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", request.Title, request.Description, request.Author, request.Genre, request.ISBN, request.Stock, request.PublishDate)

		if row.Err() != nil {
			return row.Err()
		}

		var id int
		err = row.Scan(&id)
		if err != nil {
			log.Printf("Scan error: %v", err)
			return err
		}

		return c.Status(http.StatusOK).JSON(map[string]any{
			"data": map[string]int{
				"id": id,
			},
		})
	})

	log.Fatal(app.Listen(":3000"))
}
