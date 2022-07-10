package handler

import (
	"fmt"
	"goweb/internals/transactions"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type request struct {
	Codigo   string `json:"codigo"`
	Moneda   string `json:"moneda"`
	Monto    int    `json:"monto"`
	Emisor   string `json:"emisor"`
	Receptor string `json:"receptor"`
}

type Transaction struct {
	service transactions.Service
}

func NewTransaction(t transactions.Service) *Transaction {
	return &Transaction{
		service: t,
	}
}

func (t *Transaction) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, gin.H{"error": "No tiene permisos para realizar la petición solicitada ;)"})
			return
		}
		trans, err := t.service.GetAll()
		if err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, trans)
	}
}

func (t *Transaction) GetByQuery() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, gin.H{"error": "No tiene permisos para realizar la petición solicitada ;)"})
			return
		}
		trans, err := t.service.GetAll()
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
		}
		monto, _ := strconv.Atoi(c.Query("monto"))
		for _, t := range trans {
			if t.Codigo == c.Query("codigo") {
				c.JSON(http.StatusAccepted, t)
			}
			if t.Moneda == c.Query("moneda") {
				c.JSON(http.StatusAccepted, t)
			}
			if t.Monto == monto {
				c.JSON(http.StatusAccepted, t)
			}
			if t.Emisor == c.Query("emisor") {
				c.JSON(http.StatusAccepted, t)
			}
			if t.Receptor == c.Query("receptor") {
				c.JSON(http.StatusAccepted, t)
			}
			if t.Fecha == c.Query("fecha") {
				c.JSON(http.StatusAccepted, t)
			}
		}
	}
}

func (t *Transaction) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		trans, err := t.service.GetAll()
		if err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		for _, trans := range trans {
			if trans.Id == id {
				c.JSON(http.StatusAccepted, trans)
				return
			}
		}
		c.JSON(404, gin.H{"error": "ID no existente"})
	}
}

//POST

func (t *Transaction) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, gin.H{"error": "No tiene permisos para realizar la petición solicitada ;)"})
			return
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		trans, err := t.service.Store(req.Codigo, req.Moneda, req.Monto, req.Emisor, req.Receptor)
		if err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, trans)
	}
}

//UPDATE

func (t *Transaction) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, gin.H{"error": "No tiene permisos para realizar la petición solicitada ;)"})
			return
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.Codigo == "" {
			c.JSON(400, gin.H{"error": "El código del producto es requerido"})
			return
		}
		if req.Moneda == "" {
			c.JSON(400, gin.H{"error": "El tipo de moneda del producto es requerido"})
			return
		}
		if req.Monto == 0 {
			c.JSON(400, gin.H{"error": "El monto del producto es requerido"})
			return
		}
		if req.Emisor == "" {
			c.JSON(400, gin.H{"error": "El emisor del producto es requerido"})
			return
		}
		if req.Receptor == "" {
			c.JSON(400, gin.H{"error": "El receptor del producto es requerido"})
			return
		}
		trans, err := t.service.Update(int(id), req.Codigo, req.Moneda, req.Monto, req.Emisor, req.Receptor)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, trans)
	}
}

// DELETE

func (t *Transaction) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, gin.H{"error": "No tiene permisos para realizar la petición solicitada ;)"})
			return
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(404, gin.H{"error": "invalid ID"})
			return
		}
		err = t.service.Delete(int(id))
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"data": fmt.Sprintf("El producto %d ha sido eliminado", id)})
	}
}
