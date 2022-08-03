package handler

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"C3-Testing/internal/servicio"
	"C3-Testing/pkg/web"

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
		c.JSON(401, web.NewResponse(401, nil, "No tiene permisos para realizar la peticion solicitada"))
		return false
	}
	return true
}

// ListUsers godoc
// @Summary      List users
// @Tags         Users
// @Description  get all users
// @Accept       json
// @Produce      json
// @Param        token  header    string  true  "token"
// @Success      200    {object}  interface{}
//@Failure 404  {object} interface{}
//@Failure 500 {object} interface{}
// @Router       /users [get]
func (s *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}
		allUsers, err := s.service.GetAll()
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}
		if len(allUsers) == 0 {
			ctx.JSON(404, web.NewResponse(404, nil, "No existen usuarios registrados"))
		}
		ctx.JSON(200, web.NewResponse(200, allUsers, ""))
	}
}

// CreateUser godoc
// @Summary      Create user
// @Tags         Users
// @Description  Create user
// @Accept       json
// @Produce      json
// @Param        token    header    string   true  "token"
// @Param        user  body  User{}    true  "User to add"
// @Success      200      {object}  User{}
// @Failure      404      {object}  string
// @Router       /users [post]
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
				ctx.JSON(404, web.NewResponse(404, nil, result))
				return
			}
		}
		fmt.Println(user)
		u, err := s.service.Store(user.FirstName, user.LastName, user.Email, user.Age, user.Height, user.Activo, user.CreatedAt)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}

// UpdateUser godoc
// @Summary      Update user
// @Tags         Users
// @Description  Update user
// @Accept       json
// @Produce      json
// @Param        token    header    string   true  "token"
//@Param id path    string  true    "User ID"
// @Param        product  body      User{}  true  "User to add"
// @Success      200      {object}  User{}
// @Failure      404      {object}  string
// @Router       /users/{id} [put]
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
				ctx.JSON(404, web.NewResponse(404, nil, result))
				return
			}
		}

		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			log.Fatal("Error al convertir a entero", err)
		}

		u, err := s.service.Update(id, user.FirstName, user.LastName, user.Email, user.Age, user.Height, user.Activo, user.CreatedAt)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}

// UpdateLastNameAge godoc
// @Summary      Update Last Name User
// @Tags         Products
// @Description  Update Last Name User
// @Accept       json
// @Produce      json
// @Param        token    header    string   true  "token"
//@Param id path    string  true    "User ID"
// @Param        product  body      User{}  true  "Last Name and user to update"
// @Success      200      {object}  User{}
// @Failure      404      {object}  string
// @Router       /users/{id} [patch]
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
				ctx.JSON(404, web.NewResponse(404, nil, result))
				return
			}
		}

		idParam := ctx.Param("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			log.Fatal("Error al convertir a entero", err)
		}

		u, err := s.service.UpdateLastNameAge(id, user.LastName, user.Age)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}

// Delete godoc
// @Summary      Delete user
// @Tags         Users
// @Description  Delete user
// @Accept       json
// @Produce      json
// @Param        token  header    string  true  "token"
//@Param id path    string  true    "User ID"
// @Success      200    {object}  User{}
// @Failure      404    {object}  string
// @Router       /users/{id} [delete]
func (s *User) Delete() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}

		idParam := ctx.Param("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			log.Fatal("Error al convertir a entero", err)
		}
		fmt.Println(idParam)
		err = s.service.Delete(id)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, nil, "usuario eliminado exitosamente"))
	}
}
