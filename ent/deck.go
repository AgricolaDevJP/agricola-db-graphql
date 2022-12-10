// Code generated by ent, DO NOT EDIT.

package ent

import (
	"agricoladb/ent/deck"
	"agricoladb/ent/revision"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Deck is the model entity for the Deck schema.
type Deck struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Key holds the value of the "key" field.
	Key string `json:"key,omitempty"`
	// RevisionID holds the value of the "revision_id" field.
	RevisionID int `json:"revision_id,omitempty"`
	// NameJa holds the value of the "name_ja" field.
	NameJa string `json:"name_ja,omitempty"`
	// NameEn holds the value of the "name_en" field.
	NameEn string `json:"name_en,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DeckQuery when eager-loading is set.
	Edges DeckEdges `json:"edges"`
}

// DeckEdges holds the relations/edges for other nodes in the graph.
type DeckEdges struct {
	// Cards holds the value of the cards edge.
	Cards []*Card `json:"cards,omitempty"`
	// Revision holds the value of the revision edge.
	Revision *Revision `json:"revision,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedCards map[string][]*Card
}

// CardsOrErr returns the Cards value or an error if the edge
// was not loaded in eager-loading.
func (e DeckEdges) CardsOrErr() ([]*Card, error) {
	if e.loadedTypes[0] {
		return e.Cards, nil
	}
	return nil, &NotLoadedError{edge: "cards"}
}

// RevisionOrErr returns the Revision value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeckEdges) RevisionOrErr() (*Revision, error) {
	if e.loadedTypes[1] {
		if e.Revision == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: revision.Label}
		}
		return e.Revision, nil
	}
	return nil, &NotLoadedError{edge: "revision"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Deck) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case deck.FieldID, deck.FieldRevisionID:
			values[i] = new(sql.NullInt64)
		case deck.FieldKey, deck.FieldNameJa, deck.FieldNameEn:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Deck", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Deck fields.
func (d *Deck) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case deck.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = int(value.Int64)
		case deck.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				d.Key = value.String
			}
		case deck.FieldRevisionID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field revision_id", values[i])
			} else if value.Valid {
				d.RevisionID = int(value.Int64)
			}
		case deck.FieldNameJa:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_ja", values[i])
			} else if value.Valid {
				d.NameJa = value.String
			}
		case deck.FieldNameEn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_en", values[i])
			} else if value.Valid {
				d.NameEn = value.String
			}
		}
	}
	return nil
}

// QueryCards queries the "cards" edge of the Deck entity.
func (d *Deck) QueryCards() *CardQuery {
	return (&DeckClient{config: d.config}).QueryCards(d)
}

// QueryRevision queries the "revision" edge of the Deck entity.
func (d *Deck) QueryRevision() *RevisionQuery {
	return (&DeckClient{config: d.config}).QueryRevision(d)
}

// Update returns a builder for updating this Deck.
// Note that you need to call Deck.Unwrap() before calling this method if this Deck
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Deck) Update() *DeckUpdateOne {
	return (&DeckClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the Deck entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Deck) Unwrap() *Deck {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Deck is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Deck) String() string {
	var builder strings.Builder
	builder.WriteString("Deck(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("key=")
	builder.WriteString(d.Key)
	builder.WriteString(", ")
	builder.WriteString("revision_id=")
	builder.WriteString(fmt.Sprintf("%v", d.RevisionID))
	builder.WriteString(", ")
	builder.WriteString("name_ja=")
	builder.WriteString(d.NameJa)
	builder.WriteString(", ")
	builder.WriteString("name_en=")
	builder.WriteString(d.NameEn)
	builder.WriteByte(')')
	return builder.String()
}

// NamedCards returns the Cards named value or an error if the edge was not
// loaded in eager-loading with this name.
func (d *Deck) NamedCards(name string) ([]*Card, error) {
	if d.Edges.namedCards == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := d.Edges.namedCards[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (d *Deck) appendNamedCards(name string, edges ...*Card) {
	if d.Edges.namedCards == nil {
		d.Edges.namedCards = make(map[string][]*Card)
	}
	if len(edges) == 0 {
		d.Edges.namedCards[name] = []*Card{}
	} else {
		d.Edges.namedCards[name] = append(d.Edges.namedCards[name], edges...)
	}
}

// Decks is a parsable slice of Deck.
type Decks []*Deck

func (d Decks) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
