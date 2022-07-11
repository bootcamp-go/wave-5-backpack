package handler

import (
	"GoWeb/internals/transactions"
	"GoWeb/pkg/web"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type transanction struct {
	Code     string  `json:"code"`
	Coin     string  `json:"coin"`
	Amount   float64 `json:"amount"`
	Emisor   string  `json:"emisor"`
	Receptor string  `json:"receptor"`
	Date     string  `json:"date"`
}

type Transaction struct {
	service transactions.Service
}

func NewTransaction(t transactions.Service) *Transaction {
	return &Transaction{service: t}
}

//ListTransactions godoc
//@Summary list of transactions
//@Tags Transactions
//@Description get all transactions
//@Accept json
//@Produce json
//@Param token header string true "token"
//@Success 200 {object} web.Response
//@Router /transacciones [get]
func (tt *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tran, err := tt.service.GetAll()
		if err != nil {
			ctx.JSON(204, web.NewResponse(401, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, tran)
	}
}

//StoreTransactions godoc
//@Summary Store of transactions
//@Tags Store Transactions
//@Description store transactions
//@Accept json
//@Produce json
//@Param token header string true "token"
//@Param transactions body transanction true "Transaction to store"
//@Success 200 {object} web.Response
//@Router /transacciones [post]
func (tt *Transaction) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var tr transanction
		if err := ctx.Bind(&tr); err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, err.Error()))
			return
		}
		if tr.Code == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "codigo no valido"))
			return
		}
		if tr.Coin == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "tipo de moneda invalido"))
			return
		}
		if tr.Amount == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "cantidad no valida"))
			return
		}
		if tr.Emisor == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "emisor no valido"))
			return
		}
		if tr.Receptor == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "receptor no valido"))
			return
		}
		if tr.Date == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "fecha no valida"))
			return
		}
		tran, err := tt.service.Store(tr.Code, tr.Coin, tr.Amount, tr.Emisor, tr.Receptor, tr.Date)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(200, tran, ""))
	}
}

//UpdateTransactions godoc
//@Summary Update of transactions
//@Tags Update Transactions
//@Description Update transactions
//@Accept json
//@Produce json
//@Param token header string true "token"
//@Param transactions body transanction true "Transaction to update"
//@Param id path string true "transanction id"
//@Success 200 {object} web.Response
//@Router /transacciones/{id} [put]
func (tt *Transaction) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "id invalido"))
			return
		}
		var tr transanction
		if err := ctx.Bind(&tr); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if tr.Code == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "codigo no valido"))
			return
		}
		if tr.Coin == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "tipo de moneda invalido"))
			return
		}
		if tr.Amount == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "cantidad no valida"))
			return
		}
		if tr.Emisor == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "emisor no valido"))
			return
		}
		if tr.Receptor == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "receptor no valido"))
			return
		}
		if tr.Date == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "fecha no valida"))
			return
		}

		tran, err := tt.service.Update(id, tr.Code, tr.Coin, tr.Amount, tr.Emisor, tr.Receptor, tr.Date)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		}

		ctx.JSON(http.StatusOK, web.NewResponse(200, tran, ""))
	}
}

//DeleteTransactions godoc
//@Summary Delete of transactions
//@Tags Delete Transactions
//@Description delete transactions
//@Accept json
//@Produce json
//@Param token header string true "token"
//@Param transactions body transanction true "Transaction to delete"
//@Param id path string true "transanction id"
//@Success 200 {object} web.Response
//@Router /transacciones/{id} [delete]
func (tt *Transaction) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Id invalido"))
			return
		}
		err = tt.service.Delete(id)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": fmt.Sprintf("el registro %d ha sido eliminado", id),
		})
	}
}

//UpdateCodeTransactions godoc
//@Summary UpdateCode of transactions
//@Tags UpdateCode Transactions
//@Description UpdateCode transactions
//@Accept json
//@Produce json
//@Param token header string true "token"
//@Param transactions body transanction true "Transaction to updatecode"
//@Param id path string true "transanction id"
//@Success 200 {object} web.Response
//@Router /transacciones/{id} [patch]
func (tt *Transaction) UpdateCode() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		var tr transanction
		if err := ctx.ShouldBindJSON(&tr); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if tr.Code == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "el codigo es requerido"))
			return
		}
		if tr.Amount == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "la cantidad es requerida"))
			return
		}
		tran, err := tt.service.UpdateCode(id, tr.Code, tr.Amount)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(200, tran, ""))
	}
}
