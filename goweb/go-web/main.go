package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Id       int    `json:"id"`
	Code     string `json:"code" binding:"required"`
	Moneda   string `json:"moneda" binding:"required"`
	Monto    int    `json:"monto" binding:"required"`
	Emisor   string `json:"emisor" binding:"required"`
	Receptor string `json:"receptor" binding:"required"`
	Fecha    string `json:"fecha" binding:"required"`
}

func (t Transaction) getMissingField() string {
	if t.Code == "" {
		return "code"
	}
	if t.Moneda == "" {
		return "moneda"
	}
	if t.Monto == 0 {
		return "monto"
	}
	if t.Emisor == "" {
		return "emisor"
	}
	if t.Receptor == "" {
		return "receptor"
	}
	if t.Fecha == "" {
		return "fecha"
	}
	return ""
}

func getAll() ([]Transaction, error) {
	var transactions []Transaction
	file, err := os.ReadFile("./transactions.json")
	if err != nil {
		return transactions, err
	}
	if err2 := json.Unmarshal(file, &transactions); err != nil {
		return transactions, err2
	}
	return transactions, nil
}

func filterById(id int) (Transaction, error) {
	transactions, err := getAll()
	if err != nil {
		return Transaction{}, err
	}
	for _, t := range transactions {
		if t.Id == id {
			return t, nil
		}
	}
	return Transaction{}, errors.New("No se encontro el registro")
}

func filterListById(transacciones []Transaction, id int) []Transaction {
	var transaccionesFiltradas []Transaction
	for _, t := range transacciones {
		if t.Id == id {
			transaccionesFiltradas = append(transaccionesFiltradas, t)
		}
	}
	return transaccionesFiltradas
}

func filterListByCode(transacciones []Transaction, code string) []Transaction {
	var transaccionesFiltradas []Transaction
	for _, t := range transacciones {
		if t.Code == code {
			transaccionesFiltradas = append(transaccionesFiltradas, t)
		}
	}
	return transaccionesFiltradas
}

func filterListByMoneda(transacciones []Transaction, moneda string) []Transaction {
	var transaccionesFiltradas []Transaction
	for _, t := range transacciones {
		if t.Moneda == moneda {
			transaccionesFiltradas = append(transaccionesFiltradas, t)
		}
	}
	return transaccionesFiltradas
}

func filterListByMonto(transacciones []Transaction, monto int) []Transaction {
	var transaccionesFiltradas []Transaction
	for _, t := range transacciones {
		if t.Monto == monto {
			transaccionesFiltradas = append(transaccionesFiltradas, t)
		}
	}
	return transaccionesFiltradas
}

func filterListByEmisor(transacciones []Transaction, emisor string) []Transaction {
	var transaccionesFiltradas []Transaction
	for _, t := range transacciones {
		if t.Emisor == emisor {
			transaccionesFiltradas = append(transaccionesFiltradas, t)
		}
	}
	return transaccionesFiltradas
}

func filterListByReceptor(transacciones []Transaction, receptor string) []Transaction {
	var transaccionesFiltradas []Transaction
	for _, t := range transacciones {
		if t.Receptor == receptor {
			transaccionesFiltradas = append(transaccionesFiltradas, t)
		}
	}
	return transaccionesFiltradas
}

func filterListByFecha(transacciones []Transaction, fecha string) []Transaction {
	var transaccionesFiltradas []Transaction
	for _, t := range transacciones {
		if t.Fecha == fecha {
			transaccionesFiltradas = append(transaccionesFiltradas, t)
		}
	}
	return transaccionesFiltradas
}

func getAllHandler(c *gin.Context) {
	transacciones, err := getAll()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"transacciones": transacciones,
		})
	}
}

func filterByFieldsHandler(c *gin.Context) {
	transacciones, err := getAll()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		id := c.Query("id")
		if id != "" {
			i, _ := strconv.Atoi(id)
			transacciones = filterListById(transacciones, i)
		}
		code := c.Query("code")
		if code != "" {
			transacciones = filterListByCode(transacciones, code)
		}
		moneda := c.Query("moneda")
		if moneda != "" {
			transacciones = filterListByMoneda(transacciones, moneda)
		}
		monto := c.Query("monto")
		if monto != "" {
			m, _ := strconv.Atoi(monto)
			transacciones = filterListByMonto(transacciones, m)
		}
		emisor := c.Query("emisor")
		if emisor != "" {
			transacciones = filterListByEmisor(transacciones, emisor)
		}
		receptor := c.Query("receptor")
		if receptor != "" {
			transacciones = filterListByReceptor(transacciones, receptor)
		}
		fecha := c.Query("fecha")
		if fecha != "" {
			transacciones = filterListByFecha(transacciones, fecha)
		}
		c.JSON(200, gin.H{
			"transacciones": transacciones,
		})
	}
}

func filterByIdHandler(c *gin.Context) {
	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	transaction, err := filterById(i)
	if err != nil {
		c.String(404, err.Error())
	} else {
		c.JSON(200, gin.H{
			"transaccion": transaction,
		})
	}
}

func insertHandler(c *gin.Context) {
	token := c.GetHeader("token")
	if token != TOKEN {
		c.String(401, "no tiene permisos para realizar la peticion solicitada")
		return
	}
	var req Transaction
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(400, fmt.Sprintf("el campo %s es requerido", req.getMissingField()))
		return
	}
	lastID++
	req.Id = lastID
	transaccionesTemporales = append(transaccionesTemporales, req)
	c.JSON(200, transaccionesTemporales)
}

const TOKEN = "12345"

var transaccionesTemporales []Transaction
var lastID int

func main() {
	router := gin.Default()

	router.GET("/hola", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Cristobal",
		})
	})

	router.GET("/transacciones", getAllHandler)

	transacciones := router.Group("/transaccion")
	{
		transacciones.POST("/", insertHandler)
		transacciones.GET("/", filterByFieldsHandler)
		transacciones.GET("/:id", filterByIdHandler)
	}

	router.Run()
}
