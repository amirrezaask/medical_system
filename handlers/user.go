package handlers

import (
	"medical_system/entities"
	"medical_system/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UsersHandler struct {
	srv services.UserService
}

func (h *UsersHandler) Register(e *echo.Echo) {
	users := e.Group("users")
	users.GET("/profile/:nationalNumber", h.GetProfile)
	users.POST("/login", h.Login)
	users.POST("/signup", h.SignUp)
}
func NewUsersHandler(srv services.UserService) *UsersHandler {
	return &UsersHandler{srv}
}
func (h *UsersHandler) SignUp(ctx echo.Context) error {
	var u entities.UserSignupRequest
	err := ctx.Bind(&u)

	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	err = h.srv.NewUser(u)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return nil
}
func (h *UsersHandler) Login(ctx echo.Context) error {
	var u entities.UserLoginRequest
	err := ctx.Bind(&u)

	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	token, err := h.srv.LoginUser(u)
	if err != nil {
		return ctx.NoContent(http.StatusUnauthorized)
	}
	return ctx.JSON(200, map[string]string{
		"token": token,
	})

}
func (h *UsersHandler) GetProfile(ctx echo.Context) error {
	var u entities.UserGetProfileRequest
	u.NationalNumber = ctx.Param("nationalNumber")
	user, err := h.srv.FindUser(u.NationalNumber)
	if err != nil {
		return ctx.NoContent(http.StatusUnauthorized)
	}
	return ctx.JSON(200, user)
}
