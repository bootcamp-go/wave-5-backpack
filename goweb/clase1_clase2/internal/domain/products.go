package domain

type Product struct {
	Id        int    `form:"id" json:"id"`
	Nombre    string `form:"nombre" json:"nombre" binding:"required"`
	Color     string `form:"color" json:"color" binding:"required"`
	Precio    int    `form:"precio" json:"precio" binding:"required"`
	Stock     int    `form:"stock" json:"stock" binding:"required"`
	Codigo    string `form:"codigo" json:"codigo" binding:"required"`
	Publicado bool   `form:"publicado" json:"publicado"`
	Fecha     string `form:"fecha" json:"fecha" binding:"required"`
}
