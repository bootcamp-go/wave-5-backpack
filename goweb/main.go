package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type usuarios struct {
	Id            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Apellido      string  `json:"apellido"`
	Email         string  `json:"email"`
	Edad          int     `json:"edad"`
	Altura        float64 `json:"altura"`
	Activo        bool    `json:"activo"`
	FechaCreacion string  `json:"fecha_de_creacion"`
}

var users []usuarios = []usuarios{
	{Id: 1, Nombre: "ABC", Apellido: "ABC", Email: "a@a.com", Edad: 21, Altura: 1.82, Activo: true, FechaCreacion: "2022-02-20"},
	{Id: 4, Nombre: "BCD", Apellido: "BCD", Email: "b@b.com", Edad: 30, Altura: 1.76, Activo: false, FechaCreacion: "2022-02-25"},
	{},
}

func GetAll(ctx *gin.Context) {
	ctx.JSON(200, users)
}

func GetById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	for i := 0; i < len(users); i++ {
		if id == users[i].Id {
			fmt.Println("LO ENCONTRE!")
			ctx.JSON(200, users[i])
			return
		}
	}
	fmt.Println("NO LO ENCONTRE!")
}

func FilterByName(ctx *gin.Context) {
	name := ctx.Query("name")
	result := []usuarios{}
	for i := 0; i < len(users); i++ {
		if strings.HasPrefix(users[i].Nombre, name) {
			result = append(result, users[i])
		}
	}
	ctx.JSON(200, result)
}

func main() {
	// para leer json desde archivo
	/*
		byteArr, err := os.ReadFile("./usuarios.json")
		if err != nil {
			log.Fatal("No se pudo abrir el .json")
		}
		var usuarios []usuarios
		err2 := json.Unmarshal(byteArr, &usuarios)
		if err2 != nil {
			fmt.Println(err2.Error())
			log.Fatal("No representa a usuarios")
		}
	*/
	// dummy
	/*
		router.GET("hola", func(ctx *gin.Context) {
			ctx.JSON(200, "Hola Andres!")
		})
	*/
	router := gin.Default()
	router.GET("usuarios/:id", GetById)
	router.GET("usuarios", FilterByName)
	router.Run()
}
