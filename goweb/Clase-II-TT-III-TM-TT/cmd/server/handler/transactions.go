package handler

import (
	"arquitectura/internal/transactions"
	"fmt"
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

func (t *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := os.Getenv("TOKEN")
		if token != "12345" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		t, err := t.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, t)
	}
}

func (t *Transaction) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "12345" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		t, err := t.service.Store(req.TranCode, req.Currency, req.Amount, req.Transmitter, req.Receiver, req.TranDate)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, t)
	}
}

func (t *Transaction) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "12345" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.TranCode == "" {
			ctx.JSON(400, gin.H{"error": "El código de la transacción es requerido"})
			return
		}
		if req.Receiver == "" {
			ctx.JSON(400, gin.H{"error": "El receptor de la transacción es requerido"})
			return
		}
		if req.Currency == "" {
			ctx.JSON(400, gin.H{"error": "La moneda de la transacción es requerido"})
			return
		}
		if req.TranDate == "" {
			ctx.JSON(400, gin.H{"error": "La fecha de la transacción es requerido"})
			return
		}
		if req.Amount == 0 {
			ctx.JSON(400, gin.H{"error": "El monto de la transacción es requerido"})
			return
		}
		if req.Transmitter == "" {
			ctx.JSON(400, gin.H{"error": "El emisor de la transacción es requerido"})
			return
		}

		tranc, err := t.service.Update(int(id), req.TranCode, req.Currency, req.Amount, req.Transmitter, req.Receiver, req.TranDate)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
		}

		ctx.JSON(200, tranc)

	}
}

func (t *Transaction) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "12345" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		err = t.service.Delete(id)
		if err != nil {
			ctx.JSON(404, gin.H{"error2": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{
			"error": fmt.Sprintf("La transacción %d ha sido eliminada", id),
		})

	}
}

func (t *Transaction) UpdateCodeAmount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "12345" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}
		var req requestPath
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.TranCode == "" {
			ctx.JSON(400, gin.H{"error": "El código de la transacción es requerido"})
			return
		}
		if req.Amount == 0 {
			ctx.JSON(400, gin.H{"error": "La moneda de la transacción es requerido"})
			return
		}

		tranc, err := t.service.UpdateCodeAmount(int(id), req.TranCode, req.Amount)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
		}

		ctx.JSON(200, tranc)

	}
}
