package graphql

import (
	graphqlgo "github.com/graphql-go/graphql"
)

var userType = graphqlgo.NewObject(
	graphqlgo.ObjectConfig{
		Name:        "User",
		Description: "An application user",
		Fields: graphqlgo.Fields{
			"id": &graphqlgo.Field{
				Type:        graphqlgo.NewNonNull(graphqlgo.Int),
				Description: "The id of the user",
			},
			"emailAddress": &graphqlgo.Field{
				Type:        graphqlgo.NewNonNull(graphqlgo.String),
				Description: "The email address of the user",
			},
		},
	},
)
