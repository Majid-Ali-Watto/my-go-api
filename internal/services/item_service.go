package services

import (
	"my-go-api/internal/models"
	"my-go-api/internal/repositories"
)

// GetAllItems retrieves all items from the repository.
func GetAllItems() []models.Item {
	return repositories.GetItems()
}

// GetItemByID retrieves an item by its ID.
func GetItemByID(id int) (models.Item, error) {
	item, err := repositories.GetItemByID(id)
	if err != nil {
		return models.Item{}, err // Return zero value and the error
	}
	return item, nil
}
func RemoveItemByID(id int) (models.Item, error) {
	item, err := repositories.RemoveItemByID(id)
	if err != nil {
		return models.Item{}, err // Return zero value and the error
	}
	return item, nil
}

func UpdateItemByID(id int, newItem models.Item) (models.Item, error) {
	item, err := repositories.UpdateItemByID(id, newItem)
	if err != nil {
		return models.Item{}, err // Return zero value and the error
	}
	return item, nil
}

// AddItem creates a new item in the repository.
func AddItem(item models.Item) models.Item {
	return repositories.CreateItem(item)
}
