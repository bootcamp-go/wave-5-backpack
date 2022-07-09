package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

var usuariosRegistro []User

func endpointFilter() {
	/*
		Según la temática elegida, necesitamos agregarles filtros a nuestro endpoint, el mismo se tiene que poder filtrar por todos los campos.
		Dentro del handler del endpoint, recibí del contexto los valores a filtrar.
		Luego genera la lógica de filtrado de nuestro array.
		Devolver por el endpoint el array filtrado.

	*/
	usuarios := opFileJson()
	for _, v := range usuarios {
		usuariosRegistro = append(usuariosRegistro, v)
	}
	r := gin.Default()
	r.GET("/usuarios", endpointHandler)
	r.Run()
}

func endpointHandler(ctx *gin.Context) {
	var resultsBox []User
	var searchQuerys []string
	//var user User
	idQ := ctx.Query("id")
	if idQ != "" {
		searchQuerys = append(searchQuerys, idQ)
	}
	nombreQ := ctx.Query("nombre")
	if nombreQ != "" {
		searchQuerys = append(searchQuerys, nombreQ)
	}
	apellidoQ := ctx.Query("apellido")
	if apellidoQ != "" {
		searchQuerys = append(searchQuerys, apellidoQ)
	}
	emailQ := ctx.Query("email")
	if emailQ != "" {
		searchQuerys = append(searchQuerys, emailQ)
	}
	edadQ := ctx.Query("edad")
	if edadQ != "" {
		searchQuerys = append(searchQuerys, edadQ)
	}
	alturaQ := ctx.Query("altura")
	if alturaQ != "" {
		searchQuerys = append(searchQuerys, alturaQ)
	}
	activoQ := ctx.Query("activo")
	if activoQ != "" {
		searchQuerys = append(searchQuerys, activoQ)
	}
	fechaQ := ctx.Query("createdAt")
	if fechaQ != "" {
		searchQuerys = append(searchQuerys, fechaQ)
	}

	for _, user := range usuariosRegistro {
		if user.Activo == searchQuerys[0] {
			resultsBox = append(resultsBox, user)
		}
	}
	if len(resultsBox) != 0 {
		result := fmt.Sprint(resultsBox)
		ctx.String(200, result)
		return
	}
	ctx.String(400, "Ningun usuario coincide con tu busqueda")

}

func opFileJson() []User {
	var data []User

	file, err := ioutil.ReadFile("usuarios.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
