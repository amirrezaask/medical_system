package services

import (
	"context"
	"medical_system/database/models"
	"medical_system/database/models/user"
	"medical_system/entities"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	NewUser(entities.UserCreateRequest) error
	FindUser(nationalNumber string) (*entities.User, error)
	UpdateUser(entities.User) error
	DeleteUser(nationalNumber string) error
}
type userService struct {
	db *models.UserClient
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func (srv *userService) NewUser(u entities.UserCreateRequest) error {
	passwordHash, err := HashPassword(u.Password)
	if err != nil {
		return err
	}
	_, err = srv.db.Create().
		SetName(u.Name).
		SetNationalCode(u.NationalNumber).
		SetPasswordHash(passwordHash).Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (srv *userService) FindUser(nationalNumber string) (*entities.User, error) {
	u, err := srv.db.Query().Where(user.NationalCodeEQ(nationalNumber)).First(context.Background())
	if err != nil {
		return nil, err
	}
	return &entities.User{Name: u.Name, NationalNumber: u.NationalCode}, nil
}

func (srv *userService) UpdateUser(u entities.User) error {
	return srv.db.Update().Where(user.NationalCodeEQ(u.NationalNumber)).SetName(u.Name).Exec(context.Background())
}

func (srv *userService) DeleteUser(nationalNumber string) error {
	_, err := srv.db.Delete().Where(user.NationalCodeEQ(nationalNumber)).Exec(context.Background())
	return err
}
