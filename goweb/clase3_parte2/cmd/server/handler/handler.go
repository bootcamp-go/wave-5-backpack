package handler

import (
	"fmt"
	"os"
	"strconv"

	"github.com/bootcamp-go/wave-5-backpack/tree/ospina_christian/goweb/clase3_parte2/internal/productos"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreación string  `json:"fecha_creacion"`
}

type Product struct {
	service productos.Service
}

func NewProduct(s productos.Service) *Product {
	return &Product{service: s}
}

func (p *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{"error": "Token Invalido"})
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.Nombre == "" {
			ctx.JSON(400, gin.H{"error": "El producto no puede ser vacio"})
			return
		}
		if req.Color == "" {
			ctx.JSON(400, gin.H{"error": "El color no puede ser vacio"})
			return
		}
		if req.Precio < 0 {
			ctx.JSON(400, gin.H{"error": "El precio no puede ser menor a 0"})
			return
		}
		if req.Stock < 0 {
			ctx.JSON(400, gin.H{"error": "El stock no puede ser menor a 0"})
			return
		}
		if req.Codigo == "" {
			ctx.JSON(400, gin.H{"error": "El codigo no puede ser vacio"})
			return
		}
		if req.FechaCreación == "" {
			ctx.JSON(400, gin.H{"error": "La fecha de creacion no puede ser vacio"})
			return
		}
		p, err := p.service.Store(req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.FechaCreación)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"Products": p})
	}
}

func (p *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{"error": "Token Invalido"})
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Id invalido"})
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.Nombre == "" {
			ctx.JSON(400, gin.H{"error": "El producto no puede ser vacio"})
			return
		}
		if req.Color == "" {
			ctx.JSON(400, gin.H{"error": "El color no puede ser vacio"})
			return
		}
		if req.Precio < 0 {
			ctx.JSON(400, gin.H{"error": "El precio no puede ser menor a 0"})
			return
		}
		if req.Stock < 0 {
			ctx.JSON(400, gin.H{"error": "El stock no puede ser menor a 0"})
			return
		}
		if req.Codigo == "" {
			ctx.JSON(400, gin.H{"error": "El codigo no puede ser vacio"})
			return
		}
		if req.FechaCreación == "" {
			ctx.JSON(400, gin.H{"error": "La fecha de creacion no puede ser vacio"})
			return
		}
		p, err := p.service.Update(int(id), req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.FechaCreación)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)

	}
}

func (p *Product) UpdatePrecio() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{"error": "Token Invalido"})
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Id Invalido"})
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.Precio < 0 {
			ctx.JSON(400, gin.H{"error": "el precio debe ser indicado"})
			return
		}
		p, err := p.service.UpdatePrecio(int(id), req.Precio)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{"error": "Token Invalido"})
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Id invalido"})
			return
		}
		err = p.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"data": fmt.Sprintf("El producto %d ha sido eliminado", id)})
	}
}

func (p *Product) GetForId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{"error": "Token Invalido"})
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Id invalido"})
			return
		}
		p, err := p.service.GetForId(int(id))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{"error": "Token Invalido"})
			return
		}
		p, err := p.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}
