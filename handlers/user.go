package handlers

import (
	"medical_system/services"
)

type UsersHandler struct {
	srv services.UserService
}

func NewUsersHandler(srv services.UserService) *UsersHandler {
	return &UsersHandler{srv}
}
