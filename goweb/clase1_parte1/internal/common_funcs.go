package internal

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// estructura a recibir en la URL
type user struct {
	Name      string `form:"name"`
	Last_name string `form:"last_name"`
}

// estructura escogida para el ejercicio
type producto struct {
	Id            int    `form:"id" json:"id"`
	Nombre        string `form:"nombre" json:"nombre" binding:"required"`
	Color         string `form:"color" json:"color" binding:"required"`
	Precio        int    `form:"precio" json:"precio" binding:"required"`
	Stock         int    `form:"stock" json:"stock" binding:"required"`
	Codigo        string `form:"codigo" json:"codigo" binding:"required"`
	Publicado     bool   `form:"publicado" json:"publicado"`
	FechaCreacion string `form:"fecha_creacion" json:"fecha_creacion" binding:"required"`
}

// Obtiene la lista de productos del archivo products.json
func readJSON() []producto {
	jsonFile, err := os.Open("products.json")
	if err != nil {
		panic(err)
	}

	var productos []producto
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &productos)
	jsonFile.Close()
	return productos
}

// Guarda una lista de productos en products.json
func writeJSON(product_list []producto) error {
	file, err := json.Marshal(product_list)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("products.json", file, 0644)
	if err != nil {
		return err
	}
	return nil
}
