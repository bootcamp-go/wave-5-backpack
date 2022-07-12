package handler

import (
	"errors"
	"fmt"
	"goweb/internal/transactions"
	"goweb/pkg/web"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type request struct {
	Code      string  `json:"code" binding:"required"`
	Currency  string  `json:"currency" binding:"required"`
	Amount    float64 `json:"amount" binding:"required"`
	Issuer    string  `json:"issuer" binding:"required"`
	Recipient string  `json:"recipient" binding:"required"`
	Date      string  `json:"date" binding:"required"`
}

type request2 struct {
	Code   string  `json:"code" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
}

type Transaction struct {
	service transactions.Service
}

func NewTransaction(t transactions.Service) *Transaction {
	return &Transaction{
		service: t,
	}
}

func validateFields(err *error) string {
	errAns := ""
	var ve validator.ValidationErrors
	if errors.As(*err, &ve) {
		for _, fe := range ve {
			errAns += fmt.Sprintf("el campo %s es requerido ", fe.Field())
		}
	}
	return errAns
}

// ListTransactions godoc
// @Summary List transactions
// @Tags Transactions
// @Description get transactions
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Sucess 200 {object} web.Response
// @Router /transactions [get]
func (t *Transaction) GetAll(ctx *gin.Context) {
	transactions, _ := t.service.GetAll()
	ctx.JSON(200, web.NewResponse(200, transactions, ""))
}

// CreateTransactions godoc
// @Summary Create transaction
// @Tags Transactions
// @Description create transactions
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param transaction body request true "Transaction to create"
// @Sucess 200 {object} web.Response
// @Router /transaction [post]
func (t *Transaction) Create(ctx *gin.Context) {
	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, web.NewResponse(400, nil, validateFields(&err)))
		return
	}
	transaction, err := t.service.Create(req.Code, req.Currency, req.Amount, req.Issuer, req.Recipient, req.Date)
	if err != nil {
		ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		return
	}

	ctx.JSON(200, web.NewResponse(200, transaction, ""))
}

// GetOneTransaction godoc
// @Summary Get transaction
// @Tags Transactions
// @Description get transaction
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path int true "transaction id"
// @Sucess 200 {object} web.Response
// @Router /transaction/{id} [get]
func (t *Transaction) GetOne(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, web.NewResponse(400, nil, "id invalido"))
		return
	}
	transaction, err := t.service.GetOne(id)
	if err != nil {
		ctx.JSON(404, web.NewResponse(404, nil, "Transaccion no existente"))
		return
	}
	ctx.JSON(200, web.NewResponse(200, transaction, ""))
}

// UpdateTransaction godoc
// @Summary Update transaction
// @Tags Transactions
// @Description Update transaction
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param transaction body request true "Transaction to update"
// @Param id path int true "transaction id"
// @Sucess 200 {object} web.Response
// @Router /transaction/{id} [put]
func (t *Transaction) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, web.NewResponse(400, nil, "id no vallido"))
		return
	}
	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, web.NewResponse(400, nil, validateFields(&err)))
		return
	}
	transaction, err := t.service.Update(id, req.Code, req.Currency, req.Amount, req.Issuer, req.Recipient, req.Date)
	if err != nil {
		ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		return
	}
	ctx.JSON(200, web.NewResponse(200, transaction, ""))
}

// DeleteTransaction godoc
// @Summary Delete transaction
// @Tags Transactions
// @Description Delete transaction
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path int true "transaction id"
// @Sucess 200 {object} web.Response
// @Router /transaction/{id} [delete]
func (t *Transaction) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, web.NewResponse(400, nil, "id no valido"))
		return
	}

	if err := t.service.Delete(id); err != nil {
		ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		return
	}
	ctx.JSON(200, web.NewResponse(200, "Eliminado exitosamente", ""))
}

// PartialUpdateTransaction godoc
// @Summary Update transaction
// @Tags Transactions
// @Description Update code and ammount of transaction
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param transaction body request2 true "Transaction to update"
// @Param id path int true "transaction id"
// @Sucess 200 {object} web.Response
// @Router /transaction/{id} [patch]
func (t *Transaction) Update2(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, web.NewResponse(400, nil, "id no existente"))
		return
	}
	var req request2
	ctx.ShouldBindJSON(&req)
	if req.Code == "" {
		ctx.JSON(400, web.NewResponse(400, nil, "code requerido"))
		return
	}
	if req.Amount == 0 {
		ctx.JSON(400, web.NewResponse(400, nil, "amount requerido"))
		return
	}
	transaction, err := t.service.Update2(id, req.Code, req.Amount)
	if err != nil {
		ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		return
	}
	ctx.JSON(200, web.NewResponse(200, transaction, ""))
}
