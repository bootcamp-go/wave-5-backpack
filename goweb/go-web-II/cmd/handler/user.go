package handler

import (
	"fmt"
	"goweb/go-web-II/internal/products"
	"goweb/go-web-II/pkg/web"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type request struct {
	Name    string `json:"name" binding:"required"`
	Surname string `json:"surname" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Age     int    `json:"age" binding:"required"`
	Active  bool   `json:"active"`
	Created string `json:"created" binding:"required"`
}

/*
ac
á va la data que recibimos del body
*/

type User struct {
	service products.Service
}

func NewUser(s products.Service) *User {
	return &User{service: s}
}

func tokenValidator(c *gin.Context) {
	token := c.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		c.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "token invalido"))
		return
	}
}

func (u *User) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenValidator(c)
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		err = u.service.Delete(id)
		if err != nil {
			c.JSON(404, web.NewResponse(400, nil, err.Error()))
		}

		c.JSON(200, gin.H{"data": fmt.Sprintf("el producto %d ha sido eliminado", id)})
	}
}

func (u *User) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenValidator(c)
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if req.Age == 0 {
			c.JSON(400, web.NewResponse(400, nil, "la edad es necesaria"))
		}
		if req.Created == "" {
			c.JSON(400, web.NewResponse(400, nil, "la fecha de creacion es necesaria"))
		}
		if req.Email == "" {
			c.JSON(400, web.NewResponse(400, nil, "el email es necesario"))
		}
		if req.Name == "" {
			c.JSON(400, web.NewResponse(400, nil, "el nombres es necesario"))
		}
		if req.Surname == "" {
			c.JSON(400, web.NewResponse(400, nil, "el apellido es necesario"))
		}
		u, err := u.service.Update(id, req.Age, req.Name, req.Surname, req.Email, req.Created, req.Active)
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		c.JSON(200, u)
	}
}

// ListProducts godoc
// @Summary Get a list of users
// @Tags Users endpoints
// @Description get users
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Failure 204 {object} web.Response
// @Router /users [get]
func (u *User) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenValidator(c)
		users, err := u.service.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, web.NewResponse(400, nil, err.Error()))
			return
		}
		if len(users) <= 0 {
			c.JSON(http.StatusOK, gin.H{
				"message": "no se encontraron usuarios"})
		}
		c.JSON(200, web.NewResponse(200, users, ""))
	}
}

// StoreProducts godoc
// @Summary Store user
// @Tags Users
// @Description store users
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Users to store"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Router / [post]
func (u *User) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenValidator(c)
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := u.service.Store(req.Age, req.Name, req.Surname, req.Email, req.Created, req.Active)
		if err != nil {
			c.JSON(http.StatusInternalServerError, web.NewResponse(400, nil, err.Error()))
			return
		}
		c.JSON(http.StatusOK, web.NewResponse(200, user, ""))
	}
}
