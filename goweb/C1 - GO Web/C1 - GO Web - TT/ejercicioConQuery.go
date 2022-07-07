package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type usuarios struct {
	Id       int       `json:"id" binding:"required"`
	Nombre   string    `json:"nombre"`
	Apellido string    `json:"apellido"`
	Email    string    `json:"email"`
	Edad     int       `json:"edad"`
	Altura   float64   `json:"altura"`
	Activo   bool      `json:"activo"`
	Fecha    time.Time `json:"fecha"`
}

//Leo json de usuarios
func usersJson() []usuarios {
	//Leo el json y lo envío como retorno
	jsonUsers, err := os.ReadFile("../../users.json")
	if err != nil {
		fmt.Print(err)
	}
	var users []usuarios
	err = json.Unmarshal(jsonUsers, &users)
	if err != nil {
		fmt.Print(err)
	}
	return users
}

//Este handler se encargará de responder a /.
func PaginaInicio(ctx *gin.Context) {
	ctx.String(200, "¡Hola, Bienvenido a mi primer ejercicio con Query!")
}

//Este handler se encargará de responder a /usuarios
func AllUsers(ctx *gin.Context) {
	user := usersJson()
	ctx.JSON(200, gin.H{
		fmt.Sprintf("Todos los usuarios\n"): user,
	})
}

//Ejercicio 2
//Este handler verificará si la id que pasa el usuario existe en nuestra
// base de datos, de existir, retorna la info del usuario
func BuscarUsuariosId(ctx *gin.Context) {
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
	router := server.Group("/usuarios")
	{
		//Ejercicio 2
		router.GET("/:id", BuscarUsuariosId)
		//Ejercicio 1
		router.GET("", func(ctx *gin.Context) {
			route := ctx.Request.URL.Query()
			//Variable para aplicar query
			if route != nil {
				var usersU []usuarios
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
				var usersA []usuarios
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
				var usersE []usuarios
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

			// router.GET("/activo/:activo", BuscarUsuariosActivos)
		})
	}
	server.Run(":8081")
}
