package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	users "C3ejercicioTT/internal/users"
	"C3ejercicioTT/pkg/web"

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

		u, err := u.service.Update(int(id), req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.Fecha)
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		c.JSON(200, web.NewResponse(200, u, ""))
	}
}

// Patch Users godoc
// @Summary Patch a user by Id
// @Tags Usuarios
// @ID update-user-by-id
// @Accept  json
// @Produce json
// @Param token header string true "token"
// @Param id path string true "id from user to patch"
// @Param usuario body request true "User to patch"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 409 {object} web.Response
// @Router /usuarios/{id} [patch]
func (u *Usuarios) UpdateLastAge() gin.HandlerFunc {
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

		if req.Apellido == "" {
			c.JSON(400, web.NewResponse(http.StatusUnauthorized, nil, "error, el apellido del usuario, es requerido"))
			return
		}

		if req.Edad == 0 {
			c.JSON(400, web.NewResponse(http.StatusUnauthorized, nil, "error, la edad del usuario, es requerido"))
			return
		}

		u, err := u.service.UpdateLastAge(int(id), req.Apellido, req.Edad)
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		c.JSON(200, web.NewResponse(200, u, ""))
	}
}

// DeleteUsers godoc
// @Summary Delete a user by Id
// @Tags Usuarios
// @ID delete-user-by-id
// @Produce json
// @Param token header string true "token"
// @Param id path string true "id from user to delete"
// @Success 200 {object} web.Response
// @Failure 404 {object} web.Response
// @Router /usuarios/{id} [delete]
func (u *Usuarios) Delete() gin.HandlerFunc {
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

		err = u.service.Delete(int(id))
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		c.JSON(200, web.NewResponse(200, gin.H{"data": fmt.Sprintf("El usuario %d ha sido eliminado", id)}, ""))
	}
}

// ListUsers godoc
// @Summary List users
// @Tags Usuarios
// @Description get usuarios
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /usuarios [get]
func (u *Usuarios) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "error, Token inválido"))
			return
		}

		u, err := u.service.GetAll()
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
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
			log.Println(err, req)
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

		u, err := u.service.Store(req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.Fecha)
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		c.JSON(200, web.NewResponse(200, u, ""))
	}
}
