package repositories

import (
	"errors"
	"my-go-api/internal/models"
	"strconv"
)

// In-memory item storage
var items []models.Item
var nextID = 1

// GetItems retrieves all items.
func GetItems() []models.Item {
	return items
}

// GetItemByID retrieves an item by its ID.
func GetItemByID(id int) (models.Item, error) {
	for _, item := range items {
		if item.ID == id {
			return item, nil
		}
	}
	return models.Item{}, errors.New("item with ID " + strconv.Itoa(id) + " not found")
}

// The function `RemoveItemByID` removes an item from a collection by its ID and returns the removed
// item or an error if the item is not found.
func RemoveItemByID(id int) (models.Item, error) {
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			return item, nil
		}
	}
	return models.Item{}, errors.New("item with ID " + strconv.Itoa(id) + " not found")
}

// The `UpdateItemByID` function is responsible for updating an existing item in the collection based
// on its ID. It takes two parameters: `id` which is the ID of the item to be updated, and `newItem`
// which is the updated item data.
func UpdateItemByID(id int, newItem models.Item) (models.Item, error) {
	for i, item := range items {
		if item.ID == id {
			items[i].Name = newItem.Name
			return items[i], nil
		}
	}
	return models.Item{}, errors.New("item with ID " + strconv.Itoa(id) + " not found")
}

// CreateItem adds a new item to the list.
func CreateItem(item models.Item) models.Item {
	item.ID = nextID
	nextID++
	items = append(items, item)
	return item
}
