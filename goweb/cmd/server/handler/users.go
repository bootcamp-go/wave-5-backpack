package handler

import (
	"fmt"
	"strconv"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/users"
	"github.com/bootcamp-go/wave-5-backpack/goweb/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Name         string  `json:"name" binding:"required"`
	LastName     string  `json:"last_name" binding:"required"`
	Email        string  `json:"email" binding:"required"`
	Age          int     `json:"age" binding:"required"`
	Height       float64 `json:"height" binding:"required"`
	Active       bool    `json:"active" binding:"required"`
	CreationDate string  `json:"creation_date" binding:"required"`
}

type patchRequest struct {
	LastName string `json:"last_name" binding:"required"`
	Age      int    `json:"age" binding:"required"`
}

type User struct {
	service users.Service
}

func NewUser(p users.Service) *User {
	return &User{
		service: p,
	}
}

// ListUsers godoc
// @Summary List users
// @Tags Users
// @Description get users
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @router /users [get]
func (u *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u, err := u.service.GetAll(ctx)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}

func (c *User) GetByName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println(ctx.Param("nombre"))
		u, err := c.service.GetByName(ctx, ctx.Param("nombre"))
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "no existen registros con el nombre indicado"))
			return
		}
		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}

func (u *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			if req.Name == "" {
				ctx.JSON(400, web.NewResponse(400, nil, "el nombre del usuario es requerido"))
				return
			}

			if req.LastName == "" {
				ctx.JSON(400, web.NewResponse(400, nil, "el apellido del usuario es requerido"))
				return
			}

			// validar email con regex
			if req.Email == "" {
				ctx.JSON(400, web.NewResponse(400, nil, "el email del usuario es requerido"))
				return
			}

			if req.Age < 1 {
				ctx.JSON(400, web.NewResponse(400, nil, "la edad del usuario debe ser mayor a 0"))
				return
			}

			if req.Height <= 0 {
				ctx.JSON(400, web.NewResponse(400, nil, "la altura del usuario es requerida"))
				return
			}

			// preguntar, es raro
			if !req.Active {
				ctx.JSON(400, web.NewResponse(400, nil, "es requerido saber si es activo el usuario"))
				return
			}

			if req.CreationDate == "" {
				ctx.JSON(400, web.NewResponse(400, nil, "la fecha de creacion del usuario es requerida"))
				return
			}
		}

		u, err := u.service.Store(ctx, req.Age, req.Name, req.LastName, req.Email, req.CreationDate, req.Height, req.Active)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(201, web.NewResponse(201, u, ""))
	}
}

func (u *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "id inválido"))
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			if req.Name == "" {
				ctx.JSON(400, web.NewResponse(400, nil, "el nombre del usuario es requerido"))
				return
			}

			if req.LastName == "" {
				ctx.JSON(400, web.NewResponse(400, nil, "el apellido del usuario es requerido"))
				return
			}

			// validar email con regex
			if req.Email == "" {
				ctx.JSON(400, web.NewResponse(400, nil, "el email del usuario es requerido"))
				return
			}

			if req.Age < 1 {
				ctx.JSON(400, web.NewResponse(400, nil, "la edad del usuario debe ser mayor a 0"))
				return
			}

			if req.Height <= 0 {
				ctx.JSON(400, web.NewResponse(400, nil, "la altura del usuario es requerida"))
				return
			}

			// preguntar, es raro
			if !req.Active {
				ctx.JSON(400, web.NewResponse(400, nil, "es requerido saber si es activo el usuario"))
				return
			}

			if req.CreationDate == "" {
				ctx.JSON(400, web.NewResponse(400, nil, "la fecha de creacion del usuario es requerida"))
				return
			}
		}

		u, err := u.service.Update(ctx, id, req.Age, req.Name, req.LastName, req.Email, req.CreationDate, req.Height, req.Active)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}

func (u *User) UpdateLastNameAndAge() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "id inválido"))
			return
		}

		var req patchRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			if req.LastName == "" {
				ctx.JSON(400, web.NewResponse(400, nil, "el apellido del usuario es requerido"))
				return
			}

			if req.Age < 1 {
				ctx.JSON(400, web.NewResponse(400, nil, "la edad del usuario debe ser mayor a 0"))
				return
			}
		}

		u, err := u.service.UpdateLastNameAndAge(ctx, id, req.Age, req.LastName)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}

func (u *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "id inválido"))
			return
		}

		err = u.service.Delete(ctx, id)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(204, web.NewResponse(204, nil, fmt.Sprintf("El producto %d ha sido eliminado", id)))
	}
}
