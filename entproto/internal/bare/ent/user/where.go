// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/contrib/entproto/internal/bare/ent/predicate"
	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// EmailAddress applies equality check predicate on the "email_address" field. It's identical to EmailAddressEQ.
func EmailAddress(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmailAddress, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldName, v))
}

// EmailAddressEQ applies the EQ predicate on the "email_address" field.
func EmailAddressEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmailAddress, v))
}

// EmailAddressNEQ applies the NEQ predicate on the "email_address" field.
func EmailAddressNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmailAddress, v))
}

// EmailAddressIn applies the In predicate on the "email_address" field.
func EmailAddressIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldEmailAddress, vs...))
}

// EmailAddressNotIn applies the NotIn predicate on the "email_address" field.
func EmailAddressNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldEmailAddress, vs...))
}

// EmailAddressGT applies the GT predicate on the "email_address" field.
func EmailAddressGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldEmailAddress, v))
}

// EmailAddressGTE applies the GTE predicate on the "email_address" field.
func EmailAddressGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldEmailAddress, v))
}

// EmailAddressLT applies the LT predicate on the "email_address" field.
func EmailAddressLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldEmailAddress, v))
}

// EmailAddressLTE applies the LTE predicate on the "email_address" field.
func EmailAddressLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldEmailAddress, v))
}

// EmailAddressContains applies the Contains predicate on the "email_address" field.
func EmailAddressContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldEmailAddress, v))
}

// EmailAddressHasPrefix applies the HasPrefix predicate on the "email_address" field.
func EmailAddressHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldEmailAddress, v))
}

// EmailAddressHasSuffix applies the HasSuffix predicate on the "email_address" field.
func EmailAddressHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldEmailAddress, v))
}

// EmailAddressEqualFold applies the EqualFold predicate on the "email_address" field.
func EmailAddressEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldEmailAddress, v))
}

// EmailAddressContainsFold applies the ContainsFold predicate on the "email_address" field.
func EmailAddressContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldEmailAddress, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
