package handler

import (
	"goweb/internals/transactions"
	"goweb/pkg/web"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type request struct {
	Codigo   string `json:"codigo"`
	Moneda   string `json:"moneda"`
	Monto    int    `json:"monto"`
	Emisor   string `json:"emisor"`
	Receptor string `json:"receptor"`
}

type Transaction struct {
	service transactions.Service
}

func NewTransaction(t transactions.Service) *Transaction {
	return &Transaction{
		service: t,
	}
}

//ListProducts godocs
//@Summary List of Transactions
//@Tags Transactions
//@Description Get All Transactions
//@Accept json
//@Produce json
//Success 200 {object} web.Response
//Failed 400 {object} web.Response
//@Router /transacciones [get]
func (t *Transaction) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		trans, err := t.service.GetAll()
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, trans, ""))
	}
}

//ListProducts godocs
//@Summary List of Transactions By Query
//@Tags Transactions
//@Description Get Transactions By Query
//@Accept json
//@Produce json
//Success 200 {object} web.Response
//Failed 400 {object} web.Response
//@Param query header string true "query"
//@Router /transacciones/filtros [get]
func (t *Transaction) GetByQuery() gin.HandlerFunc {
	return func(c *gin.Context) {
		trans, err := t.service.GetAll()
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
		}
		monto, _ := strconv.Atoi(c.Query("monto"))
		for _, t := range trans {
			if t.Codigo == c.Query("codigo") {
				c.JSON(http.StatusAccepted, web.NewResponse(200, t, ""))
			}
			if t.Moneda == c.Query("moneda") {
				c.JSON(http.StatusAccepted, web.NewResponse(200, t, ""))
			}
			if t.Monto == monto {
				c.JSON(http.StatusAccepted, web.NewResponse(200, t, ""))
			}
			if t.Emisor == c.Query("emisor") {
				c.JSON(http.StatusAccepted, web.NewResponse(200, t, ""))
			}
			if t.Receptor == c.Query("receptor") {
				c.JSON(http.StatusAccepted, web.NewResponse(200, t, ""))
			}
			if t.Fecha == c.Query("fecha") {
				c.JSON(http.StatusAccepted, web.NewResponse(200, t, ""))
			}
		}
	}
}

//ListProducts godocs
//@Summary Get Transaction by ID
//@Tags Transaction
//@Description Get Transaction by ID
//@Accept json
//@Produce json
//Success 200 {object} web.Response
//Failed 400 {object} web.Response
//@Param id path int true "id"
//@Router /transacciones/ID [get]
func (t *Transaction) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		trans, err := t.service.GetAll()
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		for _, trans := range trans {
			if trans.Id == id {
				c.JSON(http.StatusAccepted, web.NewResponse(200, trans, ""))
				return
			}
		}
		c.JSON(404, web.NewResponse(404, nil, "El ID no existe"))
	}
}

//ListProducts godocs
//@Summary Post a Transacction
//@Tags Transaction
//@Description Post a Transacction
//@Accept json
//@Produce json
//Success 200 {object} web.Response
//Failed 400 {object} web.Response
//@Param token header string true "token"
//@Router /transacciones [post]
func (t *Transaction) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		trans, err := t.service.Store(req.Codigo, req.Moneda, req.Monto, req.Emisor, req.Receptor)
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, trans, ""))
	}
}

//ListProducts godocs
//@Summary Update a Transacction
//@Tags Transaction
//@Description Update a Transacction
//@Accept json
//@Produce json
//Success 200 {object} web.Response
//Failed 400 {object} web.Response
//@Param token header string true "token"
//@Router /transacciones [put]
func (t *Transaction) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, web.NewResponse(401, nil, "No tiene permisos para realizar la petición solicitada ;)"))
			return
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		if req.Codigo == "" {
			c.JSON(400, web.NewResponse(400, nil, "El código del producto es requerido"))
			return
		}
		if req.Moneda == "" {
			c.JSON(400, web.NewResponse(400, nil, "El tipo de moneda es requerido"))
			return
		}
		if req.Monto == 0 {
			c.JSON(400, web.NewResponse(400, nil, "El monto del producto es requerido"))
			return
		}
		if req.Emisor == "" {
			c.JSON(400, web.NewResponse(400, nil, "El emisor del producto es requerido"))
			return
		}
		if req.Receptor == "" {
			c.JSON(400, web.NewResponse(400, nil, "El receptor del producto es requerido"))
			return
		}
		trans, err := t.service.Update(int(id), req.Codigo, req.Moneda, req.Monto, req.Emisor, req.Receptor)
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, trans, ""))
	}
}

//ListProducts godocs
//@Summary Delete a Transacction
//@Tags Transaction
//@Description Delete a Transacction
//@Accept json
//@Produce json
//Success 200 {object} web.Response
//Failed 400 {object} web.Response
//@Param token header string true "token"
//@Router /transacciones [delete]
func (t *Transaction) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, web.NewResponse(401, nil, "No tiene permisos para realizar la petición solicitada ;)"))
			return
		}
		// id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, "Invalid ID"))
			return
		}
		err = t.service.Delete(id)
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, nil, "El producto ha sido eliminado exitosamente"))
	}
}
