package handler

import (
	"clase2_2/internal/users"
	"clase2_2/pkg/web"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type request struct {
	Name       string  `json:"name"`
	LastName   string  `json:"last_name"`
	Mail       string  `json:"mail"`
	Years      int     `json:"years"`
	Tall       float64 `json:"tall"`
	Enable     bool    `json:"enable"`
	CreateDate string  `json:"create_date"`
}
type User struct {
	service users.Service
}

func NewUser(s users.Service) *User {
	return &User{service: s}
}
func (u *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
			return
		}
		users, err := u.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		}
		ctx.JSON(http.StatusOK, users)
	}
}
func (u *User) AddUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		if req.Name == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "El nombre del usuario es requerido"))
			return
		}
		if req.LastName == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "El apellido del usuario es requerido"))
			return
		}
		if req.Mail == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "El correo del usuario es requerido"))
			return
		}
		if req.Years == 0 {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "La edad del usuario es requerido"))
			return
		}
		if req.Tall == 0 {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "La altura del usuario es requerido"))
			return
		}
		user, err := u.service.AddUser(req.Name, req.LastName, req.Mail, req.CreateDate, req.Years, req.Tall, req.Enable)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, user, ""))
	}
}
func (u *User) UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "invalid ID"))
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		if req.Name == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "El nombre del usuario es requerido"))
			return
		}
		if req.LastName == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "El apellido del usuario es requerido"))
			return
		}
		if req.Mail == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "El correo del usuario es requerido"))
			return
		}
		if req.Years == 0 {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "La edad del usuario es requerido"))
			return
		}
		if req.Tall == 0 {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "La altura del usuario es requerido"))
			return
		}
		user, err := u.service.UpdateUser(req.Name, req.LastName, req.Mail, req.CreateDate, req.Years, id, req.Tall, req.Enable)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, user, ""))
	}
}
func (u *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "invalid ID"))
			return
		}
		err = u.service.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		respose := fmt.Sprintf("El producto %d ha sido eliminado", id)
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, respose, ""))
	}
}
func (u *User) UpdateUserName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "invalid ID"))
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		if req.Name == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "El nombre del usuario es requerido"))
			return
		}
		user, err := u.service.UpdateUserName(req.Name, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}
