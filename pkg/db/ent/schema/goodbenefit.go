//nolint:nolintlint,dupl
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/mixin"
	"github.com/google/uuid"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

// GoodBenefit holds the schema definition for the GoodBenefit entity.
type GoodBenefit struct {
	ent.Schema
}

func (GoodBenefit) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the GoodBenefit.
func (GoodBenefit) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("good_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			String("good_name").
			Optional().
			Default(""),
		field.
			String("amount").
			Optional().
			Default("0"),
		field.
			String("state").
			Optional().
			Default(basetypes.Result_DefaultResult.String()),
		field.
			String("message").
			Optional().
			Default(""),
		field.
			Uint32("benefit_date").
			Optional().
			Default(0),
		field.
			UUID("tx_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			Bool("notified").
			Optional().
			Default(false),
	}
}

// Edges of the GoodBenefit.
func (GoodBenefit) Edges() []ent.Edge {
	return nil
}
