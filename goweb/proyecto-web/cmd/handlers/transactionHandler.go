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

func (t *TransactionHandler) GetAll() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		if !validarToken(ctx) {
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, t.service.GetAll(), ""))
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

func (t *TransactionHandler) Create() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		if !validarToken(ctx) {
			return
		}

		var request domain.Transaction
		err := ctx.ShouldBindJSON(&request)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
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
