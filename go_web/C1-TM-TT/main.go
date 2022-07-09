package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	ID        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Email     string `json:"email"`
	Edad      int    `json:"edad"`
	Altura    string `json:"altura"`
	Activo    string `json:"activo"`
	CreatedAt string `json:"createdAt"`
}

func main() {
	/*
		Crea dentro de la carpeta go-web un archivo llamado main.go
		Crea un servidor web con Gin que te responda un JSON que tenga una clave “message”
		y diga Hola seguido por tu nombre.
		Pegale al endpoint para corroborar que la respuesta sea la correcta.

		Ya habiendo creado y probado nuestra API que nos saluda, generamos una ruta que devuelve un listado de la temática elegida.
		Dentro del “main.go”, crea una estructura según la temática con los campos correspondientes.
		Genera un endpoint cuya ruta sea /temática (en plural). Ejemplo: “/productos”
		Genera un handler para el endpoint llamado “GetAll”.
		Crea una slice de la estructura, luego devuelvela a través de nuestro endpoint.

	*/

	r := gin.Default()
	r.GET("/saludos", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hola Rodri.go",
		})
	})
	//r.Run()

	//Ejercicio 2
	/*
		{"ID":1000,"nombre":"Roko","apellido":"Moonrock","email":"roko@dogmail.com",
		"edad":3,"altura":"45cms","activo":true,"createdAt":"6/07/22"}*/

	n := gin.Default()
	n.GET("/usuarios", getAll)
	//n.Run()
	endpointFilter()

}
func getAll(ctx *gin.Context) {
	var UsersBD []User
	newUser := User{ID: 1000, Nombre: "Roko", Apellido: "Moonrock", Email: "roko@dogmail.com", Edad: 3, Altura: "45cms", Activo: "si", CreatedAt: "06/07/22"}
	UsersBD = append(UsersBD, newUser)
	ctx.JSON(200, UsersBD)
}
