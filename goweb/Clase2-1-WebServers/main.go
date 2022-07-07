package main

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type usuario struct {
	Id              int    `json:"id"`
	Nombre          string `json:"nombre"binding:"required"`
	Apellido        string `json:"apellido"binding:"required"`
	Email           string `json:"email"binding:"required"`
	Edad            int    `json:"edad"binding:"required"`
	Altura          int    `json:"altura"binding:"required"`
	Activo          bool   `json:"activo"binding:"required"`
	FechaDeCreacion string `json:"fecha_de_creacion"binding:"required"`
}

func Guardar() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := "contraseña123"
		if c.GetHeader("token") != token {
			c.JSON(401, "error: no tiene permisos para realizar la peticion solicitada")
			return
		}
		var user usuario
		if err := c.ShouldBindJSON(&user); err != nil {
			mensajeError := ""
			var ve validator.ValidationErrors
			if errors.As(err, &ve) {

				for _, field := range ve {
					mensajeError += fmt.Sprintf("el campo %s es requerido\n", field.Field())
				}
			}
			fmt.Println(user)
			c.JSON(400, gin.H{
				"error": mensajeError,
			})
			return
			//}

		}
		lastID++
		user.Id = lastID
		fmt.Println(user)
		users = append(users, user)
		c.JSON(200, users)
	}
}

var users []usuario
var lastID int

func main() {

	router := gin.Default()
	router.POST("/usuarios", Guardar())
	router.Run(":8080")
}

/*


 */
