package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/mixin"

	"github.com/google/uuid"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

// Contact holds the schema definition for the Contact entity.
type Contact struct {
	ent.Schema
}

func (Contact) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the AppContact.
func (Contact) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}),
		field.
			String("used_for").
			Optional().
			Default(basetypes.UsedFor_DefaultUsedFor.String()),
		field.
			String("sender").
			Optional().
			Default(""),
		field.
			String("account").
			Optional().
			Default(""),
		field.
			String("account_type").
			Optional().
			Default(""),
	}
}
