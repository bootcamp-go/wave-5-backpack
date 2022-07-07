package handler

import (
	"errors"
	"fmt"
	"goweb/internal/domain"
	"goweb/internal/products"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	Nombre        string  `json:"nombre" binding:"required"`
	Color         string  `json:"color" binding:"required"`
	Precio        float64 `json:"precio" binding:"required"`
	Stock         int     `json:"stock" binding:"required"`
	Codigo        string  `json:"codigo" binding:"required"`
	Publicado     bool    `json:"publicado" binding:"required"`
	FechaCreacion string  `json:"fechaCreacion" binding:"required"`
}

type Product struct {
	service products.Service
}

func NewProduct(s products.Service) *Product {
	return &Product{service: s}
}

// ==================================
// Funciones de clases anteriores
// ==================================

// Clase 1 Ejercicio 1 Parte 1
func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := p.service.GetAll()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Ejercicio 1 parte 2
		var filtrados []domain.Product
		for _, product := range products {

			var idVerif, nombreVerif, colorVerif, precioVerif, stockVerif, codigoVerif, publicadoVerif, fechaCreacionVerif bool

			idVerif = ctx.Query("id") != "" && ctx.Query("id") != fmt.Sprintf("%d", product.Id)
			nombreVerif = ctx.Query("nombre") != "" && ctx.Query("nombre") != product.Nombre
			colorVerif = ctx.Query("color") != "" && ctx.Query("color") != product.Color
			precioVerif = ctx.Query("precio") != "" && ctx.Query("precio") != fmt.Sprintf("%f", product.Precio)
			stockVerif = ctx.Query("stock") != "" && ctx.Query("stock") != fmt.Sprintf("%d", product.Stock)
			codigoVerif = ctx.Query("codigo") != "" && ctx.Query("codigo") != product.Codigo
			publicadoVerif = ctx.Query("publicado") != "" && ctx.Query("publicado") != fmt.Sprintf("%t", product.Publicado)
			fechaCreacionVerif = ctx.Query("fechaCreacion") != "" && ctx.Query("fechaCreacion") != product.FechaCreacion

			if idVerif || nombreVerif || colorVerif || precioVerif || stockVerif || codigoVerif || publicadoVerif || fechaCreacionVerif {
				continue
			}

			filtrados = append(filtrados, product)

		}

		if len(filtrados) == 0 {
			ctx.JSON(500, gin.H{"error": "No se hallaron resultados"})
			return
		}

		ctx.JSON(200, filtrados)
	}
}

// Clase 1 Ejercicio 2 Parte 2
func (p *Product) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idInt, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
		}

		producto, err := p.service.GetById(idInt)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, producto)
	}
}

// Clase 2 Ejercicio 1 Parte 1
func validateToken(ctx *gin.Context) bool {
	if token := ctx.GetHeader("token"); token != "123" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token invalido"})
		return false
	}
	return true
}

// Clase 2 Ejercicio 1 Parte 1
func (p *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req Request

		if !validateToken(ctx) {
			return
		}

		if err := ctx.ShouldBindJSON(&req); err != nil {
			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				result := ""
				for i, field := range ve {
					if i != len(ve)-1 {
						result += fmt.Sprintf("El campo %s es requerido y ", field.Field())
					} else {
						result += fmt.Sprintf("El campo %s es requerido", field.Field())
					}
				}
				ctx.JSON(404, result)
				return
			}
		}

		producto, err := p.service.Store(
			req.Nombre,
			req.Color,
			req.Precio,
			req.Stock,
			req.Codigo,
			req.Publicado,
			req.FechaCreacion,
		)

		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, producto)
	}
}
