package domain

type Products struct {
	Id            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"código"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fecha_de_creación"`
}
