// Code generated by entc, DO NOT EDIT.

package prescription

const (
	// Label holds the string label denoting the prescription type in the database.
	Label = "prescription"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDoctorID holds the string denoting the doctor_id field in the database.
	FieldDoctorID = "doctor_id"
	// FieldPatientNationalCode holds the string denoting the patient_national_code field in the database.
	FieldPatientNationalCode = "patient_national_code"
	// FieldDrugsCommaSeperated holds the string denoting the drugs_comma_seperated field in the database.
	FieldDrugsCommaSeperated = "drugs_comma_seperated"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// Table holds the table name of the prescription in the database.
	Table = "prescriptions"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "prescriptions"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "user_prescriptions"
)

// Columns holds all SQL columns for prescription fields.
var Columns = []string{
	FieldID,
	FieldDoctorID,
	FieldPatientNationalCode,
	FieldDrugsCommaSeperated,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "prescriptions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_prescriptions",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
