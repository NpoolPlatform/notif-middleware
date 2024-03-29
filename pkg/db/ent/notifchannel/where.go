// Code generated by ent, DO NOT EDIT.

package notifchannel

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// EventType applies equality check predicate on the "event_type" field. It's identical to EventTypeEQ.
func EventType(v string) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventType), v))
	})
}

// Channel applies equality check predicate on the "channel" field. It's identical to ChannelEQ.
func Channel(v string) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChannel), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.NotifChannel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.NotifChannel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.NotifChannel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.NotifChannel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.NotifChannel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.NotifChannel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntID), v))
	})
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.NotifChannel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEntID), v...))
	})
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.NotifChannel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEntID), v...))
	})
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntID), v))
	})
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntID), v))
	})
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntID), v))
	})
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntID), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.NotifChannel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.NotifChannel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// AppIDIsNil applies the IsNil predicate on the "app_id" field.
func AppIDIsNil() predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAppID)))
	})
}

// AppIDNotNil applies the NotNil predicate on the "app_id" field.
func AppIDNotNil() predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAppID)))
	})
}

// EventTypeEQ applies the EQ predicate on the "event_type" field.
func EventTypeEQ(v string) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventType), v))
	})
}

// EventTypeNEQ applies the NEQ predicate on the "event_type" field.
func EventTypeNEQ(v string) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEventType), v))
	})
}

// EventTypeIn applies the In predicate on the "event_type" field.
func EventTypeIn(vs ...string) predicate.NotifChannel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEventType), v...))
	})
}

// EventTypeNotIn applies the NotIn predicate on the "event_type" field.
func EventTypeNotIn(vs ...string) predicate.NotifChannel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEventType), v...))
	})
}

// EventTypeGT applies the GT predicate on the "event_type" field.
func EventTypeGT(v string) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEventType), v))
	})
}

// EventTypeGTE applies the GTE predicate on the "event_type" field.
func EventTypeGTE(v string) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEventType), v))
	})
}

// EventTypeLT applies the LT predicate on the "event_type" field.
func EventTypeLT(v string) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEventType), v))
	})
}

// EventTypeLTE applies the LTE predicate on the "event_type" field.
func EventTypeLTE(v string) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEventType), v))
	})
}

// EventTypeContains applies the Contains predicate on the "event_type" field.
func EventTypeContains(v string) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEventType), v))
	})
}

// EventTypeHasPrefix applies the HasPrefix predicate on the "event_type" field.
func EventTypeHasPrefix(v string) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEventType), v))
	})
}

// EventTypeHasSuffix applies the HasSuffix predicate on the "event_type" field.
func EventTypeHasSuffix(v string) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEventType), v))
	})
}

// EventTypeIsNil applies the IsNil predicate on the "event_type" field.
func EventTypeIsNil() predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEventType)))
	})
}

// EventTypeNotNil applies the NotNil predicate on the "event_type" field.
func EventTypeNotNil() predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEventType)))
	})
}

// EventTypeEqualFold applies the EqualFold predicate on the "event_type" field.
func EventTypeEqualFold(v string) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEventType), v))
	})
}

// EventTypeContainsFold applies the ContainsFold predicate on the "event_type" field.
func EventTypeContainsFold(v string) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEventType), v))
	})
}

// ChannelEQ applies the EQ predicate on the "channel" field.
func ChannelEQ(v string) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChannel), v))
	})
}

// ChannelNEQ applies the NEQ predicate on the "channel" field.
func ChannelNEQ(v string) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldChannel), v))
	})
}

// ChannelIn applies the In predicate on the "channel" field.
func ChannelIn(vs ...string) predicate.NotifChannel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldChannel), v...))
	})
}

// ChannelNotIn applies the NotIn predicate on the "channel" field.
func ChannelNotIn(vs ...string) predicate.NotifChannel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldChannel), v...))
	})
}

// ChannelGT applies the GT predicate on the "channel" field.
func ChannelGT(v string) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldChannel), v))
	})
}

// ChannelGTE applies the GTE predicate on the "channel" field.
func ChannelGTE(v string) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldChannel), v))
	})
}

// ChannelLT applies the LT predicate on the "channel" field.
func ChannelLT(v string) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldChannel), v))
	})
}

// ChannelLTE applies the LTE predicate on the "channel" field.
func ChannelLTE(v string) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldChannel), v))
	})
}

// ChannelContains applies the Contains predicate on the "channel" field.
func ChannelContains(v string) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldChannel), v))
	})
}

// ChannelHasPrefix applies the HasPrefix predicate on the "channel" field.
func ChannelHasPrefix(v string) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldChannel), v))
	})
}

// ChannelHasSuffix applies the HasSuffix predicate on the "channel" field.
func ChannelHasSuffix(v string) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldChannel), v))
	})
}

// ChannelIsNil applies the IsNil predicate on the "channel" field.
func ChannelIsNil() predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldChannel)))
	})
}

// ChannelNotNil applies the NotNil predicate on the "channel" field.
func ChannelNotNil() predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldChannel)))
	})
}

// ChannelEqualFold applies the EqualFold predicate on the "channel" field.
func ChannelEqualFold(v string) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldChannel), v))
	})
}

// ChannelContainsFold applies the ContainsFold predicate on the "channel" field.
func ChannelContainsFold(v string) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldChannel), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.NotifChannel) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.NotifChannel) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.NotifChannel) predicate.NotifChannel {
	return predicate.NotifChannel(func(s *sql.Selector) {
		p(s.Not())
	})
}
