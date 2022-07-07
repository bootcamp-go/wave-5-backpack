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
		fmt.Sprintf("Todos los usuarios"): user,
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

//Ejercicio 1
//Para resolver este ejercicio hice dos archivos
//este sin utilizar query y otro con el uso de las mismas
func BuscarUsuariosNombre(ctx *gin.Context) {
	var user_same_name []usuarios
	nombre := ctx.Param("nombre")
	user := usersJson()
	if nombre != "" {
		for _, valueUser := range user {
			//Si el nombre de algún usuario coincide
			if valueUser.Nombre == nombre {
				user_same_name = append(user_same_name, valueUser)
			}
			fmt.Println(valueUser)
		}
	}
	if len(user_same_name) > 0 {
		ctx.JSON(200, gin.H{
			fmt.Sprintf("Usuarios con nombre %s", nombre): user_same_name,
		})
	}
}

//Buscando usuarios por apellido Lucumí
func BuscarUsuariosApellido(ctx *gin.Context) {
	var user_same_last []usuarios
	apellido := ctx.Param("apellido")
	user := usersJson()
	if apellido != "" {
		for _, valueUser := range user {
			//Si el apellido de algún usuario coincide
			if valueUser.Apellido == apellido {
				user_same_last = append(user_same_last, valueUser)
			}
		}
	}
	if len(user_same_last) > 0 {
		ctx.JSON(200, gin.H{
			fmt.Sprintf("Usuarios con apellido %s", apellido): user_same_last,
		})
	}
}

//usuarios con email
func BuscarUsuariosEmail(ctx *gin.Context) {
	var user_same_email []usuarios
	email := ctx.Param("email")
	user := usersJson()
	if email != "" {
		for _, valueUser := range user {
			//Si el email de algún usuario coincide
			if valueUser.Email == email {
				user_same_email = append(user_same_email, valueUser)
			}
		}
	}
	if len(user_same_email) > 0 {
		ctx.JSON(200, gin.H{
			fmt.Sprintf("Usuarios con email %s", email): user_same_email,
		})
	}
}

//Usuarios edad
func BuscarUsuariosEdad(ctx *gin.Context) {
	var user_same_edad []usuarios
	edad, err := strconv.Atoi(ctx.Param("edad"))
	if err != nil {
		ctx.String(400, "la edad ingresada no existe")
		return
	}
	user := usersJson()
	if edad != 0 {
		for _, valueUser := range user {
			//Si la edad de algún usuario coincide
			if valueUser.Edad == edad {
				user_same_edad = append(user_same_edad, valueUser)
			}
		}
	}
	if len(user_same_edad) > 0 {
		ctx.JSON(200, gin.H{
			fmt.Sprintf("Usuarios con edad %d", edad): user_same_edad,
		})
	}
}

//Usuarios altura
func BuscarUsuariosAltura(ctx *gin.Context) {
	var user_same_altura []usuarios
	altura, err := strconv.ParseFloat((ctx.Param("altura")), 64)
	if err != nil {
		ctx.String(400, "La altura ingresado no existe")
		return
	}
	user := usersJson()
	if altura != 0.0 {
		for _, valueUser := range user {
			//Si la edad de algún usuario coincide
			if valueUser.Altura == altura {
				user_same_altura = append(user_same_altura, valueUser)
			}
		}
	}
	if len(user_same_altura) > 0 {
		ctx.JSON(200, gin.H{
			fmt.Sprintf("Usuarios con altura %f", altura): user_same_altura,
		})
	}
}

func main() {
	server := gin.Default()
	server.GET("/", PaginaInicio)
	router := server.Group("/usuarios")
	{
		router.GET("", AllUsers)
		//Ejercicio 2
		router.GET("/:id", BuscarUsuariosId)
		//Ejercicio 1
		router.GET("/name/:nombre", BuscarUsuariosNombre)
		router.GET("/last/:apellido", BuscarUsuariosApellido)
		router.GET("/email/:email", BuscarUsuariosEmail)
		router.GET("/edad/:edad", BuscarUsuariosEdad)
		router.GET("/altura/:altura", BuscarUsuariosAltura)
		// router.GET("/activo/:activo", BuscarUsuariosActivos)
	}
	server.Run(":8081")
}
