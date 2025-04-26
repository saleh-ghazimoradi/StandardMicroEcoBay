package domain

type Address struct {
	ID           int64  `json:"id"`
	AddressLine1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2"`
	City         string `json:"city"`
	PostCode     string `json:"post_code"`
	Country      string `json:"country"`
	UserId       int64  `json:"user_id"`
}
