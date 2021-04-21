package gql

import (
	"github.com/graphql-go/graphql"
	resolver "github.com/iftech-a/lookum/src/backend/internal/graphql/resolvers"
	gql_types "github.com/iftech-a/lookum/src/backend/internal/graphql/types"
)

func (s *GQLServer) UserField() *graphql.Field {
	return &graphql.Field{
		Type: gql_types.UserType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "User ID",
				Type:        graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: resolver.User(s.s.User()),
	}
}

func (s *GQLServer) ProductField() *graphql.Field {
	return &graphql.Field{
		Type: gql_types.ProductType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "Product ID",
				Type:        graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: resolver.Product(s.s.Product()),
	}
}

func (s *GQLServer) ProductsField() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(gql_types.ProductType),
		Args: graphql.FieldConfigArgument{
			"category_id": &graphql.ArgumentConfig{
				Description: "Product category id",
				Type:        graphql.ID,
			},
			"limit": &graphql.ArgumentConfig{
				Description:  "Limit result count(default: 20)",
				Type:         graphql.Int,
				DefaultValue: 20,
			},
		},
		Resolve: resolver.Products(s.s.Product()),
	}
}

func (s *GQLServer) Fields() graphql.Fields {
	return graphql.Fields{
		"user":     s.UserField(),
		"product":  s.ProductField(),
		"products": s.ProductsField(),
	}
}

func (s *GQLServer) newQuery() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:   "Query",
		Fields: s.Fields(),
	})
}
