package handler

import (
	"goweb/clase2-go-web-tt/internal/transactions"
	"strconv"

	"github.com/gin-gonic/gin"
)

type request struct {
	Id                int     `json:"id" binding:"-"`
	CodigoTransaccion string  `json:"codigo de transaccion" binding:"required"`
	Moneda            string  `json:"moneda" binding:"required"`
	Monto             float64 `json:"monto" binding:"required"`
	Emisor            string  `json:"emisor" binding:"required"`
	Receptor          string  `json:"receptor" binding:"required"`
	Fecha             string  `json:"fecha de transaccion" binding:"required"`
}

type Transaction struct {
	service transactions.Service
}

func NewTransaction(t transactions.Service) *Transaction {
	return &Transaction{
		service: t,
	}
}

func (t *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		p, err := t.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, p)
	}
}

func (t *Transaction) Ecommerce() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		t, err := t.service.Ecommerce(req.CodigoTransaccion, req.Moneda, req.Monto,
			req.Emisor, req.Receptor, req.Fecha)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, t)
	}
}

func (t *Transaction) GetOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		idParam := ctx.Param("id")
		id, errStr := strconv.Atoi(idParam)
		if errStr != nil {
			ctx.String(404, errStr.Error())
		}

		p, err := t.service.GetOne(id)
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, p)
	}
}

func PaginaPrincipal(ctx *gin.Context) {
	ctx.String(200, "¡Bienvenido! Pagina Principal 🏠 ")
}
