package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type Transactions struct {
	Id              int     `json:"id"`
	TransactionCode string  `json:"transaction_code"`
	TypeCurrency    string  `json:"type_of_currency"`
	Amount          float64 `json:"amount"`
	Transmitter     string  `json:"transmitter"`
	Receiver        string  `json:"receiver"`
	Date            string  `json:"date"`
	Completed       bool    `json:"completed"`
}

func GetAll(c *gin.Context) {
	var transaction []Transactions
	jsonFile, err := ioutil.ReadFile("transactions.json")
	if err != nil {
		fmt.Errorf("el error es %v", err)
	}

	if err := json.Unmarshal(jsonFile, &transaction); err != nil {
		c.JSON(500, gin.H{
			"msg": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"json": transaction,
	})

}

/*func GetQuery(c *gin.Context) {
	var filter []*Transactions
	for i, values := range transaction {
		if c.Query("type_of_currecncy") == values.TypeCurrency {
			filter = append(filter, &values)
		}
	}
}*/

func main() {

	router := gin.Default()
	router.GET("/transactions", GetAll)
	//router.GET("/transactions", GetQuery)
	router.Run(":3000")
}
