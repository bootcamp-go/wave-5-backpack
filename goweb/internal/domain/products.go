package domain

type Products struct {
	Id        int     `json:"id"`
	Nombre    string  `json:"nombre" binding:"required"`
	Color     string  `json:"color" binding:"required"`
	Precio    float64 `json:"precio" binding:"required"`
	Stock     int     `json:"stock" binding:"required"`
	Codigo    string  `json:"codigo" binding:"required"`
	Publicado bool    `json:"publicado" binding:"required"`
	Fecha     string  `json:"fecha_creacion" binding:"required"`
}
