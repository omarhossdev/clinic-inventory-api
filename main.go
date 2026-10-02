package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID       string `json:"id"` // Go field: ID   -> JSON key: "id"
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}

// In-memory data store for now
var inventory = []Item{
	{ID: "1", Name: "Amoxicillin 500mg", Quantity: 150, Unit: "boxes"},
	{ID: "2", Name: "Paracetamol 500mg", Quantity: 500, Unit: "tablets"},
	{ID: "3", Name: "Oral Rehydration Salts", Quantity: 200, Unit: "sachets"},
}

func main() {
	r := gin.Default()

	// Basic welcome API
	r.GET("/welcome", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Welcome to my To-Do API!",
		})
	})

	// Get Inventory
	r.GET("/inventory", func(c *gin.Context) {
		c.JSON(http.StatusOK, inventory)
	})

	// Fetch an item from inventory
	r.GET("/inventory/:id", func(c *gin.Context) {
		id := c.Param("id") // extract the id param

		// 1. loop through inventory
		// 2. find the item with the id
		// 3. if found then return it otherwise error
		for _, item := range inventory {
			if item.ID == id {
				c.JSON(http.StatusOK, item)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
	})

	// Add a new item to inventory (POST)
	r.POST("/inventory", func(c *gin.Context) {
		var newItem Item

		if err := c.ShouldBindJSON(&newItem); err != nil { // if there is an error/ the error is not empty but already found
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Add item to our inventory
		inventory = append(inventory, newItem)

		// Return the newly created item with 201 StatusCreated
		c.JSON(http.StatusCreated, newItem)
	})

	// Update Existing data (PUT)
	r.PUT("/inventory/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedItem Item

		if err := c.ShouldBindJSON(&updatedItem); err != nil { // if there is an error from the request
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if updatedItem.ID == "" {
			updatedItem.ID = id
		}

		for i, item := range inventory {
			if item.ID == id {
				inventory[i] = updatedItem

				c.JSON(http.StatusOK, updatedItem)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "no matched item"})
	})

	// Deleting an item
	r.DELETE("/inventory/:id", func(c *gin.Context) {
		id := c.Param("id")

		for i, item := range inventory {
			if item.ID == id {
				inventory = append(inventory[:i], inventory[i+1:]...) // add items from the start till that item, and include all items after that item

				c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "no matched items"})
	})

	// Starting the server on port 8080
	r.Run(":8080")
}
