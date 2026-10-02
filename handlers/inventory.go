package handlers

import (
	"clinic-inventory/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetInventory(c *gin.Context) {
	c.JSON(http.StatusOK, models.Inventory)
}

func GetItemById(c *gin.Context) {
	id := c.Param("id") // extract the id param

	// 1. loop through inventory
	// 2. find the item with the id
	// 3. if found then return it otherwise error
	for _, item := range models.Inventory {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}

// Add a new item to inventory (POST)
func CreateItem(c *gin.Context) {
	var newItem models.Item

	if err := c.ShouldBindJSON(&newItem); err != nil { // if there is an error/ the error is not empty but already found
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add item to our inventory
	models.Inventory = append(models.Inventory, newItem)

	// Return the newly created item with 201 StatusCreated
	c.JSON(http.StatusCreated, newItem)
}

func UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var updatedItem models.Item

	if err := c.ShouldBindJSON(&updatedItem); err != nil { // if there is an error from the request
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updatedItem.ID == "" {
		updatedItem.ID = id
	}

	for i, item := range models.Inventory {
		if item.ID == id {
			models.Inventory[i] = updatedItem

			c.JSON(http.StatusOK, updatedItem)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "no matched item"})
}

func DeleteItem(c *gin.Context) {
	id := c.Param("id")

	for i, item := range models.Inventory {
		if item.ID == id {
			models.Inventory = append(models.Inventory[:i], models.Inventory[i+1:]...) // add items from the start till that item, and include all items after that item

			c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "no matched items"})
}
