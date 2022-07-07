package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

//EJERCICIO 2
//Validando campos: Con binding obligo a escribir
//cada campo requerido
type usuarios struct {
	Id       int       `json:"id"`
	Nombre   string    `json:"nombre"binding:"required"`
	Apellido string    `json:"apellido"binding:"required"`
	Email    string    `json:"email"binding:"required"`
	Edad     int       `json:"edad"binding:"required"`
	Altura   float64   `json:"altura"binding:"required"`
	Activo   bool      `json:"activo"binding:"required"`
	Fecha    time.Time `json:"fecha"`
}

var users = usersJson()

//EJERCICIO 1
//Para que incremente a partir del id del
//último usuario
var incrementId = users[len(users)-1].Id

//Leo json de usuarios
func usersJson() []usuarios {
	//Leo el json y lo envío como retorno
	jsonUsers, err := os.ReadFile("./users.json")
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

func GuardarUsuario() gin.HandlerFunc {
	return func(c *gin.Context) {
		//EJERCICIO 3
		//Validando el token
		token := c.GetHeader("token")
		if token != "246810" {
			c.JSON(401, gin.H{
				"error": "no tiene permisos para realizar la petición solicitada",
			})
			return
		}
		var us usuarios
		if err := c.ShouldBindJSON(&us); err != nil {
			var validatorE validator.ValidationErrors
			if errors.As(err, &validatorE) {
				errorMessage := ""
				for i, fieldError := range validatorE {
					if i != len(validatorE)-1 {
						errorMessage += fmt.Sprintf("el campo %s es requerido y ", fieldError.Field())
					} else {
						errorMessage += fmt.Sprintf("el campo %s es requerido", fieldError.Field())
					}
				}
				c.JSON(404, gin.H{"errors": errorMessage})
			}
			return
		}
		incrementId++
		us.Id = incrementId
		us.Activo = true
		//Para añadir fecha actual cuando se crea
		//el nuevo usuario
		us.Fecha = time.Now()

		users = append(users, us)
		c.JSON(200, users)
	}
}

func main() {
	r := gin.Default()
	//Agrupo mediante path usuarios
	//Creo endpoint mediante POST
	pr := r.Group("/usuarios")
	pr.POST("/", GuardarUsuario())
	r.Run()
}
