package handler

import (
	"ejer02-TT/internal/transactions"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
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

func (t *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{
				"error": "token invalido",
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
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{
				"error": "token invalido",
			})
			return
		}
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
				ctx.JSON(400, result)
				log.Println("prueba 2")
				return
			}
			log.Println("prueba 1")
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

func (t *Transaction) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{
				"error": "token invalido",
			})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Id inválido"})
			return
		}

		err = t.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{"data": fmt.Sprintf("La transaccion %d ha sido eliminada", id)})
	}

}

func (t *Transaction) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "12345" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}

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
				ctx.JSON(400, result)
				log.Println("prueba 2")
				return
			}
			log.Println("prueba 1")
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

func (t *Transaction) UpdateCodeAndAmount() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "123456" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Id inválido"})
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if req.TranCode == "" {
			c.JSON(400, gin.H{"error": "El codigo de transaccion es requerido"})
			return
		}
		if req.Amount < 0 {
			c.JSON(400, gin.H{"error": "El monto de transaccion es requerido y positivo"})
			return
		}

		t, err := t.service.UpdateCodeAndAmount(int(id), req.TranCode, req.Amount)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, t)
	}
}
