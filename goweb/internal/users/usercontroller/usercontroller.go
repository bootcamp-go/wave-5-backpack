package usercontroller

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id, Age                               int
	FirstName, LastName, Email, CreatedAt string
	Height                                float64
	Active                                bool
}

func ReadJSONFile(fileName string) ([]User, error) {
	data, err := os.ReadFile(fileName)
	users := []User{}
	if err == nil {
		err = json.Unmarshal(data, &users)
	}
	return users, err
}

func filterUsers(ctx *gin.Context, users []User) (*[]User, error) {
	resultUsers := []User{}

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

func GetAll(ctx *gin.Context) {
	users, err := ReadJSONFile("users.json")

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "internal server error",
		})
		fmt.Printf("Error: %v\n", err)
		return
	}

	resultUsers, err := filterUsers(ctx, users)

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

	users, err := ReadJSONFile("users.json")

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "internal server error",
		})
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, user := range users {
		if Id == user.Id {
			ctx.JSON(200, user)
			return
		}
	}
	ctx.JSON(400, gin.H{
		"error": "user not found",
	})
}
