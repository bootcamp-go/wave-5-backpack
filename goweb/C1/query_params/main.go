package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Empleado struct {
	Nombre string
	Codigo string
	Edad   int
	Activo int
}

var empls = []Empleado{
	{Nombre: "Lucas", Codigo: "12", Edad: 25, Activo: 1},
	{Nombre: "Maria", Codigo: "13", Edad: 21, Activo: 0},
	{Nombre: "Juan", Codigo: "14", Edad: 22, Activo: 1},
	{Nombre: "Nico", Codigo: "15", Edad: 22, Activo: 1},
}

func main() {
	r := gin.Default()

	r.GET("filtrarempleados", FiltrarEmpleado())
	r.GET("filtrarempleadosdos", FiltrarempleadosDos())

	r.Run()
}

func FiltrarEmpleado() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		edad, _ := strconv.Atoi(ctx.Query("edad"))
		activo, _ := strconv.Atoi(ctx.Query("activo"))

		fmt.Println(edad, activo)
		var results []Empleado
		for _, value := range empls {
			if edad == value.Edad && activo == value.Activo {
				results = append(results, value)
			}
		}

		ctx.JSON(200, results)
	}
}

func FiltrarempleadosDos() gin.HandlerFunc {
	return func(c *gin.Context) {

		var query struct {
			Edad   int `form:"edad"`
			Activo int `form:"activo"`
		}
		if err := c.ShouldBindQuery(&query); err != nil { // Setea las variables obtenidas de c.Query("nombredelavariable")
			c.JSON(400, err)
			return
		}

		var results []Empleado
		for _, value := range empls {
			if query.Edad == value.Edad && query.Activo == value.Activo {
				results = append(results, value)
			}
		}

		c.JSON(http.StatusOK, results) // devovemos el array filtrado
	}
}
