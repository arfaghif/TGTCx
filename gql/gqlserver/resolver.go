package gqlserver

import (
	"log"
	"strings"

	"github.com/arfaghif/TGTCx/backend/helper"
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

func (r *Resolver) UpdateBanner() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["id"].(int)
		name := p.Args["name"].(string)
		description := p.Args["description"].(string)
		image_path := p.Args["image_path"].(string)
		start_date, err := helper.ParseTimestamp(p.Args["start_date"].(string))
		end_date, err := helper.ParseTimestamp(p.Args["end_date"].(string))

		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		err = service.UpdateBanner(
			id,
			name,
			description,
			image_path,
			start_date,
			end_date,
		)

		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		// update to use Usecase from previous session
		return nil, err
	}
}
