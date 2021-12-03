package entities

type User struct {
	Name           string `json:"name"`
	NationalNumber string `json:"national_number"`
}

type UserSignupRequest struct {
	User
	Password string `json:"password"`
}

type UserLoginRequest struct {
	NationalNumber string `json:"national_number"`
	Password       string `json:"password"`
}

type UserGetProfileRequest struct {
	NationalNumber string `json:"national_number"`
}
