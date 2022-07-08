package handler

import (
	"errors"
	"goweb/clase1_clase2/internal/products"

	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre    string `form:"nombre" json:"nombre"`
	Color     string `form:"color" json:"color"`
	Precio    int    `form:"precio" json:"precio"`
	Stock     int    `form:"stock" json:"stock"`
	Codigo    string `form:"codigo" json:"codigo"`
	Publicado bool   `form:"publicado" json:"publicado"`
	Fecha     string `form:"fecha" json:"fecha"`
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
		var req request
		token := ctx.GetHeader("token")

		if err := validateToken(token); err != nil {
			ctx.JSON(401, gin.H{"error: ": err.Error()})
			return
		}

		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		pr, err := p.service.GetAll(req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.Fecha)

		if err != nil {
			ctx.JSON(404, gin.H{
				"error: ": err.Error(),
			})
		}
		ctx.JSON(200, pr)
	}
}

func (p *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("token")

		if err := validateToken(token); err != nil {
			ctx.JSON(401, gin.H{"error: ": err.Error()})
			return
		}

		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		pr, err := p.service.Store(req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.Fecha)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, pr)
	}
}

func validateToken(token string) error {
	if token == "" {
		return errors.New("no ingresó el token y es requerido")
	}
	if token != "123456" {
		return errors.New("no tiene permisos para realizar la petición solicitada")
	}
	return nil
}

func (p *Product) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		token := ctx.GetHeader("token")

		if err := validateToken(token); err != nil {
			ctx.JSON(401, gin.H{"error: ": err.Error()})
			return
		}

		producto, err := p.service.GetById(id)
		if err != nil {
			ctx.JSON(401, gin.H{"error: ": err.Error()})
			return
		}
		ctx.JSON(200, producto)
	}
}
