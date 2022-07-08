package domain

type Product struct {
	ID            int     `json:"ID"`
	Nombre        string  `json:"Nombre"`
	Color         string  `json:"Color"`
	Precio        float64 `json:"Precio"`
	Stock         int     `json:"Stock"`
	Codigo        string  `json:"Codigo"`
	Publicado     bool    `json:"Publicado"`
	FechaCreacion string  `json:"FechaCreacion"`
}
