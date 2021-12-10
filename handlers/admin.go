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
	users.GET("/users/:type", h.GetUsers, jwtMiddleware)
	users.GET("/admin/stats", h.Stats, jwtMiddleware)
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
	count, newP, newD, err := h.srv.Stats()
	if err != nil {
		return ctx.JSON(500, err)
	}

	return ctx.JSON(200, map[string]interface{}{
		"count": count,
		"new_patients": newP,
		"new_doctors": newD,
	})
}
