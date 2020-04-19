package graphql

import (
	graphqlgo "github.com/graphql-go/graphql"
)

var rootQuery = graphqlgo.NewObject(graphqlgo.ObjectConfig{
	Name:        "RootQuery",
	Description: "Root of all queries",
	Fields: graphqlgo.Fields{
		"viewer": &graphqlgo.Field{
			Type:        userType,
			Description: "Current user",
			Resolve:     getCurrentUser,
		},
	},
})
