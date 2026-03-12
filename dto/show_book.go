package dto

type ShowBook struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ISBN        string `json:"isbn"`
	Author      string `json:"author"`
	Genre       string `json:"genre"`
	Stock       int    `json:"stock"`
	PublishDate string `json:"publish_date"`
}
