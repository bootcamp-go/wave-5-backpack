package handler

import (
	"github.com/gin-gonic/gin"
	"goweb/clase3_parte2/internal/users"
	"strconv"
	"fmt"
	"os"
)

// Se genera la estructura del request
type request struct {
	Nombre 			string	`json:"nombre"`
	Apellido 		string	`json:"apellido"`
	Email 			string	`json:"email"`
	Edad 			int		`json:"edad"`
	Altura 			float64	`json:"altura"`
	Activo			*bool	`json:"activo"`
	FechaCreacion 	string	`json:"fecha_creacion"`
}

// Se genera la estructura del controlador que tiene como campo el servicio
type User struct {
	service users.Service
}

// Se genera la función que retorna el controlador
func NewUser(u users.Service) *User {
	return &User{
		service: u,
	}
}

/* Se generan todos los métodos correspondientes a los endpoints */

func (c *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{"error": "El token no es válido"})
			return
		}

		u, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, u)
	}
}

func (c *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{ "error": "El token no es válido"})
			return
		}

		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		u, err := c.service.Store(req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.FechaCreacion)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, u)
	}
}

// Se agrega el controlador Update en el Handler de usuarios
func (c *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{ "error": "El token no es válido" })
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{ "error": "ID inválido"})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{ "error": err.Error()})
			return
		}

		var errors []string
		if req.Nombre == "" { errors = append(errors, "El nombre del usuario es requerido")}
		if req.Apellido == "" { errors = append(errors, "El apellido del usuario es requerido")}
		if req.Email == "" { errors = append(errors, "El email del usuario es requerido.")}
		if req.Edad  == 0 { errors = append(errors, "La edad del usuario es requerida.")}
		if req.Altura == 0 { errors = append(errors, "La altura del usuario es requerida.")}
		if req.Activo == nil { errors = append(errors, "El campo activo del usuario es requerido.")}
		if req.FechaCreacion == "" { errors = append(errors, "El campo fecha de creación del usuario es requerido.")}
		
		if len(errors) > 0 { 
			ctx.JSON(400, gin.H{ "errores": errors }) 
			return
		}

		u, err := c.service.Update(int(id), req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.FechaCreacion)
		if err != nil {
			ctx.JSON(404, gin.H{ "error": err.Error() })
			return
		}

		ctx.JSON(200, u)
	}
}

// Se agrega el controlador UpdateNameAge en el Handler de usuarios
func (c *User) UpdateLastNameAndAge() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{ "error": "El token no es válido" })
			return
		}
	
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{ "error": "ID inválido" })
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{ "error": err.Error() })
			return
		}

		var errors []string
		if req.Apellido == "" { errors = append(errors, "El apellido del usuario es requerido")}
		if req.Edad == 0 { errors = append(errors, "La edad del usuario es requerida")}

		if len(errors) > 0 { 
			ctx.JSON(400, gin.H{ "errores": errors }) 
			return
		}

		u, err := c.service.UpdateLastNameAndAge(int(id), req.Apellido, req.Edad)
		if err != nil {
			ctx.JSON(404, gin.H{ "error": err.Error() })
			return
		}
		ctx.JSON(200, u)
	}
}

// Se agrega el controlador Delete en el Handler de usuarios
func (c *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{ "error": "El token no es válido" })
			return
		}
		
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{ "error": "ID inválido" })
			return
		}
		
		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, gin.H{ "error": err.Error() })
			return
		}
		ctx.JSON(200, gin.H{ "data": fmt.Sprintf("El usuario %d ha sido eliminado", id) })
	}
}