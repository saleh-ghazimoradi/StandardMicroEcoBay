package dto

import "github.com/saleh-ghazimoradi/StandardMicroEcoBay/internal/helper"

type UserSignup struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ForgotPassword struct {
	Email string `json:"email"`
}

type SetPassword struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

type UserProfile struct {
	UserId    int64   `json:"-"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Email     *string `json:"email"`
	Phone     *string `json:"phone"`
	Address   Address `json:"address"`
}

type Address struct {
	AddressLine1 *string `json:"address_line1"`
	AddressLine2 *string `json:"address_line2"`
	City         *string `json:"city"`
	PostCode     *string `json:"post_code"`
	Country      *string `json:"country"`
}

func validateEmail(v *helper.Validator, email string) {
	v.Check(email != "", "email", "must be provided")
	v.Check(helper.Matches(email, helper.EmailRX), "email", "must be a valid email address")
}

func validatePassword(v *helper.Validator, password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) >= 8, "password", "must be at least 8 bytes long")
	v.Check(len(password) <= 72, "password", "must not be more than 72 bytes long")
}

func validatePhone(v *helper.Validator, phone string) {
	v.Check(phone != "", "phone", "must be provided")
}

func validToken(v *helper.Validator, token string) {
	v.Check(token != "", "token", "must be provided")
}

func ValidateUserSignup(v *helper.Validator, user *UserSignup) {
	validateEmail(v, user.Email)
	if user.Password != "" {
		validatePassword(v, user.Password)
	}
	validatePhone(v, user.Phone)
}

func ValidateUserLogin(v *helper.Validator, user *UserLogin) {
	validateEmail(v, user.Email)
	if user.Password != "" {
		validatePassword(v, user.Password)
	}
}

func ValidateForgotPassword(v *helper.Validator, forgotPassword *ForgotPassword) {
	validateEmail(v, forgotPassword.Email)
}

func ValidateSetPassword(v *helper.Validator, setPassword *SetPassword) {
	validToken(v, setPassword.Token)
	validatePassword(v, setPassword.Password)
}
