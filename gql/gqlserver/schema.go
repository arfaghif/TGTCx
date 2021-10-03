package gqlserver

import (
	"github.com/graphql-go/graphql"
)

type SchemaWrapper struct {
	productResolver *Resolver
	Schema          graphql.Schema
}

func NewSchemaWrapper() *SchemaWrapper {
	return &SchemaWrapper{}
}

func (s *SchemaWrapper) WithProductResolver(pr *Resolver) *SchemaWrapper {
	s.productResolver = pr

	return s
}

func (s *SchemaWrapper) Init() error {
	// add gql schema as needed
	schema, err := graphql.NewSchema(graphql.SchemaConfig{

		// Mutation: graphql.NewObject(graphql.ObjectConfig{}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:        "ProductCreate",
			Description: "Create a new product",
			Fields: graphql.Fields{
				"AddBannerTag": &graphql.Field{
					Type:        BannerType,
					Description: "Add Tag a Banner",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"tags": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: s.productResolver.AddBannerTags(),
				},
				"UpdateBanner": &graphql.Field{
					Type:        BannerType,
					Description: "Update Banner Information",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
						"name": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"description": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"start_date": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"end_date": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: s.productResolver.UpdateBanner(),
				},
			},
		}),
	})

	if err != nil {
		return err
	}

	s.Schema = schema

	return nil
}
