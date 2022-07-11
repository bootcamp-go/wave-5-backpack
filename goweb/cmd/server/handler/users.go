package handler

import (
	"fmt"
	"goweb/internal/users"
	"goweb/pkg/web"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type request struct {
	Name         string  `json:"name"`
	LastName     string  `json:"lastname"`
	Email        string  `json:"email"`
	Age          int     `json:"age"`
	Height       float64 `json:"height"`
	Active       bool    `json:"active"`
	CreationDate string  `json:"creation-date" `
}

type User struct {
	service users.Service
}

func NewUsers(u users.Service) *User {
	return &User{
		service: u,
	}
}

func (u *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
			return
		}
		u, err := u.service.GetAll()
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}

func (u *User) NewUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		u, _ := u.service.NewUser(req.Name, req.LastName, req.Email, req.Age, req.Height, req.Active, req.CreationDate)
		if v := validators(req); len(v) != 0 {
			ctx.JSON(400, web.NewResponse(400, nil, v))
			return
		}

		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}

func (u *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Id Invalido"))
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		if v := validators(req); len(v) != 0 {
			ctx.JSON(400, web.NewResponse(400, nil, v))
			return
		}

		u, err := u.service.Update(int(id), req.Name, req.LastName, req.Email, req.Age, req.Height, req.Active, req.CreationDate)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(200, u)
		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}

func (u *User) UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Id inválido"))
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		if req.Name == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El nombre del producto es requerido"))
			return
		}

		u, err := u.service.UpdateName(int(id), req.Name)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}

func (u *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Id inválido"))
			return
		}
		err = u.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, fmt.Sprintf("El usuario %d ha sido eliminado", id), ""))
	}

}

func validators(req request) string {
	var requiredFiles string
	if req.Name == "" {
		requiredFiles += "El campo nombre del usuario es requerido\n"
	}
	if req.LastName == "" {
		requiredFiles += "El campo apellido del usuario es requerido\n"
	}
	if req.Email == "" {
		requiredFiles += "El campo email del usuario es requerido\n"
	}
	if strconv.Itoa(req.Age) == "" {
		requiredFiles += "El campo edad del usuario es requerido\n"
	}
	if strconv.FormatFloat(req.Height, 'E', -1, 64) == "" {
		requiredFiles += "El campo altura del usuario es requerido\n"
	}
	if strconv.FormatBool(req.Active) == "" {
		requiredFiles += "El campo activo del usuario es requerido\n"
	}
	if req.CreationDate == "" {
		requiredFiles += "El campo fecha de creacion del usuario es requerido\n"
	}

	return requiredFiles
}
