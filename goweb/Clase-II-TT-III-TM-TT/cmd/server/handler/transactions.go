package handler

import (
	"fmt"
	"goweb/internal/transactions"
	"goweb/pkg/web"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type request struct {
	TranCode    string  `json:"tranCode" binding:"required"`
	Currency    string  `json:"currency" binding:"required"`
	Amount      float64 `json:"amount" binding:"required"`
	Transmitter string  `json:"transmitter" binding:"required"`
	Receiver    string  `json:"receiver" binding:"required"`
	TranDate    string  `json:"tranDate" binding:"required"`
}

type requestPath struct {
	TranCode string  `json:"tranCode" binding:"required"`
	Amount   float64 `json:"amount" binding:"required"`
}

type Transaction struct {
	service transactions.Service
}

func NewTransaction(s transactions.Service) *Transaction {
	return &Transaction{
		service: s,
	}
}

//List of Transactions
//@Summary Obtain list of transactions.
//@Tags Transactions
//@description Get all transactions.
//@Accept json
//@Produce json
//@Param token header string true "token"
//@Success 200 {object} web.Response
//@Router /transactions [get]
func (t *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := os.Getenv("TOKEN")
		if token != "12345" {
			ctx.JSON(401, web.NewResponse(401, nil, "error: Token inválido"))
			return
		}
		t, err := t.service.GetAll()
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}
		if len(t) == 0 {
			ctx.JSON(404, web.NewResponse(404, nil, "error: No hay transacciones registradas"))
			return
		}
		ctx.JSON(200, web.NewResponse(200, t, ""))
	}
}

//Store of Transactions
//@Summary Store transaction in the list of them
//@Tags Transactions
//@description Store transactions indicating its parameters.
//@Accept json
//@Produce json
//@Param token header string true "token"
//@Param transaction body request true "Product to store"
//@Success 200 {object} web.Response
//@Failed 400 {object} web.Response
//@Router /transactions [post]
func (t *Transaction) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "12345" {
			ctx.JSON(401, web.NewResponse(401, nil, "error: Token inválido"))
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		t, err := t.service.Store(req.TranCode, req.Currency, req.Amount, req.Transmitter, req.Receiver, req.TranDate)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, t, ""))
	}
}

//Update Transactions
//@Summary Update of transaction by id
//@Tags Transactions
//@Description Update transaction modifying the parameters.
//@Accept json
//@Produce json
//@Param token header string true "token"
//@Param transaction body request true "Product to store"
//@Param id path string true "id"
//@Success 200 {object} web.Response
//@Failed 400 {object} web.Response
//@Router /transactions/{id} [put]
func (t *Transaction) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "12345" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(401, nil, "error: token inválido"))
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "error: invalid ID"))
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, err.Error()))
			return
		}
		if req.TranCode == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "error: El código de la transacción es requerido"))
			return
		}
		if req.Receiver == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "error: El receptor de la transacción es requerido"))
			return
		}
		if req.Currency == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "error: La moneda de la transacción es requerido"))
			return
		}
		if req.TranDate == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "error: La fecha de la transacción es requerido"))
			return
		}
		if req.Amount == 0 {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "error: El monto de la transacción es requerido"))
			return
		}
		if req.Transmitter == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "error: El emisor de la transacción es requerido"))
			return
		}

		tranc, err := t.service.Update(int(id), req.TranCode, req.Currency, req.Amount, req.Transmitter, req.Receiver, req.TranDate)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, err.Error()))
		}

		ctx.JSON(http.StatusAccepted, web.NewResponse(200, tranc, ""))

	}
}

//Delete Transactions
//@Summary Delete one transaction by id.
//@Tags Transactions
//@Description Delete transaction typing the id of transaction to eliminate.
//@Param token header string true "token"
//@Param id path string true "transanction id"
//@Success 200 {object} web.Response
//@Failed 400 {object} web.Response
//@Router /transactions/{id} [delete]
func (t *Transaction) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "12345" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(401, nil, "error: token inválido"))
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "error: invalid ID"))
			return
		}

		err = t.service.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(404, nil, err.Error()))
			return
		}

		msg := fmt.Sprintf("La transacción %d ha sido eliminada", id)
		ctx.JSON(http.StatusAccepted, web.NewResponse(200, msg, ""))

	}
}

//Update Transactions Patch
//@Summary Update transaction with Patch by id
//@Tags Transactions
//@Description Update transaction modifying only code and amount parameters.
//@Accept json
//@Produce json
//@Param token header string true "token"
//@Param transaction body request true "Transaction to uptdate Code and Amount"
//@Param id path string true "id"
//@Success 200 {object} web.Response
//@Failed 400 {object} web.Response
//@Router /transactions/{id} [patch]
func (t *Transaction) UpdateCodeAmount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "12345" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(401, nil, "token inválido"))
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "id inválido"))
			return
		}
		var req requestPath
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, err.Error()))
			return
		}
		if req.TranCode == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "error: El código de la transacción es requerido"))
			return
		}
		if req.Amount == 0 {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "error: La moneda de la transacción es requerido"))
			return
		}

		tranc, err := t.service.UpdateCodeAmount(int(id), req.TranCode, req.Amount)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(404, nil, err.Error()))
		}

		ctx.JSON(http.StatusAccepted, web.NewResponse(200, tranc, ""))

	}
}