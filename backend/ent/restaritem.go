// Code generated by entc, DO NOT EDIT.

package ent

import (
	"backend/ent/restaritem"
	"backend/pkg/photo"
	"backend/pkg/work"
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Restaritem is the model entity for the Restaritem schema.
type Restaritem struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// OnecGUID holds the value of the "onecGUID" field.
	OnecGUID string `json:"onecGUID,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Sku holds the value of the "sku" field.
	Sku string `json:"sku,omitempty"`
	// ItemGUID holds the value of the "itemGUID" field.
	ItemGUID string `json:"itemGUID,omitempty"`
	// CharGUID holds the value of the "charGUID" field.
	CharGUID string `json:"charGUID,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Inspector holds the value of the "inspector" field.
	Inspector string `json:"inspector,omitempty"`
	// Inspection holds the value of the "inspection" field.
	Inspection []string `json:"inspection,omitempty"`
	// Photos holds the value of the "photos" field.
	Photos []photo.Photo `json:"photos,omitempty"`
	// Works holds the value of the "works" field.
	Works []work.Work `json:"works,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Restaritem) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case restaritem.FieldInspection, restaritem.FieldPhotos, restaritem.FieldWorks:
			values[i] = new([]byte)
		case restaritem.FieldID:
			values[i] = new(sql.NullInt64)
		case restaritem.FieldOnecGUID, restaritem.FieldName, restaritem.FieldSku, restaritem.FieldItemGUID, restaritem.FieldCharGUID, restaritem.FieldDescription, restaritem.FieldInspector:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Restaritem", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Restaritem fields.
func (r *Restaritem) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case restaritem.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case restaritem.FieldOnecGUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field onecGUID", values[i])
			} else if value.Valid {
				r.OnecGUID = value.String
			}
		case restaritem.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				r.Name = value.String
			}
		case restaritem.FieldSku:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sku", values[i])
			} else if value.Valid {
				r.Sku = value.String
			}
		case restaritem.FieldItemGUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field itemGUID", values[i])
			} else if value.Valid {
				r.ItemGUID = value.String
			}
		case restaritem.FieldCharGUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field charGUID", values[i])
			} else if value.Valid {
				r.CharGUID = value.String
			}
		case restaritem.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				r.Description = value.String
			}
		case restaritem.FieldInspector:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field inspector", values[i])
			} else if value.Valid {
				r.Inspector = value.String
			}
		case restaritem.FieldInspection:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field inspection", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &r.Inspection); err != nil {
					return fmt.Errorf("unmarshal field inspection: %w", err)
				}
			}
		case restaritem.FieldPhotos:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field photos", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &r.Photos); err != nil {
					return fmt.Errorf("unmarshal field photos: %w", err)
				}
			}
		case restaritem.FieldWorks:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field works", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &r.Works); err != nil {
					return fmt.Errorf("unmarshal field works: %w", err)
				}
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Restaritem.
// Note that you need to call Restaritem.Unwrap() before calling this method if this Restaritem
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Restaritem) Update() *RestaritemUpdateOne {
	return (&RestaritemClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Restaritem entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Restaritem) Unwrap() *Restaritem {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Restaritem is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Restaritem) String() string {
	var builder strings.Builder
	builder.WriteString("Restaritem(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", onecGUID=")
	builder.WriteString(r.OnecGUID)
	builder.WriteString(", name=")
	builder.WriteString(r.Name)
	builder.WriteString(", sku=")
	builder.WriteString(r.Sku)
	builder.WriteString(", itemGUID=")
	builder.WriteString(r.ItemGUID)
	builder.WriteString(", charGUID=")
	builder.WriteString(r.CharGUID)
	builder.WriteString(", description=")
	builder.WriteString(r.Description)
	builder.WriteString(", inspector=")
	builder.WriteString(r.Inspector)
	builder.WriteString(", inspection=")
	builder.WriteString(fmt.Sprintf("%v", r.Inspection))
	builder.WriteString(", photos=")
	builder.WriteString(fmt.Sprintf("%v", r.Photos))
	builder.WriteString(", works=")
	builder.WriteString(fmt.Sprintf("%v", r.Works))
	builder.WriteByte(')')
	return builder.String()
}

// Restaritems is a parsable slice of Restaritem.
type Restaritems []*Restaritem

func (r Restaritems) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
