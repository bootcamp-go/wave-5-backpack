package handler

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"ejercicioTT/internal/domain"
	users "ejercicioTT/internal/users"
	"ejercicioTT/pkg/web"

	"github.com/gin-gonic/gin"
)

//request que envío a cada JSON en swagger
type request struct {
	Id       int       `json:"id"`
	Nombre   string    `json:"nombre"`
	Apellido string    `json:"apellido"`
	Email    string    `json:"email"`
	Edad     int       `json:"edad"`
	Altura   float64   `json:"altura"`
	Activo   bool      `json:"activo"`
	Fecha    time.Time `json:"fecha"`
}

type Usuarios struct {
	service users.Service
}

func NewUser(s users.Service) *Usuarios {
	return &Usuarios{service: s}
}

// GetByName godoc
// @Summary Get a user by Name
// @Tags Usuarios
// @Produce json
// @Param token header string true "token"
// @Param nombre path string true "nombre"
// @Success 200 {object} web.Response
// @Failure 404 {object} web.Response
// @Router /usuarios/byName/{nombre} [get]
func (u *Usuarios) GetByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		usuario, err := u.service.GetByName(c.Param("nombre"))
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, "El usuario con el nombre ingresado no fue encontrado"))
			return
		}
		c.JSON(200, web.NewResponse(200, usuario, ""))
	}
}

// StoreUsers godoc
// @Summary Store usuarios
// @Tags Usuarios
// @Description store usuarios
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param usuario body request true "User to store"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 409 {object} web.Response
// @Router /usuarios [post]
func (u *Usuarios) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "error, Token inválido"))
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		//Nombre Apellido Email Edad Altura
		if req.Nombre == "" {
			c.JSON(400, web.NewResponse(400, nil, "error: el nombre del usuario es requerido"))
			return
		}

		if req.Apellido == "" {
			c.JSON(400, web.NewResponse(400, nil, "error: el apellido del usuario es requerido"))
			return
		}

		if req.Email == "" {
			c.JSON(400, web.NewResponse(400, nil, "error: el email del usuario es requerido"))
			return
		}

		if req.Edad == 0 {
			c.JSON(400, web.NewResponse(400, nil, "error: la edad del usuario es requerido"))
			return
		}

		if req.Altura == 0.0 {
			c.JSON(400, web.NewResponse(400, nil, "error: la altura del usuario es requerido"))
			return
		}

		user := domain.Usuarios(req)
		u, err := u.service.Store(user)
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		c.JSON(200, web.NewResponse(200, u, ""))
	}
}

// GetOne godoc
// @Summary Get a user by id
// @Tags Usuarios
// @Produce json
// @ID get-user-by-id
// @Param token header string true "token"
// @Param id path string true "user id"
// @Success 200 {object} web.Response
// @Failure 404 {object} web.Response
// @Router /usuarios/{id} [get]
func (u *Usuarios) GetOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, "error, Id inválido"))
			return
		}
		usuario, err := u.service.GetOne(id)
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, "El usuario con el id ingresado no fue encontrado"))
			return
		}
		c.JSON(200, web.NewResponse(200, usuario, ""))
	}
}

// Update Users godoc
// @Summary Update a user by Id
// @Tags Usuarios
// @Accept  json
// @Produce json
// @Param token header string true "token"
// @Param usuario body request true "User to update"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 409 {object} web.Response
// @Router /usuarios [put]
func (u *Usuarios) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "error, Token inválido"))
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		//Nombre Apellido Email Edad Altura
		if req.Nombre == "" {
			c.JSON(400, web.NewResponse(400, nil, "error: el nombre del usuario es requerido"))
			return
		}

		if req.Apellido == "" {
			c.JSON(400, web.NewResponse(400, nil, "error: el apellido del usuario es requerido"))
			return
		}

		if req.Email == "" {
			c.JSON(400, web.NewResponse(400, nil, "error: el email del usuario es requerido"))
			return
		}

		if req.Edad == 0 {
			c.JSON(400, web.NewResponse(400, nil, "error: la edad del usuario es requerido"))
			return
		}

		if req.Altura == 0.0 {
			c.JSON(400, web.NewResponse(400, nil, "error: la altura del usuario es requerido"))
			return
		}

		user := domain.Usuarios(req)
		u, err := u.service.Update(user)
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		c.JSON(200, web.NewResponse(200, u, ""))
	}
}
