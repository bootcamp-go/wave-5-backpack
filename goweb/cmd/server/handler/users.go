package handler

import (
	"fmt"
	"goweb/internal/users"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	Nombre        string  `json:"nombre" binding:"required"`
	Apellido      string  `json:"apellido" binding:"required"`
	Email         string  `json:"email" binding:"required,email"`
	Edad          int     `json:"edad" binding:"required"`
	Altura        float64 `json:"altura" binding:"required"`
	Activo        bool    `json:"activo" binding:"required"`
	FechaCreacion string  `json:"fechaCreacion" binding:"required"`
}

type User struct {
	service users.Service
}

func NewProduct(s users.Service) *User {
	return &User{service: s}
}

// ==================================
// Funciones de clases anteriores
// ==================================

func (u *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, err := u.service.GetAll()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		if len(users) == 0 {
			ctx.JSON(500, gin.H{"error": "No se hallaron resultados"})
			return
		}

		ctx.JSON(200, users)
	}
}

// Clase 1 Ejercicio 2 Parte 2
func (u *User) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idInt, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
		}

		user, err := u.service.GetById(idInt)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}

func validateToken(ctx *gin.Context) bool {
	if token := ctx.GetHeader("token"); token != "62c5b68a0cc23a33375c85f8" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "NO tiene permisos para realizar la petición solicitada"})
		return false
	}
	return true
}

func ValidateErrors(campo string, v validator.FieldError) string {
	switch v.Tag() {
	case "required":
		return "El campo " + campo + " es requerido"
	case "email":
		return "Direccion de correo electronico invalida"
	}
	return "Error desconoodido..."
}

func (u *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req Request

		if !validateToken(ctx) {
			return
		}

		if err := ctx.ShouldBindJSON(&req); err != nil {
			for _, fieldError := range err.(validator.ValidationErrors) {
				ctx.JSON(400, ValidateErrors(fieldError.Field(), fieldError))
			}
			return
		}

		user, err := u.service.Store(
			req.Nombre,
			req.Apellido,
			req.Email,
			req.Edad,
			req.Altura,
			req.Activo,
			req.FechaCreacion,
		)

		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, user)
	}
}

func (u *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error: ": "Ingrese un ID valido"})
			return
		}

		var req Request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			for _, fieldError := range err.(validator.ValidationErrors) {
				ctx.JSON(400, ValidateErrors(fieldError.Field(), fieldError))
			}
			return
		}

		user, err := u.service.Update(int(id), req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.FechaCreacion)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, user)
	}
}

func (u *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error: ": "Ingrese un ID valido"})
			return
		}

		err = u.service.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error: ": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, fmt.Sprintf("El usuario con el ID %d se eliminó correctamente", id))
	}
}
