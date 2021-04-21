package gql_types

import (
	"github.com/graphql-go/graphql"
)

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "User",
	Description: "For fetching user related informations",
	Fields: graphql.Fields{
		"id":            &graphql.Field{Type: graphql.NewNonNull(graphql.ID)},
		"firstName":     &graphql.Field{Type: graphql.String},
		"middleName":    &graphql.Field{Type: graphql.String},
		"lastName":      &graphql.Field{Type: graphql.String},
		"email":         &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"isAdmin":       &graphql.Field{Type: graphql.NewNonNull(graphql.Boolean)},
		"isVendor":      &graphql.Field{Type: graphql.NewNonNull(graphql.Boolean)},
		"registeredAt":  &graphql.Field{Type: graphql.DateTime},
		"lastLoginedAt": &graphql.Field{Type: graphql.DateTime},
		"intro":         &graphql.Field{Type: graphql.String},
		"profile":       &graphql.Field{Type: graphql.String},
	},
})
