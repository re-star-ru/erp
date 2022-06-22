// Code generated by entc, DO NOT EDIT.

package ent

import (
	"backend/ent/predicate"
	"backend/ent/restaritem"
	"context"
	"errors"
	"fmt"
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
	TypeRestaritem = "Restaritem"
)

// RestaritemMutation represents an operation that mutates the Restaritem nodes in the graph.
type RestaritemMutation struct {
	config
	op            Op
	typ           string
	id            *int
	onecGUID      *string
	name          *string
	sku           *string
	itemGUID      *string
	charGUID      *string
	description   *string
	inspector     *string
	inspection    *[]string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Restaritem, error)
	predicates    []predicate.Restaritem
}

var _ ent.Mutation = (*RestaritemMutation)(nil)

// restaritemOption allows management of the mutation configuration using functional options.
type restaritemOption func(*RestaritemMutation)

// newRestaritemMutation creates new mutation for the Restaritem entity.
func newRestaritemMutation(c config, op Op, opts ...restaritemOption) *RestaritemMutation {
	m := &RestaritemMutation{
		config:        c,
		op:            op,
		typ:           TypeRestaritem,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withRestaritemID sets the ID field of the mutation.
func withRestaritemID(id int) restaritemOption {
	return func(m *RestaritemMutation) {
		var (
			err   error
			once  sync.Once
			value *Restaritem
		)
		m.oldValue = func(ctx context.Context) (*Restaritem, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Restaritem.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withRestaritem sets the old Restaritem of the mutation.
func withRestaritem(node *Restaritem) restaritemOption {
	return func(m *RestaritemMutation) {
		m.oldValue = func(context.Context) (*Restaritem, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m RestaritemMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m RestaritemMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *RestaritemMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *RestaritemMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Restaritem.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetOnecGUID sets the "onecGUID" field.
func (m *RestaritemMutation) SetOnecGUID(s string) {
	m.onecGUID = &s
}

// OnecGUID returns the value of the "onecGUID" field in the mutation.
func (m *RestaritemMutation) OnecGUID() (r string, exists bool) {
	v := m.onecGUID
	if v == nil {
		return
	}
	return *v, true
}

// OldOnecGUID returns the old "onecGUID" field's value of the Restaritem entity.
// If the Restaritem object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RestaritemMutation) OldOnecGUID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldOnecGUID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldOnecGUID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOnecGUID: %w", err)
	}
	return oldValue.OnecGUID, nil
}

// ResetOnecGUID resets all changes to the "onecGUID" field.
func (m *RestaritemMutation) ResetOnecGUID() {
	m.onecGUID = nil
}

// SetName sets the "name" field.
func (m *RestaritemMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *RestaritemMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Restaritem entity.
// If the Restaritem object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RestaritemMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ClearName clears the value of the "name" field.
func (m *RestaritemMutation) ClearName() {
	m.name = nil
	m.clearedFields[restaritem.FieldName] = struct{}{}
}

// NameCleared returns if the "name" field was cleared in this mutation.
func (m *RestaritemMutation) NameCleared() bool {
	_, ok := m.clearedFields[restaritem.FieldName]
	return ok
}

// ResetName resets all changes to the "name" field.
func (m *RestaritemMutation) ResetName() {
	m.name = nil
	delete(m.clearedFields, restaritem.FieldName)
}

// SetSku sets the "sku" field.
func (m *RestaritemMutation) SetSku(s string) {
	m.sku = &s
}

// Sku returns the value of the "sku" field in the mutation.
func (m *RestaritemMutation) Sku() (r string, exists bool) {
	v := m.sku
	if v == nil {
		return
	}
	return *v, true
}

// OldSku returns the old "sku" field's value of the Restaritem entity.
// If the Restaritem object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RestaritemMutation) OldSku(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSku is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSku requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSku: %w", err)
	}
	return oldValue.Sku, nil
}

// ClearSku clears the value of the "sku" field.
func (m *RestaritemMutation) ClearSku() {
	m.sku = nil
	m.clearedFields[restaritem.FieldSku] = struct{}{}
}

// SkuCleared returns if the "sku" field was cleared in this mutation.
func (m *RestaritemMutation) SkuCleared() bool {
	_, ok := m.clearedFields[restaritem.FieldSku]
	return ok
}

// ResetSku resets all changes to the "sku" field.
func (m *RestaritemMutation) ResetSku() {
	m.sku = nil
	delete(m.clearedFields, restaritem.FieldSku)
}

// SetItemGUID sets the "itemGUID" field.
func (m *RestaritemMutation) SetItemGUID(s string) {
	m.itemGUID = &s
}

// ItemGUID returns the value of the "itemGUID" field in the mutation.
func (m *RestaritemMutation) ItemGUID() (r string, exists bool) {
	v := m.itemGUID
	if v == nil {
		return
	}
	return *v, true
}

// OldItemGUID returns the old "itemGUID" field's value of the Restaritem entity.
// If the Restaritem object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RestaritemMutation) OldItemGUID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldItemGUID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldItemGUID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldItemGUID: %w", err)
	}
	return oldValue.ItemGUID, nil
}

// ClearItemGUID clears the value of the "itemGUID" field.
func (m *RestaritemMutation) ClearItemGUID() {
	m.itemGUID = nil
	m.clearedFields[restaritem.FieldItemGUID] = struct{}{}
}

// ItemGUIDCleared returns if the "itemGUID" field was cleared in this mutation.
func (m *RestaritemMutation) ItemGUIDCleared() bool {
	_, ok := m.clearedFields[restaritem.FieldItemGUID]
	return ok
}

// ResetItemGUID resets all changes to the "itemGUID" field.
func (m *RestaritemMutation) ResetItemGUID() {
	m.itemGUID = nil
	delete(m.clearedFields, restaritem.FieldItemGUID)
}

// SetCharGUID sets the "charGUID" field.
func (m *RestaritemMutation) SetCharGUID(s string) {
	m.charGUID = &s
}

// CharGUID returns the value of the "charGUID" field in the mutation.
func (m *RestaritemMutation) CharGUID() (r string, exists bool) {
	v := m.charGUID
	if v == nil {
		return
	}
	return *v, true
}

// OldCharGUID returns the old "charGUID" field's value of the Restaritem entity.
// If the Restaritem object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RestaritemMutation) OldCharGUID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCharGUID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCharGUID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCharGUID: %w", err)
	}
	return oldValue.CharGUID, nil
}

// ClearCharGUID clears the value of the "charGUID" field.
func (m *RestaritemMutation) ClearCharGUID() {
	m.charGUID = nil
	m.clearedFields[restaritem.FieldCharGUID] = struct{}{}
}

// CharGUIDCleared returns if the "charGUID" field was cleared in this mutation.
func (m *RestaritemMutation) CharGUIDCleared() bool {
	_, ok := m.clearedFields[restaritem.FieldCharGUID]
	return ok
}

// ResetCharGUID resets all changes to the "charGUID" field.
func (m *RestaritemMutation) ResetCharGUID() {
	m.charGUID = nil
	delete(m.clearedFields, restaritem.FieldCharGUID)
}

// SetDescription sets the "description" field.
func (m *RestaritemMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the value of the "description" field in the mutation.
func (m *RestaritemMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old "description" field's value of the Restaritem entity.
// If the Restaritem object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RestaritemMutation) OldDescription(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDescription is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDescription requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescription: %w", err)
	}
	return oldValue.Description, nil
}

// ClearDescription clears the value of the "description" field.
func (m *RestaritemMutation) ClearDescription() {
	m.description = nil
	m.clearedFields[restaritem.FieldDescription] = struct{}{}
}

// DescriptionCleared returns if the "description" field was cleared in this mutation.
func (m *RestaritemMutation) DescriptionCleared() bool {
	_, ok := m.clearedFields[restaritem.FieldDescription]
	return ok
}

// ResetDescription resets all changes to the "description" field.
func (m *RestaritemMutation) ResetDescription() {
	m.description = nil
	delete(m.clearedFields, restaritem.FieldDescription)
}

// SetInspector sets the "inspector" field.
func (m *RestaritemMutation) SetInspector(s string) {
	m.inspector = &s
}

// Inspector returns the value of the "inspector" field in the mutation.
func (m *RestaritemMutation) Inspector() (r string, exists bool) {
	v := m.inspector
	if v == nil {
		return
	}
	return *v, true
}

// OldInspector returns the old "inspector" field's value of the Restaritem entity.
// If the Restaritem object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RestaritemMutation) OldInspector(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldInspector is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldInspector requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldInspector: %w", err)
	}
	return oldValue.Inspector, nil
}

// ClearInspector clears the value of the "inspector" field.
func (m *RestaritemMutation) ClearInspector() {
	m.inspector = nil
	m.clearedFields[restaritem.FieldInspector] = struct{}{}
}

// InspectorCleared returns if the "inspector" field was cleared in this mutation.
func (m *RestaritemMutation) InspectorCleared() bool {
	_, ok := m.clearedFields[restaritem.FieldInspector]
	return ok
}

// ResetInspector resets all changes to the "inspector" field.
func (m *RestaritemMutation) ResetInspector() {
	m.inspector = nil
	delete(m.clearedFields, restaritem.FieldInspector)
}

// SetInspection sets the "inspection" field.
func (m *RestaritemMutation) SetInspection(s []string) {
	m.inspection = &s
}

// Inspection returns the value of the "inspection" field in the mutation.
func (m *RestaritemMutation) Inspection() (r []string, exists bool) {
	v := m.inspection
	if v == nil {
		return
	}
	return *v, true
}

// OldInspection returns the old "inspection" field's value of the Restaritem entity.
// If the Restaritem object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RestaritemMutation) OldInspection(ctx context.Context) (v []string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldInspection is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldInspection requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldInspection: %w", err)
	}
	return oldValue.Inspection, nil
}

// ClearInspection clears the value of the "inspection" field.
func (m *RestaritemMutation) ClearInspection() {
	m.inspection = nil
	m.clearedFields[restaritem.FieldInspection] = struct{}{}
}

// InspectionCleared returns if the "inspection" field was cleared in this mutation.
func (m *RestaritemMutation) InspectionCleared() bool {
	_, ok := m.clearedFields[restaritem.FieldInspection]
	return ok
}

// ResetInspection resets all changes to the "inspection" field.
func (m *RestaritemMutation) ResetInspection() {
	m.inspection = nil
	delete(m.clearedFields, restaritem.FieldInspection)
}

// Where appends a list predicates to the RestaritemMutation builder.
func (m *RestaritemMutation) Where(ps ...predicate.Restaritem) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *RestaritemMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Restaritem).
func (m *RestaritemMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *RestaritemMutation) Fields() []string {
	fields := make([]string, 0, 8)
	if m.onecGUID != nil {
		fields = append(fields, restaritem.FieldOnecGUID)
	}
	if m.name != nil {
		fields = append(fields, restaritem.FieldName)
	}
	if m.sku != nil {
		fields = append(fields, restaritem.FieldSku)
	}
	if m.itemGUID != nil {
		fields = append(fields, restaritem.FieldItemGUID)
	}
	if m.charGUID != nil {
		fields = append(fields, restaritem.FieldCharGUID)
	}
	if m.description != nil {
		fields = append(fields, restaritem.FieldDescription)
	}
	if m.inspector != nil {
		fields = append(fields, restaritem.FieldInspector)
	}
	if m.inspection != nil {
		fields = append(fields, restaritem.FieldInspection)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *RestaritemMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case restaritem.FieldOnecGUID:
		return m.OnecGUID()
	case restaritem.FieldName:
		return m.Name()
	case restaritem.FieldSku:
		return m.Sku()
	case restaritem.FieldItemGUID:
		return m.ItemGUID()
	case restaritem.FieldCharGUID:
		return m.CharGUID()
	case restaritem.FieldDescription:
		return m.Description()
	case restaritem.FieldInspector:
		return m.Inspector()
	case restaritem.FieldInspection:
		return m.Inspection()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *RestaritemMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case restaritem.FieldOnecGUID:
		return m.OldOnecGUID(ctx)
	case restaritem.FieldName:
		return m.OldName(ctx)
	case restaritem.FieldSku:
		return m.OldSku(ctx)
	case restaritem.FieldItemGUID:
		return m.OldItemGUID(ctx)
	case restaritem.FieldCharGUID:
		return m.OldCharGUID(ctx)
	case restaritem.FieldDescription:
		return m.OldDescription(ctx)
	case restaritem.FieldInspector:
		return m.OldInspector(ctx)
	case restaritem.FieldInspection:
		return m.OldInspection(ctx)
	}
	return nil, fmt.Errorf("unknown Restaritem field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RestaritemMutation) SetField(name string, value ent.Value) error {
	switch name {
	case restaritem.FieldOnecGUID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOnecGUID(v)
		return nil
	case restaritem.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case restaritem.FieldSku:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSku(v)
		return nil
	case restaritem.FieldItemGUID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetItemGUID(v)
		return nil
	case restaritem.FieldCharGUID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCharGUID(v)
		return nil
	case restaritem.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	case restaritem.FieldInspector:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetInspector(v)
		return nil
	case restaritem.FieldInspection:
		v, ok := value.([]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetInspection(v)
		return nil
	}
	return fmt.Errorf("unknown Restaritem field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *RestaritemMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *RestaritemMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RestaritemMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Restaritem numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *RestaritemMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(restaritem.FieldName) {
		fields = append(fields, restaritem.FieldName)
	}
	if m.FieldCleared(restaritem.FieldSku) {
		fields = append(fields, restaritem.FieldSku)
	}
	if m.FieldCleared(restaritem.FieldItemGUID) {
		fields = append(fields, restaritem.FieldItemGUID)
	}
	if m.FieldCleared(restaritem.FieldCharGUID) {
		fields = append(fields, restaritem.FieldCharGUID)
	}
	if m.FieldCleared(restaritem.FieldDescription) {
		fields = append(fields, restaritem.FieldDescription)
	}
	if m.FieldCleared(restaritem.FieldInspector) {
		fields = append(fields, restaritem.FieldInspector)
	}
	if m.FieldCleared(restaritem.FieldInspection) {
		fields = append(fields, restaritem.FieldInspection)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *RestaritemMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *RestaritemMutation) ClearField(name string) error {
	switch name {
	case restaritem.FieldName:
		m.ClearName()
		return nil
	case restaritem.FieldSku:
		m.ClearSku()
		return nil
	case restaritem.FieldItemGUID:
		m.ClearItemGUID()
		return nil
	case restaritem.FieldCharGUID:
		m.ClearCharGUID()
		return nil
	case restaritem.FieldDescription:
		m.ClearDescription()
		return nil
	case restaritem.FieldInspector:
		m.ClearInspector()
		return nil
	case restaritem.FieldInspection:
		m.ClearInspection()
		return nil
	}
	return fmt.Errorf("unknown Restaritem nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *RestaritemMutation) ResetField(name string) error {
	switch name {
	case restaritem.FieldOnecGUID:
		m.ResetOnecGUID()
		return nil
	case restaritem.FieldName:
		m.ResetName()
		return nil
	case restaritem.FieldSku:
		m.ResetSku()
		return nil
	case restaritem.FieldItemGUID:
		m.ResetItemGUID()
		return nil
	case restaritem.FieldCharGUID:
		m.ResetCharGUID()
		return nil
	case restaritem.FieldDescription:
		m.ResetDescription()
		return nil
	case restaritem.FieldInspector:
		m.ResetInspector()
		return nil
	case restaritem.FieldInspection:
		m.ResetInspection()
		return nil
	}
	return fmt.Errorf("unknown Restaritem field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *RestaritemMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *RestaritemMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *RestaritemMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *RestaritemMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *RestaritemMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *RestaritemMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *RestaritemMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Restaritem unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *RestaritemMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Restaritem edge %s", name)
}
