package usuarios

import (
	"fmt"
	"testing"

	"github.com/del_rio/web-server/internal/domain"
	"github.com/stretchr/testify/assert"
)

type StubStore struct {
}

func (fs *StubStore) Write(data interface{}) error {
	return nil
}
func (fs *StubStore) Read(data interface{}) error {
	testInfo := []domain.Usuario{
		{
			Id:             1,
			Nombre:         "andrea1.7",
			Apellido:       "esquapovel",
			Email:          "an@gmail.com",
			Edad:           100,
			Altura:         153,
			Activo:         false,
			Fecha_creacion: "2022-07-12 02:35:30.982275 -0400 -04 m=+4.639375710",
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
	(*data.(*[]domain.Usuario)) = testInfo
	return nil
}

func TestGetAll(t *testing.T) {
	fmt.Println("pase")
	stubStore := StubStore{}
	repo := NewRepository(&stubStore)
	ExpectedList := []domain.Usuario{
		{
			Id:             1,
			Nombre:         "andrea1.7",
			Apellido:       "esquapovel",
			Email:          "an@gmail.com",
			Edad:           100,
			Altura:         153,
			Activo:         false,
			Fecha_creacion: "2022-07-12 02:35:30.982275 -0400 -04 m=+4.639375710",
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
	result, err := repo.GetAll()
	assert.Nil(t, err, "el read siempre es nil asi que no se deveria generar ningun error")
	assert.Equal(t, ExpectedList, result, "deben ser iguales")
}
