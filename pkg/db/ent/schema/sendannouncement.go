package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/mixin"
	"github.com/google/uuid"
)

// SendAnnouncement holds the schema definition for the SendAnnouncement entity.
type SendAnnouncement struct {
	ent.Schema
}

func (SendAnnouncement) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the SendAnnouncement.
func (SendAnnouncement) Fields() []ent.Field {
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
			UUID("announcement_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			String("channel").
			Optional().
			Default(basetypes.NotifChannel_DefaultChannel.String()),
	}
}

// Edges of the SendAnnouncement.
func (SendAnnouncement) Edges() []ent.Edge {
	return nil
}
