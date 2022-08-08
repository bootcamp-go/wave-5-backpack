package handler

import (
	"net/http"
	"os"
	"strconv"
	"time"

	users "ejercicioTT/internal/users"
	"ejercicioTT/pkg/web"

	"github.com/gin-gonic/gin"
)

//request que envío a cada JSON en swagger
type request struct {
	Nombre   string    `json:"nombre"`
	Apellido string    `json:"apellido"`
	Email    string    `json:"email"`
	Edad     int       `json:"edad"`
	Altura   float64   `json:"altura"`
	Activo   bool      `json:"activo"`
	Fecha    time.Time `json:"fecha" example:"2022-07-12T00:53:16.535668Z" format:"date-time"`
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
// @ID get-user-by-name
// @Produce json
// @Param token header string true "token"
// @Param nombre path string true "name from user to get"
// @Success 200 {object} web.Response
// @Failure 404 {object} web.Response
// @Router /usuarios/{id} [get]
func (c *UsuariosBD) GetByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		u, erro := c.service.GetByName(c.Param("nombre"))
		if erro != nil {
			c.JSON(404, web.NewResponse(404, nil, "El usuario con el nombre ingresado no fue encontrado"))
			return
		}
		c.JSON(200, web.NewResponse(200, u, ""))
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

		u, err := u.service.Store(req)
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		c.JSON(200, web.NewResponse(200, u, ""))
	}
}

// Update Users godoc
// @Summary Update a user by Id
// @Tags Usuarios
// @ID update-user-by-id
// @Accept  json
// @Produce json
// @Param token header string true "token"
// @Param id path string true "id from user to update"
// @Param usuario body request true "User to update"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 409 {object} web.Response
// @Router /usuarios/{id} [put]
func (u *Usuarios) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "error, Token inválido"))
			return
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, "error, Id inválido"))
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

		u, err := u.service.Update(req)
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		c.JSON(200, web.NewResponse(200, u, ""))
	}
}
