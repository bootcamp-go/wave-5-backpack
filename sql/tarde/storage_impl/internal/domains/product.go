package domains

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"nombre"`
	Type        string  `json:"tipo"`
	Count       int     `json:"cantidad"`
	Price       float64 `json:"precio"`
	WarehouseId int     `json:"warehouse_id" binding:"required"`
}

type Products struct {
	ID               int     `json:"id"`
	Name             string  `json:"nombre"`
	Type             string  `json:"tipo"`
	Count            int     `json:"cantidad"`
	Price            float64 `json:"precio"`
	WarehouseName    string  `json:"warehouse_nombre"`
	WarehouseAddress string  `json:"warehouse_direccion"`
}
