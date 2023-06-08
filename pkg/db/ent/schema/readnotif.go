package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/mixin"
	"github.com/google/uuid"
)

// ReadNotif holds the schema definition for the ReadNotif entity.
type ReadNotif struct {
	ent.Schema
}

func (ReadNotif) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the ReadNotif.
func (ReadNotif) Fields() []ent.Field {
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
	}
}

// Edges of the ReadNotif.
func (ReadNotif) Edges() []ent.Edge {
	return nil
}
