package graphql

import (
	graphqlgo "github.com/graphql-go/graphql"
)

var schema, _ = graphqlgo.NewSchema(graphqlgo.SchemaConfig{
	Query: rootQuery,
})
