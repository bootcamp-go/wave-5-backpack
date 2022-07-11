package handler

import (
	"errors"
	"fmt"
	"goweb/internal/domain"
	"goweb/internal/products"
	"goweb/pkg/web"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// --------------------------------------------
// --------------- Estructuras ----------------
// --------------------------------------------

type RequestRequired struct {
	Nombre        string  `json:"nombre" binding:"required"`
	Color         string  `json:"color" binding:"required"`
	Precio        float64 `json:"precio" binding:"required"`
	Stock         int     `json:"stock" binding:"required"`
	Codigo        string  `json:"codigo" binding:"required"`
	Publicado     bool    `json:"publicado" binding:"required"`
	FechaCreacion string  `json:"fechaCreacion" binding:"required"`
}

type Request struct {
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fechaCreacion"`
}

type Product struct {
	service products.Service
}

func NewProduct(s products.Service) *Product {
	return &Product{service: s}
}

// --------------------------------------------
// ------------------- CRUD -------------------
// --------------------------------------------

// Clase 1 Ejercicio 1 Parte 1
func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}

		products, err := p.service.GetAll()
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				web.NewResponse(http.StatusInternalServerError, nil, err.Error()),
			)
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
			ctx.JSON(
				http.StatusNotFound,
				web.NewResponse(http.StatusNotFound, nil, "No se hallaron resultados"),
			)
			return
		}

		ctx.JSON(
			http.StatusOK,
			web.NewResponse(http.StatusOK, filtrados, ""),
		)
	}
}

// Clase 1 Ejercicio 2 Parte 2
func (p *Product) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, "Id inválido"),
			)
			return
		}

		producto, err := p.service.GetById(id)
		if err != nil {
			ctx.JSON(
				http.StatusNotFound,
				web.NewResponse(http.StatusNotFound, nil, err.Error()),
			)
			return
		}

		ctx.JSON(
			http.StatusOK,
			web.NewResponse(http.StatusOK, producto, ""),
		)
	}
}

// Clase 2 Ejercicio 1 Parte 1
func (p *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}

		var req RequestRequired
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
				ctx.JSON(
					http.StatusBadRequest,
					web.NewResponse(http.StatusBadRequest, nil, result),
				)
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
			ctx.JSON(
				http.StatusNotFound,
				web.NewResponse(http.StatusNotFound, nil, err.Error()),
			)
			return
		}

		ctx.JSON(
			http.StatusOK,
			web.NewResponse(http.StatusOK, producto, ""),
		)
	}
}

// Clase 3 Ejercicio 1 Parte 1
func (p *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, "Id inválido"),
			)
			return
		}

		var req RequestRequired
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, err.Error()),
			)
			return
		}

		if req.Nombre == "" {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, "El nombre es requerido"),
			)
			return
		}

		if req.Color == "" {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, "El color es requerido"),
			)
			return
		}

		if req.Precio == 0 {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, "El precio es requerido"),
			)
			return
		}

		if req.Stock == 0 {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, "El stock es requerido"),
			)
			return
		}

		if req.Codigo == "" {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, "El codigo es requerido"),
			)
			return
		}

		if !req.Publicado {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, "El publicado es requerido"),
			)
			return
		}

		if req.FechaCreacion == "" {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, "La fecha de creacion es requerida"),
			)
			return
		}

		producto, err := p.service.Update(
			id,
			req.Nombre,
			req.Color,
			req.Precio,
			req.Stock,
			req.Codigo,
			req.Publicado,
			req.FechaCreacion,
		)

		if err != nil {
			ctx.JSON(
				http.StatusNotFound,
				web.NewResponse(http.StatusNotFound, nil, err.Error()),
			)
			return
		}

		ctx.JSON(
			http.StatusOK,
			web.NewResponse(http.StatusOK, producto, ""),
		)
	}
}

// Clase 3 Ejercicio 1 Parte 1
func (p *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, "Id inválido"),
			)
			return
		}

		err = p.service.Delete(id)
		if err != nil {
			ctx.JSON(
				http.StatusNotFound,
				web.NewResponse(http.StatusNotFound, nil, err.Error()),
			)
			return
		}

		ctx.JSON(
			http.StatusOK,
			web.NewResponse(http.StatusOK, fmt.Sprintf("El producto con id %d fue eliminado correctamente", id), ""),
		)
	}
}

// Clase 3 Ejercicio 1 Parte 1
func (p *Product) UpdateNombreYPrecio() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, "Id inválido"),
			)
			return
		}

		var req Request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, err.Error()),
			)
			return
		}

		if req.Nombre == "" {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, "El nombre es requerido"),
			)
			return
		}

		if req.Precio == 0 {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, "El precio es requerido"),
			)
			return
		}

		_, err = p.service.UpdateNombre(
			id,
			req.Nombre,
		)

		if err != nil {
			ctx.JSON(
				http.StatusNotFound,
				web.NewResponse(http.StatusNotFound, nil, err.Error()),
			)
			return
		}

		producto, err := p.service.UpdatePrecio(
			id,
			req.Precio,
		)

		if err != nil {
			ctx.JSON(
				http.StatusNotFound,
				web.NewResponse(http.StatusNotFound, nil, err.Error()),
			)
			return
		}

		ctx.JSON(
			http.StatusOK,
			web.NewResponse(http.StatusOK, producto, ""),
		)
	}
}

// --------------------------------------------
// ------------- Otras funciones --------------
// --------------------------------------------

// Clase 2 Ejercicio 1 Parte 1
func validateToken(ctx *gin.Context) bool {
	if token := ctx.GetHeader("token"); token != os.Getenv("TOKEN") {
		ctx.JSON(
			http.StatusUnauthorized,
			web.NewResponse(http.StatusUnauthorized, nil, "Token inválido"),
		)
		return false
	}
	return true
}
