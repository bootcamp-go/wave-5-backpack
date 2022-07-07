package internal

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Variable global, lista de productos registrados
var product_list []producto

// Ruta para enviar petición post
func ProductosPost(router *gin.Engine) {
	router.POST("/productos", newProduct)
}

// Handler para recibir un nuevo producto y agregarlo
func newProduct(c *gin.Context) {

	if err := checkToken(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var p producto
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := checkFields(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Errorf("el campo %s es requerido", err),
		})
	}

	p = saveJSON(p)
	c.JSON(http.StatusOK, p)
}

// Añade un producto a la lista y lo guarda de nuevo en products.json
func saveJSON(p producto) producto {
	product_list = readJSON()
	lastID := product_list[len(product_list)-1].Id
	p.Id = lastID + 1
	product_list = append(product_list, p)

	if err := writeJSON(product_list); err != nil {
		fmt.Println("Falló guardando el archivo")
	}
	return p
}

// Revisa que se hayan suministrado todos los campos requeridos
func checkFields(p producto) error {
	if p.Nombre == "" {
		return errors.New("Nombre")
	}
	if p.Color == "" {
		return errors.New("Color")
	}
	if p.Precio == 0 {
		return errors.New("Precio")
	}
	if p.Stock == 0 {
		return errors.New("Stock")
	}
	if p.Codigo == "" {
		return errors.New("Codigo")
	}
	/* if !p.Publicado {
		return errors.New("Publicado")
	} */
	if p.FechaCreacion == "" {
		return errors.New("FechaCreacion")
	}
	return nil
}

// Revisa que el token haya sido enviado y sea válido
func checkToken(c *gin.Context) error {
	token := c.GetHeader("token")

	if token == "" {
		fmt.Println("entra vacio")
		return errors.New("el token no fue suministrado")
	}

	if token != "12345" {
		fmt.Println("entra invalido")
		return errors.New("no tiene permisos para realizar la petición solicitada")
	}
	return nil
}
