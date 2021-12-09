package handlers

import (
	"medical_system/config"
	"medical_system/entities"
	"medical_system/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type AdminHandler struct {
	srv  services.AdminService
	srvU services.UserService
}

func (h *AdminHandler) Register(e *echo.Echo) {
	users := e.Group("admin")
	jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: config.Instance.JWTSecret,
	})
	users.GET("/profile/:username", h.GetProfile, jwtMiddleware)
	users.POST("/login", h.Login)
	users.GET("/users/:type", h.GetUsers)
	users.GET("/admin/stats", h.Stats)
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

func (h *AdminHandler) Stats(ctx echo.Context) error {
}
