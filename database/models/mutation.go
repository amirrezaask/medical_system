// Code generated by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"medical_system/database/models/predicate"
	"medical_system/database/models/user"
	"sync"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeUser = "User"
)

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	user_type     *user.UserType
	national_code *string
	password_hash *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("models: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the "name" field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// SetUserType sets the "user_type" field.
func (m *UserMutation) SetUserType(ut user.UserType) {
	m.user_type = &ut
}

// UserType returns the value of the "user_type" field in the mutation.
func (m *UserMutation) UserType() (r user.UserType, exists bool) {
	v := m.user_type
	if v == nil {
		return
	}
	return *v, true
}

// OldUserType returns the old "user_type" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUserType(ctx context.Context) (v user.UserType, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUserType is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUserType requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserType: %w", err)
	}
	return oldValue.UserType, nil
}

// ResetUserType resets all changes to the "user_type" field.
func (m *UserMutation) ResetUserType() {
	m.user_type = nil
}

// SetNationalCode sets the "national_code" field.
func (m *UserMutation) SetNationalCode(s string) {
	m.national_code = &s
}

// NationalCode returns the value of the "national_code" field in the mutation.
func (m *UserMutation) NationalCode() (r string, exists bool) {
	v := m.national_code
	if v == nil {
		return
	}
	return *v, true
}

// OldNationalCode returns the old "national_code" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldNationalCode(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldNationalCode is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldNationalCode requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNationalCode: %w", err)
	}
	return oldValue.NationalCode, nil
}

// ResetNationalCode resets all changes to the "national_code" field.
func (m *UserMutation) ResetNationalCode() {
	m.national_code = nil
}

// SetPasswordHash sets the "password_hash" field.
func (m *UserMutation) SetPasswordHash(s string) {
	m.password_hash = &s
}

// PasswordHash returns the value of the "password_hash" field in the mutation.
func (m *UserMutation) PasswordHash() (r string, exists bool) {
	v := m.password_hash
	if v == nil {
		return
	}
	return *v, true
}

// OldPasswordHash returns the old "password_hash" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldPasswordHash(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldPasswordHash is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldPasswordHash requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPasswordHash: %w", err)
	}
	return oldValue.PasswordHash, nil
}

// ResetPasswordHash resets all changes to the "password_hash" field.
func (m *UserMutation) ResetPasswordHash() {
	m.password_hash = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	if m.user_type != nil {
		fields = append(fields, user.FieldUserType)
	}
	if m.national_code != nil {
		fields = append(fields, user.FieldNationalCode)
	}
	if m.password_hash != nil {
		fields = append(fields, user.FieldPasswordHash)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldName:
		return m.Name()
	case user.FieldUserType:
		return m.UserType()
	case user.FieldNationalCode:
		return m.NationalCode()
	case user.FieldPasswordHash:
		return m.PasswordHash()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldName:
		return m.OldName(ctx)
	case user.FieldUserType:
		return m.OldUserType(ctx)
	case user.FieldNationalCode:
		return m.OldNationalCode(ctx)
	case user.FieldPasswordHash:
		return m.OldPasswordHash(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case user.FieldUserType:
		v, ok := value.(user.UserType)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserType(v)
		return nil
	case user.FieldNationalCode:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNationalCode(v)
		return nil
	case user.FieldPasswordHash:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPasswordHash(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldName:
		m.ResetName()
		return nil
	case user.FieldUserType:
		m.ResetUserType()
		return nil
	case user.FieldNationalCode:
		m.ResetNationalCode()
		return nil
	case user.FieldPasswordHash:
		m.ResetPasswordHash()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown User edge %s", name)
}
