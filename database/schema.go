package database

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Enum("user_type").Values("patient", "doctor"),
		field.String("national_code"),
		field.String("password_hash"),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("prescriptions", Prescription.Type),
	}

}

type Admin struct {
	ent.Schema
}

func (Admin) Fields() []ent.Field {
	return []ent.Field{
		field.String("username"),
		field.String("password_hash"),
	}
}

type Prescription struct {
	ent.Schema
}

func (Prescription) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("doctor_id"),
		field.String("patient_national_code"),
		field.String("drugs_comma_seperated"),
		field.Time("created_at").
			Default(time.Now),
	}
}
func (Prescription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).Ref("prescriptions").Unique(),
	}
}
