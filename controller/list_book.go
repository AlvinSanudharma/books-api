package controller

import (
	"fmt"
	"net/http"

	"github.com/AlvinSanudharma/books-api/database"
	"github.com/AlvinSanudharma/books-api/dto"
	"github.com/gofiber/fiber/v2"
)

func ListBookController(c *fiber.Ctx) error {

	rows, err := database.DB.Query("SELECT id, title, description, isbn, author, genre, stock, publish_date FROM books")
	if err != nil {
		fmt.Println("Error executing query:", err)

		return err
	}

	var response []dto.ListBook

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
