package handler

import (
	"fmt"
	"os"
	"strconv"

	"clase2-testing-prt1/internal/transactions"
	"clase2-testing-prt1/pkg/bank/web"

	"github.com/gin-gonic/gin"
)

const (
	InvalidToken = "Token inválido 🔐"
	IDinvalid    = "Invalid ID 🫠"
	FIELD_EMPTY  = "El campo * %s * es requerido para la transaccion 🤕"
)

type request struct {
	ID                int     `json:"id" binding:"-"`
	CodigoTransaccion string  `json:"codigo de transaccion" binding:"required"`
	Moneda            string  `json:"moneda" binding:"required"`
	Monto             float64 `json:"monto" binding:"required"`
	Emisor            string  `json:"emisor" binding:"required"`
	Receptor          string  `json:"receptor" binding:"required"`
	Fecha             string  `json:"fecha de transaccion" binding:"required"`
}

type requestPatch struct {
	ID                int     `json:"id" binding:"-"`
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

// ListTransactions godoc
// @Summary List of all transactions from database
// @Tags Transactions
// @Description Get all transactions from database
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Failure 204 {object} web.Response
// @Router /transactions [GET]
func (t *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, InvalidToken))
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

// EcommerceTransactions godoc
// @Summary New Ecommerce wtih transactions database
// @Tags Transactions
// @Description Ecommerce transactions
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param transaction body request true "Transaction to ecommerce"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 409 {object} web.Response
// @Router /transactions [POST]
func (t *Transaction) Ecommerce() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, InvalidToken))
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
			ctx.JSON(409, web.NewResponse(409, nil, err.Error()))
			return
		}
		ctx.JSON(200, t)
	}
}

// GetOne with id godoc
// @Summary Get a transaction with ID
// @Tags Transactions
// @Description Get transaction information using ID
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id path int true "id"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Router /transactions/{id} [GET]
func (t *Transaction) GetOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, InvalidToken))
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

// UpdateAll-id godoc
// @Summary Ecommerce transactions
// @Tags Transactions
// @Description Change all transaction information using ID but it's necessary to complete all the fields
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id path int true "id"
// @Param transaction body request true "Transaction to ecommerce"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Router /transactions/{id} [PUT]
func (t *Transaction) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, InvalidToken))
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

// UpdateTransaction godoc
// @Summary Ecommerce transactions
// @Tags Transactions
// @Description Change partial transaction information using ID, only complete the fields *monto* & *codigo de transaccion*
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id path int true "id"
// @Param transaction body request true "Transaction to ecommerce"
// @Success 200 {object} web.Response
// @Router /transactions/{id} [PATCH]
func (t *Transaction) UpdateOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, InvalidToken))
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, InvalidToken))
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
			ctx.JSON(409, web.NewResponse(409, nil, err.Error()))
			return
		}

		ctx.JSON(200, p)
	}
}

// DeleteTransaction-id godoc
// @Summary Delete transaction
// @Tags Transactions
// @Description Delete transaction
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id path int true "id"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Router /transactions/{id} [DELETE]
func (t *Transaction) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, InvalidToken))
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
