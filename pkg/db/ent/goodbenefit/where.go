// Code generated by ent, DO NOT EDIT.

package goodbenefit

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// GoodID applies equality check predicate on the "good_id" field. It's identical to GoodIDEQ.
func GoodID(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodID), v))
	})
}

// GoodName applies equality check predicate on the "good_name" field. It's identical to GoodNameEQ.
func GoodName(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodName), v))
	})
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// Message applies equality check predicate on the "message" field. It's identical to MessageEQ.
func Message(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessage), v))
	})
}

// BenefitDate applies equality check predicate on the "benefit_date" field. It's identical to BenefitDateEQ.
func BenefitDate(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBenefitDate), v))
	})
}

// TxID applies equality check predicate on the "tx_id" field. It's identical to TxIDEQ.
func TxID(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTxID), v))
	})
}

// Notified applies equality check predicate on the "notified" field. It's identical to NotifiedEQ.
func Notified(v bool) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNotified), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.GoodBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.GoodBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.GoodBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.GoodBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.GoodBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.GoodBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// GoodIDEQ applies the EQ predicate on the "good_id" field.
func GoodIDEQ(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodID), v))
	})
}

// GoodIDNEQ applies the NEQ predicate on the "good_id" field.
func GoodIDNEQ(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGoodID), v))
	})
}

// GoodIDIn applies the In predicate on the "good_id" field.
func GoodIDIn(vs ...uuid.UUID) predicate.GoodBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldGoodID), v...))
	})
}

// GoodIDNotIn applies the NotIn predicate on the "good_id" field.
func GoodIDNotIn(vs ...uuid.UUID) predicate.GoodBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldGoodID), v...))
	})
}

// GoodIDGT applies the GT predicate on the "good_id" field.
func GoodIDGT(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGoodID), v))
	})
}

// GoodIDGTE applies the GTE predicate on the "good_id" field.
func GoodIDGTE(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGoodID), v))
	})
}

// GoodIDLT applies the LT predicate on the "good_id" field.
func GoodIDLT(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGoodID), v))
	})
}

// GoodIDLTE applies the LTE predicate on the "good_id" field.
func GoodIDLTE(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGoodID), v))
	})
}

// GoodIDIsNil applies the IsNil predicate on the "good_id" field.
func GoodIDIsNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldGoodID)))
	})
}

// GoodIDNotNil applies the NotNil predicate on the "good_id" field.
func GoodIDNotNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldGoodID)))
	})
}

// GoodNameEQ applies the EQ predicate on the "good_name" field.
func GoodNameEQ(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodName), v))
	})
}

// GoodNameNEQ applies the NEQ predicate on the "good_name" field.
func GoodNameNEQ(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGoodName), v))
	})
}

// GoodNameIn applies the In predicate on the "good_name" field.
func GoodNameIn(vs ...string) predicate.GoodBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldGoodName), v...))
	})
}

// GoodNameNotIn applies the NotIn predicate on the "good_name" field.
func GoodNameNotIn(vs ...string) predicate.GoodBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldGoodName), v...))
	})
}

// GoodNameGT applies the GT predicate on the "good_name" field.
func GoodNameGT(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGoodName), v))
	})
}

// GoodNameGTE applies the GTE predicate on the "good_name" field.
func GoodNameGTE(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGoodName), v))
	})
}

// GoodNameLT applies the LT predicate on the "good_name" field.
func GoodNameLT(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGoodName), v))
	})
}

// GoodNameLTE applies the LTE predicate on the "good_name" field.
func GoodNameLTE(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGoodName), v))
	})
}

// GoodNameContains applies the Contains predicate on the "good_name" field.
func GoodNameContains(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGoodName), v))
	})
}

// GoodNameHasPrefix applies the HasPrefix predicate on the "good_name" field.
func GoodNameHasPrefix(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGoodName), v))
	})
}

// GoodNameHasSuffix applies the HasSuffix predicate on the "good_name" field.
func GoodNameHasSuffix(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGoodName), v))
	})
}

// GoodNameIsNil applies the IsNil predicate on the "good_name" field.
func GoodNameIsNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldGoodName)))
	})
}

// GoodNameNotNil applies the NotNil predicate on the "good_name" field.
func GoodNameNotNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldGoodName)))
	})
}

// GoodNameEqualFold applies the EqualFold predicate on the "good_name" field.
func GoodNameEqualFold(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGoodName), v))
	})
}

// GoodNameContainsFold applies the ContainsFold predicate on the "good_name" field.
func GoodNameContainsFold(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGoodName), v))
	})
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmount), v))
	})
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...string) predicate.GoodBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAmount), v...))
	})
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...string) predicate.GoodBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAmount), v...))
	})
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmount), v))
	})
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmount), v))
	})
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmount), v))
	})
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmount), v))
	})
}

// AmountContains applies the Contains predicate on the "amount" field.
func AmountContains(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAmount), v))
	})
}

// AmountHasPrefix applies the HasPrefix predicate on the "amount" field.
func AmountHasPrefix(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAmount), v))
	})
}

// AmountHasSuffix applies the HasSuffix predicate on the "amount" field.
func AmountHasSuffix(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAmount), v))
	})
}

// AmountIsNil applies the IsNil predicate on the "amount" field.
func AmountIsNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAmount)))
	})
}

// AmountNotNil applies the NotNil predicate on the "amount" field.
func AmountNotNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAmount)))
	})
}

// AmountEqualFold applies the EqualFold predicate on the "amount" field.
func AmountEqualFold(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAmount), v))
	})
}

// AmountContainsFold applies the ContainsFold predicate on the "amount" field.
func AmountContainsFold(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAmount), v))
	})
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldState), v))
	})
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...string) predicate.GoodBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldState), v...))
	})
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...string) predicate.GoodBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldState), v...))
	})
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldState), v))
	})
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldState), v))
	})
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldState), v))
	})
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldState), v))
	})
}

// StateContains applies the Contains predicate on the "state" field.
func StateContains(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldState), v))
	})
}

// StateHasPrefix applies the HasPrefix predicate on the "state" field.
func StateHasPrefix(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldState), v))
	})
}

// StateHasSuffix applies the HasSuffix predicate on the "state" field.
func StateHasSuffix(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldState), v))
	})
}

// StateIsNil applies the IsNil predicate on the "state" field.
func StateIsNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldState)))
	})
}

// StateNotNil applies the NotNil predicate on the "state" field.
func StateNotNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldState)))
	})
}

// StateEqualFold applies the EqualFold predicate on the "state" field.
func StateEqualFold(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldState), v))
	})
}

// StateContainsFold applies the ContainsFold predicate on the "state" field.
func StateContainsFold(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldState), v))
	})
}

// MessageEQ applies the EQ predicate on the "message" field.
func MessageEQ(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessage), v))
	})
}

// MessageNEQ applies the NEQ predicate on the "message" field.
func MessageNEQ(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMessage), v))
	})
}

// MessageIn applies the In predicate on the "message" field.
func MessageIn(vs ...string) predicate.GoodBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMessage), v...))
	})
}

// MessageNotIn applies the NotIn predicate on the "message" field.
func MessageNotIn(vs ...string) predicate.GoodBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMessage), v...))
	})
}

// MessageGT applies the GT predicate on the "message" field.
func MessageGT(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMessage), v))
	})
}

// MessageGTE applies the GTE predicate on the "message" field.
func MessageGTE(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMessage), v))
	})
}

// MessageLT applies the LT predicate on the "message" field.
func MessageLT(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMessage), v))
	})
}

// MessageLTE applies the LTE predicate on the "message" field.
func MessageLTE(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMessage), v))
	})
}

// MessageContains applies the Contains predicate on the "message" field.
func MessageContains(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMessage), v))
	})
}

// MessageHasPrefix applies the HasPrefix predicate on the "message" field.
func MessageHasPrefix(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMessage), v))
	})
}

// MessageHasSuffix applies the HasSuffix predicate on the "message" field.
func MessageHasSuffix(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMessage), v))
	})
}

// MessageIsNil applies the IsNil predicate on the "message" field.
func MessageIsNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMessage)))
	})
}

// MessageNotNil applies the NotNil predicate on the "message" field.
func MessageNotNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMessage)))
	})
}

// MessageEqualFold applies the EqualFold predicate on the "message" field.
func MessageEqualFold(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMessage), v))
	})
}

// MessageContainsFold applies the ContainsFold predicate on the "message" field.
func MessageContainsFold(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMessage), v))
	})
}

// BenefitDateEQ applies the EQ predicate on the "benefit_date" field.
func BenefitDateEQ(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBenefitDate), v))
	})
}

// BenefitDateNEQ applies the NEQ predicate on the "benefit_date" field.
func BenefitDateNEQ(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBenefitDate), v))
	})
}

// BenefitDateIn applies the In predicate on the "benefit_date" field.
func BenefitDateIn(vs ...uint32) predicate.GoodBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBenefitDate), v...))
	})
}

// BenefitDateNotIn applies the NotIn predicate on the "benefit_date" field.
func BenefitDateNotIn(vs ...uint32) predicate.GoodBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBenefitDate), v...))
	})
}

// BenefitDateGT applies the GT predicate on the "benefit_date" field.
func BenefitDateGT(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBenefitDate), v))
	})
}

// BenefitDateGTE applies the GTE predicate on the "benefit_date" field.
func BenefitDateGTE(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBenefitDate), v))
	})
}

// BenefitDateLT applies the LT predicate on the "benefit_date" field.
func BenefitDateLT(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBenefitDate), v))
	})
}

// BenefitDateLTE applies the LTE predicate on the "benefit_date" field.
func BenefitDateLTE(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBenefitDate), v))
	})
}

// BenefitDateIsNil applies the IsNil predicate on the "benefit_date" field.
func BenefitDateIsNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldBenefitDate)))
	})
}

// BenefitDateNotNil applies the NotNil predicate on the "benefit_date" field.
func BenefitDateNotNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldBenefitDate)))
	})
}

// TxIDEQ applies the EQ predicate on the "tx_id" field.
func TxIDEQ(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTxID), v))
	})
}

// TxIDNEQ applies the NEQ predicate on the "tx_id" field.
func TxIDNEQ(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTxID), v))
	})
}

// TxIDIn applies the In predicate on the "tx_id" field.
func TxIDIn(vs ...uuid.UUID) predicate.GoodBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTxID), v...))
	})
}

// TxIDNotIn applies the NotIn predicate on the "tx_id" field.
func TxIDNotIn(vs ...uuid.UUID) predicate.GoodBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTxID), v...))
	})
}

// TxIDGT applies the GT predicate on the "tx_id" field.
func TxIDGT(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTxID), v))
	})
}

// TxIDGTE applies the GTE predicate on the "tx_id" field.
func TxIDGTE(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTxID), v))
	})
}

// TxIDLT applies the LT predicate on the "tx_id" field.
func TxIDLT(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTxID), v))
	})
}

// TxIDLTE applies the LTE predicate on the "tx_id" field.
func TxIDLTE(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTxID), v))
	})
}

// TxIDIsNil applies the IsNil predicate on the "tx_id" field.
func TxIDIsNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTxID)))
	})
}

// TxIDNotNil applies the NotNil predicate on the "tx_id" field.
func TxIDNotNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTxID)))
	})
}

// NotifiedEQ applies the EQ predicate on the "notified" field.
func NotifiedEQ(v bool) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNotified), v))
	})
}

// NotifiedNEQ applies the NEQ predicate on the "notified" field.
func NotifiedNEQ(v bool) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNotified), v))
	})
}

// NotifiedIsNil applies the IsNil predicate on the "notified" field.
func NotifiedIsNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldNotified)))
	})
}

// NotifiedNotNil applies the NotNil predicate on the "notified" field.
func NotifiedNotNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldNotified)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GoodBenefit) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GoodBenefit) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
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
func Not(p predicate.GoodBenefit) predicate.GoodBenefit {
	return predicate.GoodBenefit(func(s *sql.Selector) {
		p(s.Not())
	})
}
