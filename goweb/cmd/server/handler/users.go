package handler

import (
	"goweb/internal/users"
	"github.com/gin-gonic/gin"
	"strconv"
)

type request struct {
	Id int 				`json:"id"`
	Name string			`json:"name" binding:"required"`
	LastName string		`json:"lastname" binding:"required"`			
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

func (c *User) GetAllUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// valido token
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		u, err := c.service.GetAllUsers()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, u)
	}
}

func (c *User) GetUserById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// valido token
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		id,_ := strconv.Atoi(ctx.Param("id"))
		u, err := c.service.GetUserById(id)
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, u)
	}
}

func (c *User) StoreUser() gin.HandlerFunc {
	return func(ctx *gin.Context){

		// valido token
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		// traigo los datos del post y los guardo en una variable del tipo struct request que generé arriba
		var req request
		if err := ctx.Bind(&req); err !=nil{
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		newUser, err := c.service.StoreUser(req.Name, req.LastName, req.Email, req.Age, req.Height, req.Active, req.CreatedAt)
		if err != nil{
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, newUser)
	}
}

func (c *User) UpdateTotal() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// valido token
		token := ctx.Request.Header.Get("token")
		var errores []string
		
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		// validaciones
		var req request
		if err:= ctx.ShouldBindJSON(&req); err!=nil{
			ctx.JSON(40, gin.H{"error": err.Error()})
			return
		}
		if req.Name == ""{
			errores = append(errores, "El nombre del usuario es requerido")
		}
		if req.LastName == ""{
			errores = append(errores, "El apellido del usuario es requerido")
		}
		if req.Email == ""{
			errores = append(errores, "El email del usuario es requerido")
		}
		if req.Age == 0 {
			errores = append(errores, "La edad del usuario es requerido")
		}
		if req.Height == 0 {
			errores = append(errores, "La altura del usuario es requerido")
		}
		if req.CreatedAt == ""{
			errores = append(errores, "La fecha de creacion del usuario es requerido")
		}
		if len(errores) > 0 {
			ctx.JSON(400, gin.H{"errores": errores})
			return
		}
		id,_ := strconv.Atoi(ctx.Param("id"))
		userToUpdate, err:=	 c.service.UpdateTotal(id, req.Name, req.LastName, req.Email, req.Age, req.Height, req.Active, req.CreatedAt)

		if err !=nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, userToUpdate)
	}
}