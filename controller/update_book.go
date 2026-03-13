package controller

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/AlvinSanudharma/books-api/database"
	"github.com/AlvinSanudharma/books-api/dto"
	"github.com/gofiber/fiber/v2"
)

func UpdateBookController(c *fiber.Ctx) error {
	id := c.Params("id")

	var request dto.UpdateBookRequest
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	row := database.DB.QueryRow("UPDATE books set title = $1, description = $2, author = $3, genre = $4, isbn = $5, stock = $6, publish_date = $7 WHERE id = $8 RETURNING id", request.Title, request.Description, request.Author, request.Genre, request.ISBN, request.Stock, request.PublishDate, id)
	if row.Err() != nil {
		return row.Err()
	}

	var returnedId int
	err = row.Scan(&returnedId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(http.StatusNotFound).JSON(map[string]any{
				"error": "Book not found",
			})
		}
		return err
	}

	return c.Status(http.StatusOK).JSON(map[string]any{
		"data": map[string]any{
			"id": returnedId,
		},
	})
}
