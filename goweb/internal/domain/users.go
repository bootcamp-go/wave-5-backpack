package domain

type User struct {
	Id             int     `json:"id"`
	Edad           int     `json:"edad"`
	Nombre         string  `json:"nombre"`
	Apellido       string  `json:"apellido"`
	Email          string  `json:"email"`
	Fecha_creacion string  `json:"fecha_creacion"`
	Altura         float64 `json:"altura"`
	Activo         bool    `json:"activo"`
}
