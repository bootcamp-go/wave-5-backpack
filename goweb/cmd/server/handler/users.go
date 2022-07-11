package handler

import (
	"goweb/internal/users"
	"goweb/pkg/web"
	"github.com/gin-gonic/gin"
	"strconv"
	"fmt"
	"os"
	"net/http"
	"strings"
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

func (c *User) GetAllUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// valido token
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token inválido"))
			return
		}
		u, err := c.service.GetAllUsers()
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}

func (c *User) GetUserById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// valido token
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token inválido"))
			return
		}
		id,_ := strconv.Atoi(ctx.Param("id"))
		u, err := c.service.GetUserById(id)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}

func (c *User) StoreUser() gin.HandlerFunc {
	return func(ctx *gin.Context){

		// valido token
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token inválido"))
			return
		}

		// traigo los datos del post y los guardo en una variable del tipo struct request que generé arriba
		var req request
		if err := ctx.Bind(&req); err !=nil{
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		// valido campos requeridos
		if v := validar(req); v != "" {
			ctx.JSON(400, web.NewResponse(400, nil, v))
			return
		}

		newUser, err := c.service.StoreUser(req.Name, req.LastName, req.Email, req.Age, req.Height, req.Active, req.CreatedAt)
		if err != nil{
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, newUser, ""))
	}
}

func (c *User) UpdateTotal() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// valido token
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token inválido"))
			return
		}
		
		//var errores []string
		
		id,err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "invalid ID"))
 			return
		}

		// validaciones
		var req request
		if err:= ctx.ShouldBindJSON(&req); err!=nil{
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		// valido campos requeridos
		if v := validar(req); v != "" {
			ctx.JSON(400, web.NewResponse(400, nil, v))
			return
		}

		userUpdated, err:=	 c.service.UpdateTotal(id, req.Name, req.LastName, req.Email, req.Age, req.Height, req.Active, req.CreatedAt)

		if err !=nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, userUpdated, ""))
	}
}

func (c *User) UpdatePartial() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// valido token
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token inválido"))
			return
		}

		var errores []string

		id,err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "invalid ID"))
 			return
		}

		// validaciones
		var req request
		if err:= ctx.ShouldBindJSON(&req); err!=nil{
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		if req.LastName == ""{
			errores = append(errores, "El apellido del usuario es requerido")
		}
		if req.Age == 0 {
			errores  =append(errores, "La edad del usuario es requerido")
		}
		if len(errores) > 0 {
			erroresStr := strings.Join(errores, ", ")
			ctx.JSON(400, web.NewResponse(400, nil, erroresStr))
			return
		}

		userUpdated, err:=	 c.service.UpdatePartial(id, req.LastName, req.Age)

		if err !=nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, userUpdated, ""))
	}
}

func (c *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// valido token
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token inválido"))
			return
		}
		
		id,err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "invalid ID"))
 			return
		}

		err=c.service.Delete(id)
		if err !=nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, fmt.Sprintf("El producto %d ha sido eliminado",
		id), ""))
	}
}

func validar(req request) string {
	var response []string
	if req.Name == ""{
		response = append(response, "Nombre") 
	}
	if req.LastName == ""{
		response = append(response, "Apellido")
	}
	if req.Email == ""{
		response = append(response, "Email")
	}
	if req.Age == 0 {
		response = append(response, "Edad")
	}
	if req.Height == 0 {
		response = append(response,"Altura")
	}
	if req.CreatedAt == ""{
		response = append(response, "Fecha de creación")
	}
	if len(response) > 0 {
		errores := "Por favor completar los siguientes campos: "+strings.Join(response, ", ")
		return errores
	}
	return ""
}