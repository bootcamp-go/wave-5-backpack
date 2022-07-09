package handler

import (
	"fmt"
	"strconv"

	"goweb/clase3-go-web-tm/internal/transactions"

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

type requestPatch struct {
	Id                int     `json:"id" binding:"-"`
	CodigoTransaccion string  `json:"codigo de transaccion" binding:"required"`
	Moneda            string  `json:"moneda"`
	Monto             float64 `json:"monto" binding:"required"`
	Emisor            string  `json:"emisor"`
	Receptor          string  `json:"receptor"`
	Fecha             string  `json:"fecha de transaccion"`
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
			ctx.JSON(401, gin.H{"error": "token inválido 🔐"})
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

func (t *Transaction) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token invalido 🔐"})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		//	MESSAGES error for each Field : transacciones
		if req.CodigoTransaccion == "" {
			ctx.JSON(404, gin.H{"error": "El campo *Codigo de Transaccion* es requerido"})
			return
		}
		if req.Moneda == "" {
			ctx.JSON(404, gin.H{"error": "El campo *Moneda*  es requerido"})
			return
		}
		if req.Monto <= 0 {
			ctx.JSON(404, gin.H{"error": "El campo *Monto* es requerido"})
			return
		}
		if req.Emisor == "" {
			ctx.JSON(404, gin.H{"error": "El campo *Emisor* es requerido"})
			return
		}
		if req.Receptor == "" {
			ctx.JSON(404, gin.H{"error": "El campo *Receptor* es requerido"})
			return
		}
		if req.Fecha == "" {
			ctx.JSON(404, gin.H{"error": "El campo *Fecha* es requerido"})
			return
		}

		t, err := t.service.Update(int(id), req.CodigoTransaccion, req.Moneda,
			req.Monto, req.Emisor, req.Receptor, req.Fecha)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, t)
	}
}

func (t *Transaction) UpdateOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token invalido 🔐"})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID 🫠"})
			return
		}

		var req requestPatch
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.CodigoTransaccion == "" {
			ctx.JSON(400, gin.H{"error": "El codigo de la *transaccion* es requerido"})
			return
		}
		if req.Monto <= 0 {
			ctx.JSON(400, gin.H{"error": "El monto de la *transaccion* es invalido"})
			return
		}

		p, err := t.service.UpdateOne(int(id), req.CodigoTransaccion, req.Monto)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, p)
	}
}

func (t *Transaction) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token invalido 🔐"})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID 🫠"})
			return
		}

		err = t.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"data": fmt.Sprintf("El producto con id <%d> ha sido eliminado ✅", id)})
	}
}

func PaginaPrincipal(ctx *gin.Context) {
	ctx.String(200, "¡Bienvenido! Pagina Principal 🏠 ")
}
