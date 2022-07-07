package handler

import (
	"ejer02-TT/internal/transactions"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
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

func (t *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
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
	var req request
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "12345" {
			ctx.JSON(401, gin.H{"error": " inválido"})
			return
		}

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
				ctx.JSON(404, result)
				return
			}
			ctx.JSON(404, "error")
			return
		}
		t, err := t.service.Store(req.TranCode, req.Currency, req.Amount, req.Transmitter, req.Reciever, req.TranDate)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, t)
	}
}
