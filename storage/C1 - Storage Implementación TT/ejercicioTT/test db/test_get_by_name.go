package test

import (
	"ejercicioTT/internal/domain"
	"log"
	"testing"

	"github.com/go-playground/assert"
)

func TestGetByName(t *testing.T) {
	usuario := domain.Usuarios{
		Nombre: "test",
	}
	myRepo := NewRepo()
	usuarioResult, err := myRepo.GetByName(usuario)
	if err != nil {
		log.Println(err)
	}
	assert.Equal(t, usuario.Nombre, usuarioResult.Nombre)
}
