// Code generated by ent, DO NOT EDIT.

package ent

import (
	"agricoladb/ent/card"
	"agricoladb/ent/deck"
	"agricoladb/ent/predicate"
	"agricoladb/ent/product"
	"agricoladb/ent/revision"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RevisionUpdate is the builder for updating Revision entities.
type RevisionUpdate struct {
	config
	hooks    []Hook
	mutation *RevisionMutation
}

// Where appends a list predicates to the RevisionUpdate builder.
func (ru *RevisionUpdate) Where(ps ...predicate.Revision) *RevisionUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetNameJa sets the "name_ja" field.
func (ru *RevisionUpdate) SetNameJa(s string) *RevisionUpdate {
	ru.mutation.SetNameJa(s)
	return ru
}

// SetNillableNameJa sets the "name_ja" field if the given value is not nil.
func (ru *RevisionUpdate) SetNillableNameJa(s *string) *RevisionUpdate {
	if s != nil {
		ru.SetNameJa(*s)
	}
	return ru
}

// ClearNameJa clears the value of the "name_ja" field.
func (ru *RevisionUpdate) ClearNameJa() *RevisionUpdate {
	ru.mutation.ClearNameJa()
	return ru
}

// SetNameEn sets the "name_en" field.
func (ru *RevisionUpdate) SetNameEn(s string) *RevisionUpdate {
	ru.mutation.SetNameEn(s)
	return ru
}

// SetNillableNameEn sets the "name_en" field if the given value is not nil.
func (ru *RevisionUpdate) SetNillableNameEn(s *string) *RevisionUpdate {
	if s != nil {
		ru.SetNameEn(*s)
	}
	return ru
}

// ClearNameEn clears the value of the "name_en" field.
func (ru *RevisionUpdate) ClearNameEn() *RevisionUpdate {
	ru.mutation.ClearNameEn()
	return ru
}

// AddCardIDs adds the "cards" edge to the Card entity by IDs.
func (ru *RevisionUpdate) AddCardIDs(ids ...int) *RevisionUpdate {
	ru.mutation.AddCardIDs(ids...)
	return ru
}

// AddCards adds the "cards" edges to the Card entity.
func (ru *RevisionUpdate) AddCards(c ...*Card) *RevisionUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ru.AddCardIDs(ids...)
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (ru *RevisionUpdate) AddProductIDs(ids ...int) *RevisionUpdate {
	ru.mutation.AddProductIDs(ids...)
	return ru
}

// AddProducts adds the "products" edges to the Product entity.
func (ru *RevisionUpdate) AddProducts(p ...*Product) *RevisionUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.AddProductIDs(ids...)
}

// AddDeckIDs adds the "decks" edge to the Deck entity by IDs.
func (ru *RevisionUpdate) AddDeckIDs(ids ...int) *RevisionUpdate {
	ru.mutation.AddDeckIDs(ids...)
	return ru
}

// AddDecks adds the "decks" edges to the Deck entity.
func (ru *RevisionUpdate) AddDecks(d ...*Deck) *RevisionUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ru.AddDeckIDs(ids...)
}

// Mutation returns the RevisionMutation object of the builder.
func (ru *RevisionUpdate) Mutation() *RevisionMutation {
	return ru.mutation
}

// ClearCards clears all "cards" edges to the Card entity.
func (ru *RevisionUpdate) ClearCards() *RevisionUpdate {
	ru.mutation.ClearCards()
	return ru
}

// RemoveCardIDs removes the "cards" edge to Card entities by IDs.
func (ru *RevisionUpdate) RemoveCardIDs(ids ...int) *RevisionUpdate {
	ru.mutation.RemoveCardIDs(ids...)
	return ru
}

// RemoveCards removes "cards" edges to Card entities.
func (ru *RevisionUpdate) RemoveCards(c ...*Card) *RevisionUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ru.RemoveCardIDs(ids...)
}

// ClearProducts clears all "products" edges to the Product entity.
func (ru *RevisionUpdate) ClearProducts() *RevisionUpdate {
	ru.mutation.ClearProducts()
	return ru
}

// RemoveProductIDs removes the "products" edge to Product entities by IDs.
func (ru *RevisionUpdate) RemoveProductIDs(ids ...int) *RevisionUpdate {
	ru.mutation.RemoveProductIDs(ids...)
	return ru
}

// RemoveProducts removes "products" edges to Product entities.
func (ru *RevisionUpdate) RemoveProducts(p ...*Product) *RevisionUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.RemoveProductIDs(ids...)
}

// ClearDecks clears all "decks" edges to the Deck entity.
func (ru *RevisionUpdate) ClearDecks() *RevisionUpdate {
	ru.mutation.ClearDecks()
	return ru
}

// RemoveDeckIDs removes the "decks" edge to Deck entities by IDs.
func (ru *RevisionUpdate) RemoveDeckIDs(ids ...int) *RevisionUpdate {
	ru.mutation.RemoveDeckIDs(ids...)
	return ru
}

// RemoveDecks removes "decks" edges to Deck entities.
func (ru *RevisionUpdate) RemoveDecks(d ...*Deck) *RevisionUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ru.RemoveDeckIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RevisionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RevisionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RevisionUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RevisionUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RevisionUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *RevisionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   revision.Table,
			Columns: revision.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: revision.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.NameJa(); ok {
		_spec.SetField(revision.FieldNameJa, field.TypeString, value)
	}
	if ru.mutation.NameJaCleared() {
		_spec.ClearField(revision.FieldNameJa, field.TypeString)
	}
	if value, ok := ru.mutation.NameEn(); ok {
		_spec.SetField(revision.FieldNameEn, field.TypeString, value)
	}
	if ru.mutation.NameEnCleared() {
		_spec.ClearField(revision.FieldNameEn, field.TypeString)
	}
	if ru.mutation.CardsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.CardsTable,
			Columns: []string{revision.CardsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: card.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedCardsIDs(); len(nodes) > 0 && !ru.mutation.CardsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.CardsTable,
			Columns: []string{revision.CardsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.CardsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.CardsTable,
			Columns: []string{revision.CardsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.ProductsTable,
			Columns: []string{revision.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedProductsIDs(); len(nodes) > 0 && !ru.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.ProductsTable,
			Columns: []string{revision.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.ProductsTable,
			Columns: []string{revision.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.DecksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.DecksTable,
			Columns: []string{revision.DecksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: deck.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedDecksIDs(); len(nodes) > 0 && !ru.mutation.DecksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.DecksTable,
			Columns: []string{revision.DecksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: deck.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.DecksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.DecksTable,
			Columns: []string{revision.DecksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: deck.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{revision.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// RevisionUpdateOne is the builder for updating a single Revision entity.
type RevisionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RevisionMutation
}

// SetNameJa sets the "name_ja" field.
func (ruo *RevisionUpdateOne) SetNameJa(s string) *RevisionUpdateOne {
	ruo.mutation.SetNameJa(s)
	return ruo
}

// SetNillableNameJa sets the "name_ja" field if the given value is not nil.
func (ruo *RevisionUpdateOne) SetNillableNameJa(s *string) *RevisionUpdateOne {
	if s != nil {
		ruo.SetNameJa(*s)
	}
	return ruo
}

// ClearNameJa clears the value of the "name_ja" field.
func (ruo *RevisionUpdateOne) ClearNameJa() *RevisionUpdateOne {
	ruo.mutation.ClearNameJa()
	return ruo
}

// SetNameEn sets the "name_en" field.
func (ruo *RevisionUpdateOne) SetNameEn(s string) *RevisionUpdateOne {
	ruo.mutation.SetNameEn(s)
	return ruo
}

// SetNillableNameEn sets the "name_en" field if the given value is not nil.
func (ruo *RevisionUpdateOne) SetNillableNameEn(s *string) *RevisionUpdateOne {
	if s != nil {
		ruo.SetNameEn(*s)
	}
	return ruo
}

// ClearNameEn clears the value of the "name_en" field.
func (ruo *RevisionUpdateOne) ClearNameEn() *RevisionUpdateOne {
	ruo.mutation.ClearNameEn()
	return ruo
}

// AddCardIDs adds the "cards" edge to the Card entity by IDs.
func (ruo *RevisionUpdateOne) AddCardIDs(ids ...int) *RevisionUpdateOne {
	ruo.mutation.AddCardIDs(ids...)
	return ruo
}

// AddCards adds the "cards" edges to the Card entity.
func (ruo *RevisionUpdateOne) AddCards(c ...*Card) *RevisionUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ruo.AddCardIDs(ids...)
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (ruo *RevisionUpdateOne) AddProductIDs(ids ...int) *RevisionUpdateOne {
	ruo.mutation.AddProductIDs(ids...)
	return ruo
}

// AddProducts adds the "products" edges to the Product entity.
func (ruo *RevisionUpdateOne) AddProducts(p ...*Product) *RevisionUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.AddProductIDs(ids...)
}

// AddDeckIDs adds the "decks" edge to the Deck entity by IDs.
func (ruo *RevisionUpdateOne) AddDeckIDs(ids ...int) *RevisionUpdateOne {
	ruo.mutation.AddDeckIDs(ids...)
	return ruo
}

// AddDecks adds the "decks" edges to the Deck entity.
func (ruo *RevisionUpdateOne) AddDecks(d ...*Deck) *RevisionUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ruo.AddDeckIDs(ids...)
}

// Mutation returns the RevisionMutation object of the builder.
func (ruo *RevisionUpdateOne) Mutation() *RevisionMutation {
	return ruo.mutation
}

// ClearCards clears all "cards" edges to the Card entity.
func (ruo *RevisionUpdateOne) ClearCards() *RevisionUpdateOne {
	ruo.mutation.ClearCards()
	return ruo
}

// RemoveCardIDs removes the "cards" edge to Card entities by IDs.
func (ruo *RevisionUpdateOne) RemoveCardIDs(ids ...int) *RevisionUpdateOne {
	ruo.mutation.RemoveCardIDs(ids...)
	return ruo
}

// RemoveCards removes "cards" edges to Card entities.
func (ruo *RevisionUpdateOne) RemoveCards(c ...*Card) *RevisionUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ruo.RemoveCardIDs(ids...)
}

// ClearProducts clears all "products" edges to the Product entity.
func (ruo *RevisionUpdateOne) ClearProducts() *RevisionUpdateOne {
	ruo.mutation.ClearProducts()
	return ruo
}

// RemoveProductIDs removes the "products" edge to Product entities by IDs.
func (ruo *RevisionUpdateOne) RemoveProductIDs(ids ...int) *RevisionUpdateOne {
	ruo.mutation.RemoveProductIDs(ids...)
	return ruo
}

// RemoveProducts removes "products" edges to Product entities.
func (ruo *RevisionUpdateOne) RemoveProducts(p ...*Product) *RevisionUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.RemoveProductIDs(ids...)
}

// ClearDecks clears all "decks" edges to the Deck entity.
func (ruo *RevisionUpdateOne) ClearDecks() *RevisionUpdateOne {
	ruo.mutation.ClearDecks()
	return ruo
}

// RemoveDeckIDs removes the "decks" edge to Deck entities by IDs.
func (ruo *RevisionUpdateOne) RemoveDeckIDs(ids ...int) *RevisionUpdateOne {
	ruo.mutation.RemoveDeckIDs(ids...)
	return ruo
}

// RemoveDecks removes "decks" edges to Deck entities.
func (ruo *RevisionUpdateOne) RemoveDecks(d ...*Deck) *RevisionUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ruo.RemoveDeckIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RevisionUpdateOne) Select(field string, fields ...string) *RevisionUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Revision entity.
func (ruo *RevisionUpdateOne) Save(ctx context.Context) (*Revision, error) {
	var (
		err  error
		node *Revision
	)
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RevisionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ruo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Revision)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RevisionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RevisionUpdateOne) SaveX(ctx context.Context) *Revision {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RevisionUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RevisionUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *RevisionUpdateOne) sqlSave(ctx context.Context) (_node *Revision, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   revision.Table,
			Columns: revision.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: revision.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Revision.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, revision.FieldID)
		for _, f := range fields {
			if !revision.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != revision.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.NameJa(); ok {
		_spec.SetField(revision.FieldNameJa, field.TypeString, value)
	}
	if ruo.mutation.NameJaCleared() {
		_spec.ClearField(revision.FieldNameJa, field.TypeString)
	}
	if value, ok := ruo.mutation.NameEn(); ok {
		_spec.SetField(revision.FieldNameEn, field.TypeString, value)
	}
	if ruo.mutation.NameEnCleared() {
		_spec.ClearField(revision.FieldNameEn, field.TypeString)
	}
	if ruo.mutation.CardsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.CardsTable,
			Columns: []string{revision.CardsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: card.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedCardsIDs(); len(nodes) > 0 && !ruo.mutation.CardsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.CardsTable,
			Columns: []string{revision.CardsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.CardsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.CardsTable,
			Columns: []string{revision.CardsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.ProductsTable,
			Columns: []string{revision.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedProductsIDs(); len(nodes) > 0 && !ruo.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.ProductsTable,
			Columns: []string{revision.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.ProductsTable,
			Columns: []string{revision.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.DecksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.DecksTable,
			Columns: []string{revision.DecksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: deck.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedDecksIDs(); len(nodes) > 0 && !ruo.mutation.DecksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.DecksTable,
			Columns: []string{revision.DecksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: deck.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.DecksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.DecksTable,
			Columns: []string{revision.DecksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: deck.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Revision{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{revision.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
