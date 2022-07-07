package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	ID     int     `json:"id"`
	Nombre string  `json:"nombre" binding:"required"`
	Altura float64 `json:"altura" binding:"required"`
}

var usuarios []interface{}
var LastId int

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	usuarios := router.Group("/usuarios")
	{
		usuarios.GET("", ListarUsuarios)
		usuarios.POST("", GuardarUsuario())
	}

	router.Run()
}

func ListarUsuarios(ctx *gin.Context) {
	if ctx.GetHeader("token") != "123" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "No tiene permisos para realizar la peticion solicitada",
		})
		return
	}

	if usuarios == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "No se encontraron usuarios",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"usuarios": usuarios,
		})
	}
}

func GuardarUsuario() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetHeader("token") != "123" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "No tiene permisos para realizar la peticion solicitada",
			})
			return
		}

		var req Request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			var ve validator.ValidationErrors
			var notBinded string
			if errors.As(err, &ve) {
				notBinded += "Los siguientes campos son requeridos: "
				for _, fe := range ve {
					notBinded += fmt.Sprintf("%s ", fe.Field())
				}
			}

			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": notBinded,
			})
			return
		}

		LastId++
		req.ID = LastId
		usuarios = append(usuarios, req)

		ctx.JSON(http.StatusOK, gin.H{
			"usuarios": usuarios,
		})
	}
}
