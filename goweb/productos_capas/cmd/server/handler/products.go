package handler

import (
	"errors"
	"goweb/productos_capas/internal/products"
	"net/http"
	"strconv"

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
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		var req request
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		p, err := p.service.GetAll(req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.FechaCreacion)
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
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		p, err := p.service.GetByID(id)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}

func (p *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "12345" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		if err := validateBody(req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
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

func (p *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "12345" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		if err := validateBody(req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		p, err := p.service.Update(id, req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.FechaCreacion)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, p)
	}
}

func (p *Product) UpdateNamePrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "12345" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		if req.Nombre == "" {
			ctx.JSON(400, gin.H{"error": "El nombre del producto es requerido"})
			return
		}

		if req.Precio == 0 {
			ctx.JSON(400, gin.H{"error": "El precio del producto es requerido"})
			return
		}

		p, err := p.service.UpdateNamePrice(id, req.Nombre, req.Precio)

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
		if token != "12345" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		p, err := p.service.Delete(id)

		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, p)
	}
}

func validateBody(req request) error {
	if req.Nombre == "" {
		return errors.New("el nombre del producto es requerido")
	}

	if req.Color == "" {
		return errors.New("el color del producto es requerido")
	}

	if req.Precio == 0 {
		return errors.New("el precio del producto es requerido")
	}

	if req.Stock == 0 {
		return errors.New("el stock del producto es requerido")
	}

	if req.Codigo == "" {
		return errors.New("el codigo del producto es requerido")
	}

	if req.FechaCreacion == "" {
		return errors.New("la fecha de creacion del producto es requerido")
	}
	return nil
}
