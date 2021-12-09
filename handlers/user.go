package handlers

import (
	"github.com/golang-jwt/jwt"
	"medical_system/config"
	"medical_system/entities"
	"medical_system/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type UsersHandler struct {
	srv services.UserService
}

func (h *UsersHandler) Register(e *echo.Echo) {
	jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: config.Instance.JWTSecret,
	})
	users := e.Group("users")
	users.GET("/profile", h.GetProfile, jwtMiddleware)
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
	userData := ctx.Get("user").(jwt.MapClaims)
	user, err := h.srv.FindUser(userData["national_code"].(string))
	if err != nil {
		return ctx.NoContent(http.StatusUnauthorized)
	}
	return ctx.JSON(200, user)
}
