// Code generated by ent, DO NOT EDIT.

package ent

import (
	"agricoladb/ent/card"
	"agricoladb/ent/cardtype"
	"agricoladb/ent/predicate"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CardTypeQuery is the builder for querying CardType entities.
type CardTypeQuery struct {
	config
	limit          *int
	offset         *int
	unique         *bool
	order          []OrderFunc
	fields         []string
	predicates     []predicate.CardType
	withCards      *CardQuery
	modifiers      []func(*sql.Selector)
	loadTotal      []func(context.Context, []*CardType) error
	withNamedCards map[string]*CardQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CardTypeQuery builder.
func (ctq *CardTypeQuery) Where(ps ...predicate.CardType) *CardTypeQuery {
	ctq.predicates = append(ctq.predicates, ps...)
	return ctq
}

// Limit adds a limit step to the query.
func (ctq *CardTypeQuery) Limit(limit int) *CardTypeQuery {
	ctq.limit = &limit
	return ctq
}

// Offset adds an offset step to the query.
func (ctq *CardTypeQuery) Offset(offset int) *CardTypeQuery {
	ctq.offset = &offset
	return ctq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ctq *CardTypeQuery) Unique(unique bool) *CardTypeQuery {
	ctq.unique = &unique
	return ctq
}

// Order adds an order step to the query.
func (ctq *CardTypeQuery) Order(o ...OrderFunc) *CardTypeQuery {
	ctq.order = append(ctq.order, o...)
	return ctq
}

// QueryCards chains the current query on the "cards" edge.
func (ctq *CardTypeQuery) QueryCards() *CardQuery {
	query := &CardQuery{config: ctq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ctq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ctq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(cardtype.Table, cardtype.FieldID, selector),
			sqlgraph.To(card.Table, card.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, cardtype.CardsTable, cardtype.CardsColumn),
		)
		fromU = sqlgraph.SetNeighbors(ctq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CardType entity from the query.
// Returns a *NotFoundError when no CardType was found.
func (ctq *CardTypeQuery) First(ctx context.Context) (*CardType, error) {
	nodes, err := ctq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{cardtype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ctq *CardTypeQuery) FirstX(ctx context.Context) *CardType {
	node, err := ctq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CardType ID from the query.
// Returns a *NotFoundError when no CardType ID was found.
func (ctq *CardTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ctq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{cardtype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ctq *CardTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := ctq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CardType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CardType entity is found.
// Returns a *NotFoundError when no CardType entities are found.
func (ctq *CardTypeQuery) Only(ctx context.Context) (*CardType, error) {
	nodes, err := ctq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{cardtype.Label}
	default:
		return nil, &NotSingularError{cardtype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ctq *CardTypeQuery) OnlyX(ctx context.Context) *CardType {
	node, err := ctq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CardType ID in the query.
// Returns a *NotSingularError when more than one CardType ID is found.
// Returns a *NotFoundError when no entities are found.
func (ctq *CardTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ctq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{cardtype.Label}
	default:
		err = &NotSingularError{cardtype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ctq *CardTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := ctq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CardTypes.
func (ctq *CardTypeQuery) All(ctx context.Context) ([]*CardType, error) {
	if err := ctq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ctq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ctq *CardTypeQuery) AllX(ctx context.Context) []*CardType {
	nodes, err := ctq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CardType IDs.
func (ctq *CardTypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ctq.Select(cardtype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ctq *CardTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := ctq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ctq *CardTypeQuery) Count(ctx context.Context) (int, error) {
	if err := ctq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ctq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ctq *CardTypeQuery) CountX(ctx context.Context) int {
	count, err := ctq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ctq *CardTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := ctq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ctq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ctq *CardTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := ctq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CardTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ctq *CardTypeQuery) Clone() *CardTypeQuery {
	if ctq == nil {
		return nil
	}
	return &CardTypeQuery{
		config:     ctq.config,
		limit:      ctq.limit,
		offset:     ctq.offset,
		order:      append([]OrderFunc{}, ctq.order...),
		predicates: append([]predicate.CardType{}, ctq.predicates...),
		withCards:  ctq.withCards.Clone(),
		// clone intermediate query.
		sql:    ctq.sql.Clone(),
		path:   ctq.path,
		unique: ctq.unique,
	}
}

// WithCards tells the query-builder to eager-load the nodes that are connected to
// the "cards" edge. The optional arguments are used to configure the query builder of the edge.
func (ctq *CardTypeQuery) WithCards(opts ...func(*CardQuery)) *CardTypeQuery {
	query := &CardQuery{config: ctq.config}
	for _, opt := range opts {
		opt(query)
	}
	ctq.withCards = query
	return ctq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Key string `json:"key,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CardType.Query().
//		GroupBy(cardtype.FieldKey).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ctq *CardTypeQuery) GroupBy(field string, fields ...string) *CardTypeGroupBy {
	grbuild := &CardTypeGroupBy{config: ctq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ctq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ctq.sqlQuery(ctx), nil
	}
	grbuild.label = cardtype.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Key string `json:"key,omitempty"`
//	}
//
//	client.CardType.Query().
//		Select(cardtype.FieldKey).
//		Scan(ctx, &v)
func (ctq *CardTypeQuery) Select(fields ...string) *CardTypeSelect {
	ctq.fields = append(ctq.fields, fields...)
	selbuild := &CardTypeSelect{CardTypeQuery: ctq}
	selbuild.label = cardtype.Label
	selbuild.flds, selbuild.scan = &ctq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a CardTypeSelect configured with the given aggregations.
func (ctq *CardTypeQuery) Aggregate(fns ...AggregateFunc) *CardTypeSelect {
	return ctq.Select().Aggregate(fns...)
}

func (ctq *CardTypeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ctq.fields {
		if !cardtype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ctq.path != nil {
		prev, err := ctq.path(ctx)
		if err != nil {
			return err
		}
		ctq.sql = prev
	}
	return nil
}

func (ctq *CardTypeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CardType, error) {
	var (
		nodes       = []*CardType{}
		_spec       = ctq.querySpec()
		loadedTypes = [1]bool{
			ctq.withCards != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CardType).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CardType{config: ctq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(ctq.modifiers) > 0 {
		_spec.Modifiers = ctq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ctq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ctq.withCards; query != nil {
		if err := ctq.loadCards(ctx, query, nodes,
			func(n *CardType) { n.Edges.Cards = []*Card{} },
			func(n *CardType, e *Card) { n.Edges.Cards = append(n.Edges.Cards, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range ctq.withNamedCards {
		if err := ctq.loadCards(ctx, query, nodes,
			func(n *CardType) { n.appendNamedCards(name) },
			func(n *CardType, e *Card) { n.appendNamedCards(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range ctq.loadTotal {
		if err := ctq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ctq *CardTypeQuery) loadCards(ctx context.Context, query *CardQuery, nodes []*CardType, init func(*CardType), assign func(*CardType, *Card)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*CardType)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.Card(func(s *sql.Selector) {
		s.Where(sql.InValues(cardtype.CardsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.CardTypeID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "card_type_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (ctq *CardTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ctq.querySpec()
	if len(ctq.modifiers) > 0 {
		_spec.Modifiers = ctq.modifiers
	}
	_spec.Node.Columns = ctq.fields
	if len(ctq.fields) > 0 {
		_spec.Unique = ctq.unique != nil && *ctq.unique
	}
	return sqlgraph.CountNodes(ctx, ctq.driver, _spec)
}

func (ctq *CardTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := ctq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (ctq *CardTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cardtype.Table,
			Columns: cardtype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cardtype.FieldID,
			},
		},
		From:   ctq.sql,
		Unique: true,
	}
	if unique := ctq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ctq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cardtype.FieldID)
		for i := range fields {
			if fields[i] != cardtype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ctq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ctq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ctq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ctq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ctq *CardTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ctq.driver.Dialect())
	t1 := builder.Table(cardtype.Table)
	columns := ctq.fields
	if len(columns) == 0 {
		columns = cardtype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ctq.sql != nil {
		selector = ctq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ctq.unique != nil && *ctq.unique {
		selector.Distinct()
	}
	for _, p := range ctq.predicates {
		p(selector)
	}
	for _, p := range ctq.order {
		p(selector)
	}
	if offset := ctq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ctq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedCards tells the query-builder to eager-load the nodes that are connected to the "cards"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (ctq *CardTypeQuery) WithNamedCards(name string, opts ...func(*CardQuery)) *CardTypeQuery {
	query := &CardQuery{config: ctq.config}
	for _, opt := range opts {
		opt(query)
	}
	if ctq.withNamedCards == nil {
		ctq.withNamedCards = make(map[string]*CardQuery)
	}
	ctq.withNamedCards[name] = query
	return ctq
}

// CardTypeGroupBy is the group-by builder for CardType entities.
type CardTypeGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ctgb *CardTypeGroupBy) Aggregate(fns ...AggregateFunc) *CardTypeGroupBy {
	ctgb.fns = append(ctgb.fns, fns...)
	return ctgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ctgb *CardTypeGroupBy) Scan(ctx context.Context, v any) error {
	query, err := ctgb.path(ctx)
	if err != nil {
		return err
	}
	ctgb.sql = query
	return ctgb.sqlScan(ctx, v)
}

func (ctgb *CardTypeGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range ctgb.fields {
		if !cardtype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ctgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ctgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ctgb *CardTypeGroupBy) sqlQuery() *sql.Selector {
	selector := ctgb.sql.Select()
	aggregation := make([]string, 0, len(ctgb.fns))
	for _, fn := range ctgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ctgb.fields)+len(ctgb.fns))
		for _, f := range ctgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ctgb.fields...)...)
}

// CardTypeSelect is the builder for selecting fields of CardType entities.
type CardTypeSelect struct {
	*CardTypeQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cts *CardTypeSelect) Aggregate(fns ...AggregateFunc) *CardTypeSelect {
	cts.fns = append(cts.fns, fns...)
	return cts
}

// Scan applies the selector query and scans the result into the given value.
func (cts *CardTypeSelect) Scan(ctx context.Context, v any) error {
	if err := cts.prepareQuery(ctx); err != nil {
		return err
	}
	cts.sql = cts.CardTypeQuery.sqlQuery(ctx)
	return cts.sqlScan(ctx, v)
}

func (cts *CardTypeSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(cts.fns))
	for _, fn := range cts.fns {
		aggregation = append(aggregation, fn(cts.sql))
	}
	switch n := len(*cts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		cts.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		cts.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := cts.sql.Query()
	if err := cts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
