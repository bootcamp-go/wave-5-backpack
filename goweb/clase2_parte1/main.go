package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"log"
	"encoding/json"
)

type request struct {
	Id				int		`json:"Id"`
	Nombre			string	`json:"nombre" binding:"required"`
	Apellido		string	`json:"apellido" binding:"required"`
	Email			string	`json:"email" binding:"required"`
	Edad			int		`json:"edad" binding:"required"`
	Altura			float64	`json:"altura" binding:"required"`
	Activo			*bool	`json:"activo" binding:"required"`
	FechaCreacion	string	`json:"fecha_creacion" binding:"required"`
}

var usuarios []request

func readJSON() {
	jsonUsers, err := os.ReadFile("./users.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(jsonUsers), &usuarios); err != nil {
		log.Fatal(err)
	}
}

func autoIncrementarID() int {
	var lastID int
	if len(usuarios) > 0 {
		lastID = usuarios[len(usuarios)-1].Id
	}
	lastID++
	return lastID
}

func ObtenerUsuarios() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, usuarios)
	}
}

func GuardarUsuario() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "123456" {
			c.JSON(401, gin.H{ "error": "No tiene permisos para realizar la petición solicitada"})
			return
		}

		var req request

		if err := c.ShouldBindJSON(&req); err != nil {

			var error_messages []string
			error_messages = append(error_messages, err.Error())

			if req.Nombre == "" {
				error_messages = append(error_messages, "El campo nombre es requerido.")
			}

			if req.Apellido == "" {
				error_messages = append(error_messages, "El campo apellido es requerido.")
			}

			if req.Email == "" {
				error_messages = append(error_messages, "El campo email es requerido.")
			}

			if req.Edad  == 0 {
				error_messages = append(error_messages, "El campo edad es requerido.")
			}

			if req.Altura == 0 {
				error_messages = append(error_messages, "El campo altura es requerido.")
			}

			if req.Activo == nil {
				error_messages = append(error_messages, "El campo activo es requerido.")
			}

			if req.FechaCreacion == "" {
				error_messages = append(error_messages, "El campo fecha de creación es requerido.")
			}

			c.JSON(400, error_messages)
			return
		}

		req.Id = autoIncrementarID()
		usuarios = append(usuarios, req)
		c.JSON(200, req)
	}
}

func main() {
	readJSON()
	r := gin.Default()
	ur :=  r.Group("/usuarios")
	ur.GET("/", ObtenerUsuarios())
	ur.POST("/guardar", GuardarUsuario())
	r.Run()
}