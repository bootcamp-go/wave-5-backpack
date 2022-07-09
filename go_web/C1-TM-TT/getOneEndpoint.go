package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var usuariosNode []User

func getOneEndpoint() {
	/*
		Según la temática elegida, necesitamos agregarles filtros a nuestro endpoint, el mismo se tiene que poder filtrar por todos los campos.
		Dentro del handler del endpoint, recibí del contexto los valores a filtrar.
		Luego genera la lógica de filtrado de nuestro array.
		Devolver por el endpoint el array filtrado.

	*/
	usuarios := opFileJson2()
	for _, v := range usuarios {
		usuariosNode = append(usuariosNode, v)
	}
	fmt.Println(usuariosNode)
	r := gin.Default()
	r.GET("/usuarios/:id", endpointHandler2)
	r.Run()
}

func endpointHandler2(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	var selectedUser User
	var found bool
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range usuariosNode {
		if v.ID == userId {
			selectedUser = v
			found = true
		}
	}
	if found {
		response := fmt.Sprint(selectedUser)
		ctx.String(http.StatusOK, response)
		return
	}

	ctx.String(http.StatusBadRequest, "El usuario %s no existe!!", ctx.Param("id"))

}

func opFileJson2() []User {
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
