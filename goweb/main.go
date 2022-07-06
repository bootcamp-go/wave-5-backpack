package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	Id          int
	Names       string  `json:nombre`
	LastName    string  `json:apellido`
	Age         int     `json:edad`
	DateCreated string  `json:fechaCreacion`
	Estatura    float64 `json:altura`
	Email       string  `json:email`
	IsActivo    bool    `json:activo`
}

func getAll() ([]byte, error) {
	u := Usuario{
		Id:          1,
		Names:       "Andrea",
		LastName:    "Esquivel",
		Email:       "ing.andreaesquivel@gmail.com",
		Estatura:    1.56,
		Age:         23,
		IsActivo:    true,
		DateCreated: "06/07/2022",
	}

	usuarios := []Usuario{}
	usuarios = append(usuarios, u)
	usuarios = append(usuarios, u)

	jsonData, err := json.Marshal(usuarios)
	if err != nil {
		return nil, err
	}

	return jsonData, nil

}
func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hola, Andy!",
		})
	})

	router.GET("/usuarios", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hola, Andy!",
		})

		usuariosJSON, _ := getAll()
		/*if err != nil {
			log.Fatal(err)
		}*/
		c.JSON(http.StatusOK, gin.H{
			"usuarios": string(usuariosJSON),
		})
	})

	router.Run(":8001")
}
