package domain

type Warehouse struct {
	Id        int    `json:"id"`
	Nombre    string `json:"name"`
	Direccion string `json:"address"`
}
