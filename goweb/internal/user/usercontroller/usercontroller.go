package usercontroller

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/internal/errors"
	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/internal/file"
	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/internal/user"
	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/internal/user/userhelper"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetAll(ctx *gin.Context) {
	db, err := file.ReadJSONFile("users.json")

	if err != nil {
		errors.InternalError(ctx, err)
		return
	}

	resultUsers, err := userhelper.FilterUsers(ctx, db.Users)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, *resultUsers)
}

func GetById(ctx *gin.Context) {
	Id, err := strconv.Atoi(ctx.Param("Id"))

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db, err := file.ReadJSONFile("users.json")

	if err != nil {
		errors.InternalError(ctx, err)
		return
	}

	for _, user := range db.Users {
		if Id == user.Id {
			ctx.JSON(200, user)
			return
		}
	}

	ctx.JSON(404, gin.H{
		"error": "user not found",
	})
}

func Create(ctx *gin.Context) {
	var user user.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
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

	db, err := file.ReadJSONFile("users.json")
	if err != nil {
		errors.InternalError(ctx, err)
		return
	}

	db.LastId++
	user.Id = db.LastId
	t := time.Now()
	user.CreatedAt = fmt.Sprintf("%02d/%02d/%d", t.Day(), t.Month(), t.Year())
	db.Users = append(db.Users, user)

	if err := file.WriteJSONFile("users.json", db); err != nil {
		errors.InternalError(ctx, err)
		return
	}

	ctx.JSON(201, user)
}
