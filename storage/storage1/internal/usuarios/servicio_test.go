package usuarios

import (
	"fmt"
	"testing"

	"github.com/del_rio/web-server/db"
	"github.com/del_rio/web-server/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestUpdateUsuario(t *testing.T) {
	fmt.Println("aqui voy señor")
	stubStore := StubStore{}
	dbStore := db.InitDb()
	repo := NewRepository(&stubStore, dbStore)
	servi := NewService(repo)
	ExpectedList := []domain.Usuario{
		{
			Id:             1,
			Nombre:         "angelica",
			Apellido:       "lover",
			Email:          "al@gmail.com",
			Edad:           24,
			Altura:         163,
			Activo:         false,
			Fecha_creacion: "2020-07-12 02:35:30.982275 -0400 -04 m=+4.639375710",
		},
		{
			Id:             2,
			Nombre:         "pedro.7",
			Apellido:       "van persie",
			Email:          "vp@gmail.com",
			Edad:           12,
			Altura:         189,
			Activo:         true,
			Fecha_creacion: "2022-07-25 02:35:30.982275 -0400 -04 m=+4.639375710",
		},
	}
	result, err := servi.UpdateUsuario(
		"angelica",
		"lover",
		"al@gmail.com",
		"2020-07-12 02:35:30.982275 -0400 -04 m=+4.639375710",
		1,
		24,
		163,
		false,
	)
	assert.Nil(t, err, "el read siempre es nil asi que no se deveria generar ningun error")
	assert.Equal(t, ExpectedList[0], result, "deben ser iguales")
	assert.Equal(t, ExpectedList, stubStore.db, "deben ser iguales")
	assert.True(t, stubStore.readUsed, "la funcion read tiene que ser llamada para este test")
}
func TestDeleteUsuario(t *testing.T) {
	fmt.Println("me encuentro ante ti")
	stubStore := StubStore{}
	dbStore := db.InitDb()
	repo := NewRepository(&stubStore, dbStore)
	servi := NewService(repo)
	ExpectedList := []domain.Usuario{
		{
			Id:             2,
			Nombre:         "pedro.7",
			Apellido:       "van persie",
			Email:          "vp@gmail.com",
			Edad:           12,
			Altura:         189,
			Activo:         true,
			Fecha_creacion: "2022-07-25 02:35:30.982275 -0400 -04 m=+4.639375710",
		},
	}
	err := servi.DeleteUsuario(0)
	assert.NotNil(t, err, "deveria haber un error porque ese id no existe")
	notError := servi.DeleteUsuario(1)
	assert.Nil(t, notError, "no deveria haber error pues el id ingresado si existe")
	assert.Equal(t, ExpectedList, stubStore.db, "estas 2 listas deverian ser iguales")
}
