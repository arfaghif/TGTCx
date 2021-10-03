package gqlserver

import (
	"errors"
	"log"
	"strings"

	"github.com/arfaghif/TGTCx/backend/dictionary"
	time_helper "github.com/arfaghif/TGTCx/backend/helpers"
	"github.com/arfaghif/TGTCx/backend/service"
	"github.com/graphql-go/graphql"
)

type Resolver struct {
}

func NewResolver() *Resolver {
	return &Resolver{}
}

func (r *Resolver) AddBannerTags() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, ok := p.Args["id"].(int)
		if !ok {
			return dictionary.Banner{}, errors.New("id invalid")
		}

		tags, ok := p.Args["tags"].(string)
		if !ok {
			return dictionary.Banner{}, errors.New("tag invalid")
		}
		banner, err := service.AddTagBanner(
			id,
			strings.Split(tags, ","),
		)
		if err != nil {
			log.Println(err.Error())
			return dictionary.Banner{}, err
		}
		// update to use Usecase from previous session
		return banner, err
	}
}

func (r *Resolver) UpdateBanner() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		//TODO: Nil Check: OMIT EMPTY
		id := p.Args["id"].(int)

		name, ok := p.Args["name"].(string)
		if !ok {
			name = ""
		}

		description, ok := p.Args["description"].(string)
		if !ok {
			description = ""
		}

		start_date_string, ok := p.Args["description"].(string)
		if !ok {
			start_date_string = "0001-01-01 00:00:00 +0000 UTC"
		}

		end_date_string, ok := p.Args["description"].(string)
		if !ok {
			end_date_string = "0001-01-01 00:00:00 +0000 UTC"
		}

		start_date, err := time_helper.ParseTimestamp(start_date_string)
		end_date, err := time_helper.ParseTimestamp(end_date_string)

		err = service.UpdateBanner(
			id,
			name,
			description,
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
func (r *Resolver) GetBannerUser() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, ok := p.Args["id"].(int)
		if !ok {
			return dictionary.Banner{}, errors.New("id invalid")
		}

		banners, err := service.GetBannerUser(
			id,
		)

		if err != nil {
			log.Println(err.Error())
			return []dictionary.Banner{}, err
		}

		// update to use Usecase from previous session
		return banners, err
	}
}
