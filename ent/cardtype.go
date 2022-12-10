// Code generated by ent, DO NOT EDIT.

package ent

import (
	"agricoladb/ent/cardtype"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// CardType is the model entity for the CardType schema.
type CardType struct {
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
	// The values are being populated by the CardTypeQuery when eager-loading is set.
	Edges CardTypeEdges `json:"edges"`
}

// CardTypeEdges holds the relations/edges for other nodes in the graph.
type CardTypeEdges struct {
	// Cards holds the value of the cards edge.
	Cards []*Card `json:"cards,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedCards map[string][]*Card
}

// CardsOrErr returns the Cards value or an error if the edge
// was not loaded in eager-loading.
func (e CardTypeEdges) CardsOrErr() ([]*Card, error) {
	if e.loadedTypes[0] {
		return e.Cards, nil
	}
	return nil, &NotLoadedError{edge: "cards"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CardType) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case cardtype.FieldID:
			values[i] = new(sql.NullInt64)
		case cardtype.FieldKey, cardtype.FieldNameJa, cardtype.FieldNameEn:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CardType", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CardType fields.
func (ct *CardType) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cardtype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ct.ID = int(value.Int64)
		case cardtype.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				ct.Key = value.String
			}
		case cardtype.FieldNameJa:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_ja", values[i])
			} else if value.Valid {
				ct.NameJa = value.String
			}
		case cardtype.FieldNameEn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_en", values[i])
			} else if value.Valid {
				ct.NameEn = value.String
			}
		}
	}
	return nil
}

// QueryCards queries the "cards" edge of the CardType entity.
func (ct *CardType) QueryCards() *CardQuery {
	return (&CardTypeClient{config: ct.config}).QueryCards(ct)
}

// Update returns a builder for updating this CardType.
// Note that you need to call CardType.Unwrap() before calling this method if this CardType
// was returned from a transaction, and the transaction was committed or rolled back.
func (ct *CardType) Update() *CardTypeUpdateOne {
	return (&CardTypeClient{config: ct.config}).UpdateOne(ct)
}

// Unwrap unwraps the CardType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ct *CardType) Unwrap() *CardType {
	_tx, ok := ct.config.driver.(*txDriver)
	if !ok {
		panic("ent: CardType is not a transactional entity")
	}
	ct.config.driver = _tx.drv
	return ct
}

// String implements the fmt.Stringer.
func (ct *CardType) String() string {
	var builder strings.Builder
	builder.WriteString("CardType(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ct.ID))
	builder.WriteString("key=")
	builder.WriteString(ct.Key)
	builder.WriteString(", ")
	builder.WriteString("name_ja=")
	builder.WriteString(ct.NameJa)
	builder.WriteString(", ")
	builder.WriteString("name_en=")
	builder.WriteString(ct.NameEn)
	builder.WriteByte(')')
	return builder.String()
}

// NamedCards returns the Cards named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ct *CardType) NamedCards(name string) ([]*Card, error) {
	if ct.Edges.namedCards == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ct.Edges.namedCards[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ct *CardType) appendNamedCards(name string, edges ...*Card) {
	if ct.Edges.namedCards == nil {
		ct.Edges.namedCards = make(map[string][]*Card)
	}
	if len(edges) == 0 {
		ct.Edges.namedCards[name] = []*Card{}
	} else {
		ct.Edges.namedCards[name] = append(ct.Edges.namedCards[name], edges...)
	}
}

// CardTypes is a parsable slice of CardType.
type CardTypes []*CardType

func (ct CardTypes) config(cfg config) {
	for _i := range ct {
		ct[_i].config = cfg
	}
}
