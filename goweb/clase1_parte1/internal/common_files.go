package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type producto struct {
	Id        int    `form:"id" json:"id"`
	Nombre    string `form:"nombre" json:"nombre" binding:"required"`
	Color     string `form:"color" json:"color" binding:"required"`
	Precio    int    `form:"precio" json:"precio" binding:"required"`
	Stock     int    `form:"stock" json:"stock" binding:"required"`
	Codigo    string `form:"codigo" json:"codigo" binding:"required"`
	Publicado bool   `form:"publicado" json:"publicado"`
	Fecha     string `form:"fecha" json:"fecha" binding:"required"`
}

func getJson() []producto {
	jsonFile, err := os.Open("./products.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var productos []producto
	json.Unmarshal([]byte(byteValue), &productos)
	// ctx.JSON(200, productos)
	defer jsonFile.Close()
	return productos
}
