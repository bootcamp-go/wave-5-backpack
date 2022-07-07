package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type usuarios struct {
	Id            int     `json:"id"`
	Nombre        string  `json:"nombre" binding:"required"`
	Apellido      string  `json:"apellido" binding:"required"`
	Email         string  `json:"email" binding:"required"`
	Edad          int     `json:"edad" binding:"required"`
	Altura        float64 `json:"altura" binding:"required"`
	Activo        bool    `json:"activo" binding:"required"`
	FechaCreacion string  `json:"fecha_de_creacion" binding:"required"`
}

var users []usuarios = []usuarios{
	{Id: 1, Nombre: "ABC", Apellido: "ABC", Email: "a@a.com", Edad: 21, Altura: 1.82, Activo: true, FechaCreacion: "2022-02-20"},
	{Id: 2, Nombre: "BCD", Apellido: "BCD", Email: "b@b.com", Edad: 30, Altura: 1.76, Activo: false, FechaCreacion: "2022-02-25"},
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
	// ctx.Request.URL.Query()
	name := ctx.Query("name")
	result := []usuarios{}
	for i := 0; i < len(users); i++ {
		if strings.HasPrefix(users[i].Nombre, name) {
			result = append(result, users[i])
		}
	}
	ctx.JSON(200, result)
}

func Guardar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		expectedToken := "abcd"
		if ctx.GetHeader("token") != expectedToken {
			ctx.JSON(401, "no tiene permisos para realizar la operacion solicitada")
			return
		}
		var user usuarios
		if err := ctx.ShouldBindJSON(&user); err != nil {
			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				result := ""
				for i, field := range ve {
					if i != len(ve)-1 {
						result += fmt.Sprintf("el campo %s es requerido y ", field.Field())
					} else {
						result += fmt.Sprintf("el campo %s es requerido", field.Field())
					}
				}
				ctx.JSON(404, result)
			}
			return
		}
		user.Id = users[len(users)-1].Id + 1
		users = append(users, user)
		ctx.JSON(200, user)
	}
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
	router.POST("usuarios", Guardar())
	router.Run()
}
