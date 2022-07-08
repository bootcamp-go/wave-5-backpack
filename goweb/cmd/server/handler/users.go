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
	Active                     *bool   `binding:"required"`
}

type putRequest struct {
	Age                        int     `binding:"required"`
	FirstName, LastName, Email string  `binding:"required"`
	CreatedAt                  string  `binding:"required"`
	Height                     float64 `binding:"required"`
	Active                     *bool   `binding:"required"`
}

type pathRequest struct {
	Age      int
	LastName string
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
	filters, err := querysMap(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"eror": err.Error()})
		return
	}

	users, err := c.service.GetAll(filters)

	if err != nil {
		ctx.JSON(500, gin.H{"error": "internal server error"})
		fmt.Printf("Internal Error: %v", err.Error())
		return
	}

	ctx.JSON(200, users)
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

	user, err := c.service.Store(req.Age, req.FirstName, req.LastName, req.Email, req.Height, *req.Active)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(201, user)
}

func (c *User) Update(ctx *gin.Context) {
	Id, err := strconv.Atoi(ctx.Param("Id"))

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var req putRequest

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

	user, err := c.service.Update(Id, req.Age, req.FirstName, req.LastName, req.Email, req.CreatedAt, req.Height, *req.Active)

	if err != nil {
		ctx.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, user)
}

func (c *User) UpdateAgeLastName(ctx *gin.Context) {
	Id, err := strconv.Atoi(ctx.Param("Id"))

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var req pathRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	if req.LastName == "" && req.Age == 0 {
		ctx.JSON(400, gin.H{
			"error": "Debe enviar por lo menos uno de los siguientes campos: 'LastName', 'Age'",
		})
		return
	}
	user, err := c.service.UpdateAgeLastName(Id, req.Age, req.LastName)

	if err != nil {
		ctx.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, user)
}

func (c *User) Delete(ctx *gin.Context) {
	Id, err := strconv.Atoi(ctx.Param("Id"))

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = c.service.Delete(Id)

	if err != nil {
		ctx.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Status(204)
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
