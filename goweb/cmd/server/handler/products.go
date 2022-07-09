package handler

import (
	"fmt"
	"goweb/internal/domain"
	"goweb/internal/products"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type request struct {
	Id            int     `json:"-"`
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"código"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fecha_de_creación"`
}

type Product struct {
	service products.Service
}

func InitProduct(p products.Service) *Product {
	return &Product{
		service: p,
	}
}

func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}
		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}

func (c *Product) CreateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}
		var req domain.Products

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error: ": err.Error(),
			})
			return
		}

		if req.Nombre == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "El nombre del producto es requerido",
			})
			return
		}
		if req.Color == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "El color del producto es requerido",
			})
			return
		}
		if req.Precio == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "El precio del producto es requerido",
			})
			return
		}
		if req.Stock == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "El stock del producto es requerido",
			})
			return
		}
		if req.Codigo == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "El código del producto es requerido",
			})
			return
		}
		if req.FechaCreacion == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "La fecha de creación del producto es requerida",
			})
			return
		}

		p, err := c.service.CreateProduct(req.Id, req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.FechaCreacion)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"Error: ": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"responseMessage": p,
		})

	}
}

func (c *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if req.Nombre == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "El nombre del producto es requerido",
			})
			return
		}
		if req.Color == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "El color del producto es requerido",
			})
			return
		}
		if req.Precio == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "El precio del producto es requerido",
			})
			return
		}
		if req.Stock == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "El stock del producto es requerido",
			})
			return
		}
		if req.Codigo == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "El código del producto es requerido",
			})
			return
		}
		if req.FechaCreacion == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "La fecha de creación del producto es requerida",
			})
			return
		}
		p, err := c.service.Update(int(id), req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.FechaCreacion)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"responseMessage": p,
		})
	}

}

func (c *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Id inválido",
			})
			return
		}
		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"responseMessage": fmt.Sprintf("El producto %d ha sido eliminado", id),
		})
	}

}

func (c *Product) UpdateOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Id inválido",
			})
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if req.Nombre == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "El nombre del producto es requerido",
			})
			return
		}
		if req.Precio == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "El precio del producto es requerido",
			})
			return
		}
		p, err := c.service.UpdateOne(int(id), req.Nombre, req.Precio)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"responseMessage": p,
		})
	}

}
