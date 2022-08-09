package domain

type Product struct {
	Id            int     `json:"id" db:"id"`
	Nombre        string  `json:"nombre" db:"name"`
	Color         string  `json:"color" db:"color"`
	Precio        float64 `json:"precio" db:"price"`
	Stock         int     `json:"stock" db:"stock"`
	Codigo        string  `json:"codigo" db:"code"`
	Publicado     bool    `json:"publicado" db:"publish"`
	FechaCreacion string  `json:"fecha_de_creacion" db:"create_date"`
}
