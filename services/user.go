package services

import (
	"context"
	"fmt"
	"medical_system/database/models"
	"medical_system/database/models/user"
	"medical_system/entities"
)

type UserService interface {
	NewUser(entities.UserSignupRequest) error
	FindUser(nationalNumber string) (*entities.User, error)
	UpdateUser(entities.User) error
	DeleteUser(nationalNumber string) error
	LoginUser(entities.UserLoginRequest) (string, error)
	GetUsers(typ string) ([]*entities.User, error)
}

func NewUserService(db *models.UserClient, auth *AuthService) UserService {
	return &userService{
		db:   db,
		auth: auth,
	}
}

type userService struct {
	db   *models.UserClient
	auth *AuthService
}

func (srv *userService) GetUsers(typ string) ([]*entities.User, error) {
	users, err := srv.db.Query().Where(user.UserTypeEQ(user.UserType(typ))).All(context.Background())
	if err != nil {
		return nil, err
	}
	var usersE []*entities.User
	for _, u := range users {
		usersE = append(usersE, &entities.User{Name: u.Name, NationalNumber: u.NationalCode})
	}

	return usersE, nil

}

func (srv *userService) NewUser(u entities.UserSignupRequest) error {
	passwordHash, err := srv.auth.HashPassword(u.Password)
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

func (srv *userService) LoginUser(u entities.UserLoginRequest) (string, error) {
	user, err := srv.db.Query().Where(user.NationalCode(u.NationalNumber)).First(context.Background())
	if err != nil {
		return "", err
	}
	if !srv.auth.CheckPasswordHash(u.Password, user.PasswordHash) {
		return "", fmt.Errorf("passwords do not match")
	}
	token, err := srv.auth.MakeJWT(user.NationalCode)
	if err != nil {
		return "", err
	}
	return token, err
}
