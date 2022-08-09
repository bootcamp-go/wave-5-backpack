package domain

import "time"

type Usuarios struct {
	Id       int       `json:"id"`
	Nombre   string    `json:"nombre"`
	Apellido string    `json:"apellido"`
	Email    string    `json:"email"`
	Edad     int       `json:"edad"`
	Altura   float64   `json:"altura"`
	Activo   bool      `json:"activo"`
	Fecha    time.Time `json:"fecha"`
}

type UserAndWarehouse struct {
	Usuarios
	Warehouse Warehouse `json:"warehouse"`
}

//Para pruebas con POST
// {
//     "nombre": "Martha",
//     "apellido": "Hernandez",
//     "email": "martha.hernandez",
//     "edad": 60,
//     "altura": 1.60
// }
