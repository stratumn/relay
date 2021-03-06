package starwars

import (
	"errors"

	"context"

	"github.com/graphql-go/graphql"
	"github.com/stratumn/graphql-pagination-go"
)

/**
 * This is a basic end-to-end test, designed to demonstrate the various
 * capabilities of a Relay-compliant GraphQL server.
 *
 * It is recommended that readers of this test be familiar with
 * the end-to-end test in GraphQL.js first, as this test skips
 * over the basics covered there in favor of illustrating the
 * key aspects of the Relay spec that this test is designed to illustrate.
 *
 * We will create a GraphQL schema that describes the major
 * factions and ships in the original Star Wars trilogy.
 *
 * NOTE: This may contain spoilers for the original Star
 * Wars trilogy.
 */

/**
 * Using our shorthand to describe type systems, the type system for our
 * example will be the following:
 *
 * interface Item {
 *   id: ID!
 * }
 *
 * type Faction : Item {
 *   id: ID!
 *   name: String
 *   ships: ShipList
 * }
 *
 * type Ship : Item {
 *   id: ID!
 *   name: String
 * }
 *
 * type ShipList {
 *   items: [Ship]
 *   pageInfo: PageInfo!
 * }
 *
 * type PageInfo {
 *   hasNextPage: Boolean!
 *   hasPreviousPage: Boolean!
 *   startCursor: String
 *   endCursor: String
 * }
 *
 * type Query {
 *   rebels: Faction
 *   empire: Faction
 *   item(id: ID!): Item
 * }
 *
 * input IntroduceShipInput {
 *   clientMutationID: string!
 *   shipName: string!
 *   factionId: ID!
 * }
 *
 * input IntroduceShipPayload {
 *   clientMutationID: string!
 *   ship: Ship
 *   faction: Faction
 * }
 *
 * type Mutation {
 *   introduceShip(input IntroduceShipInput!): IntroduceShipPayload
 * }
 */

// declare definitions first, and initialize them in init() to break `initialization loop`
// i.e.:
// - itemDefinitions refers to
// - shipType refers to
// - itemDefinitions

var itemDefinitions *pagination.ItemDefinitions
var shipType *graphql.Object
var factionType *graphql.Object

// Schema is the exported schema, defined in init()
var Schema graphql.Schema

func init() {

	/**
	 * We get the item interface and field from the relay library.
	 *
	 * The first method is the way we resolve an ID to its object. The second is the
	 * way we resolve an object that implements item to its type.
	 */
	itemDefinitions = pagination.NewItemDefinitions(pagination.ItemDefinitionsConfig{
		IDFetcher: func(id string, info graphql.ResolveInfo, ctx context.Context) (interface{}, error) {
			// resolve id from global id
			resolvedID := pagination.FromGlobalID(id)

			// based on id and its type, return the object
			switch resolvedID.Type {
			case "Faction":
				return GetFaction(resolvedID.ID), nil
			case "Ship":
				return GetShip(resolvedID.ID), nil
			default:
				return nil, errors.New("Unknown item type")
			}
		},
		TypeResolve: func(p graphql.ResolveTypeParams) *graphql.Object {
			// based on the type of the value, return GraphQLObjectType
			switch p.Value.(type) {
			case *Faction:
				return factionType
			default:
				return shipType
			}
		},
	})

	/**
	 * We define our basic ship type.
	 *
	 * This implements the following type system shorthand:
	 *   type Ship : Item {
	 *     id: String!
	 *     name: String
	 *   }
	 */
	shipType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Ship",
		Description: "A ship in the Star Wars saga",
		Fields: graphql.Fields{
			"id": pagination.GlobalIDField("Ship", nil),
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the ship.",
			},
		},
		Interfaces: []*graphql.Interface{
			itemDefinitions.ItemInterface,
		},
	})

	/**
	 * We define a list between a faction and its ships.
	 *
	 * listType implements the following type system shorthand:
	 *   type ShipList {
	 *     items: [Ship]
	 *     pageInfo: PageInfo!
	 *   }
	 */
	shipListDefinition := pagination.ListDefinitions(pagination.ListConfig{
		Name:     "Ship",
		ItemType: shipType,
	})

	/**
	 * We define our faction type, which implements the item interface.
	 *
	 * This implements the following type system shorthand:
	 *   type Faction : Item {
	 *     id: String!
	 *     name: String
	 *     ships: ShipList
	 *   }
	 */
	factionType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Faction",
		Description: "A faction in the Star Wars saga",
		Fields: graphql.Fields{
			"id": pagination.GlobalIDField("Faction", nil),
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the faction.",
			},
			"ships": &graphql.Field{
				Type: shipListDefinition.ListType,
				Args: pagination.ListArgs,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// convert args map[string]interface into ListArguments
					args := pagination.NewListArguments(p.Args)

					// get ship objects from current faction
					ships := []interface{}{}
					if faction, ok := p.Source.(*Faction); ok {
						for _, shipID := range faction.Ships {
							ships = append(ships, GetShip(shipID))
						}
					}
					// let relay library figure out the result, given
					// - the list of ships for this faction
					// - and the filter arguments (i.e. first, last, after, before)
					return pagination.ListFromArray(ships, args), nil
				},
			},
		},
		Interfaces: []*graphql.Interface{
			itemDefinitions.ItemInterface,
		},
	})

	/**
	 * This is the type that will be the root of our query, and the
	 * entry point into our schema.
	 *
	 * This implements the following type system shorthand:
	 *   type Query {
	 *     rebels: Faction
	 *     empire: Faction
	 *     item(id: String!): Item
	 *   }
	 */
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"rebels": &graphql.Field{
				Type: factionType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return GetRebels(), nil
				},
			},
			"empire": &graphql.Field{
				Type: factionType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return GetEmpire(), nil
				},
			},
			"item": itemDefinitions.ItemField,
		},
	})

	/**
	 * This will return a GraphQLField for our ship
	 * mutation.
	 *
	 * It creates these two types implicitly:
	 *   input IntroduceShipInput {
	 *     clientMutationID: string!
	 *     shipName: string!
	 *     factionId: ID!
	 *   }
	 *
	 *   input IntroduceShipPayload {
	 *     clientMutationID: string!
	 *     ship: Ship
	 *     faction: Faction
	 *   }
	 */
	shipMutation := pagination.MutationWithClientMutationID(pagination.MutationConfig{
		Name: "IntroduceShip",
		InputFields: graphql.InputObjectConfigFieldMap{
			"shipName": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"factionId": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		OutputFields: graphql.Fields{
			"ship": &graphql.Field{
				Type: shipType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if payload, ok := p.Source.(map[string]interface{}); ok {
						return GetShip(payload["shipId"].(string)), nil
					}
					return nil, nil
				},
			},
			"faction": &graphql.Field{
				Type: factionType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if payload, ok := p.Source.(map[string]interface{}); ok {
						return GetFaction(payload["factionId"].(string)), nil
					}
					return nil, nil
				},
			},
		},
		MutateAndGetPayload: func(inputMap map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
			// `inputMap` is a map with keys/fields as specified in `InputFields`
			// Note, that these fields were specified as non-nullables, so we can assume that it exists.
			shipName := inputMap["shipName"].(string)
			factionID := inputMap["factionId"].(string)

			// This mutation involves us creating (introducing) a new ship
			newShip := CreateShip(shipName, factionID)
			// return payload
			return map[string]interface{}{
				"shipId":    newShip.ID,
				"factionId": factionID,
			}, nil
		},
	})

	/**
	 * This is the type that will be the root of our mutations, and the
	 * entry point into performing writes in our schema.
	 *
	 * This implements the following type system shorthand:
	 *   type Mutation {
	 *     introduceShip(input IntroduceShipInput!): IntroduceShipPayload
	 *   }
	 */

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"introduceShip": shipMutation,
		},
	})

	/**
	 * Finally, we construct our schema (whose starting query type is the query
	 * type we defined above) and export it.
	 */
	var err error
	Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
	if err != nil {
		// panic if there is an error in schema
		panic(err)
	}
}
