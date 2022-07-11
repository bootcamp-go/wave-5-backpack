package handler

import (
	"net/http"
	"os"
	"strconv"

	"github.com/bootcamp-go/go-web/internal/transactions"
	"github.com/gin-gonic/gin"
)

type request struct {
	Codigo   string  `json:"codigo"`
	Monto    float64 `json:"monto"`
	Moneda   string  `json:"moneda"`
	Emisor   string  `json:"emisor"`
	Receptor string  `json:"receptor"`
}

type Transaction struct {
	service transactions.Service
}

func NewTransaction(s transactions.Service) *Transaction {
	return &Transaction{service: s}
}

func (t *Transaction) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"}) //401
			return
		}

		transactions, err := t.service.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "token inválido"}) // 500
			return
		}

		if len(transactions) <= 0 {
			c.JSON(http.StatusOK, gin.H{"message": "No se encontraron transacciones."}) // 500
			return
		}

		c.JSON(http.StatusOK, gin.H{"transactions": transactions})
	}
}

func (t *Transaction) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"}) //401
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) //400
			return
		}

		transaction, err := t.service.Store(req.Codigo, req.Moneda, req.Emisor, req.Receptor, req.Monto)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) //400
			return
		}

		c.JSON(http.StatusOK, gin.H{"transaction": transaction})
	}
}

func (t *Transaction) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"}) //401
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Id incorrecto"}) //400
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) //400
			return
		}

		if req.Codigo == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El campo codigo es requerido"}) //400
			return
		}

		if req.Emisor == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El campo emisor es requerido"}) //400
			return
		}

		if req.Receptor == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El campo receptor es requerido"}) //400
			return
		}

		if req.Moneda == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El campo moneda es requerido"}) //400
			return
		}

		if req.Monto <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El campo monto es requerido"}) //400
			return
		}

		transaction, err := t.service.Update(id, req.Codigo, req.Moneda, req.Emisor, req.Receptor, req.Monto)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()}) //404
			return
		}

		c.JSON(http.StatusOK, gin.H{"transaction": transaction})
	}
}

func (t *Transaction) UpdateReceptorYMonto() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"}) //401
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Id incorrecto"}) //400
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) //400
			return
		}

		if req.Receptor == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El campo receptor es requerido"}) //400
			return
		}

		if req.Monto <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El campo monto es requerido"}) //400
			return
		}

		transaction, err := t.service.UpdateReceptorYMonto(id, req.Receptor, req.Monto)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()}) //404
			return
		}

		c.JSON(http.StatusOK, gin.H{"transaction": transaction})
	}
}

func (t *Transaction) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"}) //401
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Id incorrecto"}) //400
			return
		}

		err = t.service.Delete(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()}) //404
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "La transacción se ha eliminado correctamente."})
	}
}
