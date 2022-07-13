package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/users"
	"github.com/bootcamp-go/wave-5-backpack/goweb/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Estructura de rerquest
type request struct {
	Nombre   string  `json:"nombre" binding:"required"`
	Apellido string  `json:"apellido" binding:"required"`
	Email    string  `json:"email" binding:"required,email"`
	Edad     int     `json:"edad"`
	Altura   float64 `json:"altura"`
}

// Estructura de rerquest
type putrequest struct {
	Nombre   string  `json:"nombre" binding:"required"`
	Apellido string  `json:"apellido" binding:"required"`
	Email    string  `json:"email" binding:"required,email"`
	Edad     int     `json:"edad" binding:"required"`
	Altura   float64 `json:"altura" binding:"required"`
}

// Estructura de patch rerquest
type patchrequest struct {
	Apellido string `json:"apellido"`
	Edad     int    `json:"edad"`
}

// Estructura para los errores
type ErroresMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type User struct {
	service users.Service
}

func NewUser(s users.Service) *User {
	return &User{
		service: s,
	}
}

// GetAll godoc
// @Summary Obtiene todos los usuarios
// @Tags Usuarios
// @Description Obtiene todos los usuarios que no estan borrados
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /users [get]
func (u *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, err := u.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, users, ""))
	}
}

// GetById godoc
// @Summary Obtieneun usuario por id
// @Tags Usuarios
// @Description Obtiene un usuario por su id que no este borrado
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path int true "id"
// @Success 200 {object} web.Response
// @Router /users/{id} [get]
func (u *User) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "invalid id"))
			return
		}

		user, errUser := u.service.GetById(int(id))
		if errUser != nil {
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, errUser.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, user, ""))
	}
}

// StoreUsers godoc
// @Summary Guarda un usuario
// @Tags Usuarios
// @Description Guarda un usuario
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param user body request true "User to store"
// @Success 200 {object} web.Response
// @Router /users [post]
func (u *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			var errores []ErroresMsg
			for _, fieldError := range err.(validator.ValidationErrors) {
				errores = append(errores, ErroresMsg{fieldError.Field(), getErrorMsg(fieldError)})
			}
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, errores))
			return
		}

		user, err := u.service.Store(req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, user, ""))
	}
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "El campo es requerido"
	case "email":
		return "El email no es válido"
	}
	return "Error desconocido"
}

// UpdateUser godoc
// @Summary Actualiza un usuario
// @Tags Usuarios
// @Description Actualiza toda la información de un usuario
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param user body putrequest true "User to update"
// @Success 200 {object} web.Response
// @Router /users [put]
func (u *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Se valida el param id
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "invalid id"))
			return
		}

		var req putrequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			var errores []ErroresMsg
			for _, fieldError := range err.(validator.ValidationErrors) {
				errores = append(errores, ErroresMsg{fieldError.Field(), getErrorMsg(fieldError)})
			}
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, errores))
			return
		}

		user, err := u.service.Update(int(id), req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, user, ""))
	}
}

// UpdateApellidoEdad godoc
// @Summary Actualiza parcialmente un usuario
// @Tags Usuarios
// @Description Actualiza el apellido y edad de un usuario
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path int true "Update user by id"
// @Param data body patchrequest true "User to update"
// @Success 200 {object} web.Response
// @Router /users/{id} [patch]
func (u *User) UpdateApellidoEdad() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "invalid id"))
			return
		}

		var req patchrequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		if (req.Apellido == "" || req.Apellido == "string") && req.Edad == 0 {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "No hay información para actualización"))
			return
		}

		user, err := u.service.UpdateApellidoEdad(int(id), req.Apellido, req.Edad)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, user, "que"))
	}
}

// DeleteUsers godoc
// @Summary Borra un usuario
// @Tags Usuarios
// @Description Borra un usuario por su id
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path int true "Delete user by id"
// @Success 200 {object} web.Response
// @Router /users/{id} [delete]
func (u *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "invalid id"))
			return
		}

		err = u.service.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, fmt.Sprintf("el usuario %d ha sido eliminado", id), ""))
	}
}

// SearchUsers godoc
// @Summary Busca usuarios
// @Tags Usuarios
// @Description Busca usuarios por algun criterio de búsqueda
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param nombre query string false "Nombre"
// @Param apellido query string false "Apellido"
// @Param email query string false "Email"
// @Param edad query string false "Edad"
// @Param altura query string false "Altura"
// @Param activo query string false "Activo"
// @Param fecha_creacion query string false "Fecha de creación"
// @Success 200 {object} web.Response
// @Router /users/search [get]
func (u *User) SearchUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Se obtienen los query params
		nombreQuery := ctx.Query("nombre")
		apellidoQuery := ctx.Query("apellido")
		emailQuery := ctx.Query("email")
		edadQuery := ctx.Query("edad")
		alturaQuery := ctx.Query("altura")
		activoQuery := ctx.Query("activo")
		fechaCreacionQuery := ctx.Query("fecha_creacion")

		if nombreQuery == "" && apellidoQuery == "" && emailQuery == "" && edadQuery == "" && alturaQuery == "" && activoQuery == "" && fechaCreacionQuery == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "debe especificar algún criterio de búsqueda"))
			return
		}

		users, err := u.service.SearchUser(nombreQuery, apellidoQuery, emailQuery, edadQuery, alturaQuery, activoQuery, fechaCreacionQuery)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, users, ""))
	}
}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	if requiredToken == "" {
		log.Fatal("no esta definido el token de seguridad")
	}

	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token requerido"))
			return
		}

		if token != requiredToken {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token incorrecto"))
			return
		}

		ctx.Next()
	}
}
