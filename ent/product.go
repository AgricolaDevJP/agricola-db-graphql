// Code generated by ent, DO NOT EDIT.

package ent

import (
	"agricoladb/ent/product"
	"agricoladb/ent/revision"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Product is the model entity for the Product schema.
type Product struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Key holds the value of the "key" field.
	Key string `json:"key,omitempty"`
	// RevisionID holds the value of the "revision_id" field.
	RevisionID int `json:"revision_id,omitempty"`
	// IsOfficialJa holds the value of the "is_official_ja" field.
	IsOfficialJa bool `json:"is_official_ja,omitempty"`
	// NameJa holds the value of the "name_ja" field.
	NameJa string `json:"name_ja,omitempty"`
	// NameEn holds the value of the "name_en" field.
	NameEn string `json:"name_en,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProductQuery when eager-loading is set.
	Edges ProductEdges `json:"edges"`
}

// ProductEdges holds the relations/edges for other nodes in the graph.
type ProductEdges struct {
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
func (e ProductEdges) CardsOrErr() ([]*Card, error) {
	if e.loadedTypes[0] {
		return e.Cards, nil
	}
	return nil, &NotLoadedError{edge: "cards"}
}

// RevisionOrErr returns the Revision value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProductEdges) RevisionOrErr() (*Revision, error) {
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
func (*Product) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case product.FieldIsOfficialJa:
			values[i] = new(sql.NullBool)
		case product.FieldID, product.FieldRevisionID:
			values[i] = new(sql.NullInt64)
		case product.FieldKey, product.FieldNameJa, product.FieldNameEn:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Product", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Product fields.
func (pr *Product) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case product.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case product.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				pr.Key = value.String
			}
		case product.FieldRevisionID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field revision_id", values[i])
			} else if value.Valid {
				pr.RevisionID = int(value.Int64)
			}
		case product.FieldIsOfficialJa:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_official_ja", values[i])
			} else if value.Valid {
				pr.IsOfficialJa = value.Bool
			}
		case product.FieldNameJa:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_ja", values[i])
			} else if value.Valid {
				pr.NameJa = value.String
			}
		case product.FieldNameEn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_en", values[i])
			} else if value.Valid {
				pr.NameEn = value.String
			}
		}
	}
	return nil
}

// QueryCards queries the "cards" edge of the Product entity.
func (pr *Product) QueryCards() *CardQuery {
	return (&ProductClient{config: pr.config}).QueryCards(pr)
}

// QueryRevision queries the "revision" edge of the Product entity.
func (pr *Product) QueryRevision() *RevisionQuery {
	return (&ProductClient{config: pr.config}).QueryRevision(pr)
}

// Update returns a builder for updating this Product.
// Note that you need to call Product.Unwrap() before calling this method if this Product
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Product) Update() *ProductUpdateOne {
	return (&ProductClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the Product entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Product) Unwrap() *Product {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Product is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Product) String() string {
	var builder strings.Builder
	builder.WriteString("Product(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("key=")
	builder.WriteString(pr.Key)
	builder.WriteString(", ")
	builder.WriteString("revision_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.RevisionID))
	builder.WriteString(", ")
	builder.WriteString("is_official_ja=")
	builder.WriteString(fmt.Sprintf("%v", pr.IsOfficialJa))
	builder.WriteString(", ")
	builder.WriteString("name_ja=")
	builder.WriteString(pr.NameJa)
	builder.WriteString(", ")
	builder.WriteString("name_en=")
	builder.WriteString(pr.NameEn)
	builder.WriteByte(')')
	return builder.String()
}

// NamedCards returns the Cards named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Product) NamedCards(name string) ([]*Card, error) {
	if pr.Edges.namedCards == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedCards[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Product) appendNamedCards(name string, edges ...*Card) {
	if pr.Edges.namedCards == nil {
		pr.Edges.namedCards = make(map[string][]*Card)
	}
	if len(edges) == 0 {
		pr.Edges.namedCards[name] = []*Card{}
	} else {
		pr.Edges.namedCards[name] = append(pr.Edges.namedCards[name], edges...)
	}
}

// Products is a parsable slice of Product.
type Products []*Product

func (pr Products) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
