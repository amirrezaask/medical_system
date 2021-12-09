package handlers

import (
	"medical_system/config"
	"medical_system/entities"
	"medical_system/services"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"

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
	users.POST("/prescriptions", h.AddPrescription)
	users.GET("/prescriptions/:patientID", h.GetPrescriptions)
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

func (h *UsersHandler) AddPrescription(ctx echo.Context) error {
	doctor := ctx.Get("user").(jwt.MapClaims)
	if doctor["user_type"] != "doctor" {
		return ctx.String(401, "only doctors can add prescription")
	}
	var req entities.AddPrescriptionRequest
	err := ctx.Bind(&req)
	if err != nil {
		ctx.JSON(400, err)
	}
	err = h.srv.AddPrescription(req.PatientNationalCode, entities.Prescription{})
	if err != nil {
		ctx.JSON(500, err)
	}
	return ctx.NoContent(204)
}

func (h *UsersHandler) GetPrescriptions(ctx echo.Context) error {
	user := ctx.Get("user").(jwt.MapClaims)
	idStr := ctx.Param("patientID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.JSON(400, err)
	}
	var ps interface{}
	if user["user_type"] == "doctor" {
		ps, err = h.srv.GetPrescriptionsForDoctor(id)
	} else if user["user_type"] == "patient" {
		ps, err = h.srv.GetPrescriptionsForPatient(id)
	} else if user["user_type"] == "admin" {
		//admin
	}
	if err != nil {
		return ctx.JSON(500, err)
	}
	return ctx.JSON(200, ps)
}
