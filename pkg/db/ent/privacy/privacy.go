// Code generated by ent, DO NOT EDIT.

package privacy

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/notif-middleware/pkg/db/ent"

	"entgo.io/ent/entql"
	"entgo.io/ent/privacy"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with allow decision.
	Allow = privacy.Allow

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with deny decision.
	Deny = privacy.Deny

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = privacy.Skip
)

// Allowf returns an formatted wrapped Allow decision.
func Allowf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Allow)...)
}

// Denyf returns an formatted wrapped Deny decision.
func Denyf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Deny)...)
}

// Skipf returns an formatted wrapped Skip decision.
func Skipf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Skip)...)
}

// DecisionContext creates a new context from the given parent context with
// a policy decision attach to it.
func DecisionContext(parent context.Context, decision error) context.Context {
	return privacy.DecisionContext(parent, decision)
}

// DecisionFromContext retrieves the policy decision from the context.
func DecisionFromContext(ctx context.Context) (error, bool) {
	return privacy.DecisionFromContext(ctx)
}

type (
	// Policy groups query and mutation policies.
	Policy = privacy.Policy

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule = privacy.QueryRule
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy = privacy.QueryPolicy

	// MutationRule defines the interface which decides whether a
	// mutation is allowed and optionally modifies it.
	MutationRule = privacy.MutationRule
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy = privacy.MutationPolicy
)

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

// MutationRuleFunc type is an adapter which allows the use of
// ordinary functions as mutation rules.
type MutationRuleFunc func(context.Context, ent.Mutation) error

// EvalMutation returns f(ctx, m).
func (f MutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return f(ctx, m)
}

// QueryMutationRule is an interface which groups query and mutation rules.
type QueryMutationRule interface {
	QueryRule
	MutationRule
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return fixedDecision{Allow}
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return fixedDecision{Deny}
}

type fixedDecision struct {
	decision error
}

func (f fixedDecision) EvalQuery(context.Context, ent.Query) error {
	return f.decision
}

func (f fixedDecision) EvalMutation(context.Context, ent.Mutation) error {
	return f.decision
}

type contextDecision struct {
	eval func(context.Context) error
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return contextDecision{eval}
}

func (c contextDecision) EvalQuery(ctx context.Context, _ ent.Query) error {
	return c.eval(ctx)
}

func (c contextDecision) EvalMutation(ctx context.Context, _ ent.Mutation) error {
	return c.eval(ctx)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op ent.Op) MutationRule {
	return MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		if m.Op().Is(op) {
			return rule.EvalMutation(ctx, m)
		}
		return Skip
	})
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op ent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		return Denyf("ent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The AnnouncementQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AnnouncementQueryRuleFunc func(context.Context, *ent.AnnouncementQuery) error

// EvalQuery return f(ctx, q).
func (f AnnouncementQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AnnouncementQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AnnouncementQuery", q)
}

// The AnnouncementMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AnnouncementMutationRuleFunc func(context.Context, *ent.AnnouncementMutation) error

// EvalMutation calls f(ctx, m).
func (f AnnouncementMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AnnouncementMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AnnouncementMutation", m)
}

// The ContactQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ContactQueryRuleFunc func(context.Context, *ent.ContactQuery) error

// EvalQuery return f(ctx, q).
func (f ContactQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ContactQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ContactQuery", q)
}

// The ContactMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ContactMutationRuleFunc func(context.Context, *ent.ContactMutation) error

// EvalMutation calls f(ctx, m).
func (f ContactMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ContactMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ContactMutation", m)
}

// The EmailTemplateQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type EmailTemplateQueryRuleFunc func(context.Context, *ent.EmailTemplateQuery) error

// EvalQuery return f(ctx, q).
func (f EmailTemplateQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.EmailTemplateQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.EmailTemplateQuery", q)
}

// The EmailTemplateMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type EmailTemplateMutationRuleFunc func(context.Context, *ent.EmailTemplateMutation) error

// EvalMutation calls f(ctx, m).
func (f EmailTemplateMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.EmailTemplateMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.EmailTemplateMutation", m)
}

// The FrontendTemplateQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type FrontendTemplateQueryRuleFunc func(context.Context, *ent.FrontendTemplateQuery) error

// EvalQuery return f(ctx, q).
func (f FrontendTemplateQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.FrontendTemplateQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.FrontendTemplateQuery", q)
}

// The FrontendTemplateMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type FrontendTemplateMutationRuleFunc func(context.Context, *ent.FrontendTemplateMutation) error

// EvalMutation calls f(ctx, m).
func (f FrontendTemplateMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.FrontendTemplateMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.FrontendTemplateMutation", m)
}

// The GoodBenefitQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GoodBenefitQueryRuleFunc func(context.Context, *ent.GoodBenefitQuery) error

// EvalQuery return f(ctx, q).
func (f GoodBenefitQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.GoodBenefitQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.GoodBenefitQuery", q)
}

// The GoodBenefitMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GoodBenefitMutationRuleFunc func(context.Context, *ent.GoodBenefitMutation) error

// EvalMutation calls f(ctx, m).
func (f GoodBenefitMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.GoodBenefitMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.GoodBenefitMutation", m)
}

// The NotifQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type NotifQueryRuleFunc func(context.Context, *ent.NotifQuery) error

// EvalQuery return f(ctx, q).
func (f NotifQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.NotifQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.NotifQuery", q)
}

// The NotifMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type NotifMutationRuleFunc func(context.Context, *ent.NotifMutation) error

// EvalMutation calls f(ctx, m).
func (f NotifMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.NotifMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.NotifMutation", m)
}

// The NotifChannelQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type NotifChannelQueryRuleFunc func(context.Context, *ent.NotifChannelQuery) error

// EvalQuery return f(ctx, q).
func (f NotifChannelQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.NotifChannelQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.NotifChannelQuery", q)
}

// The NotifChannelMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type NotifChannelMutationRuleFunc func(context.Context, *ent.NotifChannelMutation) error

// EvalMutation calls f(ctx, m).
func (f NotifChannelMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.NotifChannelMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.NotifChannelMutation", m)
}

// The NotifUserQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type NotifUserQueryRuleFunc func(context.Context, *ent.NotifUserQuery) error

// EvalQuery return f(ctx, q).
func (f NotifUserQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.NotifUserQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.NotifUserQuery", q)
}

// The NotifUserMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type NotifUserMutationRuleFunc func(context.Context, *ent.NotifUserMutation) error

// EvalMutation calls f(ctx, m).
func (f NotifUserMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.NotifUserMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.NotifUserMutation", m)
}

// The ReadAnnouncementQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ReadAnnouncementQueryRuleFunc func(context.Context, *ent.ReadAnnouncementQuery) error

// EvalQuery return f(ctx, q).
func (f ReadAnnouncementQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ReadAnnouncementQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ReadAnnouncementQuery", q)
}

// The ReadAnnouncementMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ReadAnnouncementMutationRuleFunc func(context.Context, *ent.ReadAnnouncementMutation) error

// EvalMutation calls f(ctx, m).
func (f ReadAnnouncementMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ReadAnnouncementMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ReadAnnouncementMutation", m)
}

// The SMSTemplateQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type SMSTemplateQueryRuleFunc func(context.Context, *ent.SMSTemplateQuery) error

// EvalQuery return f(ctx, q).
func (f SMSTemplateQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.SMSTemplateQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.SMSTemplateQuery", q)
}

// The SMSTemplateMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type SMSTemplateMutationRuleFunc func(context.Context, *ent.SMSTemplateMutation) error

// EvalMutation calls f(ctx, m).
func (f SMSTemplateMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.SMSTemplateMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.SMSTemplateMutation", m)
}

// The SendAnnouncementQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type SendAnnouncementQueryRuleFunc func(context.Context, *ent.SendAnnouncementQuery) error

// EvalQuery return f(ctx, q).
func (f SendAnnouncementQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.SendAnnouncementQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.SendAnnouncementQuery", q)
}

// The SendAnnouncementMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type SendAnnouncementMutationRuleFunc func(context.Context, *ent.SendAnnouncementMutation) error

// EvalMutation calls f(ctx, m).
func (f SendAnnouncementMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.SendAnnouncementMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.SendAnnouncementMutation", m)
}

// The TxNotifStateQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type TxNotifStateQueryRuleFunc func(context.Context, *ent.TxNotifStateQuery) error

// EvalQuery return f(ctx, q).
func (f TxNotifStateQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TxNotifStateQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.TxNotifStateQuery", q)
}

// The TxNotifStateMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type TxNotifStateMutationRuleFunc func(context.Context, *ent.TxNotifStateMutation) error

// EvalMutation calls f(ctx, m).
func (f TxNotifStateMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.TxNotifStateMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.TxNotifStateMutation", m)
}

// The UserAnnouncementQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type UserAnnouncementQueryRuleFunc func(context.Context, *ent.UserAnnouncementQuery) error

// EvalQuery return f(ctx, q).
func (f UserAnnouncementQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.UserAnnouncementQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.UserAnnouncementQuery", q)
}

// The UserAnnouncementMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type UserAnnouncementMutationRuleFunc func(context.Context, *ent.UserAnnouncementMutation) error

// EvalMutation calls f(ctx, m).
func (f UserAnnouncementMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.UserAnnouncementMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.UserAnnouncementMutation", m)
}

type (
	// Filter is the interface that wraps the Where function
	// for filtering nodes in queries and mutations.
	Filter interface {
		// Where applies a filter on the executed query/mutation.
		Where(entql.P)
	}

	// The FilterFunc type is an adapter that allows the use of ordinary
	// functions as filters for query and mutation types.
	FilterFunc func(context.Context, Filter) error
)

// EvalQuery calls f(ctx, q) if the query implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	fr, err := queryFilter(q)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

// EvalMutation calls f(ctx, q) if the mutation implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	fr, err := mutationFilter(m)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

var _ QueryMutationRule = FilterFunc(nil)

func queryFilter(q ent.Query) (Filter, error) {
	switch q := q.(type) {
	case *ent.AnnouncementQuery:
		return q.Filter(), nil
	case *ent.ContactQuery:
		return q.Filter(), nil
	case *ent.EmailTemplateQuery:
		return q.Filter(), nil
	case *ent.FrontendTemplateQuery:
		return q.Filter(), nil
	case *ent.GoodBenefitQuery:
		return q.Filter(), nil
	case *ent.NotifQuery:
		return q.Filter(), nil
	case *ent.NotifChannelQuery:
		return q.Filter(), nil
	case *ent.NotifUserQuery:
		return q.Filter(), nil
	case *ent.ReadAnnouncementQuery:
		return q.Filter(), nil
	case *ent.SMSTemplateQuery:
		return q.Filter(), nil
	case *ent.SendAnnouncementQuery:
		return q.Filter(), nil
	case *ent.TxNotifStateQuery:
		return q.Filter(), nil
	case *ent.UserAnnouncementQuery:
		return q.Filter(), nil
	default:
		return nil, Denyf("ent/privacy: unexpected query type %T for query filter", q)
	}
}

func mutationFilter(m ent.Mutation) (Filter, error) {
	switch m := m.(type) {
	case *ent.AnnouncementMutation:
		return m.Filter(), nil
	case *ent.ContactMutation:
		return m.Filter(), nil
	case *ent.EmailTemplateMutation:
		return m.Filter(), nil
	case *ent.FrontendTemplateMutation:
		return m.Filter(), nil
	case *ent.GoodBenefitMutation:
		return m.Filter(), nil
	case *ent.NotifMutation:
		return m.Filter(), nil
	case *ent.NotifChannelMutation:
		return m.Filter(), nil
	case *ent.NotifUserMutation:
		return m.Filter(), nil
	case *ent.ReadAnnouncementMutation:
		return m.Filter(), nil
	case *ent.SMSTemplateMutation:
		return m.Filter(), nil
	case *ent.SendAnnouncementMutation:
		return m.Filter(), nil
	case *ent.TxNotifStateMutation:
		return m.Filter(), nil
	case *ent.UserAnnouncementMutation:
		return m.Filter(), nil
	default:
		return nil, Denyf("ent/privacy: unexpected mutation type %T for mutation filter", m)
	}
}
