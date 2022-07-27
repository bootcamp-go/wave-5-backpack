package handler

import (
	"WebServer/internal/transactions"
	"WebServer/pkg/web"
	"errors"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type request struct {
	CodigoTransaccion int     `json:"codigo_de_transaccion"`
	Moneda            string  `json:"moneda"`
	Monto             float64 `json:"monto"`
	Emisor            string  `json:"emisor"`
	Receptor          string  `json:"receptor"`
	FechaTransaccion  string  `json:"fecha_de_transaccion"`
}

type Transaction struct {
	service transactions.Service
}

func NewTransaction(p transactions.Service) *Transaction {
	return &Transaction{
		service: p,
	}
}

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get transactions
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Failure 204 {object} web.Response
// @Router /transactions [get]
func (c *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token invalido"))
			return
		}
		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, p)
	}
}

// CreateTransactions godoc
// @Summary Create Transactions
// @Tags Transactions
// @Description Create transactions
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param transaction body request true "transaction to create"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 409 {object} web.Response
// @Router /products [post]
func (c *Transaction) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, web.NewResponse(401, nil, "token invalido"))
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		t, err := c.service.Create(req.CodigoTransaccion, req.Moneda, req.Monto, req.Emisor, req.Receptor, req.FechaTransaccion)
		//codigoTransaccion int, moneda string, monto float64, emisor, receptor, fechaTransaccion string
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
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64) //convert to int64 string id param , if it's not a number return an error
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Id inválido"})
			return
		}

		var req request // struct for entity without id(entire entity just will be visible for repositoy and service interfaces)
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if missingFields, err := req.validateReq(); err != nil {
			ctx.JSON(400, gin.H{"missing Fields": missingFields})
		}

		t, err := t.service.Update(int(id), req.CodigoTransaccion, req.Moneda, req.Monto, req.Emisor, req.Receptor, req.FechaTransaccion)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, t)
	}
}
func (t *Transaction) UpdatePartial() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64) //convert to int64 string id param , if it's not a number return an error
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Id inválido"})
			return
		}

		var req request // struct for entity without id(entire entity just will be visible for repositoy and service interfaces)
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if req.CodigoTransaccion == 0 || req.Monto == 0 {
			ctx.JSON(401, gin.H{"missing Fields": "Codigo de Transaccion ó Monto"})
			return
		}

		t, err := t.service.UpdatePartial(int(id), req.CodigoTransaccion, req.Monto)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, t)
	}
}
func (t *Transaction) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64) //convert to int64 string id param , if it's not a number return an error
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Id inválido"})
			return
		}

		t, err := t.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, t)
	}
}
func (r request) validateReq() ([]string, error) {
	var fields []string
	var missingField bool

	if r.CodigoTransaccion == 0 {
		fields = append(fields, "Codigo de Transaccion")
		missingField = true
	}
	if r.Moneda == "" {
		fields = append(fields, "Moneda")
		missingField = true
	}
	if r.Monto == 0 {
		fields = append(fields, "Codigo de monto")
		missingField = true
	}

	if r.Emisor == "" {
		fields = append(fields, "Codigo de emisor")
		missingField = true
	}
	if r.Receptor == "" {
		fields = append(fields, "Codigo de receptor")
		missingField = true
	}
	if r.FechaTransaccion == "" {
		fields = append(fields, "Codigo de Fecha de transaccion")
		missingField = true
	}

	if missingField {
		return nil, errors.New("Missing Fields")
	}
	return fields, nil
}
