package main

import (
	"github.com/gin-gonic/gin"
	"encoding/json"
	"log"
	"os"
	"strconv"
)

type usuario struct {
	Id 				int		`json:"id"`
	Nombre 			string	`json:"nombre"`
	Apellido 		string	`json:"apellido"`
	Email 			string	`json:"email"`
	Edad 			int		`json:"edad"`
	Altura 			float64	`json:"altura"`
	Activo			bool	`json:"activo"`
	FechaCreacion 	string	`json:"fecha_creacion"`
}

var listaUsuarios []usuario

func GetAllUsers() {
	jsonUsers, err := os.ReadFile("./users.json")
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal([]byte(jsonUsers), &listaUsuarios); err != nil {
		log.Fatal(err)
	}
}

//Este handler verificará si la id que pasa el cliente existe en nuestra base de datos.
func BuscarUsuario(ctx *gin.Context) {
	for _, usuario := range listaUsuarios {
		if strconv.Itoa(usuario.Id) == ctx.Param("id") {
			ctx.String(200, "Información del usuario %s, nombre: %s", ctx.Param("id"), usuario.Nombre)
			return
		}
	}
	ctx.String(404, "El usuario no existe.")
}

func FiltrarUsuarios(ctx *gin.Context) {
	//var filtrados = listaUsuarios
	var filtrados []usuario
	for _, usuario := range listaUsuarios {
		if ctx.Query("activo") == strconv.FormatBool(usuario.Activo) {
			filtrados = append(filtrados, usuario)
		}

		if ctx.Query("nombre") == usuario.Nombre {
			filtrados = append(filtrados, usuario)
		}

		if ctx.Query("apellido") == usuario.Apellido {
			filtrados = append(filtrados, usuario)
		}

		if ctx.Query("email") == usuario.Email {
			filtrados = append(filtrados, usuario)
		}

		if ctx.Query("edad") == strconv.Itoa(usuario.Edad) {
			filtrados = append(filtrados, usuario)
		}

		if ctx.Query("altura") == strconv.FormatFloat(usuario.Altura, 'f', 3, 32) {
			filtrados = append(filtrados, usuario)
		}

		if ctx.Query("fecha_creacion") == usuario.FechaCreacion {
			filtrados = append(filtrados, usuario)
		}
	}

	if len(filtrados) == 0 {
		ctx.String(404, "No hay resultados para esta búsqueda.")
	} else {
		ctx.JSON(200, filtrados)
	}
}

func main() {
	GetAllUsers()
	server := gin.Default()
	server.GET("/usuarios", FiltrarUsuarios)
	server.GET("/usuarios/:id", BuscarUsuario)
	server.Run()
}