// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/AgricolaDevJP/agricoladb-server/ent/card"
	"github.com/AgricolaDevJP/agricoladb-server/ent/cardspecialcolor"
	"github.com/AgricolaDevJP/agricoladb-server/ent/cardtype"
	"github.com/AgricolaDevJP/agricoladb-server/ent/deck"
	"github.com/AgricolaDevJP/agricoladb-server/ent/predicate"
	"github.com/AgricolaDevJP/agricoladb-server/ent/product"
	"github.com/AgricolaDevJP/agricoladb-server/ent/revision"
)

// CardQuery is the builder for querying Card entities.
type CardQuery struct {
	config
	ctx                  *QueryContext
	order                []OrderFunc
	inters               []Interceptor
	predicates           []predicate.Card
	withRevision         *RevisionQuery
	withProducts         *ProductQuery
	withDeck             *DeckQuery
	withCardType         *CardTypeQuery
	withCardSpecialColor *CardSpecialColorQuery
	withChildren         *CardQuery
	withAncestors        *CardQuery
	modifiers            []func(*sql.Selector)
	loadTotal            []func(context.Context, []*Card) error
	withNamedProducts    map[string]*ProductQuery
	withNamedChildren    map[string]*CardQuery
	withNamedAncestors   map[string]*CardQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CardQuery builder.
func (cq *CardQuery) Where(ps ...predicate.Card) *CardQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CardQuery) Limit(limit int) *CardQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *CardQuery) Offset(offset int) *CardQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CardQuery) Unique(unique bool) *CardQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CardQuery) Order(o ...OrderFunc) *CardQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryRevision chains the current query on the "revision" edge.
func (cq *CardQuery) QueryRevision() *RevisionQuery {
	query := (&RevisionClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(card.Table, card.FieldID, selector),
			sqlgraph.To(revision.Table, revision.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, card.RevisionTable, card.RevisionColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProducts chains the current query on the "products" edge.
func (cq *CardQuery) QueryProducts() *ProductQuery {
	query := (&ProductClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(card.Table, card.FieldID, selector),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, card.ProductsTable, card.ProductsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDeck chains the current query on the "deck" edge.
func (cq *CardQuery) QueryDeck() *DeckQuery {
	query := (&DeckClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(card.Table, card.FieldID, selector),
			sqlgraph.To(deck.Table, deck.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, card.DeckTable, card.DeckColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCardType chains the current query on the "card_type" edge.
func (cq *CardQuery) QueryCardType() *CardTypeQuery {
	query := (&CardTypeClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(card.Table, card.FieldID, selector),
			sqlgraph.To(cardtype.Table, cardtype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, card.CardTypeTable, card.CardTypeColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCardSpecialColor chains the current query on the "card_special_color" edge.
func (cq *CardQuery) QueryCardSpecialColor() *CardSpecialColorQuery {
	query := (&CardSpecialColorClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(card.Table, card.FieldID, selector),
			sqlgraph.To(cardspecialcolor.Table, cardspecialcolor.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, card.CardSpecialColorTable, card.CardSpecialColorColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildren chains the current query on the "children" edge.
func (cq *CardQuery) QueryChildren() *CardQuery {
	query := (&CardClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(card.Table, card.FieldID, selector),
			sqlgraph.To(card.Table, card.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, card.ChildrenTable, card.ChildrenPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAncestors chains the current query on the "ancestors" edge.
func (cq *CardQuery) QueryAncestors() *CardQuery {
	query := (&CardClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(card.Table, card.FieldID, selector),
			sqlgraph.To(card.Table, card.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, card.AncestorsTable, card.AncestorsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Card entity from the query.
// Returns a *NotFoundError when no Card was found.
func (cq *CardQuery) First(ctx context.Context) (*Card, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{card.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CardQuery) FirstX(ctx context.Context) *Card {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Card ID from the query.
// Returns a *NotFoundError when no Card ID was found.
func (cq *CardQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{card.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CardQuery) FirstIDX(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Card entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Card entity is found.
// Returns a *NotFoundError when no Card entities are found.
func (cq *CardQuery) Only(ctx context.Context) (*Card, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{card.Label}
	default:
		return nil, &NotSingularError{card.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CardQuery) OnlyX(ctx context.Context) *Card {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Card ID in the query.
// Returns a *NotSingularError when more than one Card ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CardQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{card.Label}
	default:
		err = &NotSingularError{card.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CardQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Cards.
func (cq *CardQuery) All(ctx context.Context) ([]*Card, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Card, *CardQuery]()
	return withInterceptors[[]*Card](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CardQuery) AllX(ctx context.Context) []*Card {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Card IDs.
func (cq *CardQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err := cq.Select(card.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CardQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CardQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CardQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CardQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CardQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CardQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CardQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CardQuery) Clone() *CardQuery {
	if cq == nil {
		return nil
	}
	return &CardQuery{
		config:               cq.config,
		ctx:                  cq.ctx.Clone(),
		order:                append([]OrderFunc{}, cq.order...),
		inters:               append([]Interceptor{}, cq.inters...),
		predicates:           append([]predicate.Card{}, cq.predicates...),
		withRevision:         cq.withRevision.Clone(),
		withProducts:         cq.withProducts.Clone(),
		withDeck:             cq.withDeck.Clone(),
		withCardType:         cq.withCardType.Clone(),
		withCardSpecialColor: cq.withCardSpecialColor.Clone(),
		withChildren:         cq.withChildren.Clone(),
		withAncestors:        cq.withAncestors.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithRevision tells the query-builder to eager-load the nodes that are connected to
// the "revision" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CardQuery) WithRevision(opts ...func(*RevisionQuery)) *CardQuery {
	query := (&RevisionClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withRevision = query
	return cq
}

// WithProducts tells the query-builder to eager-load the nodes that are connected to
// the "products" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CardQuery) WithProducts(opts ...func(*ProductQuery)) *CardQuery {
	query := (&ProductClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withProducts = query
	return cq
}

// WithDeck tells the query-builder to eager-load the nodes that are connected to
// the "deck" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CardQuery) WithDeck(opts ...func(*DeckQuery)) *CardQuery {
	query := (&DeckClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withDeck = query
	return cq
}

// WithCardType tells the query-builder to eager-load the nodes that are connected to
// the "card_type" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CardQuery) WithCardType(opts ...func(*CardTypeQuery)) *CardQuery {
	query := (&CardTypeClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withCardType = query
	return cq
}

// WithCardSpecialColor tells the query-builder to eager-load the nodes that are connected to
// the "card_special_color" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CardQuery) WithCardSpecialColor(opts ...func(*CardSpecialColorQuery)) *CardQuery {
	query := (&CardSpecialColorClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withCardSpecialColor = query
	return cq
}

// WithChildren tells the query-builder to eager-load the nodes that are connected to
// the "children" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CardQuery) WithChildren(opts ...func(*CardQuery)) *CardQuery {
	query := (&CardClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withChildren = query
	return cq
}

// WithAncestors tells the query-builder to eager-load the nodes that are connected to
// the "ancestors" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CardQuery) WithAncestors(opts ...func(*CardQuery)) *CardQuery {
	query := (&CardClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withAncestors = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		LiteralID string `json:"literal_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Card.Query().
//		GroupBy(card.FieldLiteralID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *CardQuery) GroupBy(field string, fields ...string) *CardGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CardGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = card.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		LiteralID string `json:"literal_id,omitempty"`
//	}
//
//	client.Card.Query().
//		Select(card.FieldLiteralID).
//		Scan(ctx, &v)
func (cq *CardQuery) Select(fields ...string) *CardSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &CardSelect{CardQuery: cq}
	sbuild.label = card.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CardSelect configured with the given aggregations.
func (cq *CardQuery) Aggregate(fns ...AggregateFunc) *CardSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CardQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !card.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *CardQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Card, error) {
	var (
		nodes       = []*Card{}
		_spec       = cq.querySpec()
		loadedTypes = [7]bool{
			cq.withRevision != nil,
			cq.withProducts != nil,
			cq.withDeck != nil,
			cq.withCardType != nil,
			cq.withCardSpecialColor != nil,
			cq.withChildren != nil,
			cq.withAncestors != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Card).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Card{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withRevision; query != nil {
		if err := cq.loadRevision(ctx, query, nodes, nil,
			func(n *Card, e *Revision) { n.Edges.Revision = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withProducts; query != nil {
		if err := cq.loadProducts(ctx, query, nodes,
			func(n *Card) { n.Edges.Products = []*Product{} },
			func(n *Card, e *Product) { n.Edges.Products = append(n.Edges.Products, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withDeck; query != nil {
		if err := cq.loadDeck(ctx, query, nodes, nil,
			func(n *Card, e *Deck) { n.Edges.Deck = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withCardType; query != nil {
		if err := cq.loadCardType(ctx, query, nodes, nil,
			func(n *Card, e *CardType) { n.Edges.CardType = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withCardSpecialColor; query != nil {
		if err := cq.loadCardSpecialColor(ctx, query, nodes, nil,
			func(n *Card, e *CardSpecialColor) { n.Edges.CardSpecialColor = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withChildren; query != nil {
		if err := cq.loadChildren(ctx, query, nodes,
			func(n *Card) { n.Edges.Children = []*Card{} },
			func(n *Card, e *Card) { n.Edges.Children = append(n.Edges.Children, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withAncestors; query != nil {
		if err := cq.loadAncestors(ctx, query, nodes,
			func(n *Card) { n.Edges.Ancestors = []*Card{} },
			func(n *Card, e *Card) { n.Edges.Ancestors = append(n.Edges.Ancestors, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range cq.withNamedProducts {
		if err := cq.loadProducts(ctx, query, nodes,
			func(n *Card) { n.appendNamedProducts(name) },
			func(n *Card, e *Product) { n.appendNamedProducts(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range cq.withNamedChildren {
		if err := cq.loadChildren(ctx, query, nodes,
			func(n *Card) { n.appendNamedChildren(name) },
			func(n *Card, e *Card) { n.appendNamedChildren(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range cq.withNamedAncestors {
		if err := cq.loadAncestors(ctx, query, nodes,
			func(n *Card) { n.appendNamedAncestors(name) },
			func(n *Card, e *Card) { n.appendNamedAncestors(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range cq.loadTotal {
		if err := cq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *CardQuery) loadRevision(ctx context.Context, query *RevisionQuery, nodes []*Card, init func(*Card), assign func(*Card, *Revision)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Card)
	for i := range nodes {
		fk := nodes[i].RevisionID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(revision.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "revision_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *CardQuery) loadProducts(ctx context.Context, query *ProductQuery, nodes []*Card, init func(*Card), assign func(*Card, *Product)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Card)
	nids := make(map[int]map[*Card]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(card.ProductsTable)
		s.Join(joinT).On(s.C(product.FieldID), joinT.C(card.ProductsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(card.ProductsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(card.ProductsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := int(values[0].(*sql.NullInt64).Int64)
			inValue := int(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*Card]struct{}{byID[outValue]: {}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "products" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (cq *CardQuery) loadDeck(ctx context.Context, query *DeckQuery, nodes []*Card, init func(*Card), assign func(*Card, *Deck)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Card)
	for i := range nodes {
		fk := nodes[i].DeckID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(deck.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "deck_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *CardQuery) loadCardType(ctx context.Context, query *CardTypeQuery, nodes []*Card, init func(*Card), assign func(*Card, *CardType)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Card)
	for i := range nodes {
		fk := nodes[i].CardTypeID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(cardtype.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "card_type_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *CardQuery) loadCardSpecialColor(ctx context.Context, query *CardSpecialColorQuery, nodes []*Card, init func(*Card), assign func(*Card, *CardSpecialColor)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Card)
	for i := range nodes {
		fk := nodes[i].CardSpecialColorID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(cardspecialcolor.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "card_special_color_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *CardQuery) loadChildren(ctx context.Context, query *CardQuery, nodes []*Card, init func(*Card), assign func(*Card, *Card)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Card)
	nids := make(map[int]map[*Card]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(card.ChildrenTable)
		s.Join(joinT).On(s.C(card.FieldID), joinT.C(card.ChildrenPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(card.ChildrenPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(card.ChildrenPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := int(values[0].(*sql.NullInt64).Int64)
			inValue := int(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*Card]struct{}{byID[outValue]: {}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "children" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (cq *CardQuery) loadAncestors(ctx context.Context, query *CardQuery, nodes []*Card, init func(*Card), assign func(*Card, *Card)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Card)
	nids := make(map[int]map[*Card]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(card.AncestorsTable)
		s.Join(joinT).On(s.C(card.FieldID), joinT.C(card.AncestorsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(card.AncestorsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(card.AncestorsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := int(values[0].(*sql.NullInt64).Int64)
			inValue := int(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*Card]struct{}{byID[outValue]: {}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "ancestors" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (cq *CardQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CardQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   card.Table,
			Columns: card.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: card.FieldID,
			},
		},
		From:   cq.sql,
		Unique: true,
	}
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, card.FieldID)
		for i := range fields {
			if fields[i] != card.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CardQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(card.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = card.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedProducts tells the query-builder to eager-load the nodes that are connected to the "products"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (cq *CardQuery) WithNamedProducts(name string, opts ...func(*ProductQuery)) *CardQuery {
	query := (&ProductClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if cq.withNamedProducts == nil {
		cq.withNamedProducts = make(map[string]*ProductQuery)
	}
	cq.withNamedProducts[name] = query
	return cq
}

// WithNamedChildren tells the query-builder to eager-load the nodes that are connected to the "children"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (cq *CardQuery) WithNamedChildren(name string, opts ...func(*CardQuery)) *CardQuery {
	query := (&CardClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if cq.withNamedChildren == nil {
		cq.withNamedChildren = make(map[string]*CardQuery)
	}
	cq.withNamedChildren[name] = query
	return cq
}

// WithNamedAncestors tells the query-builder to eager-load the nodes that are connected to the "ancestors"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (cq *CardQuery) WithNamedAncestors(name string, opts ...func(*CardQuery)) *CardQuery {
	query := (&CardClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if cq.withNamedAncestors == nil {
		cq.withNamedAncestors = make(map[string]*CardQuery)
	}
	cq.withNamedAncestors[name] = query
	return cq
}

// CardGroupBy is the group-by builder for Card entities.
type CardGroupBy struct {
	selector
	build *CardQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CardGroupBy) Aggregate(fns ...AggregateFunc) *CardGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CardGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CardQuery, *CardGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CardGroupBy) sqlScan(ctx context.Context, root *CardQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CardSelect is the builder for selecting fields of Card entities.
type CardSelect struct {
	*CardQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CardSelect) Aggregate(fns ...AggregateFunc) *CardSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CardSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CardQuery, *CardSelect](ctx, cs.CardQuery, cs, cs.inters, v)
}

func (cs *CardSelect) sqlScan(ctx context.Context, root *CardQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
