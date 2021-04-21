package gql_types

import "github.com/graphql-go/graphql"

var ProductType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Product",
	Description: "For fetching product related informations",
	Fields: graphql.Fields{
		"id":             &graphql.Field{Type: graphql.NewNonNull(graphql.ID)},
		"title":          &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"meta_title":     &graphql.Field{Type: graphql.String},
		"slug":           &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"summary":        &graphql.Field{Type: graphql.String},
		"type":           &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
		"sku":            &graphql.Field{Type: graphql.String},
		"price":          &graphql.Field{Type: graphql.NewNonNull(graphql.Float)},
		"discount":       &graphql.Field{Type: graphql.Float},
		"quantity":       &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
		"shop_available": &graphql.Field{Type: graphql.Boolean},
		"content":        &graphql.Field{Type: graphql.String},
		"created_at":     &graphql.Field{Type: graphql.NewNonNull(graphql.DateTime)},
		"updated_at":     &graphql.Field{Type: graphql.DateTime},
		"published_at":   &graphql.Field{Type: graphql.DateTime},
		"starts_at":      &graphql.Field{Type: graphql.DateTime},
		"ends_at":        &graphql.Field{Type: graphql.DateTime},
		"images":         &graphql.Field{Type: graphql.NewList(ImageType)},
	},
})

var ImageType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Image",
	Description: "For fetching image info of product",
	Fields: graphql.Fields{
		"id":       &graphql.Field{Type: graphql.NewNonNull(graphql.ID)},
		"uri":      &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"filename": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"main":     &graphql.Field{Type: graphql.Boolean},
	},
})
