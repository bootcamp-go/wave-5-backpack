package handler

import (
	"goweb/productos_capas/internal/products"
	"net/http"

	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre        string `form:"nombre" json:"nombre"`
	Color         string `form:"color" json:"color"`
	Precio        int    `form:"precio" json:"precio"`
	Stock         int    `form:"stock" json:"stock"`
	Codigo        string `form:"codigo" json:"codigo"`
	Publicado     bool   `form:"publicado" json:"publicado"`
	FechaCreacion string `form:"fecha_creacion" json:"fecha_creacion"`
}

type Product struct {
	service products.Service
}

func NewProduct(s products.Service) *Product {
	return &Product{
		service: s,
	}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "12345" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		var req request
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		p, err := p.service.GetAll(req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.FechaCreacion)
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, p)
	}
}

func (p *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "12345" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		p, err := p.service.Store(req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.FechaCreacion)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}

func (p *Product) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "12345" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		id := ctx.Param("id")
		p, err := p.service.GetByID(id)

		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}
