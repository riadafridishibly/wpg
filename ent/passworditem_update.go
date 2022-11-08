// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/riadafridishibly/vault-app/ent/passworditem"
	"github.com/riadafridishibly/vault-app/ent/predicate"
)

// PasswordItemUpdate is the builder for updating PasswordItem entities.
type PasswordItemUpdate struct {
	config
	hooks    []Hook
	mutation *PasswordItemMutation
}

// Where appends a list predicates to the PasswordItemUpdate builder.
func (piu *PasswordItemUpdate) Where(ps ...predicate.PasswordItem) *PasswordItemUpdate {
	piu.mutation.Where(ps...)
	return piu
}

// SetDescription sets the "description" field.
func (piu *PasswordItemUpdate) SetDescription(s string) *PasswordItemUpdate {
	piu.mutation.SetDescription(s)
	return piu
}

// SetSiteName sets the "site_name" field.
func (piu *PasswordItemUpdate) SetSiteName(s string) *PasswordItemUpdate {
	piu.mutation.SetSiteName(s)
	return piu
}

// SetUsername sets the "username" field.
func (piu *PasswordItemUpdate) SetUsername(s string) *PasswordItemUpdate {
	piu.mutation.SetUsername(s)
	return piu
}

// SetPassword sets the "password" field.
func (piu *PasswordItemUpdate) SetPassword(s string) *PasswordItemUpdate {
	piu.mutation.SetPassword(s)
	return piu
}

// SetTags sets the "tags" field.
func (piu *PasswordItemUpdate) SetTags(s []string) *PasswordItemUpdate {
	piu.mutation.SetTags(s)
	return piu
}

// AppendTags appends s to the "tags" field.
func (piu *PasswordItemUpdate) AppendTags(s []string) *PasswordItemUpdate {
	piu.mutation.AppendTags(s)
	return piu
}

// ClearTags clears the value of the "tags" field.
func (piu *PasswordItemUpdate) ClearTags() *PasswordItemUpdate {
	piu.mutation.ClearTags()
	return piu
}

// SetCreatedAt sets the "created_at" field.
func (piu *PasswordItemUpdate) SetCreatedAt(t time.Time) *PasswordItemUpdate {
	piu.mutation.SetCreatedAt(t)
	return piu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (piu *PasswordItemUpdate) SetNillableCreatedAt(t *time.Time) *PasswordItemUpdate {
	if t != nil {
		piu.SetCreatedAt(*t)
	}
	return piu
}

// SetUpdatedAt sets the "updated_at" field.
func (piu *PasswordItemUpdate) SetUpdatedAt(t time.Time) *PasswordItemUpdate {
	piu.mutation.SetUpdatedAt(t)
	return piu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (piu *PasswordItemUpdate) SetNillableUpdatedAt(t *time.Time) *PasswordItemUpdate {
	if t != nil {
		piu.SetUpdatedAt(*t)
	}
	return piu
}

// Mutation returns the PasswordItemMutation object of the builder.
func (piu *PasswordItemUpdate) Mutation() *PasswordItemMutation {
	return piu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (piu *PasswordItemUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(piu.hooks) == 0 {
		if err = piu.check(); err != nil {
			return 0, err
		}
		affected, err = piu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PasswordItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = piu.check(); err != nil {
				return 0, err
			}
			piu.mutation = mutation
			affected, err = piu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(piu.hooks) - 1; i >= 0; i-- {
			if piu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = piu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, piu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (piu *PasswordItemUpdate) SaveX(ctx context.Context) int {
	affected, err := piu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (piu *PasswordItemUpdate) Exec(ctx context.Context) error {
	_, err := piu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (piu *PasswordItemUpdate) ExecX(ctx context.Context) {
	if err := piu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (piu *PasswordItemUpdate) check() error {
	if v, ok := piu.mutation.Description(); ok {
		if err := passworditem.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "PasswordItem.description": %w`, err)}
		}
	}
	return nil
}

func (piu *PasswordItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   passworditem.Table,
			Columns: passworditem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: passworditem.FieldID,
			},
		},
	}
	if ps := piu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := piu.mutation.Description(); ok {
		_spec.SetField(passworditem.FieldDescription, field.TypeString, value)
	}
	if value, ok := piu.mutation.SiteName(); ok {
		_spec.SetField(passworditem.FieldSiteName, field.TypeString, value)
	}
	if value, ok := piu.mutation.Username(); ok {
		_spec.SetField(passworditem.FieldUsername, field.TypeString, value)
	}
	if value, ok := piu.mutation.Password(); ok {
		_spec.SetField(passworditem.FieldPassword, field.TypeString, value)
	}
	if value, ok := piu.mutation.Tags(); ok {
		_spec.SetField(passworditem.FieldTags, field.TypeJSON, value)
	}
	if value, ok := piu.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, passworditem.FieldTags, value)
		})
	}
	if piu.mutation.TagsCleared() {
		_spec.ClearField(passworditem.FieldTags, field.TypeJSON)
	}
	if value, ok := piu.mutation.CreatedAt(); ok {
		_spec.SetField(passworditem.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := piu.mutation.UpdatedAt(); ok {
		_spec.SetField(passworditem.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, piu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{passworditem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// PasswordItemUpdateOne is the builder for updating a single PasswordItem entity.
type PasswordItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PasswordItemMutation
}

// SetDescription sets the "description" field.
func (piuo *PasswordItemUpdateOne) SetDescription(s string) *PasswordItemUpdateOne {
	piuo.mutation.SetDescription(s)
	return piuo
}

// SetSiteName sets the "site_name" field.
func (piuo *PasswordItemUpdateOne) SetSiteName(s string) *PasswordItemUpdateOne {
	piuo.mutation.SetSiteName(s)
	return piuo
}

// SetUsername sets the "username" field.
func (piuo *PasswordItemUpdateOne) SetUsername(s string) *PasswordItemUpdateOne {
	piuo.mutation.SetUsername(s)
	return piuo
}

// SetPassword sets the "password" field.
func (piuo *PasswordItemUpdateOne) SetPassword(s string) *PasswordItemUpdateOne {
	piuo.mutation.SetPassword(s)
	return piuo
}

// SetTags sets the "tags" field.
func (piuo *PasswordItemUpdateOne) SetTags(s []string) *PasswordItemUpdateOne {
	piuo.mutation.SetTags(s)
	return piuo
}

// AppendTags appends s to the "tags" field.
func (piuo *PasswordItemUpdateOne) AppendTags(s []string) *PasswordItemUpdateOne {
	piuo.mutation.AppendTags(s)
	return piuo
}

// ClearTags clears the value of the "tags" field.
func (piuo *PasswordItemUpdateOne) ClearTags() *PasswordItemUpdateOne {
	piuo.mutation.ClearTags()
	return piuo
}

// SetCreatedAt sets the "created_at" field.
func (piuo *PasswordItemUpdateOne) SetCreatedAt(t time.Time) *PasswordItemUpdateOne {
	piuo.mutation.SetCreatedAt(t)
	return piuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (piuo *PasswordItemUpdateOne) SetNillableCreatedAt(t *time.Time) *PasswordItemUpdateOne {
	if t != nil {
		piuo.SetCreatedAt(*t)
	}
	return piuo
}

// SetUpdatedAt sets the "updated_at" field.
func (piuo *PasswordItemUpdateOne) SetUpdatedAt(t time.Time) *PasswordItemUpdateOne {
	piuo.mutation.SetUpdatedAt(t)
	return piuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (piuo *PasswordItemUpdateOne) SetNillableUpdatedAt(t *time.Time) *PasswordItemUpdateOne {
	if t != nil {
		piuo.SetUpdatedAt(*t)
	}
	return piuo
}

// Mutation returns the PasswordItemMutation object of the builder.
func (piuo *PasswordItemUpdateOne) Mutation() *PasswordItemMutation {
	return piuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (piuo *PasswordItemUpdateOne) Select(field string, fields ...string) *PasswordItemUpdateOne {
	piuo.fields = append([]string{field}, fields...)
	return piuo
}

// Save executes the query and returns the updated PasswordItem entity.
func (piuo *PasswordItemUpdateOne) Save(ctx context.Context) (*PasswordItem, error) {
	var (
		err  error
		node *PasswordItem
	)
	if len(piuo.hooks) == 0 {
		if err = piuo.check(); err != nil {
			return nil, err
		}
		node, err = piuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PasswordItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = piuo.check(); err != nil {
				return nil, err
			}
			piuo.mutation = mutation
			node, err = piuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(piuo.hooks) - 1; i >= 0; i-- {
			if piuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = piuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, piuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*PasswordItem)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PasswordItemMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (piuo *PasswordItemUpdateOne) SaveX(ctx context.Context) *PasswordItem {
	node, err := piuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (piuo *PasswordItemUpdateOne) Exec(ctx context.Context) error {
	_, err := piuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (piuo *PasswordItemUpdateOne) ExecX(ctx context.Context) {
	if err := piuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (piuo *PasswordItemUpdateOne) check() error {
	if v, ok := piuo.mutation.Description(); ok {
		if err := passworditem.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "PasswordItem.description": %w`, err)}
		}
	}
	return nil
}

func (piuo *PasswordItemUpdateOne) sqlSave(ctx context.Context) (_node *PasswordItem, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   passworditem.Table,
			Columns: passworditem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: passworditem.FieldID,
			},
		},
	}
	id, ok := piuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PasswordItem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := piuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, passworditem.FieldID)
		for _, f := range fields {
			if !passworditem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != passworditem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := piuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := piuo.mutation.Description(); ok {
		_spec.SetField(passworditem.FieldDescription, field.TypeString, value)
	}
	if value, ok := piuo.mutation.SiteName(); ok {
		_spec.SetField(passworditem.FieldSiteName, field.TypeString, value)
	}
	if value, ok := piuo.mutation.Username(); ok {
		_spec.SetField(passworditem.FieldUsername, field.TypeString, value)
	}
	if value, ok := piuo.mutation.Password(); ok {
		_spec.SetField(passworditem.FieldPassword, field.TypeString, value)
	}
	if value, ok := piuo.mutation.Tags(); ok {
		_spec.SetField(passworditem.FieldTags, field.TypeJSON, value)
	}
	if value, ok := piuo.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, passworditem.FieldTags, value)
		})
	}
	if piuo.mutation.TagsCleared() {
		_spec.ClearField(passworditem.FieldTags, field.TypeJSON)
	}
	if value, ok := piuo.mutation.CreatedAt(); ok {
		_spec.SetField(passworditem.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := piuo.mutation.UpdatedAt(); ok {
		_spec.SetField(passworditem.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &PasswordItem{config: piuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, piuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{passworditem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
