package domain

type Product struct {
	ID        int     `json:"id"`
	Name      string  `json:"name" binding:"required"`
	Type      string  `json:"type" binding:"required"`
	Count     int     `json:"count" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
	Warehouse int     `json:"warehouse" binding:"required"`
}

type ProductAndWarehouse struct {
	Product
	Warehouse Warehouse `json:"warehouse"`
}
