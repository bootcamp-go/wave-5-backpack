package handler

import (
	"fmt"
	"strconv"

	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/internal/domain"
	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/internal/users"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type request struct {
	Age                        int     `binding:"required"`
	FirstName, LastName, Email string  `binding:"required"`
	Height                     float64 `binding:"required"`
	Active                     bool    `binding:"required"`
}

type User struct {
	service users.Service
}

func NewUser(u users.Service) *User {
	return &User{
		service: u,
	}
}

func (c *User) GetAll(ctx *gin.Context) {
	users, err := c.service.GetAll()

	if err != nil {
		ctx.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	filters, err := querysMap(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"eror": err.Error()})
		return
	}

	resultUsers, err := c.service.FilterUsers(filters, users)

	if err != nil {
		ctx.JSON(500, gin.H{"error": "internal server error"})
		return
	}

	ctx.JSON(200, *resultUsers)
}

func (c *User) GetById(ctx *gin.Context) {
	Id, err := strconv.Atoi(ctx.Param("Id"))

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := c.service.GetById(Id)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	empty := domain.User{}
	if user == empty {
		ctx.JSON(404, gin.H{
			"error": "user not found",
		})
		return
	}

	ctx.JSON(200, user)
}

func (c *User) Store(ctx *gin.Context) {
	var req request

	if err := ctx.ShouldBindJSON(&req); err != nil {
		errs := err.(validator.ValidationErrors)
		for _, valError := range errs {
			if valError.Tag() == "required" {
				ctx.JSON(400, gin.H{
					"error": fmt.Sprintf("el campo '%s' es requerido", valError.Field()),
				})
				return
			}
		}
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := c.service.Store(req.Age, req.FirstName, req.LastName, req.Email, req.Height, req.Active)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(201, user)
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
