package main

import (
	"net/http"
	"time"

	"github.com/arfaghif/TGTCx/backend/handlers"
	"github.com/arfaghif/TGTCx/backend/server"
	"github.com/gorilla/mux"
)

func main() {

	// Init database connection
	// database.InitDB()

	// Init serve HTTP
	router := mux.NewRouter()

	// routes http

	router.HandleFunc("/update-product", handlers.UpdateProduct).Methods(http.MethodPatch)
	router.HandleFunc("/delete-product", handlers.DeleteProduct).Methods(http.MethodDelete)

	router.HandleFunc("/banner", handlers.UploadBanner).Methods(http.MethodPost)

	serverConfig := server.Config{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Port:         8000,
	}
	server.Serve(serverConfig, router)
}
