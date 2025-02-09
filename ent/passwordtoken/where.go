// Code generated by ent, DO NOT EDIT.

package passwordtoken

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/brian-dlee/lab/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldLTE(FieldID, id))
}

// Hash applies equality check predicate on the "hash" field. It's identical to HashEQ.
func Hash(v string) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldEQ(FieldHash, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldEQ(FieldCreatedAt, v))
}

// HashEQ applies the EQ predicate on the "hash" field.
func HashEQ(v string) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldEQ(FieldHash, v))
}

// HashNEQ applies the NEQ predicate on the "hash" field.
func HashNEQ(v string) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldNEQ(FieldHash, v))
}

// HashIn applies the In predicate on the "hash" field.
func HashIn(vs ...string) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldIn(FieldHash, vs...))
}

// HashNotIn applies the NotIn predicate on the "hash" field.
func HashNotIn(vs ...string) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldNotIn(FieldHash, vs...))
}

// HashGT applies the GT predicate on the "hash" field.
func HashGT(v string) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldGT(FieldHash, v))
}

// HashGTE applies the GTE predicate on the "hash" field.
func HashGTE(v string) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldGTE(FieldHash, v))
}

// HashLT applies the LT predicate on the "hash" field.
func HashLT(v string) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldLT(FieldHash, v))
}

// HashLTE applies the LTE predicate on the "hash" field.
func HashLTE(v string) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldLTE(FieldHash, v))
}

// HashContains applies the Contains predicate on the "hash" field.
func HashContains(v string) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldContains(FieldHash, v))
}

// HashHasPrefix applies the HasPrefix predicate on the "hash" field.
func HashHasPrefix(v string) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldHasPrefix(FieldHash, v))
}

// HashHasSuffix applies the HasSuffix predicate on the "hash" field.
func HashHasSuffix(v string) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldHasSuffix(FieldHash, v))
}

// HashEqualFold applies the EqualFold predicate on the "hash" field.
func HashEqualFold(v string) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldEqualFold(FieldHash, v))
}

// HashContainsFold applies the ContainsFold predicate on the "hash" field.
func HashContainsFold(v string) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldContainsFold(FieldHash, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.PasswordToken {
	return predicate.PasswordToken(sql.FieldLTE(FieldCreatedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.PasswordToken {
	return predicate.PasswordToken(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.PasswordToken {
	return predicate.PasswordToken(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PasswordToken) predicate.PasswordToken {
	return predicate.PasswordToken(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PasswordToken) predicate.PasswordToken {
	return predicate.PasswordToken(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PasswordToken) predicate.PasswordToken {
	return predicate.PasswordToken(sql.NotPredicates(p))
}
