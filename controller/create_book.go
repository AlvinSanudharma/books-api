package controller

import (
	"net/http"

	"github.com/AlvinSanudharma/books-api/database"
	"github.com/AlvinSanudharma/books-api/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CreateBookController(c *fiber.Ctx) error {
	var request dto.CreateBookRequest

	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	validate := validator.New()

	err = validate.Struct(request)
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
		return err
	}

	return c.Status(http.StatusOK).JSON(map[string]any{
		"data": map[string]int{
			"id": id,
		},
	})
}
