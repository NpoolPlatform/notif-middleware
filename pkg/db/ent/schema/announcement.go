package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/mixin"
	"github.com/google/uuid"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

// Announcement holds the schema definition for the Announcement entity.
type Announcement struct {
	ent.Schema
}

func (Announcement) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the Announcement.
func (Announcement) Fields() []ent.Field {
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
			UUID("lang_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			String("title").
			Optional().
			Default(""),
		field.
			Text("content").
			Optional().
			Default(""),
		field.
			String("channel").
			Optional().
			Default(basetypes.NotifChannel_DefaultChannel.String()),
		field.
			Uint32("start_at").
			Optional().
			Default(0),
		field.
			Uint32("end_at").
			Optional().
			Default(0),
		field.
			String("type").
			Optional().
			Default(basetypes.NotifType_DefaultNotifType.String()),
	}
}

// Edges of the Announcement.
func (Announcement) Edges() []ent.Edge {
	return nil
}
