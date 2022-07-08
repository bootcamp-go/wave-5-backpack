package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Definición estructura de json

type transacciones struct {
	Id                 int     `json:"id" binding:"-"`
	Codigo_transaccion string  `json:"codigo_transaccion" binding:"required"`
	Moneda             string  `json:"moneda" binding:"required"`
	Monto              float64 `json:"monto" binding:"required"`
	Emisor             string  `json:"emisor" binding:"required"`
	Receptor           string  `json:"receptor" binding:"required"`
	Fecha_transaccion  string  `json:"fecha_transaccion" binding:"required"`
}

// Se define variable [] de transacciones para transformar []byte data a formato transacciones struct

var t []transacciones
var oneTransaction transacciones
var lastID int

// Se crea el handler GetAll

func GetAll(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": t,
	})
}

// Handler para obtener por filtro con query

func GetByFilter(ctx *gin.Context) {

	var filterSlice []transacciones

	var (
		codigo_transaccion = ctx.Query("codigo_transaccion")
		moneda             = ctx.Query("moneda")
		monto, _           = strconv.ParseFloat(ctx.Query("monto"), 64)
		emisor             = ctx.Query("emisor")
		receptor           = ctx.Query("receptor")
		fecha_transaccion  = ctx.Query("fecha_transaccion")
	)

	fmt.Println(moneda)

	for _, value := range t {
		if value.Moneda == moneda {
			filterSlice = append(filterSlice, value)
		} else if value.Codigo_transaccion == codigo_transaccion {
			filterSlice = append(filterSlice, value)
		} else if value.Monto == monto {
			filterSlice = append(filterSlice, value)
		} else if value.Emisor == emisor {
			filterSlice = append(filterSlice, value)
		} else if value.Receptor == receptor {
			filterSlice = append(filterSlice, value)
		} else if value.Fecha_transaccion == fecha_transaccion {
			filterSlice = append(filterSlice, value)
		}
	}

	ctx.JSON(200, gin.H{
		"transactions": filterSlice,
	})

}

// Handler para una transaccion por id

func GetTransactionById(ctx *gin.Context) {

	id, ok := strconv.Atoi(ctx.Param("id"))
	if ok != nil {
		ctx.String(http.StatusBadRequest, "El id no fue encontrado")
	}

	for _, transaction := range t {
		if id == transaction.Id {
			oneTransaction = transaction
		}
	}

	ctx.JSON(200, gin.H{
		"transactions": oneTransaction,
	})

}

// Endpoint Post Clase II - TM

func PostTransaccion() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("token")
		if token != "12345" {
			ctx.JSON(401, gin.H{"error": "no tiene permisos para realizar la petición solicitada ;)"})
			return
		}

		var req transacciones
		if err := ctx.ShouldBindJSON(&req); err != nil {
			if req.Codigo_transaccion == "" {
				ctx.JSON(404, gin.H{"error": "El campo código es requerido"})
			}
			if req.Moneda == "" {
				ctx.JSON(404, gin.H{"error": "El campo moneda es requerido"})
			}
			if req.Monto == 0 {
				ctx.JSON(404, gin.H{"error": "El campo monto es requerido"})
			}
			if req.Emisor == "" {
				ctx.JSON(404, gin.H{"error": "El campo emisor es requerido"})
			}
			if req.Receptor == "" {
				ctx.JSON(404, gin.H{"error": "El campo receptor es requerido"})
			}
			if req.Fecha_transaccion == "" {
				ctx.JSON(404, gin.H{"error": "El campo fecha es requerido"})
			}
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		idu := len(t) - 1
		nuevoId := t[idu].Id + 1

		req.Id = nuevoId

		t = append(t, req)
		ctx.JSON(200, gin.H{
			"transactions": t,
		})

	}
}

func main() {

	// Se obtiene la data del archivo creado

	data, _ := os.ReadFile("./transacciones.json")

	if err := json.Unmarshal(data, &t); err != nil {
		log.Fatal(err)
	}

	// Se crea el routuer y los end points para generar el saludo

	router := gin.Default()

	router.GET("/saludo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola ! ",
		})
	})

	// Se crea el end point transacciones

	router.GET("/transacciones", GetAll)
	router.GET("/transaccionesfilter", GetByFilter)
	router.GET("/transacciones/:id", GetTransactionById)
	router.POST("/transacciones", PostTransaccion())

	router.Run()
}
