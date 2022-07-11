package handler

import (
	"fmt"
	"os"
	"strconv"

	"goweb/clase4-go-web-tm/internal/transactions"
	"goweb/clase4-go-web-tm/pkg/bank/web"

	"github.com/gin-gonic/gin"
)

const (
	INVALID     = "Token inválido 🔐"
	IDinvalid   = "Invalid ID 🫠"
	FIELD_EMPTY = "El campo * %s * es requerido para la transaccion 🤕"
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
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, INVALID))
			return
		}
		p, err := t.service.GetAll()
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, p)
	}
}

func (t *Transaction) Ecommerce() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, INVALID))
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		t, err := t.service.Ecommerce(req.CodigoTransaccion, req.Moneda, req.Monto,
			req.Emisor, req.Receptor, req.Fecha)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, t)
	}
}

func (t *Transaction) GetOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, INVALID))
			return
		}

		detect := ctx.Request.URL.String()
		fmt.Println("Detectado: ", detect[14:], " [byte]")

		values := ctx.Request.URL.Query()
		for k, v := range values {
			fmt.Println(k, " => ", v)
		}

		idParam := ctx.Param("id")

		id, errStr := strconv.Atoi(idParam)
		if errStr != nil {
			ctx.String(404, errStr.Error())
		}

		p, err := t.service.GetOne(id)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, p)
	}
}

func (t *Transaction) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, INVALID))
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		errs := validate(&req)
		if len(errs) > 0 {
			ctx.JSON(400, gin.H{
				"errors": errs,
			})
			return
		}

		t, err := t.service.Update(int(id), req.CodigoTransaccion, req.Moneda,
			req.Monto, req.Emisor, req.Receptor, req.Fecha)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		ctx.JSON(200, t)
	}
}

func (t *Transaction) UpdateOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, INVALID))
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, INVALID))
			return
		}

		var req requestPatch
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		errs := validatePatch(&req)
		if len(errs) > 0 {
			ctx.JSON(400, gin.H{
				"errors": errs,
			})
			return
		}

		p, err := t.service.UpdateOne(int(id), req.CodigoTransaccion, req.Monto)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(200, p)
	}
}

func (t *Transaction) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, INVALID))
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, IDinvalid))
			return
		}

		err = t.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, gin.H{"data": fmt.Sprintf("La transaccion con id: %d, ha sido eliminado ✅", id)})
	}
}

func PaginaPrincipal(ctx *gin.Context) {
	ctx.String(200, "¡Bienvenido! Pagina Principal 🏠 ")
}

func validate(req *request) []string {
	var errors []string
	//	MESSAGES error for each Field : transacciones
	if req.CodigoTransaccion == "" {
		errors = append(errors, fmt.Errorf(FIELD_EMPTY, "Codigo de Trransaccion").Error())
	}

	if req.Moneda == "" {
		errors = append(errors, fmt.Errorf(FIELD_EMPTY, "Moneda").Error())
	}
	if req.Monto <= 0 {
		errors = append(errors, fmt.Errorf(FIELD_EMPTY, "Monto").Error())
	}
	if req.Emisor == "" {
		errors = append(errors, fmt.Errorf(FIELD_EMPTY, "Emisor").Error())
	}
	if req.Receptor == "" {
		errors = append(errors, fmt.Errorf(FIELD_EMPTY, "Receptor").Error())
	}
	if req.Fecha == "" {
		errors = append(errors, fmt.Errorf(FIELD_EMPTY, "Fecha").Error())
	}
	return errors
}

func validatePatch(req *requestPatch) []string {
	var errors []string
	//	MESSAGES error for each Field : transacciones
	if req.CodigoTransaccion == "" {
		errors = append(errors, fmt.Errorf(FIELD_EMPTY, "Codigo de Trransaccion").Error())
	}

	if req.Monto <= 0 {
		errors = append(errors, fmt.Errorf(FIELD_EMPTY, "Monto").Error())
	}

	return errors
}
