package handler

import (
	"C2-TT/internal/usuarios"
	"net/http"

	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre        string `json:"nombre"`
	Apellido      string `json:"apellido"`
	Email         string `json:"email"`
	Edad          int    `json:"edad"`
	Altura        int    `json:"altura"`
	Activo        bool   `json:"activo"`
	FechaCreacion string `json:"fecha_creacion"`
}

type Usuario struct {
	service usuarios.Service
}

func NewUsuario(s usuarios.Service) *Usuario {
	return &Usuario{service: s}
}

func (u *Usuario) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"}) //401
			return
		}

		usuarios, err := u.service.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "token inválido"}) // 500
			return
		}

		if len(usuarios) <= 0 {
			c.JSON(http.StatusOK, gin.H{"message": "No se encontraron transacciones."}) // 500
			return
		}

		c.JSON(http.StatusOK, gin.H{"usuarios": usuarios})
	}
}

func (u *Usuario) Registrar() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"}) //401
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) //400
			return
		}

		usuario, err := u.service.Registrar(req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.FechaCreacion)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) //400
			return
		}

		c.JSON(http.StatusOK, gin.H{"usuario": usuario})
	}
}
