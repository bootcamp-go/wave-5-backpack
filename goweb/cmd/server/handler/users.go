package handler

import (
	"fmt"
	"goweb/internal/users"
	"goweb/pkg/web"
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
type RequestPatch struct {
	Apellido string `json:"apellido" binding:"required"`
	Edad     int    `json:"edad" binding:"required"`
}

type User struct {
	service users.Service
}

func NewUser(s users.Service) *User {
	return &User{service: s}
}

// ListUsers godoc
// @Summary List Users
// @Tags Users
// @Description List users
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /users [get]
func (u *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, err := u.service.GetAll()
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}

		if len(users) == 0 {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "No se hallaron resultados"))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, users, ""))
	}
}

func (u *User) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idInt, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
		}

		user, err := u.service.GetById(idInt)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, user, ""))
	}
}

// StoreUsers godoc
// @Summary Store users
// @Tags Users
// @Description store users
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body Request true "User to store"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Router /users [post]
func (u *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req Request

		/* if !validateToken(ctx) {
			return
		} */

		if !validateFields(ctx, &req) {
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
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, user, ""))
	}
}

func (u *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		/* if !validateToken(ctx) {
			return
		} */

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Ingrese un ID valido"))
			return
		}

		var req Request
		if !validateFields(ctx, &req) {
			return
		}

		user, err := u.service.Update(int(id), req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.FechaCreacion)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, user, ""))
	}
}

func (u *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		/* if !validateToken(ctx) {
			return
		} */

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Ingrese un ID valido"))
			return
		}

		err = u.service.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, fmt.Sprintf("El usuario con el ID %d se eliminó correctamente", id), ""))
	}
}

func (u *User) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		/* if !validateToken(ctx) {
			return
		} */

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Ingrese un ID valido"))
			return
		}

		var req RequestPatch
		if !validatePatch(ctx, &req) {
			return
		}

		user, err := u.service.Patch(int(id), req.Apellido, req.Edad)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, user, ""))
	}
}

/*
========================================
FUNCTIONS VALIDATORS				   =
========================================
*/
/* func validateToken(ctx *gin.Context) bool {
	if token := ctx.GetHeader("token"); token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "NO tiene permisos para realizar la petición solicitada"))
		return false
	}
	return true
} */

func ValidateErrors(campo string, v validator.FieldError) string {
	switch v.Tag() {
	case "required":
		return "El campo " + campo + " es requerido"
	case "email":
		return "Direccion de correo electronico invalida"
	}
	return "Error desconocido..."
}

func validatePatch(ctx *gin.Context, req *RequestPatch) bool {
	if err := ctx.ShouldBindJSON(&req); err != nil {
		for _, fieldError := range err.(validator.ValidationErrors) {
			ctx.JSON(400, web.NewResponse(400, nil, ValidateErrors(fieldError.Field(), fieldError)))
		}
		return false
	}
	return true
}

func validateFields(ctx *gin.Context, req *Request) bool {
	if err := ctx.ShouldBindJSON(&req); err != nil {
		for _, fieldError := range err.(validator.ValidationErrors) {
			ctx.JSON(400, web.NewResponse(400, nil, ValidateErrors(fieldError.Field(), fieldError)))
		}
		return false
	}
	return true
}
