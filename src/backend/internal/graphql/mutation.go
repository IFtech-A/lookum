package gql

import (
	"github.com/graphql-go/graphql"
	resolver "github.com/iftech-a/lookum/src/backend/internal/graphql/resolvers"
	gql_types "github.com/iftech-a/lookum/src/backend/internal/graphql/types"
)

func (s *GQLServer) userCreate() *graphql.Field {
	return &graphql.Field{
		Name:    "createUser",
		Type:    gql_types.UserType,
		Resolve: resolver.UserCreate(s.s.User()),
	}
}

func (s *GQLServer) mutationFields() graphql.Fields {
	return graphql.Fields{
		"userCreate": s.userCreate(),
	}
}

func (s *GQLServer) newMutation() *graphql.Object {

	return graphql.NewObject(graphql.ObjectConfig{
		Name:   "Mutation",
		Fields: s.mutationFields(),
	})
}
