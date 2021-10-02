package product

import (
	"context"
	"errors"

	"github.com/arfaghif/TGTCx/backend/dictionary"
)

var (
	mProduct map[int64]dictionary.Product = make(map[int64]dictionary.Product)
	idx                                   = int64(1)
)

func AddProduct(ctx context.Context, p dictionary.Product) dictionary.Product {
	p.ID = idx
	idx++

	mProduct[p.ID] = p
	return p
}

func GetProduct(ctx context.Context, id int64) (dictionary.Product, error) {
	if _, ok := mProduct[id]; ok {
		return mProduct[id], nil
	}
	return dictionary.Product{}, errors.New("product nil")
}

func DeleteProduct(ctx context.Context, id int64) {
	delete(mProduct, id)
}

func UpdateProduct(ctx context.Context, p dictionary.Product) error {
	if _, ok := mProduct[p.ID]; !ok {
		return errors.New("product nil")
	}

	mProduct[p.ID] = p
	return nil
}
