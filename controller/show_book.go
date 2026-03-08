package controller

import (
	"fmt"
	"net/http"

	"github.com/AlvinSanudharma/books-api/database"
	"github.com/AlvinSanudharma/books-api/dto"
	"github.com/gofiber/fiber/v2"
)

func ShowBookController(c *fiber.Ctx) error {
	id := c.Params("id")

	var response dto.BookResponse

	row := database.DB.QueryRow("SELECT id, title, description FROM books WHERE id = $1", id)
	if row.Err() != nil {
		fmt.Println("Error executing query:", row.Err())

		return row.Err()
	}

	err := row.Scan(&response.ID, &response.Title, &response.Description)
	if err != nil {
		message := err.Error()

		if message == "sql: no rows in result set" {
			return c.Status(http.StatusNotFound).JSON(map[string]any{
				"error": "Book not found",
			})
		}

		return err
	}

	return c.Status(http.StatusOK).JSON(map[string]any{
		"data": response,
	})
}
