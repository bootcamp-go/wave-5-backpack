package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//Esta función nos permite ver la anatomía de un mensaje Request de una
func handler(context *gin.Context) {
	//El body, header y method están contenidos en el contexto de gin.

	body := context.Request.Body
	header := context.Request.Header
	metodo := context.Request.Method

	fmt.Println("¡He recibido algo!")
	fmt.Printf("\tMetodo: %s\n", metodo)
	fmt.Printf("\tContenido del header:\n")

	for key, value := range header {
		fmt.Printf("\t\t%s -> %s\n", key, value)
	}
	fmt.Printf("\tEl body es un io.ReadCloser:(%s), y para trabajar con el vamos a tener que leerlo luego\n", body)
	fmt.Println("¡Yay!")
	context.String(200, "¡Lo recibí!") //Respondemos al cliente con 200 OK y un mensaje.
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.GET("/", handler)
	router.Run()

}
