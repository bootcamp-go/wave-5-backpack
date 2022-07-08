package handler

import (
	"GoWeb/internals/transactions"
	"net/http"

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

func (tt *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "Token invalido",
			})
			return
		}
		tran, err := tt.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, tran)
	}
}

func (tt *Transaction) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token invalido",
			})
			return
		}
		var tr transanction
		if err := ctx.Bind(&tr); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		tran, err := tt.service.Store(tr.Code, tr.Coin, tr.Amount, tr.Emisor, tr.Receptor, tr.Date)
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, tran)
	}
}
