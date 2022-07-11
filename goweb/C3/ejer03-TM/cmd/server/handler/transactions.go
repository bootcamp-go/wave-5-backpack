package handler

import (
	"ejer02-TT/internal/transactions"
	"ejer02-TT/pkg/web"
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type request struct {
	TranCode    string  `json:"tranCode" binding:"required"`
	Currency    string  `json:"currency" binding:"required"`
	Amount      float64 `json:"amount" binding:"required"`
	Transmitter string  `json:"transmitter" binding:"required"`
	Reciever    string  `json:"reciever" binding:"required"`
	TranDate    string  `json:"tranDate" binding:"required"`
}

type Transaction struct {
	service transactions.Service
}

func NewTransaction(s transactions.Service) *Transaction {
	return &Transaction{
		service: s,
	}
}

//ListTransactions godoc
//@Sumary List transactions
//@Tags Transactions
//@Description get transactions
//@Accept json
//@Transactions json
//@Param token header string true "token"
//@Success 200 {object} web.Response
//@Router /transactions [get]
func (t *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		t, err := t.service.GetAll()
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}
		if len(t) == 0 {
			ctx.JSON(404, web.NewResponse(404, nil, "No hay transacciones"))
		}
		ctx.JSON(200, web.NewResponse(200, t, ""))
	}
}

//StoreTransactions godoc
//@Sumary Store transactions
//@Tags Transactions
//@Description store transactions
//@Accept json
//@Transactions json
//@Param token header string true "token"
//@Param transaction body request true "Transaction to store"
//@Success 200 {object} web.Response
//@Router /transactions [post]
func (t *Transaction) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req request
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
				ctx.JSON(400, web.NewResponse(400, nil, result))
				return
			}

		}
		t, err := t.service.Store(req.TranCode, req.Currency, req.Amount, req.Transmitter, req.Reciever, req.TranDate)
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, t, ""))
	}

}

//DeleteTransactions godoc
//@Sumary Delete transactions
//@Tags Transactions
//@Description delete transactions
//@Accept json
//@Transactions json
//@Param token header string true "token"
//@Param transaction body request true "Transaction deleted"
//@Success 200 {object} web.Response
//@Router /transactions/{id} [delete]
func (t *Transaction) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Id inválido"))
			return
		}

		err = t.service.Delete(int(id))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Id inexistente"))
			return
		}

		ctx.JSON(200, web.NewResponse(200, nil, fmt.Sprintf("La transaccion %d ha sido eliminada", id)))
	}

}

//UpdateTransactions godoc
//@Sumary Update transactions
//@Tags Transactions
//@Description update transactions
//@Accept json
//@Transactions json
//@Param token header string true "token"
//@Param transaction body request true "Transaction updated"
//@Success 200 {object} web.Response
//@Router /transactions/{id} [PUT]
func (t *Transaction) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req request
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
				ctx.JSON(400, web.NewResponse(400, nil, result))
				return
			}

		}

		t, err := t.service.Store(req.TranCode, req.Currency, req.Amount, req.Transmitter, req.Reciever, req.TranDate)
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, t, ""))
	}
}

//UpdateCodeAnAmountTransactions godoc
//@Sumary Update transactions
//@Tags Transactions
//@Description update code and amount transactions
//@Accept json
//@Transactions json
//@Param token header string true "token"
//@Param transaction body request true "Transaction updated"
//@Success 200 {object} web.Response
//@Router /transactions/{id} [PATCH]
func (t *Transaction) UpdateCodeAndAmount() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Id inválido"))
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}

		if req.TranCode == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El codigo de transaccion es requerido"))
			return
		}
		if req.Amount < 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "El codigo de transaccion es requerido y debe ser positivo"))
			return
		}

		t, err := t.service.UpdateCodeAndAmount(int(id), req.TranCode, req.Amount)
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, t, ""))
	}
}
