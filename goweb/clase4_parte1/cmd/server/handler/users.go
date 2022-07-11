package handler

import (
	"github.com/gin-gonic/gin"
	"goweb/clase4_parte1/internal/users"
	"goweb/clase4_parte1/pkg/web"
	"strconv"
	"os"
	"fmt"
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
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "El token no es válido"))
			return
		}

		u, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}

		if len(u) == 0 {
			ctx.JSON(404, web.NewResponse(404, nil, "No hay usuarios registrados"))
			return
		}

		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}

func (c *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "El token no es válido"))
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		errors := validateRequest(req)
		if errors != "" { 
			ctx.JSON(400, web.NewResponse(400, nil, errors)) 
			return
		}

		u, err := c.service.Store(req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.FechaCreacion)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}

// Se agrega el controlador Update en el Handler de usuarios
func (c *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "El token no es válido"))
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "ID inválido"))
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		errors := validateRequest(req)
		if errors != "" { 
			ctx.JSON(400, web.NewResponse(400, nil, errors)) 
			return
		}

		u, err := c.service.Update(int(id), req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.FechaCreacion)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}

// Se agrega el controlador UpdateNameAge en el Handler de usuarios
func (c *User) UpdateLastNameAndAge() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "El token no es válido"))
			return
		}
	
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "ID inválido"))
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		var errors string
		if req.Apellido == "" { errors += "El apellido del usuario es requerido. " }
		if req.Edad == 0 { errors += "La edad del usuario es requerida." }

		if errors != "" { 
			ctx.JSON(400, web.NewResponse(400, nil, errors)) 
			return
		}

		u, err := c.service.UpdateLastNameAndAge(int(id), req.Apellido, req.Edad)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}

// Se agrega el controlador Delete en el Handler de usuarios
func (c *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "El token no es válido"))
			return
		}
		
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "ID inválido"))
			return
		}
		
		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, fmt.Sprintf("El usuario %d ha sido eliminado", id), ""))
		//ctx.JSON(200, gin.H{ "data": fmt.Sprintf("El usuario %d ha sido eliminado", id) })
	}
}

// Validación de los campos del request
func validateRequest(req request) string {
	var errors string
	if req.Nombre == "" { errors += "El nombre del usuario es requerido. "}
	if req.Apellido == "" { errors += "El apellido del usuario es requerido. "}
	if req.Email == "" { errors += "El email del usuario es requerido. "}
	if req.Edad  == 0 { errors += "La edad del usuario es requerida. "}
	if req.Altura == 0 { errors += "La altura del usuario es requerida. "}
	if req.Activo == nil { errors += "El campo activo del usuario es requerido. "}
	if req.FechaCreacion == "" { errors += "El campo fecha de creación del usuario es requerido."}
	return errors
}