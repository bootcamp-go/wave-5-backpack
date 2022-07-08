package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"web-server/internal/products"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type request struct {
	Nombre    string  `json:"nombre" binding:"required" `
	Color     string  `json:"color" binding:"required"`
	Precio    float64 `json:"precio" binding:"required"`
	Stock     int     `json:"stock" binding:"required"`
	Codigo    string  `json:"codigo" binding:"required"`
	Publicado bool    `json:"publicado"`
	Fecha     string  `json:"fecha_creacion" binding:"required"`
}

type Products struct {
	service products.Service
}

func NewProduct(p products.Service) *Products {
	return &Products{
		service: p,
	}
}

func (c *Products) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, p)
	}
}

func (c *Products) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		p, err := c.service.Store(req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.Fecha)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}

func (p *Products) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}
		id, erro := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if erro != nil {
			ctx.JSON(400, gin.H{"error": "Id inválido"})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			errs := err.(validator.ValidationErrors)
			for _, valError := range errs {
				if valError.Tag() == "required" {
					ctx.JSON(400, gin.H{
						"error": fmt.Sprintf("el campo '%s' es requerido", valError.Field()),
					})
					return
				}
			}
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		p, err := p.service.Update(int(id), req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.Fecha)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, p)

	}
}

func (p *Products) UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}
		id, erro := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if erro != nil {
			ctx.JSON(400, gin.H{"error": "Id inválido"})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			errs := err.(validator.ValidationErrors)
			for _, valError := range errs {
				if valError.Tag() == "required" {
					ctx.JSON(400, gin.H{
						"error": fmt.Sprintf("el campo '%s' es requerido", valError.Field()),
					})
					return
				}
			}
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		p, err := p.service.UpdateName(int(id), req.Nombre)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, p)
	}
}

func (p *Products) UpdatePrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}
		id, erro := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if erro != nil {
			ctx.JSON(400, gin.H{"error": "Id inválido"})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			errs := err.(validator.ValidationErrors)
			for _, valError := range errs {
				if valError.Tag() == "required" {
					ctx.JSON(400, gin.H{
						"error": fmt.Sprintf("el campo '%s' es requerido", valError.Field()),
					})
					return
				}
			}
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		p, err := p.service.UpdatePrice(int(id), req.Precio)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, p)
	}
}

func (p *Products) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "123456" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Id inválido"})
			return
		}

		err = p.service.Delete(int(id))
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": fmt.Sprintf("El producto %d ha sido eliminado", id)})
	}
}
