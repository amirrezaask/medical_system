package services

import (
	"context"
	"fmt"
	"medical_system/database/models"
	"medical_system/database/models/prescription"
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
	AddPrescription(patientNationalCode string, u entities.Prescription) error
	GetPrescriptionsForPatient(patientID int) ([]entities.PrescriptionForPatient, error)
	GetPrescriptionsForDoctor(patientID int) ([]entities.PrescriptionForDoctor, error)
	GetPrescriptionsForAdmin(patientID int) ([]entities.PrescriptionForAdmin, error)
}

func NewUserService(db *models.UserClient, auth *AuthService, pDB *models.PrescriptionClient) UserService {
	return &userService{
		userDB:         db,
		auth:           auth,
		prescriptionDB: pDB,
	}
}

type userService struct {
	userDB         *models.UserClient
	prescriptionDB *models.PrescriptionClient
	auth           *AuthService
}

func (srv *userService) GetUsers(typ string) ([]*entities.User, error) {
	users, err := srv.userDB.Query().Where(user.UserTypeEQ(user.UserType(typ))).All(context.Background())
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
	_, err = srv.userDB.Create().
		SetName(u.Name).
		SetNationalCode(u.NationalNumber).
		SetPasswordHash(passwordHash).Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (srv *userService) FindUser(nationalNumber string) (*entities.User, error) {
	u, err := srv.userDB.Query().Where(user.NationalCodeEQ(nationalNumber)).First(context.Background())
	if err != nil {
		return nil, err
	}
	return &entities.User{Name: u.Name, NationalNumber: u.NationalCode}, nil
}

func (srv *userService) UpdateUser(u entities.User) error {
	return srv.userDB.Update().Where(user.NationalCodeEQ(u.NationalNumber)).SetName(u.Name).Exec(context.Background())
}

func (srv *userService) DeleteUser(nationalNumber string) error {
	_, err := srv.userDB.Delete().Where(user.NationalCodeEQ(nationalNumber)).Exec(context.Background())
	return err
}

func (srv *userService) LoginUser(u entities.UserLoginRequest) (string, error) {
	user, err := srv.userDB.Query().Where(user.NationalCode(u.NationalCode)).First(context.Background())
	if err != nil {
		return "", err
	}
	if !srv.auth.CheckPasswordHash(u.Password, user.PasswordHash) {
		return "", fmt.Errorf("passwords do not match")
	}
	token, err := srv.auth.MakeJWT(user.NationalCode, user.UserType.String(), u.NationalCode)
	if err != nil {
		return "", err
	}
	return token, err
}

func (srv *userService) AddPrescription(patientNationalCode string, u entities.Prescription) error {
	user, err := srv.userDB.Query().Where(user.NationalCodeEQ(patientNationalCode)).First(context.Background())
	if err != nil {
		return err
	}
	return srv.userDB.UpdateOneID(user.ID).AddPrescriptions(&models.Prescription{
		DoctorID:            int64(u.DoctorID),
		PatientNationalCode: u.PatientNationalCode,
		DrugsCommaSeperated: u.Drugs,
	}).Exec(context.Background())
}

func (srv *userService) GetPrescriptionsForDoctor(patientID int) ([]entities.PrescriptionForDoctor, error) {
	user, err := srv.userDB.Query().Where(user.IDEQ(patientID)).WithPrescriptions().First(context.Background())
	if err != nil {
		return nil, err
	}
	var ps []entities.PrescriptionForDoctor
	for _, p := range user.Edges.Prescriptions {
		ps = append(ps, entities.PrescriptionForDoctor{
			Patient: &entities.User{
				Name:           user.Name,
				NationalNumber: user.NationalCode,
			},
			Prescription: entities.Prescription{
				PatientNationalCode: p.PatientNationalCode,
				DoctorID:            int(p.DoctorID),
				Drugs:               p.DrugsCommaSeperated,
			},
		})
	}
	return ps, nil
}
func (srv *userService) GetPrescriptionsForAdmin(patientID int) ([]entities.PrescriptionForAdmin, error) {
	press, err := srv.prescriptionDB.Query().Where(prescription.IDEQ(patientID)).WithUsers().All(context.Background())
	if err != nil {
		return nil, err
	}
	patient, err := srv.userDB.Query().Where(user.IDEQ(patientID)).WithPrescriptions().First(context.Background())
	if err != nil {
		return nil, err
	}
	var ps []entities.PrescriptionForAdmin
	for _, p := range press {
		doctor, err := srv.userDB.Query().Where(user.IDEQ(int(p.DoctorID))).First(context.Background())
		if err != nil {
			return nil, err
		}
		ps = append(ps, entities.PrescriptionForAdmin{
			Doctor:  &entities.User{Name: doctor.Name, NationalNumber: doctor.NationalCode},
			Patient: &entities.User{Name: patient.Name, NationalNumber: patient.NationalCode},
			Prescription: entities.Prescription{
				PatientNationalCode: p.Edges.Users.NationalCode,
				DoctorID:            doctor.ID,
				Drugs:               p.DrugsCommaSeperated,
			},
		})
	}
	return ps, nil
}

func (srv *userService) GetPrescriptionsForPatient(patientID int) ([]entities.PrescriptionForPatient, error) {
	press, err := srv.prescriptionDB.Query().Where(prescription.IDEQ(patientID)).WithUsers().All(context.Background())
	if err != nil {
		return nil, err
	}
	var ps []entities.PrescriptionForPatient
	for _, p := range press {
		doctor, err := srv.userDB.Query().Where(user.IDEQ(int(p.DoctorID))).First(context.Background())
		if err != nil {
			return nil, err
		}
		ps = append(ps, entities.PrescriptionForPatient{
			Doctor: &entities.User{Name: doctor.Name, NationalNumber: doctor.NationalCode},
			Prescription: entities.Prescription{
				PatientNationalCode: p.Edges.Users.NationalCode,
				DoctorID:            doctor.ID,
				Drugs:               p.DrugsCommaSeperated,
			},
		})
	}
	return ps, nil
}
