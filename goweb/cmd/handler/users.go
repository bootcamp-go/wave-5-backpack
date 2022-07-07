package handler

import (
	"goweb/internal/users"
	"github.com/gin-gonic/gin"
)

type request struct {
	Id int 				`json:"id"`
	Name string			`json:"name"`
	LastName string		`json:"lastname"`			
	Email string		`json:"email"`
	Age int				`json:"age"`
	Height float32		`json:"height"`
	Active bool			`json:"active"`
	CreatedAt string	`json:"createdat"`
}

type User struct {
	service users.Service
}

func NewUser(u users.Service) *User{
	return &User{
		service: u,
	}
}