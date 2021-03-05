// Code generated by entc, DO NOT EDIT.

package ent

import (
	"abc/ent/test"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Test is the model entity for the Test schema.
type Test struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Note holds the value of the "note" field.
	Note string `json:"note,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Test) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case test.FieldID:
			values[i] = &sql.NullInt64{}
		case test.FieldNote:
			values[i] = &sql.NullString{}
		case test.FieldCreatedAt, test.FieldUpdatedAt:
			values[i] = &sql.NullTime{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Test", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Test fields.
func (t *Test) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case test.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = uint64(value.Int64)
		case test.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case test.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		case test.FieldNote:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field note", values[i])
			} else if value.Valid {
				t.Note = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Test.
// Note that you need to call Test.Unwrap() before calling this method if this Test
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Test) Update() *TestUpdateOne {
	return (&TestClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Test entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Test) Unwrap() *Test {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Test is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Test) String() string {
	var builder strings.Builder
	builder.WriteString("Test(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", note=")
	builder.WriteString(t.Note)
	builder.WriteByte(')')
	return builder.String()
}

// Tests is a parsable slice of Test.
type Tests []*Test

func (t Tests) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
