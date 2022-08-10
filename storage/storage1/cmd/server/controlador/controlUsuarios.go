package controlador

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/del_rio/web-server/internal/domain"
	"github.com/del_rio/web-server/internal/usuarios"
	"github.com/del_rio/web-server/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre         string `json:"nombre"`
	Apellido       string `json:"apellido"`
	Email          string `json:"email"`
	Fecha_creacion string `json:"fecha_creacion"`
	Edad           int    `json:"edad"`
	Altura         int    `json:"altura"`
	Activo         *bool  `json:"activo"`
}
type Usuario struct {
	service usuarios.Servicio //service = servicio no lo cambie porque ya lo implemente en mas de 3 lugares
}

func NewControlador(us usuarios.Servicio) *Usuario {
	return &Usuario{
		service: us,
	}
}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")
	if requiredToken == "" {
		log.Fatal("no se encontro  el .env")
	}
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token == "" {
			ctx.AbortWithStatusJSON(401, web.NewResponse(401, nil, "no hay token en cabecera genio"))
			return
		}
		if token != requiredToken {
			ctx.AbortWithStatusJSON(401, web.NewResponse(401, nil, "token incorrecto game over!!"))
		}
		ctx.Next()
	}
}

// ListUsuarios godoc
// @Summary List usuarios
// @Tags Usuarios
// @Description get usuarios
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /usuarios [get]
func (u *Usuario) VerUsuarios() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		usuarios, err := u.service.GetAll()
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "problema obteniendo datos: "+err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, usuarios, ""))
		return
	}
}

// StoreUsuarios godoc
// @Summary Store usuarios
// @Tags Usuarios
// @Description store usuarios
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param Usuario body request true "Usuario to store"
// @Success 200 {object} web.Response
// @Router /usuarios [post]
func (u *Usuario) AgregarUsuarios() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "problema usando bind: "+err.Error()))
			return
		}
		if req.Nombre == "" || req.Apellido == "" || req.Email == "" || req.Edad < 0 || req.Altura <= 0 || req.Activo == nil {
			ctx.JSON(404, web.NewResponse(404, nil, "uno o mas atributos son invalidos"))
			return
		}
		usuario := domain.Usuario{
			Id:             -1,
			Nombre:         req.Nombre,
			Apellido:       req.Apellido,
			Email:          req.Email,
			Edad:           req.Edad,
			Altura:         req.Altura,
			Activo:         *req.Activo,
			Fecha_creacion: "",
		}
		usuario, err := u.service.Save(usuario)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "problema en funcion save : "+err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, usuario, ""))
		return
	}
}

func (u *Usuario) ActualizarUsuario() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(404, "invalid id")
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		usuario, err := u.service.UpdateUsuario(ctx, req.Nombre, req.Apellido, req.Email, req.Fecha_creacion, id, req.Edad, req.Altura, req.Activo)
		if err != nil {
			ctx.JSON(404, err.Error())
			return
		}
		ctx.JSON(200, usuario)
	}
}

func (u *Usuario) ActualizarAtribUsuario() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(404, "invalid id")
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		usuario, err := u.service.UpdateAtributos(req.Nombre, req.Apellido, req.Email, req.Fecha_creacion, id, req.Edad, req.Altura, req.Activo)
		if err != nil {
			ctx.JSON(404, err.Error())
			return
		}
		ctx.JSON(200, usuario)
	}
}

func (u *Usuario) BorrarUsuario() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(404, "invalid id")
			return
		}
		err = u.service.DeleteUsuario(id)
		if err != nil {
			fmt.Println("aqui paso algo " + err.Error())
			ctx.JSON(404, err.Error())
			return
		} else {
			ctx.JSON(200, fmt.Sprint("se borro el usuario de id ", id))
			return
		}

	}
}
