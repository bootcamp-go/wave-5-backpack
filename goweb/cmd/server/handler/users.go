package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/bootcamp-go/wave-5-backpack/internal/users"
	"github.com/bootcamp-go/wave-5-backpack/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Name       string  `json: "name"`
	Lastname   string  `json: "lastname"`
	Email      string  `json: "email"`
	Age        int     `json: "age"`
	Height     float32 `json: "height"`
	Active     bool    `json: "active"`
	DoCreation string  `json: "doCreation"`
}

type User struct {
	service users.Service
}

func NewUser(u users.Service) *User {
	return &User{
		service: u,
	}
}

func (c *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := c.service.ValidateToken(ctx.Request.Header.Get("token")); err != nil {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Invalid Token"))
			return
		}
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

func (c *User) StoreUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := c.service.ValidateToken(ctx.Request.Header.Get("token")); err != nil {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Invalid Token"))
			return
		}
		var req request
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"ERROR": err.Error(),
			})
			return
		}

		/* 		errMsg := c.service.ValidateReq()
		 */

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
			ctx.JSON(400, gin.H{"errores": errMsg})
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
		if err := c.service.ValidateToken(ctx.Request.Header.Get("token")); err != nil {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Invalid Token"))
			return
		}
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
		if err := c.service.ValidateToken(ctx.Request.Header.Get("token")); err != nil {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Invalid Token"))
			return
		}
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

		if v := fieldsValidator(req); v != "" {
			ctx.JSON(400, web.NewResponse(400, nil, v))
			return
		}
	
		updatedUser, err := c.service.UpdateUser(id, req.Name, req.Lastname, req.Email, req.Age, req.Height, req.Active, req.DoCreation)
		if err != nil {
			ctx.JSON(404, gin.H{
				"ERROR": err.Error(),
			})
			return
		}
		ctx.JSON(200, updatedUser)
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
		fullMsg := "Por favor completar los siguientes campos: "+strings.Join(errMsg, ", ")
		return fullMsg
	}
	return ""
}