package controller

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/AlvinSanudharma/books-api/database"
	"github.com/AlvinSanudharma/books-api/dto"
	"github.com/gofiber/fiber/v2"
)

func ShowBookController(c *fiber.Ctx) error {
	id := c.Params("id")

	var response dto.ShowBook

	row := database.DB.QueryRow("SELECT id, title, description, isbn, author, genre, stock, publish_date FROM books WHERE id = $1", id)
	if row.Err() != nil {
		fmt.Println("Error executing query:", row.Err())

		return row.Err()
	}

	err := row.Scan(&response.ID, &response.Title, &response.Description, &response.ISBN, &response.Author, &response.Genre, &response.Stock, &response.PublishDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(http.StatusNotFound).JSON(map[string]any{
				"message": "Book not found",
			})
		}

		return err
	}

	return c.Status(http.StatusOK).JSON(map[string]any{
		"data": response,
	})
}
