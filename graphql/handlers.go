package graphql

import "github.com/graphql-go/handler"

// Handler handles GraphQL
var Handler = handler.New(&handler.Config{
	Schema:   &schema,
	Pretty:   true,
	GraphiQL: true,
})
