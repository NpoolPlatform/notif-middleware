// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent/contact"
	"github.com/google/uuid"
)

// Contact is the model entity for the Contact schema.
type Contact struct {
	config `json:"-"`
	// ID of the ent.
	ID uint32 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// EntID holds the value of the "ent_id" field.
	EntID uuid.UUID `json:"ent_id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// UsedFor holds the value of the "used_for" field.
	UsedFor string `json:"used_for,omitempty"`
	// Sender holds the value of the "sender" field.
	Sender string `json:"sender,omitempty"`
	// Account holds the value of the "account" field.
	Account string `json:"account,omitempty"`
	// AccountType holds the value of the "account_type" field.
	AccountType string `json:"account_type,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Contact) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case contact.FieldID, contact.FieldCreatedAt, contact.FieldUpdatedAt, contact.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case contact.FieldUsedFor, contact.FieldSender, contact.FieldAccount, contact.FieldAccountType:
			values[i] = new(sql.NullString)
		case contact.FieldEntID, contact.FieldAppID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Contact", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Contact fields.
func (c *Contact) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case contact.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = uint32(value.Int64)
		case contact.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = uint32(value.Int64)
			}
		case contact.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = uint32(value.Int64)
			}
		case contact.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				c.DeletedAt = uint32(value.Int64)
			}
		case contact.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				c.EntID = *value
			}
		case contact.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				c.AppID = *value
			}
		case contact.FieldUsedFor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field used_for", values[i])
			} else if value.Valid {
				c.UsedFor = value.String
			}
		case contact.FieldSender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sender", values[i])
			} else if value.Valid {
				c.Sender = value.String
			}
		case contact.FieldAccount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field account", values[i])
			} else if value.Valid {
				c.Account = value.String
			}
		case contact.FieldAccountType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field account_type", values[i])
			} else if value.Valid {
				c.AccountType = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Contact.
// Note that you need to call Contact.Unwrap() before calling this method if this Contact
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Contact) Update() *ContactUpdateOne {
	return (&ContactClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Contact entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Contact) Unwrap() *Contact {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Contact is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Contact) String() string {
	var builder strings.Builder
	builder.WriteString("Contact(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", c.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", c.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", c.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", c.EntID))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", c.AppID))
	builder.WriteString(", ")
	builder.WriteString("used_for=")
	builder.WriteString(c.UsedFor)
	builder.WriteString(", ")
	builder.WriteString("sender=")
	builder.WriteString(c.Sender)
	builder.WriteString(", ")
	builder.WriteString("account=")
	builder.WriteString(c.Account)
	builder.WriteString(", ")
	builder.WriteString("account_type=")
	builder.WriteString(c.AccountType)
	builder.WriteByte(')')
	return builder.String()
}

// Contacts is a parsable slice of Contact.
type Contacts []*Contact

func (c Contacts) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
