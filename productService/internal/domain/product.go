package domain

type Product struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	CategoryId  int64   `json:"category_id"`
	Price       float64 `json:"price"`
	Stock       uint    `json:"stock"`
	ImageURL    string  `json:"image_url"`
	Status      string  `json:"status"`
	CreatedAt   int64   `json:"created_at"`
	UpdatedAt   int64   `json:"updated_at"`
	Version     int     `json:"version"`
}
