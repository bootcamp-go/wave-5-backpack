package domain

type Usuario struct {
	Id          int     `json:"id"`
	Names       string  `json:"nombre"`
	LastName    string  `json:"apellido"`
	Age         int     `json:"edad"`
	DateCreated string  `json:"fechaCreacion"`
	Estatura    float64 `json:"altura"`
	Email       string  `json:"email"`
	IsActivo    bool    `json:"activo"`
}

type UserResult struct {
	Usuario  Usuario
	Posicion int
}

type Usuarios struct {
	Users []Usuario
}
