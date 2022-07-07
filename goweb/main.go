package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Transanction struct {
	Id       int     `json:"id"`
	Code     string  `json:"code"`
	Coin     string  `json:"coin"`
	Amount   float64 `json:"amount"`
	Emisor   string  `json:"emisor"`
	Receptor string  `json:"receptor"`
	Date     string  `json:"date"`
}

var Transancion = []Transanction{
	{Id: 1, Code: "QWE123", Coin: "COP", Amount: 1000000, Emisor: "Juan David", Receptor: "MeLi", Date: "06-06-2022"},
	{Id: 2, Code: "ASD123", Coin: "USD", Amount: 4000, Emisor: "Tulio", Receptor: "MePa", Date: "07-06-2022"},
	{Id: 3, Code: "ZXC123", Coin: "USD", Amount: 7000, Emisor: "Julio", Receptor: "MePa", Date: "06-06-2022"},
}

func Principal(ctx *gin.Context) {
	ctx.String(200, "Bienvenido a la pagina principal")
}

func GetAll(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, Transancion)
}

//Punto 1
func FiltrarCoin(ctx *gin.Context) {

	var filtrados []Transanction
	for _, t := range Transancion {
		if ctx.Query("coin") == t.Coin {
			filtrados = append(filtrados, t)

		}
	}
	ctx.JSON(http.StatusOK, filtrados)
}

//Punto 2
func BuscarId(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	for _, t := range Transancion {
		if t.Id == id {
			ctx.JSON(http.StatusOK, t)
			return
		}
	}

}

func main() {

	router := gin.Default()
	router.GET("/", Principal)
	router.GET("/todas-las-transacciones", GetAll)
	router.GET("/transacciones", FiltrarCoin)
	router.GET("/transacciones/:id", BuscarId)
	router.Run()

}

/*
sonFile, err := os.Open("usuarios.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
*/

/* primer forma de hacerlo :

// creo un router con Gin
router := gin.Default()

// captura de solicitud GET
router.GET("/saludo", func(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hola Juan David"})
})

// Corre nuestro servidor en el puerto :8080 por defecto
router.Run()
*/
