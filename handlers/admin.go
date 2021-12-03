package handlers

import (
	"medical_system/entities"
	"medical_system/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	srv  services.AdminService
	srvU services.UserService
}

func (h *AdminHandler) Register(e *echo.Echo) {
	users := e.Group("admin")
	users.GET("/profile/:username", h.GetProfile)
	users.POST("/login", h.Login)
	users.GET("/users/:type", h.GetUsers)
}
func NewAdminHandler(srv services.AdminService, userSrv services.UserService) *AdminHandler {
	return &AdminHandler{srv, userSrv}
}
func (h *AdminHandler) GetUsers(ctx echo.Context) error {
	typ := ctx.Param("type")
	users, err := h.srvU.GetUsers(typ)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(200, users)
}
func (h *AdminHandler) Login(ctx echo.Context) error {
	var u entities.AdminLoginRequest
	err := ctx.Bind(&u)

	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	token, err := h.srv.LoginAdmin(u)
	if err != nil {
		return ctx.NoContent(http.StatusUnauthorized)
	}
	return ctx.JSON(200, map[string]string{
		"token": token,
	})

}
func (h *AdminHandler) GetProfile(ctx echo.Context) error {
	user, err := h.srv.FindAdmin(ctx.Param("username"))
	if err != nil {
		return ctx.NoContent(http.StatusUnauthorized)
	}
	return ctx.JSON(200, user)
}
