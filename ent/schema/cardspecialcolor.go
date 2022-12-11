package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type CardSpecialColor struct {
	ent.Schema
}

func (CardSpecialColor) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").Immutable().NotEmpty().Unique(),
		field.String("name_ja").Optional(),
		field.String("name_en").Optional(),
	}
}

func (CardSpecialColor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cards", Card.Type),
	}
}
