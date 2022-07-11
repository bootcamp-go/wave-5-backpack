package handlers

import (
	"net/http"
	"proyecto-web/internal/domain"
	"proyecto-web/internal/transaction"
	"strconv"

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

func (t *TransactionHandler) GetAll() gin.HandlerFunc {

	//var transacciones = c.service.GetAll()

	return func(ctx *gin.Context) {
		if !validarToken(ctx) {
			return
		}
		ctx.JSON(http.StatusAccepted, t.service.GetAll())
	}

	// codigo_transaccion := c.Query("codigo_transaccion")
	// moneda := c.Query("moneda")
	// emisor := c.Query("emisor")
	// receptor := c.Query("receptor")
	// fechaTransaccion := c.Query("fecha_transaccion")
	// monto := c.Query("monto")
	// id := c.Query("id")
	// idInt, errId := strconv.Atoi(id)
	// montoFloat, errMonto := strconv.ParseFloat(monto, 64)

	// if errId != nil && id != "" {
	// 	c.JSON(http.StatusBadRequest, gin.H{"message": "id inválido"})
	// 	return
	// }

	// if errMonto != nil && monto != "" {
	// 	c.JSON(http.StatusBadRequest, gin.H{"message": "monto inválido"})
	// 	return
	// }

	// transaccionesFiltradas := []domain.Transaccion{}

	// // El filtrado es con "o lógico". O sea, cualquier condición que se cumpla, se devuelve como resultado
	// for _, transaccion := range transaciones {
	// 	if transaccion.Id == idInt || transaccion.CodigoTransaccion == codigo_transaccion || transaccion.Moneda == moneda || transaccion.Monto == montoFloat || transaccion.Emisor == emisor || transaccion.Receptor == receptor || transaccion.FechaTransaccion == fechaTransaccion {
	// 		transaccionesFiltradas = append(transaccionesFiltradas, transaccion)
	// 	}
	// }
	// c.IndentedJSON(http.StatusOK, transaciones)

	//c.IndentedJSON(http.StatusOK, transaccionesFiltradas)
}

func (t *TransactionHandler) GetById() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		id, err1 := strconv.Atoi(ctx.Param("id"))
		if err1 != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Id inválido"})
			return
		}
		transaccion, err := t.service.GetById(id)

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, transaccion)
	}
}

func (t *TransactionHandler) Create() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		if !validarToken(ctx) {
			return
		}

		var request domain.Transaction
		err := ctx.ShouldBindJSON(&request)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var newTransaction = t.service.Create(request.Id, request.CodigoTransaccion, request.Moneda, request.Monto, request.Emisor, request.Receptor, request.FechaTransaccion)

		ctx.JSON(http.StatusOK, newTransaction)
	}
}

func (t *TransactionHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err1 := strconv.Atoi(ctx.Param("id"))
		if err1 != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Id inválido"})
			return
		}

		if !validarToken(ctx) {
			return
		}

		var request domain.Transaction
		err := ctx.ShouldBindJSON(&request)

		updatedTransaction, err := t.service.Update(id, request.CodigoTransaccion, request.Moneda, request.Monto, request.Emisor, request.Receptor, request.FechaTransaccion)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, updatedTransaction)
	}

}

func (t *TransactionHandler) UpdateParcial() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err1 := strconv.Atoi(ctx.Param("id"))
		if err1 != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Id inválido"})
			return
		}

		if !validarToken(ctx) {
			return
		}

		var request domain.Transaction
		err := ctx.ShouldBindJSON(&request)

		updatedTransaction, err := t.service.UpdateParcial(id, request.CodigoTransaccion, request.Monto)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, updatedTransaction)
	}

}

func validarToken(ctx *gin.Context) bool {
	token := ctx.GetHeader("Token")

	if token != "123456" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No autorizado"})
		return false
	}
	return true
}
