package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type usuario struct {
	id       int
	nombre   string
	apellido string
	email    string
	altura   float64
	activo   bool
	fecha    string
}

func main() {

	var usuarios []usuario
	data, err := ioutil.ReadFile("usuarios.json")
	json.Unmarshal([]byte(data), &usuarios)
	if err != nil {
		fmt.Println(err)
	}

	r := gin.Default()
	r.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hola Diego",
		})
	})
	r.GET("/usuarios", func(GetAll *gin.Context) {
		GetAll.JSON(http.StatusOK, &usuarios)
	})
	r.Run()

}
