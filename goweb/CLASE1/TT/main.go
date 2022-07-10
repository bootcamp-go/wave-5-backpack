package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type users struct {
	Id       int       `json:"id" binding:"required"`
	Nombre   string    `json:"nombre"`
	Apellido string    `json:"apellido"`
	Email    string    `json:"email"`
	Edad     int       `json:"edad"`
	Altura   float64   `json:"altura"`
	Activo   bool      `json:"activo"`
	Fecha    time.Time `json:"fecha"`
}

func usersJson() []users {
	jsonUsers, err := os.ReadFile("./users.json")
	if err != nil {
		fmt.Print(err)
	}
	var users []users
	err = json.Unmarshal(jsonUsers, &users)
	if err != nil {
		fmt.Print(err)
	}
	return users
}

func PaginaInicio(ctx *gin.Context) {
	ctx.String(200, "¡Hola, Bienvenido a mi primer ejercicio con Query!")
}

func AllUsers(ctx *gin.Context) {
	user := usersJson()
	ctx.JSON(200, gin.H{
		fmt.Sprintf("Todos los users\n"): user,
	})
}

//Ejercicio 2

func GetId(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(400, "El id ingresado no existe")
		return
	}
	user := usersJson()
	for _, valueUser := range user {
		if valueUser.Id == id {
			ctx.JSON(200, gin.H{
				"Usuario encontrado": valueUser,
			})
			return
		}
	}
	ctx.String(400, "No existe ningún usuario con el id ingresado")
}

func main() {

	server := gin.Default()
	server.GET("/", PaginaInicio)
	router := server.Group("/users")
	{
		//Ejercicio 2
		router.GET("/:id", GetId)
		//Ejercicio 1
		router.GET("", func(ctx *gin.Context) {
			route := ctx.Request.URL.Query()
			//Variable para aplicar query
			if route != nil {
				var usersU []users
				user := usersJson()
				for _, valueUser := range user {
					if route.Get("nombre") == valueUser.Nombre {
						usersU = append(usersU, valueUser)
					}
				}
				if len(usersU) > 0 {
					ctx.JSON(200, gin.H{
						"Filtro por nombre": usersU,
					})
				}
				var usersA []users
				userA := usersJson()
				for _, valueUser := range userA {
					if route.Get("apellido") == valueUser.Apellido {
						usersA = append(usersA, valueUser)
					}
				}
				if len(usersA) > 0 {
					ctx.JSON(200, gin.H{
						"Filtro por apellido": usersA,
					})
				}
				var usersE []users
				userE := usersJson()
				for _, valueUser := range userE {
					if route.Get("email") == valueUser.Email {
						usersE = append(usersE, valueUser)
					}
				}
				if len(usersE) > 0 {
					ctx.JSON(200, gin.H{
						"Filtro por email": usersE,
					})
				}
			} else {
				router.GET("", AllUsers)
			}

			// router.GET("/activo/:activo", BuscarUsersusersActivos)
		})
	}
	server.Run(":8081")
}
