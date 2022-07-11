package domain

// Se debe crear la estructura de la entidad
type User struct {
	ID 				int		`json:"id"`
	Nombre 			string	`json:"nombre"`
	Apellido 		string	`json:"apellido"`
	Email 			string	`json:"email"`
	Edad 			int		`json:"edad"`
	Altura 			float64	`json:"altura"`
	Activo			bool	`json:"activo"`
	FechaCreacion 	string	`json:"fecha_creacion"`
}