package domain

type Usuario struct {
	Id             int    `json:"id"`
	Nombre         string `json:"nombre"`
	Apellido       string `json:"apellido"`
	Email          string `json:"email"`
	Edad           int    `json:"edad"`
	Altura         int    `json:"altura"`
	Activo         bool   `json:"activo"`
	Fecha_creacion string `json:"fecha_creacion"`
}
