package domain

type Product struct {
	ID              int     `json:"id"`
	Name            string  `json:"nombre" binding:"required"`
	Type            string  `json:"tipo" binding:"required"`
	Count           int     `json:"cantidad" binding:"required"`
	Price           float64 `json:"precio" binding:"required"`
	Warehouse       string  `json:"warehouse,omitempty"`
	WarehouseAdress string  `json:"warehouse_address,omitempty"`
}
