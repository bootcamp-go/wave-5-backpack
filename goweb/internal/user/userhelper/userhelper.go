package userhelper

import (
	"strconv"

	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/internal/user"
	"github.com/gin-gonic/gin"
)

func FilterUsers(ctx *gin.Context, users []user.User) (*[]user.User, error) {
	resultUsers := []user.User{}

	for _, user := range users {
		if ctx.Query("Id") != "" {
			Id, err := strconv.Atoi(ctx.Query("Id"))
			if err != nil {
				return &resultUsers, err
			}
			if Id != user.Id {
				continue
			}
		}
		if ctx.Query("Age") != "" {
			Age, err := strconv.Atoi(ctx.Query("Age"))
			if err != nil {
				return &resultUsers, err
			}
			if Age != user.Age {
				continue
			}
		}
		if (ctx.Query("FirstName") != "" && ctx.Query("FirstName") != user.FirstName) ||
			(ctx.Query("LastName") != "" && ctx.Query("LastName") != user.LastName) ||
			(ctx.Query("Email") != "" && ctx.Query("Email") != user.Email) ||
			(ctx.Query("CreatedAt") != "" && ctx.Query("CreatedAt") != user.CreatedAt) {
			continue
		}
		if ctx.Query("Height") != "" {
			Height, err := strconv.ParseFloat(ctx.Query("Height"), 64)
			if err != nil {
				return &resultUsers, err
			}
			if Height != user.Height {
				continue
			}
		}
		if ctx.Query("Active") != "" {
			Active, err := strconv.ParseBool(ctx.Query("Active"))

			if err != nil {
				return &resultUsers, err
			}
			if Active != user.Active {
				continue
			}
		}
		resultUsers = append(resultUsers, user)
	}

	return &resultUsers, nil
}
