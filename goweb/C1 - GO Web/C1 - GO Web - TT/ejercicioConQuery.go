package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	var exist bool
	if err != nil {
		ctx.String(400, "El id ingresado no existe")
		return
	}
	user := usersJson()
	for _, valueUser := range user {
		if valueUser.Id == id {
			exist = true
			ctx.JSON(200, gin.H{
				"Usuario encontrado": valueUser,
			})
			return
		}
	}
	if exist == false {
		ctx.String(400, "El id ingresado no existe")
		return
	}
}

//Filtro por todos los campos con Query
func FiltrarUsuarios(c *gin.Context) {
	listaUsuarios := usersJson()
	var u usuarios
	if c.ShouldBindQuery(&u) == nil {
		// Setea las variables del c.Query("nombrevar")
		log.Println(u.Id)
		log.Println(u.Nombre)
		log.Println(u.Apellido)
		log.Println(u.Email)
		log.Println(u.Edad)
		log.Println(u.Altura)
		log.Println(u.Activo)
		log.Println(u.Fecha)
	}

	var filtradoUsuarios []*usuarios
	for _, us := range listaUsuarios {
		// filtrado por todos los campos
		if u.Id == us.Id {
			filtradoUsuarios = append(filtradoUsuarios, us)
		}
		if u.Nombre == us.Nombre {
			filtradoUsuarios = append(filtradoUsuarios, us)
		}
		if u.Apellido == us.Apellido {
			filtradoUsuarios = append(filtradoUsuarios, us)
		}
		if u.Email == us.Email {
			filtradoUsuarios = append(filtradoUsuarios, us)
		}
		if u.Edad == us.Edad {
			filtradoUsuarios = append(filtradoUsuarios, us)
		}
		if u.Altura == us.Altura {
			filtradoUsuarios = append(filtradoUsuarios, us)
		}
		if u.Activo == us.Activo {
			filtradoUsuarios = append(filtradoUsuarios, us)
		}
		if u.Fecha == us.Fecha {
			filtradoUsuarios = append(filtradoUsuarios, us)
		}
	}

	c.JSON(http.StatusOK, filtradoUsuarios)
	// devovemos el array filtrado
}

func usersJson() []*usuarios {
	//Leo el json y lo envío como retorno
	jsonUsers, err := os.ReadFile("./users.json")
	if err != nil {
		fmt.Print(err)
	}
	var users []*usuarios
	err = json.Unmarshal(jsonUsers, &users)
	if err != nil {
		fmt.Print(err)
	}
	return users
}

func main() {
	server := gin.Default()
	server.GET("/", PaginaInicio)
	router := server.Group("/usuarios")
	router.GET("/", AllUsers)
	//Ejercicio 2
	router.GET("/:id", BuscarUsuariosId)
	//Ejercicio 1
	//Filtro de todos los campos con Query
	router.GET("/filtrarusuario", FiltrarUsuarios)

	server.Run()
}
