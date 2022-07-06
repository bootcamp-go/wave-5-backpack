package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Usuarios struct {
	Id       int     `json:"id"`
	Nombre   string  `json:"nombre"`
	Apellido string  `json:"apellido"`
	Email    string  `json:"email"`
	Edad     int     `json:"edad"`
	Altura   float64 `json:"altura"`
	Activo   bool    `json:"activo"`
	Creacion string  `json:"creacion"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Saludo estandar
	greet := router.Group("/greet")
	{
		greet.GET("", Greet)
		greet.GET("/:nombre", GreetName)
	}

	users := router.Group("/usuarios")
	{
		users.GET("", ListUsers)
		users.GET("/:id", GetUser)
	}

	router.Run()
}

// Handlers
func Greet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hola mundo!",
	})
}

func GreetName(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Hola %s!", ctx.Param("nombre")),
	})
}

func ListUsers(ctx *gin.Context) {
	activo := ctx.Query("activo")

	usuarios := OpenJsonFile("1/tm/usuarios.json")

	var usuarios_filtrados []Usuarios
	if activo != "" {
		for _, e := range usuarios {
			if strconv.FormatBool(e.Activo) == activo {
				usuarios_filtrados = append(usuarios_filtrados, e)
			}
		}
	} else {
		usuarios_filtrados = usuarios
	}

	ctx.JSON(http.StatusOK, gin.H{
		"usuarios": usuarios_filtrados,
	})
}

func GetUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}

	usuarios := OpenJsonFile("1/tm/usuarios.json")

	for _, v := range usuarios {
		if v.Id == id {
			ctx.JSON(http.StatusOK, gin.H{
				"usuario": v,
			})
			return
		}
	}

	ctx.JSON(http.StatusNoContent, gin.H{
		"message": fmt.Sprintf("No se ha encontrado un usuario con id %d", id),
	})
}

func OpenJsonFile(path string) []Usuarios {
	jsonData, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var u []Usuarios
	err = json.Unmarshal(jsonData, &u)
	if err != nil {
		panic(err)
	}

	return u
}
