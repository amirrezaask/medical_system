package entities

type User struct {
	Name           string `json:"name"`
	NationalNumber string `json:"national_number"`
}

type UserCreateRequest struct {
	User
	Password string `json:"password"`
}
