package dto

type UpdateBookRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	ISBN        string `json:"isbn" validate:"required"`
	Author      string `json:"author" validate:"required"`
	Genre       string `json:"genre" validate:"required"`
	Stock       int    `json:"stock" validate:"min=1"`
	PublishDate string `json:"publish_date" validate:"required"`
}
