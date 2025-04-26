package domain

type User struct {
	ID         int64   `json:"id"`
	Email      string  `json:"email"`
	Password   string  `json:"-"`
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	Phone      string  `json:"phone"`
	ResetToken string  `json:"reset_token"`
	Address    Address `json:"address"`
}
