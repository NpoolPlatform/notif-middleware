//nolint:dupl
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/mixin"
	"github.com/google/uuid"
)

// SMSTemplate holds the schema definition for the SMSTemplate entity.
type SMSTemplate struct {
	ent.Schema
}

func (SMSTemplate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the AppSMSTemplate.
func (SMSTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}),
		field.
			UUID("lang_id", uuid.UUID{}),
		field.
			String("used_for").
			Optional().
			Default(basetypes.UsedFor_DefaultUsedFor.String()),
		field.
			String("subject").
			Optional().
			Default(""),
		field.
			String("message").
			Optional().
			Default(""),
	}
}
