package domain

type Productos struct {
	Id            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreación string  `json:"fecha_creacion"`
	IdWarehouse   int     `json:"id_warehouse"`
}

type ProductosWarehouse struct {
	Id            int       `json:"id"`
	Nombre        string    `json:"nombre"`
	Color         string    `json:"color"`
	Precio        float64   `json:"precio"`
	Stock         int       `json:"stock"`
	Codigo        string    `json:"codigo"`
	Publicado     bool      `json:"publicado"`
	FechaCreación string    `json:"fecha_creacion"`
	IdWarehouse   int       `json:"id_warehouse"`
	Warehouses    Warehouse `json:"warehouse"`
}
