// nolint
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/mixin"
	"github.com/google/uuid"
)

// UserNotif holds the schema definition for the UserNotif entity.
type UserNotif struct {
	ent.Schema
}

func (UserNotif) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the UserNotif.
func (UserNotif) Fields() []ent.Field {
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

// Edges of the UserNotif.
func (UserNotif) Edges() []ent.Edge {
	return nil
}
