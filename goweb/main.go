package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Usuarios struct {
	Id, Edad                                int
	Nombre, Apellido, Email, Fecha_creacion string
	Altura                                  float64
	Activo                                  bool
}

func main() {
	router := gin.Default()

	router.GET("/HolaNombre", HolaNombre)

	router.GET("/GetAll", GetAll)

	router.GET("/getuser/:id", getOneUser)
	router.GET("/filtrarusuarios", filterUsers)

	router.Run()
}

func HolaNombre(c *gin.Context) {
	c.JSON(200, gin.H{
		"mesage": "Hola Jose",
	})
}

func GetAll(ctx *gin.Context) {
	var u []Usuarios
	j, _ := os.ReadFile("./usuarios.json")
	if err := json.Unmarshal(j, &u); err != nil {
		log.Println(string(j))
		log.Fatal(err)
	}
	ctx.JSON(200, gin.H{
		"usuario": u,
	})
}

func getOneUser(ctex *gin.Context) {
	var u []Usuarios
	var u2 Usuarios
	j, _ := os.ReadFile("./usuarios.json")
	if err := json.Unmarshal(j, &u); err != nil {
		log.Println(string(j))
		log.Fatal(err)
	}
	id, _ := strconv.Atoi(ctex.Param("id"))
	for _, value := range u {
		if value.Id == id {
			u2 = value
			break
		}
	}
	ctex.JSON(200, gin.H{
		"usuario": u2,
	})
}

func filterUsers(ctx *gin.Context) {
	var u []Usuarios
	j, _ := os.ReadFile("./usuarios.json")
	if err := json.Unmarshal(j, &u); err != nil {
		log.Println(string(j))
		log.Fatal(err)
	}
	var u2 Usuarios
	ctx.ShouldBindQuery(&u2)

	var filtrado []Usuarios
	for _, t := range u {
		if u2.Id == t.Id && u2.Nombre == t.Nombre && u2.Apellido == t.Apellido && u2.Email == t.Email && u2.Edad == t.Edad && u2.Altura == t.Altura && u2.Activo == t.Activo && u2.Fecha_creacion == t.Fecha_creacion {
			filtrado = append(filtrado, t)
		}
	}

	ctx.JSON(http.StatusOK, filtrado)
}
