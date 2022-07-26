package handler

import (
	"fmt"
	"strconv"

	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/internal/users"
	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type request struct {
	Age       int     `json:"age" binding:"required"`
	FirstName string  `json:"firstName" binding:"required"`
	LastName  string  `json:"lastName" binding:"required"`
	Email     string  `json:"email" binding:"required"`
	Height    float64 `json:"height" binding:"required"`
	Active    *bool   `json:"active" binding:"required"`
}

type putRequest struct {
	Age       int     `json:"age" binding:"required"`
	FirstName string  `json:"firstName" binding:"required"`
	LastName  string  `json:"lastName" binding:"required"`
	Email     string  `json:"email" binding:"required"`
	CreatedAt string  `json:"createdAt" binding:"required"`
	Height    float64 `json:"height" binding:"required"`
	Active    *bool   `json:"active" binding:"required"`
}

type patchRequest struct {
	Age      int    `json:"age"`
	LastName string `json:"lastName"`
}

type User struct {
	service users.Service
}

func NewUser(u users.Service) *User {
	return &User{
		service: u,
	}
}

// GetUsers godoc
// @Summary Get all users
// @Tags Users
// @Description get users
// @Produce json
// @Success 200 {object} web.Response
// @Router /users [get]
func (c *User) GetAll(ctx *gin.Context) {
	filters, err := querysMap(ctx)
	if err != nil {
		ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		return
	}

	users, err := c.service.GetAll(filters)

	if err != nil {
		ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
		return
	}

	ctx.JSON(200, web.NewResponse(200, users, ""))
}

// GetUsers godoc
// @Summary Get user by id
// @Tags Users
// @Description get user by id
// @Produce json
// @Success 200 {object} web.Response
// @Param id path string true "User id"
// @Router /users/{id} [get]
func (c *User) GetById(ctx *gin.Context) {
	Id, err := strconv.Atoi(ctx.Param("Id"))

	if err != nil {
		ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		return
	}

	user, err := c.service.GetById(Id)

	if err != nil {
		ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		return
	}

	ctx.JSON(200, web.NewResponse(200, user, ""))
}

// StoreUsers godoc
// @Summary Store users
// @Tags Users
// @Description store users
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param user body request true "User to store"
// @Success 200 {object} web.Response
// @Router /users [post]
func (c *User) Store(ctx *gin.Context) {
	var req request

	if err := ctx.ShouldBindJSON(&req); err != nil {
		errs := err.(validator.ValidationErrors)
		for _, valError := range errs {
			if valError.Tag() == "required" {
				ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("el campo '%s' es requerido", valError.Field())))
				return
			}
		}
		ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		return
	}

	user, err := c.service.Store(req.Age, req.FirstName, req.LastName, req.Email, req.Height, *req.Active)

	if err != nil {
		ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		return
	}

	ctx.JSON(201, web.NewResponse(201, user, ""))
}

// UpadteUsers godoc
// @Summary Update users
// @Tags Users
// @Description update users by id
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path string true "User id"
// @Param user body putRequest true "User to update"
// @Success 201 {object} web.Response
// @Router /users/{id} [put]
func (c *User) Update(ctx *gin.Context) {
	Id, err := strconv.Atoi(ctx.Param("Id"))

	if err != nil {
		ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		return
	}

	var req putRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		errs := err.(validator.ValidationErrors)
		for _, valError := range errs {
			if valError.Tag() == "required" {
				ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("el campo '%s' es requerido", valError.Field())))
				return
			}
		}
		ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		return
	}

	user, err := c.service.Update(Id, req.Age, req.FirstName, req.LastName, req.Email, req.CreatedAt, req.Height, *req.Active)

	if err != nil {
		ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		return
	}

	ctx.JSON(200, web.NewResponse(200, user, ""))
}

// UpadteUsersAgeLastName godoc
// @Summary Update user Age or LastName by id
// @Tags Users
// @Description update user Age or LastName by id
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path string true "User id"
// @Param user body patchRequest true "Fields to update"
// @Success 201 {object} web.Response
// @Router /users/{id} [patch]
func (c *User) UpdateAgeLastName(ctx *gin.Context) {
	Id, err := strconv.Atoi(ctx.Param("Id"))

	if err != nil {
		ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		return
	}

	var req patchRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		return
	}
	if req.LastName == "" && req.Age == 0 {
		ctx.JSON(400, web.NewResponse(400, nil, "Debe enviar por lo menos uno de los siguientes campos: 'LastName', 'Age'"))
		return
	}
	user, err := c.service.UpdateAgeLastName(Id, req.Age, req.LastName)

	if err != nil {
		ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		return
	}

	ctx.JSON(200, web.NewResponse(200, user, ""))
}

// DeleteUser godoc
// @Summary Delete user by id
// @Tags Users
// @Description delete user by id
// @Produce json
// @Param token header string true "token"
// @Param id path string true "User id"
// @Success 204 {object} web.Response
// @Router /users/{id} [delete]
func (c *User) Delete(ctx *gin.Context) {
	Id, err := strconv.Atoi(ctx.Param("Id"))

	if err != nil {
		ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		return
	}

	err = c.service.Delete(Id)

	if err != nil {
		ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		return
	}

	ctx.JSON(204, web.NewResponse(204, nil, ""))
}

func querysMap(ctx *gin.Context) (map[string]interface{}, error) {
	querys := make(map[string]interface{})

	if ctx.Query("Id") != "" {
		id, err := strconv.Atoi(ctx.Query("Id"))
		if err != nil {
			return querys, err
		}
		querys["Id"] = id
	}
	if ctx.Query("Age") != "" {
		age, err := strconv.Atoi(ctx.Query("Age"))
		if err != nil {
			return querys, err
		}
		querys["Age"] = age
	}
	if ctx.Query("Heigth") != "" {
		height, err := strconv.ParseFloat(ctx.Query("Heigth"), 64)
		if err != nil {
			return querys, err
		}
		querys["Height"] = height
	}
	if ctx.Query("Active") != "" {
		active, err := strconv.ParseBool(ctx.Query("Active"))
		if err != nil {
			return querys, err
		}
		querys["Active"] = active
	}
	if ctx.Query("FirstName") != "" {
		querys["FirstName"] = ctx.Query("FirstName")
	}
	if ctx.Query("LastName") != "" {
		querys["LastName"] = ctx.Query("LastName")
	}
	if ctx.Query("Email") != "" {
		querys["Email"] = ctx.Query("Email")
	}
	if ctx.Query("CreatedAt") != "" {
		querys["CreatedAt"] = ctx.Query("CreatedAt")
	}

	return querys, nil
}
