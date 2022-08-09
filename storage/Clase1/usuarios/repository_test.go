package usuarios

import (
	"log"
	"storage/Clase1/db"
	"storage/Clase1/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	usuarioEsperado := domain.Usuario{
		Nombre:   "German",
		Apellido: "Rodriguez",
		Edad:     27,
		Altura:   180,
	}
	database := db.StorageDB
	myRepo := NewRepo(database)
	userResult, err := myRepo.Store(usuarioEsperado)
	if err != nil {
		log.Println(err)
	}
	assert.Equal(t, usuarioEsperado.Nombre, userResult.Nombre)
	assert.Equal(t, usuarioEsperado.Altura, userResult.Altura)

}

func TestGetByName(t *testing.T) {
	database := db.StorageDB
	myRepo := NewRepo(database)
	userResult, err := myRepo.GetByName("German")
	if err != nil {
		log.Println(err)
	}
	assert.Equal(t, "German", userResult.Nombre)
	assert.NotNil(t, userResult.Apellido)

}
