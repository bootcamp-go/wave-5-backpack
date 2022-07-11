package handler

import (
	"net/http"
	"os"
	"strconv"

	"github.com/bootcamp-go/wave-5-backpack/tree/olivera_sebastian/goweb/clase2_parte2+siguientes/internal/transactions"
	"github.com/bootcamp-go/wave-5-backpack/tree/olivera_sebastian/goweb/clase2_parte2+siguientes/pkg/web"

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

// ListTransactions godoc
// @Summary List transactions
// @Tags Transactions
// @Description get transactions
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.NewResponse
// @Router /transactions [get]
func (t *Transaction) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "Token inválido")) //401
			return
		}

		transactions, err := t.service.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, web.NewResponse(500, nil, err.Error())) // 500
			return
		}

		if len(transactions) <= 0 {
			c.JSON(http.StatusOK, web.NewResponse(500, nil, "No se encontraron transacciones")) // 500
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(200, transactions, ""))
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
		id, err := strconv.Atoi(c.Param("Id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "id inválido"})
			return
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.Codigo == "" {
			c.JSON(400, gin.H{"error": "El código de la transacción es requerido"})
			return
		}
		if req.Emisor == "" {
			c.JSON(400, gin.H{"error": "El emisor de la transación es requerido"})
			return
		}
		if req.Moneda == "" {
			c.JSON(400, gin.H{"error": "La moneda es requerida"})
			return
		}
		if req.Monto == 0 {
			c.JSON(400, gin.H{"error": "El monto es requerido"})
			return
		}
		if req.Receptor == "" {
			c.JSON(400, gin.H{"error": "El receptor es requerido"})
			return
		}
		t, err := t.service.Update(int64(id), req.Monto, req.Codigo, req.Emisor, req.Receptor, req.Moneda)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, t)
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
