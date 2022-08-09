package domain

// Warehouse ...
type Warehouse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

// WarehouseAndProducts
type WarehouseAndProducts struct {
	Warehouse
	Transactions []Transaction `json:"transactions"`
}
