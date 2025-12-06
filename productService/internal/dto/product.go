package dto

type CreateProduct struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	CategoryId  int64   `json:"category_Id"`
	Price       float64 `json:"price"`
	Stock       uint    `json:"stock"`
	ImageURL    string  `json:"imageUrl"`
	Status      string  `json:"status"`
}

type UpdateProduct struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price"`
	Stock       *uint    `json:"stock"`
	ImageURL    *string  `json:"image_url"`
	Status      *string  `json:"status"`
}
