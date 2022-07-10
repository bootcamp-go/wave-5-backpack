package domain

type Users struct {
	Id       int     `json:"id"`
	Nombre   string  `json:"nombre"`
	Apellido string  `json:"apellido"`
	Edad     int     `json:"edad"`
	Altura   float64 `json:"altura"`
}
