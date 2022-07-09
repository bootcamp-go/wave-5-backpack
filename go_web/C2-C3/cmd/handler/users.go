package handler

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/rodrigoeshard/goweb/Practica2.2/internal/servicio"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type user struct {
	FirstName string  `json:"firstName" binding:"required"`
	LastName  string  `json:"lastName" binding:"required"`
	Email     string  `json:"email" binding:"required"`
	Age       int     `json:"age" binding:"required"`
	Height    float64 `json:"height" binding:"required"`
	Activo    bool    `json:"activo" binding:"required"`
	CreatedAt string  `json:"createdAt" binding:"required"`
}

type UserPatch struct {
	LastName string `json:"lastName" binding:"required"`
	Age      int    `json:"age" binding:"required"`
}

type User struct {
	service servicio.Service
}

func NewUser(s servicio.Service) *User {
	return &User{
		service: s,
	}
}

func validateToken(c *gin.Context) bool {
	tokenv := os.Getenv("TOKEN")
	if token := c.GetHeader("token"); token != tokenv {
		c.JSON(401, gin.H{
			"error": "No tiene permisos para realizar la peticion solicitada",
		})
		return false
	}
	return true
}

func (s *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}
		allUsers, err := s.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, allUsers)
	}
}

func (s *User) CreateUser() gin.HandlerFunc {
	var user user
	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}
		if err := ctx.ShouldBindJSON(&user); err != nil {
			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				result := ""
				for i, field := range ve {
					if i != len(ve)-1 {
						result += fmt.Sprintf("El campo %s es requerido y ", field.ActualTag())
					} else {
						result += fmt.Sprintf("El campo %s es requerido", field.ActualTag())
					}
				}
				ctx.JSON(404, result)
				return
			}
		}
		u, err := s.service.Store(user.FirstName, user.LastName, user.Email, user.Age, user.Height, user.Activo, user.CreatedAt)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, u)
	}
}

func (s *User) UpdateUser() gin.HandlerFunc {
	var user user
	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}

		if err := ctx.ShouldBindJSON(&user); err != nil {
			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				result := ""
				for i, field := range ve {
					if i != len(ve)-1 {
						result += fmt.Sprintf("El campo %s es requerido y ", field.ActualTag())
					} else {
						result += fmt.Sprintf("El campo %s es requerido", field.ActualTag())
					}
				}
				ctx.JSON(404, result)
				return
			}
		}

		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			fmt.Errorf("Error al convertir a entero%v", err)
		}

		u, err := s.service.Update(id, user.FirstName, user.LastName, user.Email, user.Age, user.Height, user.Activo, user.CreatedAt)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, u)
	}
}
func (s *User) UpdateLastNameAge() gin.HandlerFunc {
	var user UserPatch
	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}

		if err := ctx.ShouldBindJSON(&user); err != nil {
			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				result := ""
				for i, field := range ve {
					if i != len(ve)-1 {
						result += fmt.Sprintf("El campo %s es requerido y ", field.ActualTag())
					} else {
						result += fmt.Sprintf("El campo %s es requerido", field.ActualTag())
					}
				}
				ctx.JSON(404, result)
				return
			}
		}

		idParam := ctx.Param("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			fmt.Errorf("Error al convertir a entero%v", err)
		}

		u, err := s.service.UpdateLastNameAge(id, user.LastName, user.Age)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, u)
	}
}
func (s *User) Delete() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}

		idParam := ctx.Param("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			fmt.Errorf("Error al convertir a entero%v", err)
		}

		err = s.service.Delete(id)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, nil)
	}
}
