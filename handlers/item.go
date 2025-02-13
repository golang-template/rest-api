package handlers

import (
	"REST-API/utils"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
)

var mu sync.Mutex

var items = []Item{
	{ID: "1", Name: "Item one"},
	{ID: "2", Name: "Item two"},
	{ID: "3", Name: "Item three"},
}

type Item struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

func GetItems(c echo.Context) error {
	return utils.JSONResponse(c, http.StatusOK, items, "Items retrieved successfully")
}

func CreateItem(c echo.Context) error {
	item := new(Item)
	if err := c.Bind(item); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
	}

	if err := c.Validate(item); err != nil {
		return utils.ErrorResponse(c, http.StatusUnprocessableEntity, "Validation failed")
	}

	mu.Lock()
	items = append(items, *item)
	mu.Unlock()

	return utils.JSONResponse(c, http.StatusCreated, item, "Item created successfully")
}

func UpdateItem(c echo.Context) error {
	id := c.Param("id")
	return utils.JSONResponse(c, http.StatusOK, map[string]string{"id": id}, "Item updated successfully")
}

func DeleteItem(c echo.Context) error {
	id := c.Param("id")
	return utils.JSONResponse(c, http.StatusOK, map[string]string{"id": id}, "Item deleted successfully")
}
