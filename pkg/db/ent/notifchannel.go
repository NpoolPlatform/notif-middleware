// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent/notifchannel"
	"github.com/google/uuid"
)

// NotifChannel is the model entity for the NotifChannel schema.
type NotifChannel struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// EventType holds the value of the "event_type" field.
	EventType string `json:"event_type,omitempty"`
	// Channel holds the value of the "channel" field.
	Channel string `json:"channel,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*NotifChannel) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case notifchannel.FieldCreatedAt, notifchannel.FieldUpdatedAt, notifchannel.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case notifchannel.FieldEventType, notifchannel.FieldChannel:
			values[i] = new(sql.NullString)
		case notifchannel.FieldID, notifchannel.FieldAppID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type NotifChannel", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the NotifChannel fields.
func (nc *NotifChannel) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case notifchannel.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				nc.ID = *value
			}
		case notifchannel.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				nc.CreatedAt = uint32(value.Int64)
			}
		case notifchannel.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				nc.UpdatedAt = uint32(value.Int64)
			}
		case notifchannel.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				nc.DeletedAt = uint32(value.Int64)
			}
		case notifchannel.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				nc.AppID = *value
			}
		case notifchannel.FieldEventType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field event_type", values[i])
			} else if value.Valid {
				nc.EventType = value.String
			}
		case notifchannel.FieldChannel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field channel", values[i])
			} else if value.Valid {
				nc.Channel = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this NotifChannel.
// Note that you need to call NotifChannel.Unwrap() before calling this method if this NotifChannel
// was returned from a transaction, and the transaction was committed or rolled back.
func (nc *NotifChannel) Update() *NotifChannelUpdateOne {
	return (&NotifChannelClient{config: nc.config}).UpdateOne(nc)
}

// Unwrap unwraps the NotifChannel entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (nc *NotifChannel) Unwrap() *NotifChannel {
	_tx, ok := nc.config.driver.(*txDriver)
	if !ok {
		panic("ent: NotifChannel is not a transactional entity")
	}
	nc.config.driver = _tx.drv
	return nc
}

// String implements the fmt.Stringer.
func (nc *NotifChannel) String() string {
	var builder strings.Builder
	builder.WriteString("NotifChannel(")
	builder.WriteString(fmt.Sprintf("id=%v, ", nc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", nc.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", nc.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", nc.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", nc.AppID))
	builder.WriteString(", ")
	builder.WriteString("event_type=")
	builder.WriteString(nc.EventType)
	builder.WriteString(", ")
	builder.WriteString("channel=")
	builder.WriteString(nc.Channel)
	builder.WriteByte(')')
	return builder.String()
}

// NotifChannels is a parsable slice of NotifChannel.
type NotifChannels []*NotifChannel

func (nc NotifChannels) config(cfg config) {
	for _i := range nc {
		nc[_i].config = cfg
	}
}
