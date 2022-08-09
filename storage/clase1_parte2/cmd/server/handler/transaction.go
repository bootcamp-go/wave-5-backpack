package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"goweb/internal/dto/requestdto"
	"goweb/internal/transactions"
	"goweb/pkg/web"
)



type Transaction struct {
	service transactions.IService
}

func NewHandler(transactionService transactions.IService) *Transaction {
	return &Transaction{service:transactionService}
}

// DeleteTransaction godoc
// @Summary Delete transaction by ID
// @Tags Transactions
// @Description delete transaction
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path int true "Transaction ID"
// @Router /transactions/{id} [delete]
func (transaction *Transaction) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Valido el TOKEN
		if err := validateToken(ctx); !err {
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "id invalido"))
			return
		}
		err = transaction.service.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("El producto con id: %d se ha eliminado correctamente", id)})
	}
}

// ListTransactions godoc
// @Summary Get transactions
// @Tags Transactions
// @Description Get all transactions
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Router /transactions [get]
func (transaction *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Valido el TOKEN
		if err := validateToken(ctx); !err {
			return
		}

		//Envio la peticion al servicio y almaceno las respuestas que me envien en 2 variables
		p, err := transaction.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
	}
}

// TransactionById godoc
// @Summary Get transaction by sender
// @Tags Transactions
// @Description Get transaction by sender
// @Accept json
// @Produce json
// @Param sender path string true "Transaction sender"
// @Param token header string true "token"
// @Router /transactions [get]
func (transaction *Transaction) GetBySender() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Valido el TOKEN
		if err := validateToken(ctx); !err {
			return
		}

		sender := ctx.Param("sender")

		t, err := transaction.service.GetBySender(sender)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, t, ""))
	}
}

// StoreTransactions godoc
// @Summary Store transactions
// @Tags Transactions
// @Description store new Transactions
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param token body requestdto.TransactionRequest true "Transaction to store"
// @Router /transactions [post]
func (transaction *Transaction) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Valido el TOKEN
		if err := validateToken(ctx); !err {
			return
		}

		var req requestdto.TransactionRequest
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		t, err := transaction.service.Store(req.CodTransaction, req.Currency, req.Amount, req.Sender, req.Receiver, req.DateOrder)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, t, ""))
	}
}

// UpdateTransaction godoc
// @Summary update transaction by ID
// @Tags Transactions
// @Description update transaction
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path int true "Transaction ID"
// @Param token body requestdto.TransactionRequest true "Transaction to store"
// @Router /transactions/{id} [put]
func (transaction *Transaction) Upddate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Valido el TOKEN
		if err := validateToken(ctx); !err {return}

		//Valido que el id sea correcto (un numero valido)
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "id invalido"))
			return
		}

		//Instancio una variable de tipo request para la transaccion y recibo la soliciutd enviada por URL en una variable de tipo Request
		var transactionRequest requestdto.TransactionRequest
		if err := ctx.ShouldBindJSON(&transactionRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		}

		//Valido que los campos vayan llenos uno por uno ANTES de almacenarlos
		if transactionRequest.CodTransaction == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "el codigo de transaccion es requerido"))
			return
		}
		if transactionRequest.Currency == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "la moneda es requerida"))
			return
		}
		if transactionRequest.Amount <= 0 {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "el valor debe ser mayor o diferente a 0"))
			return
		}
		if transactionRequest.Sender == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "el remitente es requerido"))
			return
		}
		if transactionRequest.Receiver == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "El receptor es requerido"))
			return
		}
		if transactionRequest.DateOrder == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "La fecha de la orden es requerido"))
			return
		}

		//Envio la soliciutd al servicio y espero la respuesta del respositiorio en 2 variables
		t, err := transaction.service.Update(int(id), transactionRequest.CodTransaction, transactionRequest.Currency, transactionRequest.Amount, transactionRequest.Sender, transactionRequest.Receiver, transactionRequest.DateOrder)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, t, ""))

	}
}

// UpdateAmount godoc
// @Summary update transaction (only amount) by ID
// @Tags Transactions
// @Description update transaction (only amount)
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path int true "Transaction ID"
// @Param token body requestdto.TransactionRequest true "Transaction to store"
// @Router /transactions/{id} [patch]
func (transaction *Transaction) UpdateAmount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Valido el TOKEN
		if err := validateToken(ctx); !err {return}

		//Valido que el id sea correcto (un numero valido)
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "id invalido"))
			return
		}

		//Instancio una variable de tipo request para la transaccion y recibo la soliciutd enviada por URL en una variable de tipo Request
		var transactionRequest requestdto.TransactionRequest
		if err := ctx.ShouldBindJSON(&transactionRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		}

		//Valido que los campos vayan llenos uno por uno ANTES de almacenarlos
		if transactionRequest.Amount <= 0 {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "El valor debe ser mayor o diferente a 0"))
			return
		}

		//Envio la soliciutd al servicio y espero la respuesta del respositiorio en 2 variables
		t, err := transaction.service.UpdateAmount(int(id), transactionRequest.Amount)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, t, ""))
	}
}

/*
/-------------------/
ADDITIONAL FUNCTIONS
/-------------------/
*/

func validateToken(ctx *gin.Context) bool {
	//Valido que venga el Token correcto
	token := ctx.Request.Header.Get("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token invalido"))
		return false
	}
	return true
}