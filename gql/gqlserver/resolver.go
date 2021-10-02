package gqlserver

import (
	"log"
	"strings"

	"github.com/arfaghif/TGTCx/backend/service"
	"github.com/graphql-go/graphql"
)

type Resolver struct {
}

func NewResolver() *Resolver {
	return &Resolver{}
}

func (r *Resolver) GetProduct() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["product_id"].(int)
		product, err := service.GetProduct(
			id,
		)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		// update to use Usecase from previous session
		return product, err
	}
}

func (r *Resolver) AddBannerTags() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["id"].(int)
		tags := p.Args["tags"].(string)
		err := service.AddTagBanner(
			id,
			strings.Split(tags, ","),
		)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		// update to use Usecase from previous session
		return nil, err
	}
}
