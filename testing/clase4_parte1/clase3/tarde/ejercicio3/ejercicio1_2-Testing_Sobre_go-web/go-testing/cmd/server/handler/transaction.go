package handler

import (
	"net/http"
	"os"
	"strconv"

	"github.com/bootcamp-go/internal/transactions"
	"github.com/bootcamp-go/pkg/web"
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
// @Summary List Transactions
// @Tags Transactions
// @Descripcion get transactions
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /transactions [get]
func (t *Transaction) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "token inválido")) //401
			return
		}

		transactions, err := t.service.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, web.NewResponse(500, nil, err.Error())) // 500
			return
		}

		if len(transactions) <= 0 {
			c.JSON(http.StatusOK, web.NewResponse(200, "No se encontraron transacciones", "")) // 500
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(200, transactions, ""))
	}
}

// Store Transactions
// @Summary Store Transactions
// @Tags Transactions
// @Descripcion store transactions
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param transaction body request true "Transaction to store"
// @Success 200 {object} web.Response
// @Router /transactions [post]
func (t *Transaction) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "token inválido")) //401
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, err.Error())) //400
			return
		}

		transaction, err := t.service.Store(req.Codigo, req.Moneda, req.Emisor, req.Receptor, req.Monto)
		if err != nil {
			c.JSON(http.StatusInternalServerError, web.NewResponse(500, nil, err.Error())) //500
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(200, transaction, ""))
	}
}

// Update Transactions
// @Summary Update Transactions
// @Tags Transactions
// @Descripcion update transactions
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path int true "Transaction id"
// @Param transaction body request true "Transaction to store"
// @Success 200 {object} web.Response
// @Router /transactions/{id} [put]
func (t *Transaction) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "token inválido")) //401
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "")) //400
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "")) //400
			return
		}

		if req.Codigo == "" {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "El campo codigo es requerido")) //400
			return
		}

		if req.Emisor == "" {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "El campo emisor es requerido")) //400
			return
		}

		if req.Receptor == "" {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "El campo receptor es requerido")) //400
			return
		}

		if req.Moneda == "" {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "El campo moneda es requerido")) //400
			return
		}

		if req.Monto <= 0 {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "El campo monto es requerido")) //400
			return
		}

		transaction, err := t.service.Update(id, req.Codigo, req.Moneda, req.Emisor, req.Receptor, req.Monto)
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(404, nil, err.Error())) //404
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(200, transaction, ""))
	}
}

// Update Receptor and Monto Transactions
// @Summary Update Receptor and Monto Transactions
// @Tags Transactions
// @Descripcion update receptor and monto transactions
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path int true "Transaction id"
// @Param transaction body request true "Transaction to update receptor and monto"
// @Success 200 {object} web.Response
// @Router /transactions/{id} [patch]
func (t *Transaction) UpdateReceptorYMonto() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "token inválido")) //401
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "")) //400
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "")) //400
			return
		}

		if req.Receptor == "" {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "")) //400
			return
		}

		if req.Monto <= 0 {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "")) //400
			return
		}

		transaction, err := t.service.UpdateReceptorYMonto(id, req.Receptor, req.Monto)
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(404, nil, "")) //404
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(200, transaction, ""))
	}
}

// Delete Transactions
// @Summary Delete Transactions
// @Tags Transactions
// @Descripcion delete transactions
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path int true "Transaction id"
// @Success 200 {object} web.Response
// @Router /transactions/{id} [delete]
func (t *Transaction) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "token inválido")) //401
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, err.Error())) //400
			return
		}

		err = t.service.Delete(id)
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(404, nil, err.Error())) //404
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(200, nil, ""))
	}
}
