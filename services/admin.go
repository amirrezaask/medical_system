package services

import (
	"context"
	"fmt"
	"medical_system/database/models"
	"medical_system/database/models/admin"
	"medical_system/database/models/user"
	"medical_system/entities"
	"time"
)

type AdminService interface {
	FindAdmin(string) (*entities.Admin, error)
	LoginAdmin(entities.AdminLoginRequest) (string, error)
	Stats() (prescriptionsCount int64, newPatients []entities.User, newDoctors []entities.User, err error)
}

func NewAdminService(db *models.AdminClient, auth *AuthService, userDB *models.UserClient, prescriptionDB *models.PrescriptionClient) AdminService {
	return &adminService{
		adminDB:        db,
		auth:           auth,
		userDB:         userDB,
		prescriptionDB: prescriptionDB,
	}
}

type adminService struct {
	adminDB        *models.AdminClient
	userDB         *models.UserClient
	prescriptionDB *models.PrescriptionClient
	auth           *AuthService
}

func (srv *adminService) FindAdmin(username string) (*entities.Admin, error) {
	u, err := srv.adminDB.Query().Where(admin.Username(username)).First(context.Background())
	if err != nil {
		return nil, err
	}
	return &entities.Admin{Username: u.Username}, nil
}

func (srv *adminService) LoginAdmin(u entities.AdminLoginRequest) (string, error) {
	admin, err := srv.adminDB.Query().Where(admin.Username(u.Username)).First(context.Background())
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

func (srv *adminService) Stats() (prescriptionsCount int64, newPatients []entities.User, newDoctors []entities.User, err error) {
	count, err := srv.prescriptionDB.Query().Count(context.Background())
	if err != nil {
		return -1, nil, nil, err
	}
	y, m, d := time.Now().Date()
	timeStr := fmt.Sprintf("%d-%d-%dT00:00:00+00:00", y, m, d)
	t, _ := time.Parse(time.RFC3339, timeStr)
	ps, err := srv.userDB.Query().Where(user.CreatedAtLTE(t), user.UserTypeEQ("patient")).All(context.Background())
	if err != nil {
		return -1, nil, nil, err
	}
	ds, err := srv.userDB.Query().Where(user.CreatedAtLTE(t), user.UserTypeEQ("doctor")).All(context.Background())
	if err != nil {
		return -1, nil, nil, err
	}
	return count, ps, ds, nil
}
