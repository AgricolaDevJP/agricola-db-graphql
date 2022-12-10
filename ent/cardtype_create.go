// Code generated by ent, DO NOT EDIT.

package ent

import (
	"agricoladb/ent/card"
	"agricoladb/ent/cardtype"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CardTypeCreate is the builder for creating a CardType entity.
type CardTypeCreate struct {
	config
	mutation *CardTypeMutation
	hooks    []Hook
}

// SetKey sets the "key" field.
func (ctc *CardTypeCreate) SetKey(s string) *CardTypeCreate {
	ctc.mutation.SetKey(s)
	return ctc
}

// SetNameJa sets the "name_ja" field.
func (ctc *CardTypeCreate) SetNameJa(s string) *CardTypeCreate {
	ctc.mutation.SetNameJa(s)
	return ctc
}

// SetNillableNameJa sets the "name_ja" field if the given value is not nil.
func (ctc *CardTypeCreate) SetNillableNameJa(s *string) *CardTypeCreate {
	if s != nil {
		ctc.SetNameJa(*s)
	}
	return ctc
}

// SetNameEn sets the "name_en" field.
func (ctc *CardTypeCreate) SetNameEn(s string) *CardTypeCreate {
	ctc.mutation.SetNameEn(s)
	return ctc
}

// SetNillableNameEn sets the "name_en" field if the given value is not nil.
func (ctc *CardTypeCreate) SetNillableNameEn(s *string) *CardTypeCreate {
	if s != nil {
		ctc.SetNameEn(*s)
	}
	return ctc
}

// AddCardIDs adds the "cards" edge to the Card entity by IDs.
func (ctc *CardTypeCreate) AddCardIDs(ids ...int) *CardTypeCreate {
	ctc.mutation.AddCardIDs(ids...)
	return ctc
}

// AddCards adds the "cards" edges to the Card entity.
func (ctc *CardTypeCreate) AddCards(c ...*Card) *CardTypeCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ctc.AddCardIDs(ids...)
}

// Mutation returns the CardTypeMutation object of the builder.
func (ctc *CardTypeCreate) Mutation() *CardTypeMutation {
	return ctc.mutation
}

// Save creates the CardType in the database.
func (ctc *CardTypeCreate) Save(ctx context.Context) (*CardType, error) {
	var (
		err  error
		node *CardType
	)
	if len(ctc.hooks) == 0 {
		if err = ctc.check(); err != nil {
			return nil, err
		}
		node, err = ctc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CardTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ctc.check(); err != nil {
				return nil, err
			}
			ctc.mutation = mutation
			if node, err = ctc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ctc.hooks) - 1; i >= 0; i-- {
			if ctc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ctc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ctc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CardType)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CardTypeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ctc *CardTypeCreate) SaveX(ctx context.Context) *CardType {
	v, err := ctc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ctc *CardTypeCreate) Exec(ctx context.Context) error {
	_, err := ctc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctc *CardTypeCreate) ExecX(ctx context.Context) {
	if err := ctc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ctc *CardTypeCreate) check() error {
	if _, ok := ctc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`ent: missing required field "CardType.key"`)}
	}
	if v, ok := ctc.mutation.Key(); ok {
		if err := cardtype.KeyValidator(v); err != nil {
			return &ValidationError{Name: "key", err: fmt.Errorf(`ent: validator failed for field "CardType.key": %w`, err)}
		}
	}
	return nil
}

func (ctc *CardTypeCreate) sqlSave(ctx context.Context) (*CardType, error) {
	_node, _spec := ctc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ctc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ctc *CardTypeCreate) createSpec() (*CardType, *sqlgraph.CreateSpec) {
	var (
		_node = &CardType{config: ctc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: cardtype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cardtype.FieldID,
			},
		}
	)
	if value, ok := ctc.mutation.Key(); ok {
		_spec.SetField(cardtype.FieldKey, field.TypeString, value)
		_node.Key = value
	}
	if value, ok := ctc.mutation.NameJa(); ok {
		_spec.SetField(cardtype.FieldNameJa, field.TypeString, value)
		_node.NameJa = value
	}
	if value, ok := ctc.mutation.NameEn(); ok {
		_spec.SetField(cardtype.FieldNameEn, field.TypeString, value)
		_node.NameEn = value
	}
	if nodes := ctc.mutation.CardsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cardtype.CardsTable,
			Columns: []string{cardtype.CardsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: card.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CardTypeCreateBulk is the builder for creating many CardType entities in bulk.
type CardTypeCreateBulk struct {
	config
	builders []*CardTypeCreate
}

// Save creates the CardType entities in the database.
func (ctcb *CardTypeCreateBulk) Save(ctx context.Context) ([]*CardType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ctcb.builders))
	nodes := make([]*CardType, len(ctcb.builders))
	mutators := make([]Mutator, len(ctcb.builders))
	for i := range ctcb.builders {
		func(i int, root context.Context) {
			builder := ctcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CardTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, ctcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ctcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ctcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ctcb *CardTypeCreateBulk) SaveX(ctx context.Context) []*CardType {
	v, err := ctcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ctcb *CardTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := ctcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctcb *CardTypeCreateBulk) ExecX(ctx context.Context) {
	if err := ctcb.Exec(ctx); err != nil {
		panic(err)
	}
}
