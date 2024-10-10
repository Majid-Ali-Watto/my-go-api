package handlers

import (
	"encoding/json"
	"my-go-api/internal/models"
	"my-go-api/internal/services"
	"my-go-api/internal/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetItems handles the GET request to retrieve all items.
func GetItems(w http.ResponseWriter, r *http.Request) {
	items := services.GetAllItems()
	response := utils.Response{
		Status: "success",
		Data:   items,
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

// GetItemByID handles the GET request to retrieve a single item by ID.
func GetItemByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Get path variables
	idStr := vars["id"] // Extract ID from path
	itemID, err := strconv.Atoi(idStr)
	if err != nil || itemID < 0 {
		response := utils.Response{
			Status:  "error",
			Message: "Invalid ID",
		}
		utils.SendJSONResponse(w, http.StatusBadRequest, response)
		return
	}
	item, err := services.GetItemByID(itemID)
	if err != nil {
		response := utils.Response{
			Status:  "error",
			Message: err.Error(),
		}
		utils.SendJSONResponse(w, http.StatusNotFound, response)
		return
	}
	response := utils.Response{
		Status: "success",
		Data:   item,
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

// The function `RemoveItemByID` handles the removal of an item by ID from a server
func RemoveItemByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Get path variables
	idStr := vars["id"] // Extract ID from path
	itemID, err := strconv.Atoi(idStr)
	if err != nil || itemID < 0 {
		response := utils.Response{
			Status:  "error",
			Message: "Invalid ID",
		}
		utils.SendJSONResponse(w, http.StatusBadRequest, response)
		return
	}
	item, err := services.RemoveItemByID(itemID)
	if err != nil {
		response := utils.Response{
			Status:  "error",
			Message: err.Error(),
		}
		utils.SendJSONResponse(w, http.StatusNotFound, response)
		return
	}
	response := utils.Response{
		Status: "success",
		Data:   item,
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

// This function updates an item by ID based on the request data and sends appropriate JSON
// responses.
func UpdateItemByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Get path variables
	idStr := vars["id"] // Extract ID from path
	itemID, err := strconv.Atoi(idStr)
	if err != nil || itemID < 0 {
		response := utils.Response{
			Status:  "error",
			Message: "Invalid ID",
		}
		utils.SendJSONResponse(w, http.StatusBadRequest, response)
		return
	}
	var newItem models.Item
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		response := utils.Response{
			Status:  "error",
			Message: "Invalid body",
		}
		utils.SendJSONResponse(w, http.StatusBadRequest, response)
		return
	}
	item, err := services.UpdateItemByID(itemID, newItem)
	if err != nil {
		response := utils.Response{
			Status:  "error",
			Message: err.Error(),
		}
		utils.SendJSONResponse(w, http.StatusNotFound, response)
		return
	}
	response := utils.Response{
		Status: "success",
		Data:   item,
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

// CreateItem handles the POST request to create a new item.
func CreateItem(w http.ResponseWriter, r *http.Request) {
	var newItem models.Item
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		response := utils.Response{
			Status:  "error",
			Message: "Invalid body",
		}
		utils.SendJSONResponse(w, http.StatusBadRequest, response)
		return
	}
	createdItem := services.AddItem(newItem)

	response := utils.Response{
		Status:  "success",
		Message: "Item created successfully",
		Data:    createdItem,
	}
	utils.SendJSONResponse(w, http.StatusCreated, response)
}
