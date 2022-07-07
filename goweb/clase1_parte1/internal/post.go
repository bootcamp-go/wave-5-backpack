package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

var productos []producto

// post

/* func Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req producto

		token := ctx.GetHeader("token")

		if err := validateToken(token); err != nil {
			ctx.JSON(401, gin.H{"error: ": err.Error()})
			return
		}

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := validateFields(req); err != nil {
			ctx.JSON(400, gin.H{"error: ": err.Error()})
			return
		}
		req = saveProduct(req)
		ctx.JSON(200, req)
	}
} */

func Post(ctx *gin.Context) {
	var req producto

	token := ctx.GetHeader("token")

	if err := validateToken(token); err != nil {
		ctx.JSON(401, gin.H{"error: ": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := validateFields(req); err != nil {
		ctx.JSON(400, gin.H{"error: ": err.Error()})
		return
	}
	req = saveProduct(req)
	ctx.JSON(200, req)
}

func saveProduct(req producto) producto {
	productos = getJson()
	req.Id = productos[len(productos)-1].Id + 1
	productos = append(productos, req)
	jsonFile, err := json.Marshal(productos)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("products.json", jsonFile, 0644)
	if err != nil {
		fmt.Println(err)
	}
	return req
}

func validateFields(req producto) error {
	if req.Nombre == "" {
		return errors.New("el campo nombre es requerido")
	}
	if req.Color == "" {
		return errors.New("el campo color es requerido")
	}
	if req.Precio == 0 {
		return errors.New("el campo precio es requerido")
	}
	if req.Stock == 0 {
		return errors.New("el campo stock es requerido")
	}
	if req.Codigo == "" {
		return errors.New("el campo codigo es requerido")
	}
	if req.Fecha == "" {
		return errors.New("el campo fecha es requerido")
	}
	return nil
}

func validateToken(token string) error {
	if token == "" {
		return errors.New("no ingresó el token y es requerido")
	}
	if token != "123456" {
		return errors.New("no tiene permisos para realizar la petición solicitada")
	}
	return nil
}
