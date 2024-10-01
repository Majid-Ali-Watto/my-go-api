package v1

import (
	"my-go-api/internal/handlers"

	"github.com/gorilla/mux"
)

// SetupRoutes initializes the API routes.
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/items/create", handlers.CreateItem).Methods("POST")     //C
	router.HandleFunc("/items", handlers.GetItems).Methods("GET")               //R
	router.HandleFunc("/items/{id}", handlers.GetItemByID).Methods("GET")       //R
	router.HandleFunc("/items/{id}", handlers.UpdateItemByID).Methods("PATCH")  //U
	router.HandleFunc("/items/{id}", handlers.RemoveItemByID).Methods("DELETE") //D
	return router
}
