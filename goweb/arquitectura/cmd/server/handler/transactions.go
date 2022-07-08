package handler

import (
	"arquitectura/internal/transactions"
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type simplerequest struct {
	TranCode string  `json:"tranCode"`
	Amount   float64 `json:"amount"`
}

func (r simplerequest) validate() bool {
	return r.TranCode != "" || r.Amount > 0
}

type request struct {
	TranCode    string  `json:"tranCode" binding:"required"`
	Currency    string  `json:"currency" binding:"required"`
	Amount      float64 `json:"amount" binding:"required"`
	Transmitter string  `json:"transmitter" binding:"required"`
	Reciever    string  `json:"reciever" binding:"required"`
	TranDate    string  `json:"tranDate" binding:"required"`
}

func (r request) getMissingField() error {
	if r.TranCode == "" {
		return errors.New("el campo tranCode es requerido")
	}
	if r.Currency == "" {
		return errors.New("el campo currency es requerido")
	}
	if r.Amount == 0 {
		return errors.New("el campo amount es requerido")
	}
	if r.Transmitter == "" {
		return errors.New("el campo transmitter es requerido")
	}
	if r.Reciever == "" {
		return errors.New("el campo receiver es requerido")
	}
	if r.TranDate == "" {
		return errors.New("el campo tranDate es requerido")
	}

	return nil
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
		t, err := t.service.Store(req.TranCode, req.Currency, req.Amount, req.Transmitter, req.Reciever, req.TranDate)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, t)
	}
}

func (t *Transaction) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "12345" {
			c.JSON(401, gin.H{"error": "Token inválido"})
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(404, gin.H{"error": "Id inválido"})
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, gin.H{"error": req.getMissingField().Error()})
			return
		}

		t, err := t.service.Update(int(id), req.TranCode, req.Currency, req.Amount, req.Transmitter, req.Reciever, req.TranDate)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, t)
	}
}

func (t *Transaction) UpdateFields() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "12345" {
			c.JSON(401, gin.H{"error": "Token inválido"})
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(404, gin.H{"error": "Id inválido"})
			return
		}

		var req simplerequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		if !req.validate() {
			c.JSON(404, gin.H{"error": "los campos tranCode y amount no son validos"})
			return
		}

		t, err := t.service.UpdateFields(int(id), req.TranCode, req.Amount)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, t)
	}
}

func (t *Transaction) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "12345" {
			c.JSON(401, gin.H{"error": "Token inválido"})
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Id inválido"})
			return
		}

		err = t.service.Delete(int(id))
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": fmt.Sprintf("la transaccion con id %d ha sido eliminada", id)})
	}
}
