package handler

import (
	"GoWeb/internals/transactions"
	"fmt"
	"net/http"
	"os"
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

func (tt *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
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
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
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

func (tt *Transaction) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{
				"error": "token no valido",
			})
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "id invalido",
			})
			return
		}
		var tr transanction
		if err := ctx.Bind(&tr); err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		if tr.Code == "" {
			ctx.JSON(400, gin.H{"error": "codigo no valido"})
			return
		}
		if tr.Coin == "" {
			ctx.JSON(400, gin.H{"error": "tipo de moneda invalido"})
			return
		}
		if tr.Amount == 0 {
			ctx.JSON(400, gin.H{"error": "cantidad no valida"})
			return
		}
		if tr.Emisor == "" {
			ctx.JSON(400, gin.H{"error": "emisor no valido"})
			return
		}
		if tr.Receptor == "" {
			ctx.JSON(400, gin.H{"error": "receptor no valido"})
			return
		}
		if tr.Date == "" {
			ctx.JSON(400, gin.H{"error": "fecha no valida"})
			return
		}

		tran, err := tt.service.Update(id, tr.Code, tr.Coin, tr.Amount, tr.Emisor, tr.Receptor, tr.Date)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
		}

		ctx.JSON(http.StatusOK, tran)
	}
}

func (tt *Transaction) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{
				"error": "toke invalido",
			})
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "Id invalido",
			})
			return
		}
		err = tt.service.Delete(id)
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": fmt.Sprintf("el registro %d ha sido eliminado", id),
		})
	}
}

func (tt *Transaction) UpdateCode() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{
				"error": "token no valido",
			})
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		var tr transanction
		if err := ctx.ShouldBindJSON(&tr); err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		if tr.Code == "" {
			ctx.JSON(400, gin.H{"error": "el codigo es requerido"})
			return
		}
		if tr.Amount == 0 {
			ctx.JSON(400, gin.H{"error": "la cantidad es requerida"})
			return
		}
		tran, err := tt.service.UpdateCode(id, tr.Code, tr.Amount)
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, tran)
	}
}
