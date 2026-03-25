package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/AlvinSanudharma/books-api/database"
	"github.com/AlvinSanudharma/books-api/dto"
	"github.com/gofiber/fiber/v2"
)

func ListBookController(c *fiber.Ctx) error {
	var request dto.ListBookRequest

	err := c.QueryParser(&request)
	if err != nil {
		return err
	}

	query := "SELECT id, title, description, isbn, author, genre, stock, publish_date FROM books"
	var args []any
	if request.Search != "" {
		query += " WHERE LOWER(title) LIKE $1"
		filter := fmt.Sprintf("%%%s%%", strings.ToLower(request.Search))

		args = append(args, filter)
	}

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		fmt.Println("Error executing query:", err)

		return err
	}

	response := make([]dto.ListBook, 0)

	for rows.Next() {
		var book dto.ListBook

		err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.ISBN, &book.Author, &book.Genre, &book.Stock, &book.PublishDate)
		if err != nil {
			return err
		}

		response = append(response, book)
	}

	return c.Status(http.StatusOK).JSON(map[string]any{
		"data": response,
	})
}
