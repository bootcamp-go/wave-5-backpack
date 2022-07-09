/*---------------------------------------------------------------------------------*

     Assignment:	Practica #1
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Web

	Description:
		‣	Exercise 1 - Let's filter our endpoint
		‣	Exercise 2 - Get one endpoint

	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------------------------------*/

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type transaccion struct {
	Id                int     `json:"id" binding:"-"`
	CodigoTransaccion string  `json:"codigo de transaccion" binding:"required"`
	Moneda            string  `json:"moneda" binding:"required"`
	Monto             float64 `json:"monto" binding:"required"`
	Emisor            string  `json:"emisor" binding:"required"`
	Receptor          string  `json:"receptor" binding:"required"`
	Fecha             string  `json:"fecha de transaccion" binding:"required"`
}

var lastId int = 5

//Este handler se encargará de responder a '/.'
func PaginaPrincipal(ctx *gin.Context) {
	ctx.String(200, "¡Bienvenido! Pagina Principal 🏠 ")
}

// Este handler para error 404
func error404(ctx *gin.Context) {
	ctx.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
}

//Este handler verificará si la id que pasa el cliente existe en nuestra base de datos.
func GetAll(ctx *gin.Context) {
	dataFile, err := getData()
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"transacciones": dataFile,
		})
	}
}

func GetOne(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, errStr := strconv.Atoi(idParam)
	if errStr != nil {
		ctx.String(404, errStr.Error())
	}
	transaccion, err := filtrarByIdOne(id)
	if err != nil {
		ctx.String(404, err.Error())
	} else {
		ctx.JSON(200, gin.H{
			"transaccion": transaccion,
		})
	}
}

func filtrarByIdOne(id int) (transaccion, error) {
	dataFile, err := getData()
	if err != nil {
		return transaccion{}, err
	}
	for _, transaction := range dataFile {
		if id == transaction.Id {
			return transaction, nil
		}
	}
	return transaccion{}, errors.New("> error. No hay ninguna transaccion con el Id ingresado")
}

func filtrarByTagHandler(ctx *gin.Context) {
	dataFile, err := getData()
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		detect := ctx.Request.URL.String()
		fmt.Println("Detectado: ", detect, " [byte]")
		datosFiltrados := filtrarDatosID(dataFile, "id")
		ctx.JSON(200, gin.H{
			"lista-id": datosFiltrados,
		})

	}
}
func filtrarDatosID(data []transaccion, id string) []interface{} {
	var datosFiltrados []interface{}

	type tmpId struct {
		Id int `json:"id"`
	}

	for _, i := range data {
		tmpId := tmpId{}
		tmpId.Id = i.Id
		datosFiltrados = append(datosFiltrados, tmpId)
	}
	return datosFiltrados
}

func getData() ([]transaccion, error) {
	var dataTransacciones []transaccion
	file, err := os.ReadFile("../transacciones.json")
	if err != nil {
		return dataTransacciones, err
	}
	if err := json.Unmarshal(file, &dataTransacciones); err != nil {
		return dataTransacciones, err
	}
	return dataTransacciones, nil
}

//	FUNCTION : POST
func Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req transaccion // create request | type: transaccion

		//	CHECK token for validate operation : POST
		token := ctx.GetHeader("token")
		if token != "12345" {
			ctx.JSON(401, gin.H{"error": "No tiene permisos para realizar la petición solicitada 😞😞😞"})
			return
		}

		//	GET DATA from File
		dataFile, err := getData()
		if err != nil { // Error Data get
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
		}
		if err := ctx.ShouldBindJSON(&req); err != nil { // Errors in JSON
			//	MESSAGES error for each Field : transacciones
			if req.CodigoTransaccion == "" {
				ctx.JSON(404, gin.H{"error": "El campo *Codigo de Transaccion* es requerido"})
			}
			if req.Moneda == "" {
				ctx.JSON(404, gin.H{"error": "El campo *Moneda*  es requerido"})
			}
			if req.Monto <= 0 {
				ctx.JSON(404, gin.H{"error": "El campo *Monto* es requerido"})
			}
			if req.Emisor == "" {
				ctx.JSON(404, gin.H{"error": "El campo *Emisor* es requerido"})
			}
			if req.Receptor == "" {
				ctx.JSON(404, gin.H{"error": "El campo *Receptor* es requerido"})
			}
			if req.Fecha == "" {
				ctx.JSON(404, gin.H{"error": "El campo *Fecha* es requerido"})
			}
			/* ctx.JSON(404, gin.H{
				"error": err.Error(),
			}) */
			return
		} else { // Continue : POST

			//	Get the last ID integer
			maxId := 0
			for _, value := range dataFile {
				if value.Id > maxId {
					maxId = value.Id
				}
			}
			maxId = maxId + 1
			req.Id = maxId
			//	POST request
			dataFile = append(dataFile, req)
			ctx.JSON(http.StatusAccepted, dataFile)
		}

	}
}

//	MAIN PROGRAM
func main() {

	server := gin.Default()

	server.GET("/hola", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hola Israel ! 👋",
		})
	})

	groupEndPoint := server.Group("/")
	{
		groupEndPoint.GET("/", PaginaPrincipal)
		groupEndPoint.GET("/transacciones", GetAll)
		groupEndPoint.GET("/transacciones/id", filtrarByTagHandler)
		groupEndPoint.GET("/transacciones/:id", GetOne)
		groupEndPoint.POST("/transacciones", Post())
	}

	server.NoRoute(error404)
	server.Run(":8080")
}
