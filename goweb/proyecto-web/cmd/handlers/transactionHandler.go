package handlers

import (
	"net/http"
	"os"
	"proyecto-web/internal/domain"
	"proyecto-web/internal/transaction"
	"strconv"

	"proyecto-web/pkg/web"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	service transaction.ITransactionService
}

func NewTransactionHandler(t transaction.ITransactionService) *TransactionHandler {
	return &TransactionHandler{
		service: t,
	}
}

// @Summary Lista todos las transacciones
// @Tags Transacciones
// @Description get transaction
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /transacciones [get]
func (t *TransactionHandler) GetAll() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		if !validarToken(ctx) {
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, t.service.GetAll(), ""))
	}
}

// @Summary Obtiene una transacción por su ID
// @Tags Transacciones
// @Description get transaction
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path int true "id"
// @Success 200 {object} web.Response
// @Router /transacciones/{id} [get]
func (t *TransactionHandler) GetById() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		id, err1 := strconv.Atoi(ctx.Param("id"))
		if err1 != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "id inválido"))
			return
		}
		transaccion, err := t.service.GetById(id)

		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, transaccion, ""))
	}
}

// @Summary Crea una nueva transacción
// @Tags Transacciones
// @Description post transaction
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param transaction body domain.Transaction true "Transaccion a crear"
// @Success 200 {object} web.Response
// @Router /transacciones [post]
// @Failure 400 {object} web.Response "error"
func (t *TransactionHandler) Create() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		if !validarToken(ctx) {
			return
		}

		var request domain.Transaction
		err := ctx.ShouldBindJSON(&request)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		valido := validarCampos(request, ctx)
		if !valido {
			return
		}

		var newTransaction = t.service.Create(request.Id, request.CodigoTransaccion, request.Moneda, request.Monto, request.Emisor, request.Receptor, request.FechaTransaccion)

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, newTransaction, ""))
	}
}

// UpdateTransactions godoc
// @Summary Actualiza una transacción
// @Tags Transacciones
// @Description put transaction
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param transaction body domain.Transaction true "Transaccion a crear"
// @Param id path int true "id"
// @Success 200 {object} web.Response
// @Router /transacciones/{id} [put]
func (t *TransactionHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err1 := strconv.Atoi(ctx.Param("id"))
		if err1 != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "id inválido"))
			return
		}

		if !validarToken(ctx) {
			return
		}

		var request domain.Transaction
		err := ctx.ShouldBindJSON(&request)

		updatedTransaction, err := t.service.Update(id, request.CodigoTransaccion, request.Moneda, request.Monto, request.Emisor, request.Receptor, request.FechaTransaccion)

		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, updatedTransaction, ""))
	}

}

// @Summary Actualiza parcialmente una transacción
// @Description patch transaction
// @Tags Transacciones
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param transaction body domain.Transaction true "Transaccion a crear"
// @Param id path int true "id"
// @Success 200 {object} web.Response
// @Router /transacciones/{id} [patch]
func (t *TransactionHandler) UpdateParcial() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err1 := strconv.Atoi(ctx.Param("id"))
		if err1 != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "id inválido"))
			return
		}

		if !validarToken(ctx) {
			return
		}

		var request domain.Transaction
		err := ctx.ShouldBindJSON(&request)

		updatedTransaction, err := t.service.UpdateParcial(id, request.CodigoTransaccion, request.Monto)

		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, updatedTransaction, ""))
	}
}

// @Summary Elimina una transacción
// @Description delete transaction
// @Tags Transacciones
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path int true "id"
// @Success 200 {object} web.Response
// @Router /transacciones/{id} [delete]
func (t *TransactionHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err1 := strconv.Atoi(ctx.Param("id"))
		if err1 != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "id inválido"))
			return
		}

		if !validarToken(ctx) {
			return
		}

		err := t.service.Delete(id)

		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, "Delete exitoso", ""))
		return
	}
}

func validarToken(ctx *gin.Context) bool {
	token := ctx.GetHeader("Token")

	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "no autorizado"))
		return false
	}
	return true
}

func validarCampos(req domain.Transaction, c *gin.Context) bool {
	if req.CodigoTransaccion == "" {
		c.JSON(400, web.NewResponse(400, nil, "El código de transacción es requerido"))
		return false
	}
	if req.Emisor == "" {
		c.JSON(400, web.NewResponse(400, nil, "El emisor es requerido"))
		return false
	}
	if req.Monto == 0 {
		c.JSON(400, web.NewResponse(400, nil, "El monto es requerida"))
		return false
	}
	if req.Moneda == "" {
		c.JSON(400, web.NewResponse(400, nil, "La moneda es requerida"))
		return false
	}
	if req.Receptor == "" {
		c.JSON(400, web.NewResponse(400, nil, "El receptor es requerido"))
		return false
	}
	if req.FechaTransaccion == "" {
		c.JSON(400, web.NewResponse(400, nil, "La fecha es requerida"))
		return false
	}
	return true
}
