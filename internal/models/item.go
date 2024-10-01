package models

// Item represents a single item in the inventory.
type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
