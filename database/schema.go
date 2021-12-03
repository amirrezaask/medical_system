package database

import (
	"entgo.io/ent"
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
