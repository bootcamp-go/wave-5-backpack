package handler

import (
	"Clase2-1/internal/usuarios"
	"Clase2-1/pkg/web"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre          string `json:"nombre"`
	Apellido        string `json:"apellido"`
	Email           string `json:"email"`
	Edad            int    `json:"edad"`
	Altura          int    `json:"altura"`
	Activo          bool   `json:"activo"`
	FechaDeCreacion string `json:"fecha_de_creacion"`
}

type Usuario struct {
	service usuarios.Service
}

func NewUser(user usuarios.Service) *Usuario {
	return &Usuario{
		service: user,
	}
}

// StoreUsers godoc
// @Summary Store users
// @Tags Users
// @Description store users
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param user body request true "User to store"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 409 {object} web.Response
// @Router /usuarios [post]
func (u *Usuario) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token inválido"))
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		stringErrores := ""
		if req.Nombre == "" {
			stringErrores = stringErrores + "El nombre del usuario es requerido\n"
		}

		if req.Apellido == "" {
			stringErrores = stringErrores + "El apellido del usuario es requerido\n"
		}

		if req.Email == "" {
			stringErrores = stringErrores + "El email del usuario es requerido\n"
		}

		if req.Edad == 0 {
			stringErrores = stringErrores + "La edad es requerida\n"
		}
		if req.Altura == 0 {
			stringErrores = stringErrores + "La altura es requerida\n"
		}

		if len(stringErrores) != 0 {
			ctx.JSON(400, web.NewResponse(400, nil, stringErrores))
			return
		}

		user, err := u.service.Store(req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, user, ""))
	}
}

// ListUsers godoc
// @Summary List users
// @Tags Users
// @Description get users
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Failure 204 {object} web.Response
// @Router /usuarios [get]
func (u *Usuario) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token inválido"))
			return
		}
		users, err := u.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, web.NewResponse(200, users, ""))
	}
}

// UpdateUsers godoc
// @Summary Update users
// @Tags Users
// @Description update users
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param user body request true "User to update"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 409 {object} web.Response
// @Router /usuarios/:id [put]
func (user *Usuario) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token inválido"))
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, "Id inválido"))
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, err.Error()))
			return
		}
		stringErrores := ""
		if req.Nombre == "" {
			stringErrores = stringErrores + "El nombre del usuario es requerido\n"
		}

		if req.Apellido == "" {
			stringErrores = stringErrores + "El apellido del usuario es requerido\n"
		}

		if req.Email == "" {
			stringErrores = stringErrores + "El email del usuario es requerido\n"
		}

		if req.Edad == 0 {
			stringErrores = stringErrores + "La edad es requerida\n"
		}
		if req.Altura == 0 {
			stringErrores = stringErrores + "La altura es requerida\n"
		}

		if len(stringErrores) != 0 {
			ctx.JSON(400, web.NewResponse(400, nil, stringErrores))
			return
		}
		u, err := user.service.Update(int(id), req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}

// StoreUsers godoc
// @Summary Update SurnameAndAge users
// @Tags Users
// @Description update surname and age users
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param user body request true "surname and age to update"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 409 {object} web.Response
// @Router /usuarios [patch]
func (user *Usuario) UpdateSurnameAndAge() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token inválido"))
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(401, nil, "Id inválido"))
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(401, nil, err.Error()))
			return
		}

		stringErrores := ""
		if req.Apellido == "" {
			stringErrores = stringErrores + "El apellido del usuario es requerido\n"
		}

		if req.Edad == 0 {
			stringErrores = stringErrores + "La edad del usuario es requerida\n"
		}
		if len(stringErrores) != 0 {
			ctx.JSON(400, web.NewResponse(400, nil, stringErrores))
			return
		}

		u, err := user.service.UpdateSurnameAndAge(int(id), req.Apellido, req.Edad)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}

// StoreUsers godoc
// @Summary Delete users
// @Tags Users
// @Description delete users
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 409 {object} web.Response
// @Router /usuarios [delete]
func (u *Usuario) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token inválido"))
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Id inválido"))
			return
		}

		err = u.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, web.NewResponse(200, fmt.Sprintf("El producto %d ha sido eliminado", id), ""))
	}
}
