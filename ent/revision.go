// Code generated by ent, DO NOT EDIT.

package ent

import (
	"agricoladb/ent/revision"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Revision is the model entity for the Revision schema.
type Revision struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Key holds the value of the "key" field.
	Key string `json:"key,omitempty"`
	// NameJa holds the value of the "name_ja" field.
	NameJa string `json:"name_ja,omitempty"`
	// NameEn holds the value of the "name_en" field.
	NameEn string `json:"name_en,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RevisionQuery when eager-loading is set.
	Edges RevisionEdges `json:"edges"`
}

// RevisionEdges holds the relations/edges for other nodes in the graph.
type RevisionEdges struct {
	// Cards holds the value of the cards edge.
	Cards []*Card `json:"cards,omitempty"`
	// Products holds the value of the products edge.
	Products []*Product `json:"products,omitempty"`
	// Decks holds the value of the decks edge.
	Decks []*Deck `json:"decks,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedCards    map[string][]*Card
	namedProducts map[string][]*Product
	namedDecks    map[string][]*Deck
}

// CardsOrErr returns the Cards value or an error if the edge
// was not loaded in eager-loading.
func (e RevisionEdges) CardsOrErr() ([]*Card, error) {
	if e.loadedTypes[0] {
		return e.Cards, nil
	}
	return nil, &NotLoadedError{edge: "cards"}
}

// ProductsOrErr returns the Products value or an error if the edge
// was not loaded in eager-loading.
func (e RevisionEdges) ProductsOrErr() ([]*Product, error) {
	if e.loadedTypes[1] {
		return e.Products, nil
	}
	return nil, &NotLoadedError{edge: "products"}
}

// DecksOrErr returns the Decks value or an error if the edge
// was not loaded in eager-loading.
func (e RevisionEdges) DecksOrErr() ([]*Deck, error) {
	if e.loadedTypes[2] {
		return e.Decks, nil
	}
	return nil, &NotLoadedError{edge: "decks"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Revision) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case revision.FieldID:
			values[i] = new(sql.NullInt64)
		case revision.FieldKey, revision.FieldNameJa, revision.FieldNameEn:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Revision", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Revision fields.
func (r *Revision) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case revision.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case revision.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				r.Key = value.String
			}
		case revision.FieldNameJa:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_ja", values[i])
			} else if value.Valid {
				r.NameJa = value.String
			}
		case revision.FieldNameEn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_en", values[i])
			} else if value.Valid {
				r.NameEn = value.String
			}
		}
	}
	return nil
}

// QueryCards queries the "cards" edge of the Revision entity.
func (r *Revision) QueryCards() *CardQuery {
	return (&RevisionClient{config: r.config}).QueryCards(r)
}

// QueryProducts queries the "products" edge of the Revision entity.
func (r *Revision) QueryProducts() *ProductQuery {
	return (&RevisionClient{config: r.config}).QueryProducts(r)
}

// QueryDecks queries the "decks" edge of the Revision entity.
func (r *Revision) QueryDecks() *DeckQuery {
	return (&RevisionClient{config: r.config}).QueryDecks(r)
}

// Update returns a builder for updating this Revision.
// Note that you need to call Revision.Unwrap() before calling this method if this Revision
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Revision) Update() *RevisionUpdateOne {
	return (&RevisionClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Revision entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Revision) Unwrap() *Revision {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Revision is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Revision) String() string {
	var builder strings.Builder
	builder.WriteString("Revision(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("key=")
	builder.WriteString(r.Key)
	builder.WriteString(", ")
	builder.WriteString("name_ja=")
	builder.WriteString(r.NameJa)
	builder.WriteString(", ")
	builder.WriteString("name_en=")
	builder.WriteString(r.NameEn)
	builder.WriteByte(')')
	return builder.String()
}

// NamedCards returns the Cards named value or an error if the edge was not
// loaded in eager-loading with this name.
func (r *Revision) NamedCards(name string) ([]*Card, error) {
	if r.Edges.namedCards == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := r.Edges.namedCards[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (r *Revision) appendNamedCards(name string, edges ...*Card) {
	if r.Edges.namedCards == nil {
		r.Edges.namedCards = make(map[string][]*Card)
	}
	if len(edges) == 0 {
		r.Edges.namedCards[name] = []*Card{}
	} else {
		r.Edges.namedCards[name] = append(r.Edges.namedCards[name], edges...)
	}
}

// NamedProducts returns the Products named value or an error if the edge was not
// loaded in eager-loading with this name.
func (r *Revision) NamedProducts(name string) ([]*Product, error) {
	if r.Edges.namedProducts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := r.Edges.namedProducts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (r *Revision) appendNamedProducts(name string, edges ...*Product) {
	if r.Edges.namedProducts == nil {
		r.Edges.namedProducts = make(map[string][]*Product)
	}
	if len(edges) == 0 {
		r.Edges.namedProducts[name] = []*Product{}
	} else {
		r.Edges.namedProducts[name] = append(r.Edges.namedProducts[name], edges...)
	}
}

// NamedDecks returns the Decks named value or an error if the edge was not
// loaded in eager-loading with this name.
func (r *Revision) NamedDecks(name string) ([]*Deck, error) {
	if r.Edges.namedDecks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := r.Edges.namedDecks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (r *Revision) appendNamedDecks(name string, edges ...*Deck) {
	if r.Edges.namedDecks == nil {
		r.Edges.namedDecks = make(map[string][]*Deck)
	}
	if len(edges) == 0 {
		r.Edges.namedDecks[name] = []*Deck{}
	} else {
		r.Edges.namedDecks[name] = append(r.Edges.namedDecks[name], edges...)
	}
}

// Revisions is a parsable slice of Revision.
type Revisions []*Revision

func (r Revisions) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
