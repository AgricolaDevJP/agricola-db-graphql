// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (c *Card) Revision(ctx context.Context) (*Revision, error) {
	result, err := c.Edges.RevisionOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryRevision().Only(ctx)
	}
	return result, err
}

func (c *Card) Products(ctx context.Context) (result []*Product, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedProducts(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.ProductsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryProducts().All(ctx)
	}
	return result, err
}

func (c *Card) Deck(ctx context.Context) (*Deck, error) {
	result, err := c.Edges.DeckOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryDeck().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (c *Card) CardType(ctx context.Context) (*CardType, error) {
	result, err := c.Edges.CardTypeOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryCardType().Only(ctx)
	}
	return result, err
}

func (c *Card) CardSpecialColor(ctx context.Context) (*CardSpecialColor, error) {
	result, err := c.Edges.CardSpecialColorOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryCardSpecialColor().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (c *Card) Children(ctx context.Context) (result []*Card, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedChildren(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.ChildrenOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryChildren().All(ctx)
	}
	return result, err
}

func (c *Card) Ancestors(ctx context.Context) (result []*Card, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedAncestors(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.AncestorsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryAncestors().All(ctx)
	}
	return result, err
}

func (csc *CardSpecialColor) Cards(ctx context.Context) (result []*Card, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = csc.NamedCards(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = csc.Edges.CardsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = csc.QueryCards().All(ctx)
	}
	return result, err
}

func (ct *CardType) Cards(ctx context.Context) (result []*Card, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = ct.NamedCards(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = ct.Edges.CardsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = ct.QueryCards().All(ctx)
	}
	return result, err
}

func (d *Deck) Cards(ctx context.Context) (result []*Card, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = d.NamedCards(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = d.Edges.CardsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = d.QueryCards().All(ctx)
	}
	return result, err
}

func (d *Deck) Revision(ctx context.Context) (*Revision, error) {
	result, err := d.Edges.RevisionOrErr()
	if IsNotLoaded(err) {
		result, err = d.QueryRevision().Only(ctx)
	}
	return result, err
}

func (pr *Product) Cards(ctx context.Context) (result []*Card, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pr.NamedCards(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pr.Edges.CardsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pr.QueryCards().All(ctx)
	}
	return result, err
}

func (pr *Product) Revision(ctx context.Context) (*Revision, error) {
	result, err := pr.Edges.RevisionOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryRevision().Only(ctx)
	}
	return result, err
}

func (r *Revision) Cards(ctx context.Context) (result []*Card, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = r.NamedCards(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = r.Edges.CardsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = r.QueryCards().All(ctx)
	}
	return result, err
}

func (r *Revision) Products(ctx context.Context) (result []*Product, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = r.NamedProducts(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = r.Edges.ProductsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = r.QueryProducts().All(ctx)
	}
	return result, err
}

func (r *Revision) Decks(ctx context.Context) (result []*Deck, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = r.NamedDecks(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = r.Edges.DecksOrErr()
	}
	if IsNotLoaded(err) {
		result, err = r.QueryDecks().All(ctx)
	}
	return result, err
}
