// Code generated by ent, DO NOT EDIT.

package ent

import (
	"agricoladb/ent/card"
	"agricoladb/ent/cardspecialcolor"
	"agricoladb/ent/cardtype"
	"agricoladb/ent/deck"
	"agricoladb/ent/revision"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Card is the model entity for the Card schema.
type Card struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// LiteralID holds the value of the "literal_id" field.
	LiteralID string `json:"literal_id,omitempty"`
	// RevisionID holds the value of the "revision_id" field.
	RevisionID int `json:"revision_id,omitempty"`
	// PrintedID holds the value of the "printed_id" field.
	PrintedID string `json:"printed_id,omitempty"`
	// PlayAgricolaCardID holds the value of the "play_agricola_card_id" field.
	PlayAgricolaCardID string `json:"play_agricola_card_id,omitempty"`
	// DeckID holds the value of the "deck_id" field.
	DeckID int `json:"deck_id,omitempty"`
	// CardTypeID holds the value of the "card_type_id" field.
	CardTypeID int `json:"card_type_id,omitempty"`
	// CardSpecialColorID holds the value of the "card_special_color_id" field.
	CardSpecialColorID int `json:"card_special_color_id,omitempty"`
	// NameJa holds the value of the "name_ja" field.
	NameJa string `json:"name_ja,omitempty"`
	// NameEn holds the value of the "name_en" field.
	NameEn string `json:"name_en,omitempty"`
	// MinPlayersNumber holds the value of the "min_players_number" field.
	MinPlayersNumber int `json:"min_players_number,omitempty"`
	// Prerequisite holds the value of the "prerequisite" field.
	Prerequisite string `json:"prerequisite,omitempty"`
	// Cost holds the value of the "cost" field.
	Cost string `json:"cost,omitempty"`
	// FunctionText holds the value of the "function_text" field.
	FunctionText string `json:"function_text,omitempty"`
	// IsOfficialJa holds the value of the "is_official_ja" field.
	IsOfficialJa bool `json:"is_official_ja,omitempty"`
	// VictoryPoint holds the value of the "victory_point" field.
	VictoryPoint int `json:"victory_point,omitempty"`
	// IsVariableVictoryPoint holds the value of the "is_variable_victory_point" field.
	IsVariableVictoryPoint bool `json:"is_variable_victory_point,omitempty"`
	// HasArrrow holds the value of the "has_arrrow" field.
	HasArrrow bool `json:"has_arrrow,omitempty"`
	// HasBonusPointIcon holds the value of the "has_bonus_point_icon" field.
	HasBonusPointIcon bool `json:"has_bonus_point_icon,omitempty"`
	// HasNegativeBonusPointIcon holds the value of the "has_negative_bonus_point_icon" field.
	HasNegativeBonusPointIcon bool `json:"has_negative_bonus_point_icon,omitempty"`
	// HasPanIcon holds the value of the "has_pan_icon" field.
	HasPanIcon bool `json:"has_pan_icon,omitempty"`
	// HasBreadIcon holds the value of the "has_bread_icon" field.
	HasBreadIcon bool `json:"has_bread_icon,omitempty"`
	// HasFarmPlannerIcon holds the value of the "has_farm_planner_icon" field.
	HasFarmPlannerIcon bool `json:"has_farm_planner_icon,omitempty"`
	// HasActionsBoosterIcon holds the value of the "has_actions_booster_icon" field.
	HasActionsBoosterIcon bool `json:"has_actions_booster_icon,omitempty"`
	// HasPointsProviderIcon holds the value of the "has_points_provider_icon" field.
	HasPointsProviderIcon bool `json:"has_points_provider_icon,omitempty"`
	// HasGoodsProviderIcon holds the value of the "has_goods_provider_icon" field.
	HasGoodsProviderIcon bool `json:"has_goods_provider_icon,omitempty"`
	// HasFoodProviderIcon holds the value of the "has_food_provider_icon" field.
	HasFoodProviderIcon bool `json:"has_food_provider_icon,omitempty"`
	// HasCropProviderIcon holds the value of the "has_crop_provider_icon" field.
	HasCropProviderIcon bool `json:"has_crop_provider_icon,omitempty"`
	// HasBuildingResourceProviderIcon holds the value of the "has_building_resource_provider_icon" field.
	HasBuildingResourceProviderIcon bool `json:"has_building_resource_provider_icon,omitempty"`
	// HasLivestockProviderIcon holds the value of the "has_livestock_provider_icon" field.
	HasLivestockProviderIcon bool `json:"has_livestock_provider_icon,omitempty"`
	// HasCutPeatIcon holds the value of the "has_cut_peat_icon" field.
	HasCutPeatIcon bool `json:"has_cut_peat_icon,omitempty"`
	// HasFellTreesIcon holds the value of the "has_fell_trees_icon" field.
	HasFellTreesIcon bool `json:"has_fell_trees_icon,omitempty"`
	// HasSlashAndBurnIcon holds the value of the "has_slash_and_burn_icon" field.
	HasSlashAndBurnIcon bool `json:"has_slash_and_burn_icon,omitempty"`
	// HasHiringFareIcon holds the value of the "has_hiring_fare_icon" field.
	HasHiringFareIcon bool `json:"has_hiring_fare_icon,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CardQuery when eager-loading is set.
	Edges CardEdges `json:"edges"`
}

// CardEdges holds the relations/edges for other nodes in the graph.
type CardEdges struct {
	// Revision holds the value of the revision edge.
	Revision *Revision `json:"revision,omitempty"`
	// Products holds the value of the products edge.
	Products []*Product `json:"products,omitempty"`
	// Deck holds the value of the deck edge.
	Deck *Deck `json:"deck,omitempty"`
	// CardType holds the value of the card_type edge.
	CardType *CardType `json:"card_type,omitempty"`
	// CardSpecialColor holds the value of the card_special_color edge.
	CardSpecialColor *CardSpecialColor `json:"card_special_color,omitempty"`
	// Children holds the value of the children edge.
	Children []*Card `json:"children,omitempty"`
	// Ancestors holds the value of the ancestors edge.
	Ancestors []*Card `json:"ancestors,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [7]bool
	// totalCount holds the count of the edges above.
	totalCount [7]map[string]int

	namedProducts  map[string][]*Product
	namedChildren  map[string][]*Card
	namedAncestors map[string][]*Card
}

// RevisionOrErr returns the Revision value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CardEdges) RevisionOrErr() (*Revision, error) {
	if e.loadedTypes[0] {
		if e.Revision == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: revision.Label}
		}
		return e.Revision, nil
	}
	return nil, &NotLoadedError{edge: "revision"}
}

// ProductsOrErr returns the Products value or an error if the edge
// was not loaded in eager-loading.
func (e CardEdges) ProductsOrErr() ([]*Product, error) {
	if e.loadedTypes[1] {
		return e.Products, nil
	}
	return nil, &NotLoadedError{edge: "products"}
}

// DeckOrErr returns the Deck value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CardEdges) DeckOrErr() (*Deck, error) {
	if e.loadedTypes[2] {
		if e.Deck == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: deck.Label}
		}
		return e.Deck, nil
	}
	return nil, &NotLoadedError{edge: "deck"}
}

// CardTypeOrErr returns the CardType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CardEdges) CardTypeOrErr() (*CardType, error) {
	if e.loadedTypes[3] {
		if e.CardType == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: cardtype.Label}
		}
		return e.CardType, nil
	}
	return nil, &NotLoadedError{edge: "card_type"}
}

// CardSpecialColorOrErr returns the CardSpecialColor value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CardEdges) CardSpecialColorOrErr() (*CardSpecialColor, error) {
	if e.loadedTypes[4] {
		if e.CardSpecialColor == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: cardspecialcolor.Label}
		}
		return e.CardSpecialColor, nil
	}
	return nil, &NotLoadedError{edge: "card_special_color"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e CardEdges) ChildrenOrErr() ([]*Card, error) {
	if e.loadedTypes[5] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// AncestorsOrErr returns the Ancestors value or an error if the edge
// was not loaded in eager-loading.
func (e CardEdges) AncestorsOrErr() ([]*Card, error) {
	if e.loadedTypes[6] {
		return e.Ancestors, nil
	}
	return nil, &NotLoadedError{edge: "ancestors"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Card) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case card.FieldIsOfficialJa, card.FieldIsVariableVictoryPoint, card.FieldHasArrrow, card.FieldHasBonusPointIcon, card.FieldHasNegativeBonusPointIcon, card.FieldHasPanIcon, card.FieldHasBreadIcon, card.FieldHasFarmPlannerIcon, card.FieldHasActionsBoosterIcon, card.FieldHasPointsProviderIcon, card.FieldHasGoodsProviderIcon, card.FieldHasFoodProviderIcon, card.FieldHasCropProviderIcon, card.FieldHasBuildingResourceProviderIcon, card.FieldHasLivestockProviderIcon, card.FieldHasCutPeatIcon, card.FieldHasFellTreesIcon, card.FieldHasSlashAndBurnIcon, card.FieldHasHiringFareIcon:
			values[i] = new(sql.NullBool)
		case card.FieldID, card.FieldRevisionID, card.FieldDeckID, card.FieldCardTypeID, card.FieldCardSpecialColorID, card.FieldMinPlayersNumber, card.FieldVictoryPoint:
			values[i] = new(sql.NullInt64)
		case card.FieldLiteralID, card.FieldPrintedID, card.FieldPlayAgricolaCardID, card.FieldNameJa, card.FieldNameEn, card.FieldPrerequisite, card.FieldCost, card.FieldFunctionText:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Card", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Card fields.
func (c *Card) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case card.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case card.FieldLiteralID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field literal_id", values[i])
			} else if value.Valid {
				c.LiteralID = value.String
			}
		case card.FieldRevisionID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field revision_id", values[i])
			} else if value.Valid {
				c.RevisionID = int(value.Int64)
			}
		case card.FieldPrintedID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field printed_id", values[i])
			} else if value.Valid {
				c.PrintedID = value.String
			}
		case card.FieldPlayAgricolaCardID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field play_agricola_card_id", values[i])
			} else if value.Valid {
				c.PlayAgricolaCardID = value.String
			}
		case card.FieldDeckID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deck_id", values[i])
			} else if value.Valid {
				c.DeckID = int(value.Int64)
			}
		case card.FieldCardTypeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field card_type_id", values[i])
			} else if value.Valid {
				c.CardTypeID = int(value.Int64)
			}
		case card.FieldCardSpecialColorID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field card_special_color_id", values[i])
			} else if value.Valid {
				c.CardSpecialColorID = int(value.Int64)
			}
		case card.FieldNameJa:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_ja", values[i])
			} else if value.Valid {
				c.NameJa = value.String
			}
		case card.FieldNameEn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_en", values[i])
			} else if value.Valid {
				c.NameEn = value.String
			}
		case card.FieldMinPlayersNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field min_players_number", values[i])
			} else if value.Valid {
				c.MinPlayersNumber = int(value.Int64)
			}
		case card.FieldPrerequisite:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field prerequisite", values[i])
			} else if value.Valid {
				c.Prerequisite = value.String
			}
		case card.FieldCost:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cost", values[i])
			} else if value.Valid {
				c.Cost = value.String
			}
		case card.FieldFunctionText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field function_text", values[i])
			} else if value.Valid {
				c.FunctionText = value.String
			}
		case card.FieldIsOfficialJa:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_official_ja", values[i])
			} else if value.Valid {
				c.IsOfficialJa = value.Bool
			}
		case card.FieldVictoryPoint:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field victory_point", values[i])
			} else if value.Valid {
				c.VictoryPoint = int(value.Int64)
			}
		case card.FieldIsVariableVictoryPoint:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_variable_victory_point", values[i])
			} else if value.Valid {
				c.IsVariableVictoryPoint = value.Bool
			}
		case card.FieldHasArrrow:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field has_arrrow", values[i])
			} else if value.Valid {
				c.HasArrrow = value.Bool
			}
		case card.FieldHasBonusPointIcon:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field has_bonus_point_icon", values[i])
			} else if value.Valid {
				c.HasBonusPointIcon = value.Bool
			}
		case card.FieldHasNegativeBonusPointIcon:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field has_negative_bonus_point_icon", values[i])
			} else if value.Valid {
				c.HasNegativeBonusPointIcon = value.Bool
			}
		case card.FieldHasPanIcon:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field has_pan_icon", values[i])
			} else if value.Valid {
				c.HasPanIcon = value.Bool
			}
		case card.FieldHasBreadIcon:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field has_bread_icon", values[i])
			} else if value.Valid {
				c.HasBreadIcon = value.Bool
			}
		case card.FieldHasFarmPlannerIcon:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field has_farm_planner_icon", values[i])
			} else if value.Valid {
				c.HasFarmPlannerIcon = value.Bool
			}
		case card.FieldHasActionsBoosterIcon:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field has_actions_booster_icon", values[i])
			} else if value.Valid {
				c.HasActionsBoosterIcon = value.Bool
			}
		case card.FieldHasPointsProviderIcon:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field has_points_provider_icon", values[i])
			} else if value.Valid {
				c.HasPointsProviderIcon = value.Bool
			}
		case card.FieldHasGoodsProviderIcon:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field has_goods_provider_icon", values[i])
			} else if value.Valid {
				c.HasGoodsProviderIcon = value.Bool
			}
		case card.FieldHasFoodProviderIcon:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field has_food_provider_icon", values[i])
			} else if value.Valid {
				c.HasFoodProviderIcon = value.Bool
			}
		case card.FieldHasCropProviderIcon:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field has_crop_provider_icon", values[i])
			} else if value.Valid {
				c.HasCropProviderIcon = value.Bool
			}
		case card.FieldHasBuildingResourceProviderIcon:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field has_building_resource_provider_icon", values[i])
			} else if value.Valid {
				c.HasBuildingResourceProviderIcon = value.Bool
			}
		case card.FieldHasLivestockProviderIcon:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field has_livestock_provider_icon", values[i])
			} else if value.Valid {
				c.HasLivestockProviderIcon = value.Bool
			}
		case card.FieldHasCutPeatIcon:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field has_cut_peat_icon", values[i])
			} else if value.Valid {
				c.HasCutPeatIcon = value.Bool
			}
		case card.FieldHasFellTreesIcon:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field has_fell_trees_icon", values[i])
			} else if value.Valid {
				c.HasFellTreesIcon = value.Bool
			}
		case card.FieldHasSlashAndBurnIcon:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field has_slash_and_burn_icon", values[i])
			} else if value.Valid {
				c.HasSlashAndBurnIcon = value.Bool
			}
		case card.FieldHasHiringFareIcon:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field has_hiring_fare_icon", values[i])
			} else if value.Valid {
				c.HasHiringFareIcon = value.Bool
			}
		}
	}
	return nil
}

// QueryRevision queries the "revision" edge of the Card entity.
func (c *Card) QueryRevision() *RevisionQuery {
	return (&CardClient{config: c.config}).QueryRevision(c)
}

// QueryProducts queries the "products" edge of the Card entity.
func (c *Card) QueryProducts() *ProductQuery {
	return (&CardClient{config: c.config}).QueryProducts(c)
}

// QueryDeck queries the "deck" edge of the Card entity.
func (c *Card) QueryDeck() *DeckQuery {
	return (&CardClient{config: c.config}).QueryDeck(c)
}

// QueryCardType queries the "card_type" edge of the Card entity.
func (c *Card) QueryCardType() *CardTypeQuery {
	return (&CardClient{config: c.config}).QueryCardType(c)
}

// QueryCardSpecialColor queries the "card_special_color" edge of the Card entity.
func (c *Card) QueryCardSpecialColor() *CardSpecialColorQuery {
	return (&CardClient{config: c.config}).QueryCardSpecialColor(c)
}

// QueryChildren queries the "children" edge of the Card entity.
func (c *Card) QueryChildren() *CardQuery {
	return (&CardClient{config: c.config}).QueryChildren(c)
}

// QueryAncestors queries the "ancestors" edge of the Card entity.
func (c *Card) QueryAncestors() *CardQuery {
	return (&CardClient{config: c.config}).QueryAncestors(c)
}

// Update returns a builder for updating this Card.
// Note that you need to call Card.Unwrap() before calling this method if this Card
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Card) Update() *CardUpdateOne {
	return (&CardClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Card entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Card) Unwrap() *Card {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Card is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Card) String() string {
	var builder strings.Builder
	builder.WriteString("Card(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("literal_id=")
	builder.WriteString(c.LiteralID)
	builder.WriteString(", ")
	builder.WriteString("revision_id=")
	builder.WriteString(fmt.Sprintf("%v", c.RevisionID))
	builder.WriteString(", ")
	builder.WriteString("printed_id=")
	builder.WriteString(c.PrintedID)
	builder.WriteString(", ")
	builder.WriteString("play_agricola_card_id=")
	builder.WriteString(c.PlayAgricolaCardID)
	builder.WriteString(", ")
	builder.WriteString("deck_id=")
	builder.WriteString(fmt.Sprintf("%v", c.DeckID))
	builder.WriteString(", ")
	builder.WriteString("card_type_id=")
	builder.WriteString(fmt.Sprintf("%v", c.CardTypeID))
	builder.WriteString(", ")
	builder.WriteString("card_special_color_id=")
	builder.WriteString(fmt.Sprintf("%v", c.CardSpecialColorID))
	builder.WriteString(", ")
	builder.WriteString("name_ja=")
	builder.WriteString(c.NameJa)
	builder.WriteString(", ")
	builder.WriteString("name_en=")
	builder.WriteString(c.NameEn)
	builder.WriteString(", ")
	builder.WriteString("min_players_number=")
	builder.WriteString(fmt.Sprintf("%v", c.MinPlayersNumber))
	builder.WriteString(", ")
	builder.WriteString("prerequisite=")
	builder.WriteString(c.Prerequisite)
	builder.WriteString(", ")
	builder.WriteString("cost=")
	builder.WriteString(c.Cost)
	builder.WriteString(", ")
	builder.WriteString("function_text=")
	builder.WriteString(c.FunctionText)
	builder.WriteString(", ")
	builder.WriteString("is_official_ja=")
	builder.WriteString(fmt.Sprintf("%v", c.IsOfficialJa))
	builder.WriteString(", ")
	builder.WriteString("victory_point=")
	builder.WriteString(fmt.Sprintf("%v", c.VictoryPoint))
	builder.WriteString(", ")
	builder.WriteString("is_variable_victory_point=")
	builder.WriteString(fmt.Sprintf("%v", c.IsVariableVictoryPoint))
	builder.WriteString(", ")
	builder.WriteString("has_arrrow=")
	builder.WriteString(fmt.Sprintf("%v", c.HasArrrow))
	builder.WriteString(", ")
	builder.WriteString("has_bonus_point_icon=")
	builder.WriteString(fmt.Sprintf("%v", c.HasBonusPointIcon))
	builder.WriteString(", ")
	builder.WriteString("has_negative_bonus_point_icon=")
	builder.WriteString(fmt.Sprintf("%v", c.HasNegativeBonusPointIcon))
	builder.WriteString(", ")
	builder.WriteString("has_pan_icon=")
	builder.WriteString(fmt.Sprintf("%v", c.HasPanIcon))
	builder.WriteString(", ")
	builder.WriteString("has_bread_icon=")
	builder.WriteString(fmt.Sprintf("%v", c.HasBreadIcon))
	builder.WriteString(", ")
	builder.WriteString("has_farm_planner_icon=")
	builder.WriteString(fmt.Sprintf("%v", c.HasFarmPlannerIcon))
	builder.WriteString(", ")
	builder.WriteString("has_actions_booster_icon=")
	builder.WriteString(fmt.Sprintf("%v", c.HasActionsBoosterIcon))
	builder.WriteString(", ")
	builder.WriteString("has_points_provider_icon=")
	builder.WriteString(fmt.Sprintf("%v", c.HasPointsProviderIcon))
	builder.WriteString(", ")
	builder.WriteString("has_goods_provider_icon=")
	builder.WriteString(fmt.Sprintf("%v", c.HasGoodsProviderIcon))
	builder.WriteString(", ")
	builder.WriteString("has_food_provider_icon=")
	builder.WriteString(fmt.Sprintf("%v", c.HasFoodProviderIcon))
	builder.WriteString(", ")
	builder.WriteString("has_crop_provider_icon=")
	builder.WriteString(fmt.Sprintf("%v", c.HasCropProviderIcon))
	builder.WriteString(", ")
	builder.WriteString("has_building_resource_provider_icon=")
	builder.WriteString(fmt.Sprintf("%v", c.HasBuildingResourceProviderIcon))
	builder.WriteString(", ")
	builder.WriteString("has_livestock_provider_icon=")
	builder.WriteString(fmt.Sprintf("%v", c.HasLivestockProviderIcon))
	builder.WriteString(", ")
	builder.WriteString("has_cut_peat_icon=")
	builder.WriteString(fmt.Sprintf("%v", c.HasCutPeatIcon))
	builder.WriteString(", ")
	builder.WriteString("has_fell_trees_icon=")
	builder.WriteString(fmt.Sprintf("%v", c.HasFellTreesIcon))
	builder.WriteString(", ")
	builder.WriteString("has_slash_and_burn_icon=")
	builder.WriteString(fmt.Sprintf("%v", c.HasSlashAndBurnIcon))
	builder.WriteString(", ")
	builder.WriteString("has_hiring_fare_icon=")
	builder.WriteString(fmt.Sprintf("%v", c.HasHiringFareIcon))
	builder.WriteByte(')')
	return builder.String()
}

// NamedProducts returns the Products named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Card) NamedProducts(name string) ([]*Product, error) {
	if c.Edges.namedProducts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedProducts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Card) appendNamedProducts(name string, edges ...*Product) {
	if c.Edges.namedProducts == nil {
		c.Edges.namedProducts = make(map[string][]*Product)
	}
	if len(edges) == 0 {
		c.Edges.namedProducts[name] = []*Product{}
	} else {
		c.Edges.namedProducts[name] = append(c.Edges.namedProducts[name], edges...)
	}
}

// NamedChildren returns the Children named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Card) NamedChildren(name string) ([]*Card, error) {
	if c.Edges.namedChildren == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedChildren[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Card) appendNamedChildren(name string, edges ...*Card) {
	if c.Edges.namedChildren == nil {
		c.Edges.namedChildren = make(map[string][]*Card)
	}
	if len(edges) == 0 {
		c.Edges.namedChildren[name] = []*Card{}
	} else {
		c.Edges.namedChildren[name] = append(c.Edges.namedChildren[name], edges...)
	}
}

// NamedAncestors returns the Ancestors named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Card) NamedAncestors(name string) ([]*Card, error) {
	if c.Edges.namedAncestors == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedAncestors[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Card) appendNamedAncestors(name string, edges ...*Card) {
	if c.Edges.namedAncestors == nil {
		c.Edges.namedAncestors = make(map[string][]*Card)
	}
	if len(edges) == 0 {
		c.Edges.namedAncestors[name] = []*Card{}
	} else {
		c.Edges.namedAncestors[name] = append(c.Edges.namedAncestors[name], edges...)
	}
}

// Cards is a parsable slice of Card.
type Cards []*Card

func (c Cards) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
