package main

import (
	"clinic-inventory/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/inventory", handlers.GetInventory)
	r.GET("/inventory/:id", handlers.GetItemById)
	r.POST("/inventory", handlers.CreateItem)
	r.PUT("/inventory/:id", handlers.UpdateItem)
	r.DELETE("/inventory/:id", handlers.DeleteItem)

	// Starting the server on port 8080
	r.Run(":8080")
}
