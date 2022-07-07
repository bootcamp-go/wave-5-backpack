package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type transaccion struct {
	Id                int     `json:"id"`
	CodigoTransaccion string  `json:"codigo de transaccion"`
	Moneda            string  `json:"moneda"`
	Monto             float64 `json:"monto"`
	Emisor            string  `json:"emisor"`
	Receptor          string  `json:"receptor"`
	Fecha             string  `json:"fecha de transaccion"`
}

//Este handler se encargará de responder a /.
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
	}

	server.NoRoute(error404)
	server.Run(":8080")
}
