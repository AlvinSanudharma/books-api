package main

import (
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/AlvinSanudharma/books-api/controller"
	"github.com/AlvinSanudharma/books-api/database"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(
		fiber.Config{
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				validationErrs, ok := err.(validator.ValidationErrors)

				if ok {
					errRes := make(map[string]string)
					for _, validationErr := range validationErrs {
						tag := validationErr.Tag()
						message := "Unknown error"

						switch tag {
						case "required":
							message = "Required"
						case "min":
							minimum := validationErr.Param()

							message = "Minimum should be " + minimum
						}

						re := regexp.MustCompile("([a-z])([A-Z])")
						snakeCase := re.ReplaceAllString(validationErr.Field(), "${1}_${2}")
						errRes[strings.ToLower(snakeCase)] = message
					}

					return c.Status(http.StatusBadRequest).JSON(map[string]any{
						"error": errRes,
					})
				}

				return c.Status(http.StatusInternalServerError).JSON(map[string]any{
					"message": "Internal Server Error",
				})
			},
		},
	)

	database.InitDb()

	app.Post("/api/books", controller.CreateBookController)
	app.Get("/api/books/:id", controller.ShowBookController)
	app.Delete("/api/books/:id", controller.DeleteBookController)
	app.Put("/api/books/:id", controller.UpdateBookController)
	app.Get("/api/books", controller.ListBookController)

	log.Fatal(app.Listen(":3000"))
}
