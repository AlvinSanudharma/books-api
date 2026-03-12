package controller

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/AlvinSanudharma/books-api/database"
	"github.com/gofiber/fiber/v2"
)

func DeleteBookController(c *fiber.Ctx) error {
	id := c.Params("id")

	row := database.DB.QueryRow("DELETE FROM books WHERE id = $1 RETURNING id", id)
	if row.Err() != nil {
		return row.Err()
	}

	var returnedID int
	err := row.Scan(&returnedID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(http.StatusNotFound).JSON(map[string]any{
				"message": "Book not found",
			})
		}
	}

	return c.Status(http.StatusOK).JSON(map[string]any{
		"message": "Book deleted successfully",
	})
}
