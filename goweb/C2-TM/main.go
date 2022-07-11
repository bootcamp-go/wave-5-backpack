package main

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Users struct {
	Usuarios []Usuario `json:"usuarios"`
}

type Usuario struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre" binding:"required"`
	Apellido      string `json:"apellido" binding:"required"`
	Email         string `json:"email" binding:"required"`
	Edad          int    `json:"edad" binding:"required"`
	Altura        int    `json:"altura" binding:"required"`
	Activo        bool   `json:"activo" binding:"required"`
	FechaCreacion string `json:"fecha_creacion" binding:"required"`
}

var usuarios Users

func main() {
	router := gin.Default()
	pr := router.Group("/usuarios")
	pr.POST("/", Guardar())
	fmt.Println(usuarios.Usuarios)

	pr.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": usuarios.Usuarios,
		})
	})

	router.Run()
}

func Guardar() gin.HandlerFunc {
	var lastID int
	if len(usuarios.Usuarios) > 0 {
		lastID = usuarios.Usuarios[len(usuarios.Usuarios)-1].ID
	} else {
		lastID = 0
	}
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "123456" {
			c.JSON(401, gin.H{
				"error": "token invalido",
			})
		} else {
			var req Usuario
			var result string
			if err := c.ShouldBindJSON(&req); err != nil {
				var ve validator.ValidationErrors
				if errors.As(err, &ve) {
					for i, field := range ve {
						if i != len(ve)-1 {
							result += fmt.Sprintf("El campo %s es requerido y", field.Field())
						} else {
							result += fmt.Sprintf("El campo %s es requerido", field.Field())
						}
					}
				}
				c.JSON(400, result)
				return
			}
			lastID++
			req.ID = lastID
			usuarios.Usuarios = append(usuarios.Usuarios, req)
			c.JSON(200, req)
		}
	}
}
