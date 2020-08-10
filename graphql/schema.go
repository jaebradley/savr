package graphql

import (
	"context"
	"errors"

	graphqlgo "github.com/graphql-go/graphql"
	relay "github.com/graphql-go/relay"
	"github.com/jaebradley/savr/database"
)

var nodeDefinitions *relay.NodeDefinitions
var userType *graphqlgo.Object
var userResourceType *graphqlgo.Object
var resourceType *graphqlgo.Object
var schema graphqlgo.Schema

func init() {
	nodeDefinitions = relay.NewNodeDefinitions(relay.NodeDefinitionsConfig{
		IDFetcher: func(id string, info graphqlgo.ResolveInfo, ctx context.Context) (interface{}, error) {
			// resolve id from global id
			resolvedID := relay.FromGlobalID(id)

			// based on id and its type, return the object
			switch resolvedID.Type {
			case "User":
				return database.GetUserByID(resolvedID.ID), nil
			case "UserResource":
				return database.GetUserResourceByID(resolvedID.ID), nil
			default:
				return nil, errors.New("Unknown node type")
			}
		},
		TypeResolve: func(p graphqlgo.ResolveTypeParams) *graphqlgo.Object {
			switch p.Value.(type) {
			case *database.User:
				return userType
			case *database.UserResource:
				return userResourceType
			default:
				return nil
			}
		},
	})

	resourceType = graphqlgo.NewObject(graphqlgo.ObjectConfig{
		Name:        "Resource",
		Description: "A resource that could be an article, a gif, an image, etc.",
		Fields: graphqlgo.Fields{
			"id": relay.GlobalIDField("Resource", nil),
			"location": &graphqlgo.Field{
				Type:        graphqlgo.String,
				Description: "The location of the resource (i.e. it's URL)",
			},
		},
		Interfaces: []*graphqlgo.Interface{
			nodeDefinitions.NodeInterface,
		},
	})

	resourceConnectionDefinition := relay.ConnectionDefinitions(relay.ConnectionConfig{
		Name:     "Resource",
		NodeType: resourceType,
	})

	userType = graphqlgo.NewObject(graphqlgo.ObjectConfig{
		Name:        "User",
		Description: "An application user",
		Fields: graphqlgo.Fields{
			"id": relay.GlobalIDField("User", nil),
			"emailAddress": &graphqlgo.Field{
				Type:        graphqlgo.NewNonNull(graphqlgo.String),
				Description: "The email address of the user",
			},
			"resources": &graphqlgo.Field{
				Type: resourceConnectionDefinition.ConnectionType,
				Args: relay.ConnectionArgs,
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					// convert args map[string]interface into ConnectionArguments
					args := relay.NewConnectionArguments(p.Args)

					resources := []interface{}{}
					if user, ok := p.Source.(*User); ok {
						for _, resourceID := range user.resources {
							resources = append(resources, database.GetUserResourceByID(resourceID))
						}
					}
					return relay.ConnectionFromArray(resources, args), nil
				},
			},
		},
		Interfaces: []*graphqlgo.Interface{
			nodeDefinitions.NodeInterface,
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
	userResourceMutation := relay.MutationWithClientMutationID(relay.MutationConfig{
		Name: "AddResource",
		InputFields: graphqlgo.InputObjectConfigFieldMap{
			"location": &graphqlgo.InputObjectFieldConfig{
				Type: graphqlgo.NewNonNull(graphqlgo.String),
			},
		},
		OutputFields: graphqlgo.Fields{
			"resource": &graphqlgo.Field{
				Type: resourceType,
				Resolve: func(p graphqlgo.ResolveParams) (interface{}, error) {
					if payload, ok := p.Source.(map[string]interface{}); ok {
						return GetResource(payload["resourceId"].(string)), nil
					}
					return nil, nil
				},
			},
		},
		MutateAndGetPayload: func(inputMap map[string]interface{}, info graphqlgo.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
			// `inputMap` is a map with keys/fields as specified in `InputFields`
			// Note, that these fields were specified as non-nullables, so we can assume that it exists.
			location := inputMap["location"].(string)

			// This mutation involves us creating (introducing) a new ship
			newResource := database.createUserResource()
			// return payload
			return map[string]interface{}{
				"resource": newResource,
			}, nil
		},
	})

	schema, _ = graphqlgo.NewSchema(graphqlgo.SchemaConfig{
		Query: rootQuery,
	})
}
