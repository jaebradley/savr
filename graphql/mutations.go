package graphql

import (
	graphqlgo "github.com/graphql-go/graphql"
)

var rootMutation = graphqlgo.NewObject(graphqlgo.ObjectConfig{
	Name:        "RootMutation",
	Description: "Root of all mutations",
	Fields:      graphqlgo.Fields{},
})
