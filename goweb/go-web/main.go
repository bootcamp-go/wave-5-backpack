package main

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Id       string `json:"id"`
	Code     string `json:"code"`
	Moneda   string `json:"moneda"`
	Monto    int    `json:"monto"`
	Emisor   string `json:"emisor"`
	Receptor string `json:"receptor"`
	Fecha    string `json:"fecha"`
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

func filterById(id string) (Transaction, error) {
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

func filterListById(transacciones []Transaction, id string) []Transaction {
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
			transacciones = filterListById(transacciones, id)
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
	transaction, err := filterById(id)
	if err != nil {
		c.String(404, err.Error())
	} else {
		c.JSON(200, gin.H{
			"transaccion": transaction,
		})
	}
}

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
		transacciones.GET("/", filterByFieldsHandler)
		transacciones.GET("/:id", filterByIdHandler)
	}

	router.Run()
}
