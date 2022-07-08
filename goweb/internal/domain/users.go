package domain

import "time"

// Estructura de usuarios
type ModelUser struct {
	Id            int       `json:"id"`
	Nombre        string    `json:"nombre"`
	Apellido      string    `json:"apellido"`
	Email         string    `json:"email"`
	Edad          int       `json:"edad"`
	Altura        float64   `json:"altura"`
	Activo        bool      `json:"activo"`
	FechaCreacion time.Time `json:"fecha_creacion"`
}
