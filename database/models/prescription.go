// Code generated by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"medical_system/database/models/prescription"
	"medical_system/database/models/user"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Prescription is the model entity for the Prescription schema.
type Prescription struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// DoctorID holds the value of the "doctor_id" field.
	DoctorID int64 `json:"doctor_id,omitempty"`
	// PatientNationalCode holds the value of the "patient_national_code" field.
	PatientNationalCode string `json:"patient_national_code,omitempty"`
	// DrugsCommaSeperated holds the value of the "drugs_comma_seperated" field.
	DrugsCommaSeperated string `json:"drugs_comma_seperated,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PrescriptionQuery when eager-loading is set.
	Edges              PrescriptionEdges `json:"edges"`
	user_prescriptions *int
}

// PrescriptionEdges holds the relations/edges for other nodes in the graph.
type PrescriptionEdges struct {
	// Users holds the value of the users edge.
	Users *User `json:"users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PrescriptionEdges) UsersOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Users == nil {
			// The edge users was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Prescription) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case prescription.FieldID, prescription.FieldDoctorID:
			values[i] = new(sql.NullInt64)
		case prescription.FieldPatientNationalCode, prescription.FieldDrugsCommaSeperated:
			values[i] = new(sql.NullString)
		case prescription.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case prescription.ForeignKeys[0]: // user_prescriptions
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Prescription", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Prescription fields.
func (pr *Prescription) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case prescription.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case prescription.FieldDoctorID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field doctor_id", values[i])
			} else if value.Valid {
				pr.DoctorID = value.Int64
			}
		case prescription.FieldPatientNationalCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field patient_national_code", values[i])
			} else if value.Valid {
				pr.PatientNationalCode = value.String
			}
		case prescription.FieldDrugsCommaSeperated:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field drugs_comma_seperated", values[i])
			} else if value.Valid {
				pr.DrugsCommaSeperated = value.String
			}
		case prescription.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		case prescription.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_prescriptions", value)
			} else if value.Valid {
				pr.user_prescriptions = new(int)
				*pr.user_prescriptions = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUsers queries the "users" edge of the Prescription entity.
func (pr *Prescription) QueryUsers() *UserQuery {
	return (&PrescriptionClient{config: pr.config}).QueryUsers(pr)
}

// Update returns a builder for updating this Prescription.
// Note that you need to call Prescription.Unwrap() before calling this method if this Prescription
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Prescription) Update() *PrescriptionUpdateOne {
	return (&PrescriptionClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the Prescription entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Prescription) Unwrap() *Prescription {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("models: Prescription is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Prescription) String() string {
	var builder strings.Builder
	builder.WriteString("Prescription(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", doctor_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.DoctorID))
	builder.WriteString(", patient_national_code=")
	builder.WriteString(pr.PatientNationalCode)
	builder.WriteString(", drugs_comma_seperated=")
	builder.WriteString(pr.DrugsCommaSeperated)
	builder.WriteString(", created_at=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Prescriptions is a parsable slice of Prescription.
type Prescriptions []*Prescription

func (pr Prescriptions) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
