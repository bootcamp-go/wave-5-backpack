package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Empleado struct {
	// Una etiqueta de struct se cierra con caracteres de acento grave `
	Nombre   string `form:"name" json:"name"`
	Password string `form:"password" json:"password"`
	Id       string `form:"id" json:"id"`
	Activo   string `form:"active" json:"activa" binding:"required"`
}

func AutorizarEmpleado(ctx *gin.Context) {

	var empleado Empleado
	// el metodo ShouldBindJSON de nuestro context, asociará el contenido del body
	// a los campos de la estructura que le proveamos
	if err := ctx.ShouldBindJSON(&empleado); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if empleado.Nombre != "user1" || empleado.Password != "123" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "no esta autorizado"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "autorizado"})
}
