package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/arfaghif/TGTCx/backend/dictionary"
	"github.com/arfaghif/TGTCx/backend/domain/product"
	"github.com/arfaghif/TGTCx/backend/service"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong\n")
}

func AddProduct(w http.ResponseWriter, r *http.Request) {

	var p dictionary.Product

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	if err := service.AddProduct(p); err != nil {
		log.Println(err.Error())
		http.Error(w, "unprocessibility entity", 422)
		return
	}

	fmt.Fprintf(w, fmt.Sprint("success, id product : ", p.ID))
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	idstring := r.URL.Query().Get("id")

	idInt64, err := strconv.ParseInt(idstring, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	p, err := service.GetProduct(int(idInt64))
	if err != nil {
		// log.Fatal(err)
		log.Println(err.Error())
		fmt.Fprintf(w, err.Error())
	}

	val, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, string(val))
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var p dictionary.Product

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	product.DeleteProduct(context.Background(), p.ID)
	fmt.Fprintf(w, "success")
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var p dictionary.Product

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	product.UpdateProduct(context.Background(), p)

	fmt.Fprintf(w, "success")
}
