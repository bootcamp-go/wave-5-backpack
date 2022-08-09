package usuarios

import (
	"context"
	"fmt"
	"log"
	"storage/Clase2-TM/db"
	"storage/Clase2-TM/internal/domain"
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

func TestGetAll(t *testing.T) {
	database := db.StorageDB
	myRepo := NewRepo(database)
	usersResult, err := myRepo.GetAll()
	if err != nil {
		log.Println(err)
	}
	assert.Nil(t, err)
	assert.NotNil(t, usersResult)
	fmt.Println(usersResult)
}
func TestDelete(t *testing.T) {
	database := db.StorageDB
	myRepo := NewRepo(database)
	result := myRepo.Delete(1)
	assert.Nil(t, result)
}
func TestUpdate(t *testing.T) {
	ctx := context.Background()
	usuarioEsperado := domain.Usuario{
		Id:       4,
		Nombre:   "Pedro",
		Apellido: "Rodriguez",
		Edad:     29,
		Altura:   180,
	}
	database := db.StorageDB
	myRepo := NewRepo(database)
	userResult, err := myRepo.Update(ctx, usuarioEsperado)
	if err != nil {
		log.Println(err)
	}
	assert.Equal(t, usuarioEsperado.Nombre, userResult.Nombre)
	assert.Equal(t, usuarioEsperado.Edad, userResult.Edad)

}
