package domain

type Usuario struct {
	Id          int     `json:"id"`
	Names       string  `json:"names"`
	LastName    string  `json:"last_name"`
	Age         int     `json:"age"`
	DateCreated string  `json:"date_created"`
	Estatura    float64 `json:"height"`
	Email       string  `json:"email"`
	IsActivo    bool    `json:"is_active"`
}

type UserResult struct {
	Usuario  Usuario
	Posicion int
}

type Usuarios struct {
	Users []Usuario
}
