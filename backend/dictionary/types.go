package dictionary

import (
	"time"
)

type Product struct {
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	ShopName     string  `json:"shop_name"`
	ProductPrice float64 `json:"product_price"`
	ImageURL     string  `json:"image_url"`
}

type Banner struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImgPath     string    `json:"image_path"`
	Tags        []string  `json:"tags"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}

type Tag struct {
	ID  int    `json:"id"`
	Tag string `json:"tag"`
}

type User struct {
	ID            int    `json:"id"`
	Age           string `json:"age"`
	Region        string `json:"region"`
	Gender        string `json:"gender"`
	Tier          string `json:"tier"`
	WalletBalance int    `json:"wallet_balance"`
	ProductPref   string `json:"product_pref"`
}

type APIResponse struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error_message"`
}
