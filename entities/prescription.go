package entities

type Prescription struct {
	PatientNationalCode string `json:"patient_national_code"`
	DoctorID            int    `json:"doctor_id"`
	Drugs               string `json:"drugs"`
}

type PrescriptionForDoctor struct {
	Patient *User
	Prescription
}

type PrescriptionForPatient struct {
	Doctor *User
	Prescription
}
