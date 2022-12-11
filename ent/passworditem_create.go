// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/riadafridishibly/vault-app/ent/passworditem"
)

// PasswordItemCreate is the builder for creating a PasswordItem entity.
type PasswordItemCreate struct {
	config
	mutation *PasswordItemMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (pic *PasswordItemCreate) SetCreateTime(t time.Time) *PasswordItemCreate {
	pic.mutation.SetCreateTime(t)
	return pic
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pic *PasswordItemCreate) SetNillableCreateTime(t *time.Time) *PasswordItemCreate {
	if t != nil {
		pic.SetCreateTime(*t)
	}
	return pic
}

// SetUpdateTime sets the "update_time" field.
func (pic *PasswordItemCreate) SetUpdateTime(t time.Time) *PasswordItemCreate {
	pic.mutation.SetUpdateTime(t)
	return pic
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pic *PasswordItemCreate) SetNillableUpdateTime(t *time.Time) *PasswordItemCreate {
	if t != nil {
		pic.SetUpdateTime(*t)
	}
	return pic
}

// SetAvatar sets the "avatar" field.
func (pic *PasswordItemCreate) SetAvatar(s string) *PasswordItemCreate {
	pic.mutation.SetAvatar(s)
	return pic
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (pic *PasswordItemCreate) SetNillableAvatar(s *string) *PasswordItemCreate {
	if s != nil {
		pic.SetAvatar(*s)
	}
	return pic
}

// SetDescription sets the "description" field.
func (pic *PasswordItemCreate) SetDescription(s string) *PasswordItemCreate {
	pic.mutation.SetDescription(s)
	return pic
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pic *PasswordItemCreate) SetNillableDescription(s *string) *PasswordItemCreate {
	if s != nil {
		pic.SetDescription(*s)
	}
	return pic
}

// SetSiteName sets the "site_name" field.
func (pic *PasswordItemCreate) SetSiteName(s string) *PasswordItemCreate {
	pic.mutation.SetSiteName(s)
	return pic
}

// SetNillableSiteName sets the "site_name" field if the given value is not nil.
func (pic *PasswordItemCreate) SetNillableSiteName(s *string) *PasswordItemCreate {
	if s != nil {
		pic.SetSiteName(*s)
	}
	return pic
}

// SetSiteURL sets the "site_url" field.
func (pic *PasswordItemCreate) SetSiteURL(s string) *PasswordItemCreate {
	pic.mutation.SetSiteURL(s)
	return pic
}

// SetNillableSiteURL sets the "site_url" field if the given value is not nil.
func (pic *PasswordItemCreate) SetNillableSiteURL(s *string) *PasswordItemCreate {
	if s != nil {
		pic.SetSiteURL(*s)
	}
	return pic
}

// SetUsername sets the "username" field.
func (pic *PasswordItemCreate) SetUsername(s string) *PasswordItemCreate {
	pic.mutation.SetUsername(s)
	return pic
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (pic *PasswordItemCreate) SetNillableUsername(s *string) *PasswordItemCreate {
	if s != nil {
		pic.SetUsername(*s)
	}
	return pic
}

// SetUsernameType sets the "username_type" field.
func (pic *PasswordItemCreate) SetUsernameType(s string) *PasswordItemCreate {
	pic.mutation.SetUsernameType(s)
	return pic
}

// SetNillableUsernameType sets the "username_type" field if the given value is not nil.
func (pic *PasswordItemCreate) SetNillableUsernameType(s *string) *PasswordItemCreate {
	if s != nil {
		pic.SetUsernameType(*s)
	}
	return pic
}

// SetPassword sets the "password" field.
func (pic *PasswordItemCreate) SetPassword(s string) *PasswordItemCreate {
	pic.mutation.SetPassword(s)
	return pic
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (pic *PasswordItemCreate) SetNillablePassword(s *string) *PasswordItemCreate {
	if s != nil {
		pic.SetPassword(*s)
	}
	return pic
}

// SetTags sets the "tags" field.
func (pic *PasswordItemCreate) SetTags(s []string) *PasswordItemCreate {
	pic.mutation.SetTags(s)
	return pic
}

// Mutation returns the PasswordItemMutation object of the builder.
func (pic *PasswordItemCreate) Mutation() *PasswordItemMutation {
	return pic.mutation
}

// Save creates the PasswordItem in the database.
func (pic *PasswordItemCreate) Save(ctx context.Context) (*PasswordItem, error) {
	var (
		err  error
		node *PasswordItem
	)
	pic.defaults()
	if len(pic.hooks) == 0 {
		if err = pic.check(); err != nil {
			return nil, err
		}
		node, err = pic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PasswordItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pic.check(); err != nil {
				return nil, err
			}
			pic.mutation = mutation
			if node, err = pic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pic.hooks) - 1; i >= 0; i-- {
			if pic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pic.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pic.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (pic *PasswordItemCreate) SaveX(ctx context.Context) *PasswordItem {
	v, err := pic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pic *PasswordItemCreate) Exec(ctx context.Context) error {
	_, err := pic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pic *PasswordItemCreate) ExecX(ctx context.Context) {
	if err := pic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pic *PasswordItemCreate) defaults() {
	if _, ok := pic.mutation.CreateTime(); !ok {
		v := passworditem.DefaultCreateTime()
		pic.mutation.SetCreateTime(v)
	}
	if _, ok := pic.mutation.UpdateTime(); !ok {
		v := passworditem.DefaultUpdateTime()
		pic.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pic *PasswordItemCreate) check() error {
	if _, ok := pic.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "PasswordItem.create_time"`)}
	}
	if _, ok := pic.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "PasswordItem.update_time"`)}
	}
	return nil
}

func (pic *PasswordItemCreate) sqlSave(ctx context.Context) (*PasswordItem, error) {
	_node, _spec := pic.createSpec()
	if err := sqlgraph.CreateNode(ctx, pic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pic *PasswordItemCreate) createSpec() (*PasswordItem, *sqlgraph.CreateSpec) {
	var (
		_node = &PasswordItem{config: pic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: passworditem.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: passworditem.FieldID,
			},
		}
	)
	if value, ok := pic.mutation.CreateTime(); ok {
		_spec.SetField(passworditem.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := pic.mutation.UpdateTime(); ok {
		_spec.SetField(passworditem.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := pic.mutation.Avatar(); ok {
		_spec.SetField(passworditem.FieldAvatar, field.TypeString, value)
		_node.Avatar = &value
	}
	if value, ok := pic.mutation.Description(); ok {
		_spec.SetField(passworditem.FieldDescription, field.TypeString, value)
		_node.Description = &value
	}
	if value, ok := pic.mutation.SiteName(); ok {
		_spec.SetField(passworditem.FieldSiteName, field.TypeString, value)
		_node.SiteName = &value
	}
	if value, ok := pic.mutation.SiteURL(); ok {
		_spec.SetField(passworditem.FieldSiteURL, field.TypeString, value)
		_node.SiteURL = &value
	}
	if value, ok := pic.mutation.Username(); ok {
		_spec.SetField(passworditem.FieldUsername, field.TypeString, value)
		_node.Username = &value
	}
	if value, ok := pic.mutation.UsernameType(); ok {
		_spec.SetField(passworditem.FieldUsernameType, field.TypeString, value)
		_node.UsernameType = &value
	}
	if value, ok := pic.mutation.Password(); ok {
		_spec.SetField(passworditem.FieldPassword, field.TypeString, value)
		_node.Password = &value
	}
	if value, ok := pic.mutation.Tags(); ok {
		_spec.SetField(passworditem.FieldTags, field.TypeJSON, value)
		_node.Tags = value
	}
	return _node, _spec
}

// PasswordItemCreateBulk is the builder for creating many PasswordItem entities in bulk.
type PasswordItemCreateBulk struct {
	config
	builders []*PasswordItemCreate
}

// Save creates the PasswordItem entities in the database.
func (picb *PasswordItemCreateBulk) Save(ctx context.Context) ([]*PasswordItem, error) {
	specs := make([]*sqlgraph.CreateSpec, len(picb.builders))
	nodes := make([]*PasswordItem, len(picb.builders))
	mutators := make([]Mutator, len(picb.builders))
	for i := range picb.builders {
		func(i int, root context.Context) {
			builder := picb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PasswordItemMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, picb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, picb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, picb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (picb *PasswordItemCreateBulk) SaveX(ctx context.Context) []*PasswordItem {
	v, err := picb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (picb *PasswordItemCreateBulk) Exec(ctx context.Context) error {
	_, err := picb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (picb *PasswordItemCreateBulk) ExecX(ctx context.Context) {
	if err := picb.Exec(ctx); err != nil {
		panic(err)
	}
}
