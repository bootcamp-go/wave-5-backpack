package main

import (
	"encoding/json"
	"log"
	"os"

	"strconv"

	"github.com/gin-gonic/gin"
)

var usuarios Usuarios

type Usuarios struct {
	Users []Usuario `json:"usuarios"`
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

func main() {
	router := gin.Default()
	LeerJson()

	userGroup := router.Group("/usuarios")
	{
		userGroup.GET("/", GetAll)
		userGroup.GET("/:id", idHandler)
		userGroup.GET("/search", searchHandler)
	}

	router.Run()
}

func LeerJson() {
	jsonData, _ := os.ReadFile("usuarios.json")

	// fmt.Println(jsonData)

	if err := json.Unmarshal([]byte(jsonData), &usuarios); err != nil {
		log.Fatal(err)
	}

	// fmt.Println(u)
}

func GetAll(c *gin.Context) {
	c.JSON(200, gin.H{"message": usuarios})
}

func idHandler(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(422, gin.H{"error": "id must be int"})
		return
	} else {
		for _, val := range usuarios.Users {
			if id == val.ID {
				c.JSON(200, gin.H{"message": val})
			}
		}
	}
}

func searchHandler(c *gin.Context) {
	var result []Usuario
	for _, user := range usuarios.Users {
		if c.Query("nombre") == user.Nombre ||
			c.Query("apellido") == user.Apellido ||
			c.Query("email") == user.Email ||
			c.Query("fecha_creacion") == user.FechaCreacion {
			result = append(result, user)
		}

		if c.Query("edad") != "" {
			if edad, err := strconv.Atoi(c.Query("edad")); err != nil {
				c.JSON(422, gin.H{"error": "edad must be int"})
				return
			} else if edad == user.Edad {
				result = append(result, user)
			}
		}

		if c.Query("altura") != "" {
			if altura, err := strconv.Atoi(c.Query("altura")); err != nil {
				c.JSON(422, gin.H{"error": "altura must be int"})
				return
			} else if altura == user.Altura {
				result = append(result, user)
			}
		}

		if c.Query("activo") != "" {
			if activo, err := strconv.ParseBool(c.Query("activo")); err != nil {
				c.JSON(422, gin.H{"error": "activo must be bool"})
				return
			} else if activo == user.Activo {
				result = append(result, user)
			}
		}
	}

	c.JSON(200, gin.H{"message": result})
}
