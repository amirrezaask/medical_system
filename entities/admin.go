package entities

type AdminLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
