package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	users "ejercicioTT/internal/users"

	"github.com/gin-gonic/gin"
)

type request struct {
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

func (u *Usuarios) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		//Ejercicio 1
		//Agregando variable de entorno y
		//accediendo a su valor
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Id inválido"})
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if req.Nombre == "" {
			c.JSON(400, gin.H{"error": "El nombre del usuario es requerido"})
			return
		}

		if req.Apellido == "" {
			c.JSON(400, gin.H{"error": "El apellido del usuario es requerido"})
			return
		}

		if req.Email == "" {
			c.JSON(400, gin.H{"error": "El email del usuario es requerido"})
			return
		}

		if req.Edad == 0 {
			c.JSON(400, gin.H{"error": "La edad del usuario es requerida"})
			return
		}

		if req.Altura == 0.0 {
			c.JSON(400, gin.H{"error": "La altura del usuario es requerida"})
			return
		}

		u, err := u.service.Update(int(id), req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.Fecha)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, u)
	}
}

func (u *Usuarios) UpdateLastAge() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		//Ejercicio 1
		//Agregando variable de entorno y
		//accediendo a su valor
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Id inválido"})
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if req.Apellido == "" {
			c.JSON(400, gin.H{"error": "El apellido del usuario, es requerido"})
			return
		}

		if req.Edad == 0 {
			c.JSON(400, gin.H{"error": "La edad del usuario es requerido"})
			return
		}

		u, err := u.service.UpdateLastAge(int(id), req.Apellido, req.Edad)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, u)
	}
}

func (u *Usuarios) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		//Ejercicio 1
		//Agregando variable de entorno y
		//accediendo a su valor
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Id inválido"})
			return
		}

		err = u.service.Delete(int(id))
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": fmt.Sprintf("El usuario %d ha sido eliminado", id)})
	}
}

func (u *Usuarios) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		//Ejercicio 1
		//Agregando variable de entorno y
		//accediendo a su valor
		if token != os.Getenv("TOKEN") {
			c.JSON(401, gin.H{"error": "Token inválido"})
			return
		}

		u, err := u.service.GetAll()
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, u)
	}
}

func (u *Usuarios) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		//Ejercicio 1
		//Agregando variable de entorno y
		//accediendo a su valor
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if req.Nombre == "" {
			c.JSON(400, gin.H{"error": "El nombre del usuario es requerido"})
			return
		}

		if req.Apellido == "" {
			c.JSON(400, gin.H{"error": "El apellido del usuario es requerido"})
			return
		}

		if req.Email == "" {
			c.JSON(400, gin.H{"error": "El email del usuario es requerido"})
			return
		}

		if req.Edad == 0 {
			c.JSON(400, gin.H{"error": "La edad del usuario es requerida"})
			return
		}

		if req.Altura == 0.0 {
			c.JSON(400, gin.H{"error": "La altura del usuario es requerida"})
			return
		}

		u, err := u.service.Store(req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.Fecha)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"usuarios": u})
	}
}
