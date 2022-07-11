package main

//FIRST EJERCICIO - CREAR ENTIDAD

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var products []Productos

type Productos struct {
	Id            int     `json:"id"`
	Nombre        string  `json:"nombre" binding:"required"`
	Color         string  `json:"color"  binding:"required"`
	Precio        float64 `json:"precio" binding:"required"`
	Stock         int     `json:"stock"  binding:"required"`
	Codigo        string  `json:"codigo"  binding:"required"`
	Publicado     bool    `json:"publicado"  binding:"required"`
	FechaCreación string  `json:"fecha_creacion" binding:"required"`
}

func validatePost(c *gin.Context) {
	var req Productos
	if err := c.ShouldBindJSON(&req); err != nil {
		if req.Id == 0 || req.Nombre == "" || req.Color == "" || req.Precio == 0 || req.Stock <= -1 || req.Codigo == "" || req.FechaCreación == "" {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			result := ""
			for i, field := range ve {
				if i != len(ve)-1 {
					result += fmt.Sprintf("El campo %s es requerido y ", field.Field())
				} else {
					result += fmt.Sprintf("El campo %s es requerido", field.Field())
				}
			}
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
	}
	if !tokenValidate(c) {
		return
	}
	req.Id = generateId(products)
	fmt.Println(req.Id)
	products = append(products, req)
	c.JSON(200, req)
}

//Funcion para validar Token
func tokenValidate(c *gin.Context) bool {
	if token := c.GetHeader("token"); token != "0101" {
		c.JSON(401, gin.H{
			"error": "No esta autorizado.",
		})
		return false
	}
	return true
}

//Funcion para autoincrementar el id.
func generateId(p []Productos) int {
	idSaves := []int{}
	if len(p) > 0 {
		for _, n := range p {
			idSaves = append(idSaves, n.Id)
		}
		numMay := idSaves[0]
		for _, num := range idSaves {
			if num > numMay {
				numMay = num
			}
		}
		return numMay + 1
	}
	return 1
}

func main() {
	router := gin.Default()
	pr := router.Group("/productos")
	pr.POST("/create", validatePost)
	router.Run()
}
