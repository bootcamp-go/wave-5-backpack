package handler

import (
	"net/http"
	"strconv"

	"github.com/abelardolugo/go-web/internal/products"

	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre   string  `json:"nombre"`
	Cantidad int     `json:"cantidad"`
	Precio   float64 `json:"precio"`
}

type Product struct {
	service products.Service
}

func NewProduct(s products.Service) *Product {
	return &Product{service: s}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != "123456" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"}) // 401
			return
		}

		productos, err := p.service.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ha ocurrido un error inesperado"}) // 500
			return
		}

		if len(productos) <= 0 {
			c.JSON(http.StatusOK, gin.H{"message": "No se encontraron transacciones"}) // 200
			return
		}

		c.JSON(http.StatusOK, gin.H{"productos": productos})
	}
}

func (p *Product) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != "123456" {
			c.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		var req request
		if err := c.Bind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // 400
			return
		}

		producto, err := p.service.Store(req.Nombre, req.Cantidad, req.Precio)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // 400
			return
		}
		c.JSON(http.StatusOK, gin.H{"producto": producto}) // 200
	}
}

func (p *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != "123456" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"}) // 401
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // 400
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // 400
			return
		}

		if req.Nombre == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El campo nombre es requerido"}) // 400
			return
		}

		if req.Cantidad >= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El campo cantidad es reqeuerido"}) // 400
			return
		}

		if req.Precio >= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El campo precio es reqeuerido"}) // 400
			return
		}

		producto, err := p.service.Update(id, req.Nombre, req.Cantidad, req.Precio)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()}) // 404
			return
		}

		c.JSON(http.StatusOK, gin.H{"producto": producto})

	}
}

func (p *Product) UpdateName() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "123456" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Id incorrecto"}) // 400
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // 400
			return
		}

		if req.Nombre == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		producto, err := p.service.UpdateName(id, req.Nombre)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, producto)

	}
}

func (p *Product) UpdateNamePrice() gin.HandlerFunc {
	return func(c *gin.Context) {

		if token := c.GetHeader("token"); token != "123456" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error:": "Id Inválido"})
			return
		}

		var req request

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.Nombre == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El campo monto es requerido"}) // 400
			return
		}

		if req.Precio < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El campo monto es requerido"}) // 400
			return
		}

		producto, err := p.service.UpdateName(id, req.Nombre)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, producto)
	}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != "123456" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"}) // 401
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		err = p.service.Delete(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()}) // 404
			return
		}

		c.JSON(http.StatusNoContent, "") // 204
	}
}
