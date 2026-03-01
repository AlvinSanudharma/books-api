package dto

type CreateBookRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ISBN        string `json:"isbn"`
	Author      string `json:"author"`
	Genre       string `json:"genre"`
	Stock       int    `json:"stock"`
	PublishDate string `json:"publish_date"`
}
