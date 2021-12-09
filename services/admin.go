package services

import (
	"context"
	"fmt"
	"medical_system/database/models"
	"medical_system/database/models/admin"
	"medical_system/entities"
)

type AdminService interface {
	FindAdmin(string) (*entities.Admin, error)
	LoginAdmin(entities.AdminLoginRequest) (string, error)
}

func NewAdminService(db *models.AdminClient, auth *AuthService) AdminService {
	return &adminService{
		db:   db,
		auth: auth,
	}
}

type adminService struct {
	db   *models.AdminClient
	auth *AuthService
}

func (srv *adminService) FindAdmin(username string) (*entities.Admin, error) {
	u, err := srv.db.Query().Where(admin.Username(username)).First(context.Background())
	if err != nil {
		return nil, err
	}
	return &entities.Admin{Username: u.Username}, nil
}

func (srv *adminService) LoginAdmin(u entities.AdminLoginRequest) (string, error) {
	admin, err := srv.db.Query().Where(admin.Username(u.Username)).First(context.Background())
	if err != nil {
		return "", err
	}
	if !srv.auth.CheckPasswordHash(u.Password, admin.PasswordHash) {
		return "", fmt.Errorf("passwords do not match")
	}
	token, err := srv.auth.MakeJWT(admin.Username, "admin", "")
	if err != nil {
		return "", err
	}
	return token, err
}
