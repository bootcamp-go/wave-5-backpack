package handler

import (
	"errors"
	"goweb/internal/users"
	"goweb/pkg/web"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type request struct {
	Name     string  `json:"name" binding:"required"`
	LastName string  `json:"last_name" binding:"required"`
	Email    string  `json:"email" binding:"required"`
	Age      int     `json:"age" binding:"required"`
	Height   float64 `json:"height" binding:"required"`
	Active   bool    `json:"active" binding:"required"`
}

type User struct {
	service users.Service
}

func NewUser(u users.Service) *User {
	return &User{
		service: u,
	}
}

func verifyToken(ctx *gin.Context) bool {
	token := ctx.GetHeader("Authorization")
	return token == os.Getenv("TOKEN")
}

func getId(ctx *gin.Context) (int, error) {
	paramId, exist := ctx.Params.Get("id")
	if !exist {
		return 0, errors.New("send valid Id")
	}

	id, err := strconv.Atoi(paramId)

	if err != nil {
		ctx.JSON(web.NewResponse(http.StatusBadRequest, nil, "send valid Id"))
		return 0, errors.New("send valid Id")
	}

	return id, nil
}

func (u *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, err := u.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, gin.H{
			"data": users,
		})
	}
}

func (u *User) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !verifyToken(ctx) {
			ctx.JSON(web.NewResponse(http.StatusUnauthorized, nil, "access denied: token unauthorized"))
			return
		}

		request := request{}

		if err := ctx.ShouldBind(&request); err != nil {
			ctx.JSON(web.NewResponse(http.StatusBadRequest, nil, "error has ocurred"))
			return
		}

		user, err := u.service.Store(request.Name, request.LastName, request.Email, request.Age, request.Height, request.Active)
		if err != nil {
			ctx.JSON(web.NewResponse(http.StatusBadRequest, nil, "error: can't create user"))
			return
		}

		ctx.JSON(web.NewResponse(http.StatusCreated, user, ""))

	}
}

func (u *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !verifyToken(ctx) {
			ctx.JSON(web.NewResponse(http.StatusUnauthorized, nil, "access denied: token unauthorized"))
			return
		}

		id, err := getId(ctx)
		if err != nil {
			ctx.JSON(web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		userRequest := request{}

		if err := ctx.ShouldBind(&userRequest); err != nil {
			ctx.JSON(web.NewResponse(http.StatusBadRequest, nil, "bad request"))
			return
		}

		user, err := u.service.Update(id, userRequest.Name, userRequest.LastName, userRequest.Email, userRequest.Age, userRequest.Height, userRequest.Active)
		if err != nil {
			ctx.JSON(web.NewResponse(http.StatusBadRequest, nil, "bad request"))
			return
		}

		ctx.JSON(web.NewResponse(http.StatusOK, user, ""))
	}
}

func (u *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !verifyToken(ctx) {
			ctx.JSON(web.NewResponse(http.StatusUnauthorized, nil, "access denied: token unauthorized"))
			return
		}

		id, err := getId(ctx)
		if err != nil {
			ctx.JSON(web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		}

		if err := u.service.Delete(id); err != nil {
			ctx.JSON(web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		ctx.JSON(web.NewResponse(http.StatusAccepted, "user was deleted", ""))
	}
}

func (u *User) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !verifyToken(ctx) {
			ctx.JSON(web.NewResponse(http.StatusUnauthorized, nil, "access denied: token unauthorized"))
			return
		}

		id, err := getId(ctx)
		if err != nil {
			ctx.JSON(web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		}

		result, err := u.service.GetById(id)

		if err != nil {
			ctx.JSON(web.NewResponse(http.StatusNotFound, nil, "user not found"))
			return
		}

		ctx.JSON(web.NewResponse(http.StatusOK, result, ""))

	}
}
