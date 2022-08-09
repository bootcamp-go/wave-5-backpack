package handler

import (
	"fmt"
	"github.com/bootcamp-go/wave-5-backpack/internal/users"
	"github.com/bootcamp-go/wave-5-backpack/pkg/web"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type request struct {
	Name       string  `json:"name"`
	Lastname   string  `json:"lastname"`
	Email      string  `json:"email"`
	Age        int     `json:"age"`
	Height     float32 `json:"height"`
	Active     bool    `json:"active"`
	DoCreation string  `json:"doCreation"`
}

type User struct {
	service users.Service
}

func NewUser(u users.Service) *User {
	return &User{
		service: u,
	}
}
// ListUsers godoc
// @Summary List users
// @Tags users
// @Description get users
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /users [get]
func (c *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		allUsers, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"ERROR": err.Error(),
			})
			return
		}
		ctx.JSON(200, web.NewResponse(200, allUsers, ""))
	}
}
// StoreUsers godoc
// @Summary Store users
// @Tags Users
// @Description store users
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param Users body request true "User to store"
// @Success 200 {object} web.Response
// @Router /users [post]
func (c *User) StoreUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req request
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"ERROR": err.Error(),
			})
			return
		}

		if validated := fieldsValidator(req); validated != "" {
			ctx.JSON(400, web.NewResponse(400, nil, validated))
			return
		}

		newUser, err := c.service.StoreUser(req.Name, req.Lastname, req.Email, req.Age, req.Height, req.Active, req.DoCreation)
		if err != nil {
			ctx.JSON(404, gin.H{
				"ERROR": err.Error(),
			})
			return
		}
		ctx.JSON(200, newUser)

	}
}
func (c *User) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, _ := strconv.Atoi(ctx.Param("id"))
		userFound, err := c.service.GetById(id)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))

			return
		}
		ctx.JSON(200, web.NewResponse(200, userFound, ""))
	}
}

func (c *User) UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Id"})
			return
		}
		var req request
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		if validated := fieldsValidator(req); validated != "" {
			ctx.JSON(400, web.NewResponse(400, nil, validated))
			return
		}

		updatedUser, err := c.service.UpdateUser(id, req.Name, req.Lastname, req.Email, req.Age, req.Height, req.Active, req.DoCreation)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, updatedUser, ""))
	}
}
func (c *User) UpdateLastnameAndAge() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "error: Invalid Id"))
			return
		}
		var req request
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		var errMsg []string

		if req.Lastname == "" {
			errMsg = append(errMsg, "Lastname required")
		}
		if req.Age == 0 {
			errMsg = append(errMsg, "Age required")
		}

		if len(errMsg) > 0 {
			fullMsg := "Missing data, please complete " + strings.Join(errMsg, ", ")
			ctx.JSON(400, web.NewResponse(400, nil, fullMsg))
			return
		}

		updatedUser, err := c.service.UpdateLastnameAndAge(id, req.Lastname, req.Age)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, updatedUser, ""))
	}
}

func fieldsValidator(req request) string {
	var errMsg []string

	if req.Name == "" {
		errMsg = append(errMsg, "Name required")
	}
	if req.Lastname == "" {
		errMsg = append(errMsg, "Lastname required")
	}
	if req.Email == "" {
		errMsg = append(errMsg, "Email required")
	}
	if req.Age == 0 {
		errMsg = append(errMsg, "Age required")
	}
	if req.Height == 0 {
		errMsg = append(errMsg, "Height required")
	}
	if req.DoCreation == "" {
		errMsg = append(errMsg, "Date of creation required")
	}

	if len(errMsg) > 0 {
		fullMsg := "Por favor completar los siguientes campos: " + strings.Join(errMsg, ", ")
		return fullMsg
	}
	return ""
}

func (c *User) DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// valido token

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "invalid ID"))
			return
		}

		err = c.service.DeleteUser(id)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, fmt.Sprintf("El producto %v ha sido eliminado",
			id), ""))
	}
}


func (c *User) GetByName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u, err := c.service.GetByName(ctx.Param("nombre"))
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "no existen registros con el nombre indicado"))
			return
		}
		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}