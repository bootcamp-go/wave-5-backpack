package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/users"
	"github.com/gin-gonic/gin"
)

// Estructura de rerquest
type request struct {
	Nombre   string  `json:"nombre"`
	Apellido string  `json:"apellido"`
	Email    string  `json:"email"`
	Edad     int     `json:"edad"`
	Altura   float64 `json:"altura"`
	Activo   bool    `json:"activo"`
}

type User struct {
	service users.Service
}

func NewUser(s users.Service) *User {
	return &User{
		service: s,
	}
}

func (u *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}

		users, err := u.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, users)
	}
}

func (u *User) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		user, errUser := u.service.GetById(int(id))
		if errUser != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": errUser.Error()})
			return
		}

		ctx.JSON(http.StatusOK, user)
	}
}

func (u *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}

		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		user, err := u.service.Store(req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, user)
	}
}

func (u *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}

		// Se valida el param id
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if req.Nombre == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "El nombre es requerido"})
			return
		}
		if req.Apellido == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "El apellido es requerido"})
			return
		}
		if req.Email == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "El email es requerido"})
			return
		}
		if req.Edad == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "La edad es requerida"})
			return
		}
		if req.Altura == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "La altura es requerida"})
			return
		}
		if req.Activo {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "El activo es requerido"})
			return
		}

		user, err := u.service.Update(int(id), req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, user)
	}
}

func (u *User) UpdateApellidoEdad() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if req.Apellido == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "El apellido es requerido"})
			return
		}
		if req.Edad == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "La edad es requerida"})
			return
		}

		user, err := u.service.UpdateApellidoEdad(int(id), req.Apellido, req.Edad)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, user)
	}
}

func (u *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		err = u.service.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("el usuario %d ha sido eliminado", id)})
	}
}
