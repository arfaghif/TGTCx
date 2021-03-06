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
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:        "ProductGetter",
			Description: "All query related to getting product data",
			Fields: graphql.Fields{
				"ProductDetail": &graphql.Field{
					Type:        ProductType,
					Description: "Get product by ID",
					Args: graphql.FieldConfigArgument{
						"product_id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					// To Be Changed
					Resolve: s.productResolver.AddBannerTags(),
				},
				"Products": &graphql.Field{
					Type:        graphql.NewList(ProductType),
					Description: "Get product by ID",
					// To Be Changed
					Resolve: s.productResolver.AddBannerTags(),
				},
				"BannerUsers": &graphql.Field{
					Type:        graphql.NewList(BannerType),
					Description: "Get product by ID",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: s.productResolver.GetBannerUser(),
				},
			},
		}),

		// uncomment this and add objects for mutation
		// Mutation: graphql.NewObject(graphql.ObjectConfig{}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:        "BannerUpdate",
			Description: "Update Banner Data",
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
						"image_path": &graphql.ArgumentConfig{
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
