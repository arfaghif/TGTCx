package service

import (
	"fmt"

	"github.com/arfaghif/TGTCx/backend/database"
	"github.com/arfaghif/TGTCx/backend/dictionary"
)

func SampleFunction() {
	fmt.Printf("My Service!")

	// // you can connect and
	// // get current database connection
	// db := database.GetDB()

	// // construct query
	// query := `
	// SELECT something FROM table_something WHERE id = $1
	// `
	// // actual query process
	// row = db.QueryRow(query, paramID)

	// // read query result, and assign to variable(s)
	// err = row.Scan(&ID, &name)
}

func AddProduct(product dictionary.Product) error {
	// // you can connect and
	// // get current database connection
	db := database.GetDB()

	// // construct query
	query := `
		INSERT INTO products (product_name, product_price, product_image, shop_name) VALUES($1, $2, $3, $4)
	`
	// actual query process

	// // read query result, and assign to variable(s)
	_, err := db.Exec(query, product.Name, product.ProductPrice, product.ImageURL, product.ShopName)
	return err
}

func GetProduct(id int) (dictionary.Product, error) {

	// // you can connect and
	// // get current database connection
	db := database.GetDB()

	// // construct query
	query := `
	SELECT product_id, product_name, product_price, product_image, shop_name
	FROM products 
	WHERE product_id = $1
	`
	// defer db.Close()
	// // actual query process
	row := db.QueryRow(query, id)

	product := dictionary.Product{}
	// // read query result, and assign to variable(s)
	err := row.Scan(&product.ID, &product.Name, &product.ProductPrice, &product.ImageURL, &product.ShopName)
	return product, err
}

// func GetProducts() ([]dictionary.Product, error) {

// 	// // you can connect and
// 	// // get current database connection
// 	db := database.GetDB()
// 	query := `
// 	SELECT product_id, product_name, product_price, product_image, shop_name
// 	FROM products
// 	`
// 	// defer db.Close()
// 	// // actual query process
// 	rows, _ := db.Query(query)

// 	defer rows.Close()

// 	// An album slice to hold data from returned rows.
// 	var product []dictionary.Product

// 	// Loop through rows, using Scan to assign column data to struct fields.
// 	for rows.Next() {
// 		var alb Album
// 		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist,
// 			&alb.Price, &alb.Quantity); err != nil {
// 			return albums, err
// 		}
// 		albums = append(albums, album)
// 	}
// 	if err = rows.Err(); err != nil {
// 		return albums, err
// 	}

// 	product := dictionary.Product{}
// 	// // read query result, and assign to variable(s)
// 	err := row.Scan(&product.ID, &product.Name, &product.ProductPrice, &product.ImageURL, &product.ShopName)
// 	return product, err
// }
