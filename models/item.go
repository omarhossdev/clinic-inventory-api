package models

type Item struct {
	ID       string `json:"id"` // Go field: ID   -> JSON key: "id"
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}

// Capitalize Inventory so other packages can read/write to it!
// In-memory data store for now
var Inventory = []Item{
	{ID: "1", Name: "Amoxicillin 500mg", Quantity: 150, Unit: "boxes"},
	{ID: "2", Name: "Paracetamol 500mg", Quantity: 500, Unit: "tablets"},
	{ID: "3", Name: "Oral Rehydration Salts", Quantity: 200, Unit: "sachets"},
}
