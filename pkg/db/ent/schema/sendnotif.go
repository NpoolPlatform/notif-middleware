// nolint
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/mixin"
	"github.com/google/uuid"
)

// SendNotif holds the schema definition for the SendNotif entity.
type SendNotif struct {
	ent.Schema
}

func (SendNotif) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the SendNotif.
func (SendNotif) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			UUID("user_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			UUID("notif_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			String("channel").
			Optional().
			Default(basetypes.NotifChannel_DefaultChannel.String()),
	}
}

// Edges of the SendNotif.
func (SendNotif) Edges() []ent.Edge {
	return nil
}
