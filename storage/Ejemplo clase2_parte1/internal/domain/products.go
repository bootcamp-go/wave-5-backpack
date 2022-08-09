package domain

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"nombre" binding:"required"`
	Type  string  `json:"tipo" binding:"required"`
	Count int     `json:"cantidad" binding:"required"`
	Price float64 `json:"precio" binding:"required"`
}

type ProductAndWarehouse struct {
	Product
	Warehouse Warehouse `json:"warehouse"`
}
