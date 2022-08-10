package domain

type Warehouse struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type WarehouseAndUser struct {
	Warehouse
	Users []User `json:"users"`
}