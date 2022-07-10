package main

import (
	"goweb/cmd/handler"
	"goweb/internals/transactions"
	"goweb/pkg/store"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ocurrió un error al leer el archivo .env")
	}

	db := store.NewStore("transacciones.json")
	if err := db.Ping(); err != nil {
		log.Fatal("error al intentar cargar archivo")
	}

	repo := transactions.NewRepository(db)
	service := transactions.NewService(repo)
	h := handler.NewTransaction(service)

	router := gin.Default()

	transacciones := router.Group("transacciones")
	transacciones.GET("/", h.GetAll())
	transacciones.GET("/filtros", h.GetByQuery())
	transacciones.GET("/:id", h.GetByID())
	transacciones.POST("/", h.Store())
	transacciones.PUT("/", h.Update())
	transacciones.DELETE("/", h.Delete())

	router.Run(":8080")
}

//Ejercicio 2
//Crea dentro de la carpeta go-web un archivo llamado main.go
//Crea un servidor web con Gin que te responda un JSON que tenga una clave “message” y diga Hola seguido por tu nombre.
//Pegale al endpoint para corroborar que la respuesta sea la correcta.

//Ejercicio 3
//Ya habiendo creado y probado nuestra API que nos saluda, generamos una ruta que devuelve un listado de la temática elegida.
//Dentro del “main.go”, crea una estructura según la temática con los campos correspondientes.
//Genera un endpoint cuya ruta sea /temática (en plural). Ejemplo: “/productos”
//Genera un handler para el endpoint llamado “GetAll”.
//Crea una slice de la estructura, luego devuelvela a través de nuestro endpoint.
