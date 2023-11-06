// Code generated by ent, DO NOT EDIT.

package goodbenefit

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the goodbenefit type in the database.
	Label = "good_benefit"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldEntID holds the string denoting the ent_id field in the database.
	FieldEntID = "ent_id"
	// FieldGoodID holds the string denoting the good_id field in the database.
	FieldGoodID = "good_id"
	// FieldGoodName holds the string denoting the good_name field in the database.
	FieldGoodName = "good_name"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldMessage holds the string denoting the message field in the database.
	FieldMessage = "message"
	// FieldBenefitDate holds the string denoting the benefit_date field in the database.
	FieldBenefitDate = "benefit_date"
	// FieldTxID holds the string denoting the tx_id field in the database.
	FieldTxID = "tx_id"
	// FieldGenerated holds the string denoting the generated field in the database.
	FieldGenerated = "generated"
	// Table holds the table name of the goodbenefit in the database.
	Table = "good_benefits"
)

// Columns holds all SQL columns for goodbenefit fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldGoodID,
	FieldGoodName,
	FieldAmount,
	FieldState,
	FieldMessage,
	FieldBenefitDate,
	FieldTxID,
	FieldGenerated,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/NpoolPlatform/notif-middleware/pkg/db/ent/runtime"
//
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() uint32
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() uint32
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() uint32
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt func() uint32
	// DefaultEntID holds the default value on creation for the "ent_id" field.
	DefaultEntID func() uuid.UUID
	// DefaultGoodID holds the default value on creation for the "good_id" field.
	DefaultGoodID func() uuid.UUID
	// DefaultGoodName holds the default value on creation for the "good_name" field.
	DefaultGoodName string
	// DefaultAmount holds the default value on creation for the "amount" field.
	DefaultAmount string
	// DefaultState holds the default value on creation for the "state" field.
	DefaultState string
	// DefaultMessage holds the default value on creation for the "message" field.
	DefaultMessage string
	// DefaultBenefitDate holds the default value on creation for the "benefit_date" field.
	DefaultBenefitDate uint32
	// DefaultTxID holds the default value on creation for the "tx_id" field.
	DefaultTxID func() uuid.UUID
	// DefaultGenerated holds the default value on creation for the "generated" field.
	DefaultGenerated bool
)
