package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var items []Item
var nextID = 1

// Response structure for API responses
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// Handler to get all items
func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Status: "success",
		Data:   items,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Handler to create a new item
func createItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		response := Response{
			Status:  "error",
			Message: "Invalid input",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	newItem.ID = nextID
	nextID++
	items = append(items, newItem)

	response := Response{
		Status:  "success",
		Message: "Item created successfully",
		Data:    newItem,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// Handler for not found
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Status:  "error",
		Message: "Endpoint not found",
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(response)
}

func main() {
	items = []Item{}

	http.HandleFunc("/items", getItems)
	http.HandleFunc("/items/create", createItem)
	http.HandleFunc("/", notFoundHandler)

	// Start the server and log the status
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
