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
					Resolve: s.productResolver.GetProduct(),
				},
				"Products": &graphql.Field{
					Type:        graphql.NewList(ProductType),
					Description: "Get product by ID",
					// Args: graphql.FieldConfigArgument{
					// 	"product_id": &graphql.ArgumentConfig{
					// 		Type: graphql.Int,
					// 	},
					// },
					Resolve: s.productResolver.GetProduct(),
				},
			},
		}),

		// uncomment this and add objects for mutation
		// Mutation: graphql.NewObject(graphql.ObjectConfig{}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:        "ProductCreate",
			Description: "Create a new product",
			Fields: graphql.Fields{
				"CreateProducts": &graphql.Field{
					Type:        ProductType,
					Description: "Product Create",
					Args: graphql.FieldConfigArgument{
						"product_id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"name": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"product_price": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"image_url": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"shop_name": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: s.productResolver.GetProduct(),
				},
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
			},
		}),
	})

	if err != nil {
		return err
	}

	s.Schema = schema

	return nil
}
