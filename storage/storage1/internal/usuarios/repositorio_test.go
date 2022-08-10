package usuarios

import (
	"fmt"
	"testing"

	"github.com/del_rio/web-server/db"
	"github.com/del_rio/web-server/internal/domain"
	"github.com/stretchr/testify/assert"
)

type StubStore struct {
	db         []domain.Usuario
	readUsed   bool
	forceError error
}

func (fs *StubStore) Write(data interface{}) error {
	if fs.forceError != nil {
		return fs.forceError
	}
	fs.db = data.([]domain.Usuario)
	return nil
}
func (fs *StubStore) Read(data interface{}) error {
	if fs.forceError != nil {
		return fs.forceError
	}
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
	fs.readUsed = true
	return nil
}
func InitTest() (*StubStore, *Repository, []domain.Usuario) {
	stubStore := StubStore{}
	dbStore := db.InitDb()
	repo := NewRepository(&stubStore, dbStore)
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
	return &stubStore, &repo, ExpectedList
}
func ClasicTestReview(t *testing.T, ExpectedUsuario, result domain.Usuario, ExpectedList []domain.Usuario, realList []domain.Usuario, err error) {
	assert.Nil(t, err, "el read siempre es nil asi que no se deveria generar ningun error")
	assert.Equal(t, ExpectedUsuario, result, "deben ser iguales")
	assert.Equal(t, ExpectedList, realList, "deben ser iguales")
}
func TestGetAll(t *testing.T) {
	fmt.Println("pase")
	stubStore := StubStore{}
	dbStore := db.InitDb()
	repo := NewRepository(&stubStore, dbStore)
	result, err := repo.GetAll()
	assert.Nil(t, err, "el read siempre es nil asi que no se deveria generar ningun error")
	assert.NotZero(t, result, "deveria dar algo diferente a una lista vacia")
}
func TestUpdateAtribute(t *testing.T) {
	fmt.Println("entro aqui tmb")
	stubStore := StubStore{}
	dbStore := db.InitDb()
	repo := NewRepository(&stubStore, dbStore)
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
			Nombre:         "alicia",
			Apellido:       "van persie",
			Email:          "vp@gmail.com",
			Edad:           17,
			Altura:         175,
			Activo:         true,
			Fecha_creacion: "2022-07-25 02:35:30.982275 -0400 -04 m=+4.639375710",
		},
	}
	result, err := repo.UpdateAtributos("alicia", "", "", "", 2, 17, 175, nil)
	fmt.Println(result)
	assert.Nil(t, err, "el read siempre es nil asi que no se deveria generar ningun error")
	assert.Equal(t, ExpectedList[1], result, "deben ser iguales")
	assert.Equal(t, ExpectedList, stubStore.db, "deben ser iguales")
	assert.True(t, stubStore.readUsed, "la funcion read tiene que ser llamada para este test")
}

func TestStoreHappyPath(t *testing.T) {
	_, pointRepo, _ := InitTest()
	newName := "daniel"
	newSecondName := "van de vans"
	newEmail := "dvdv@gmail.com"
	newAge := 2
	newHeight := 130
	newActivition := true
	newCreationDate := "2025-07-25 02:35:30"
	usuario, err := (*pointRepo).Store(domain.Usuario{
		Id:             0,
		Nombre:         newName,
		Apellido:       newSecondName,
		Email:          newEmail,
		Edad:           newAge,
		Altura:         newHeight,
		Activo:         newActivition,
		Fecha_creacion: newCreationDate,
	})
	assert.Nil(t, err, "deberia ser nulo")
	expectUsuario, _ := (*pointRepo).GetById(usuario.Id)
	assert.Equal(t, expectUsuario, usuario, "ambos usuarios deberian ser iguales")
}
func TestGetByNameHappyPath(t *testing.T) {
	_, pointRepo, _ := InitTest()
	newName := "Margaret"
	newSecondName := "van de vans"
	newEmail := "dvdv@gmail.com"
	newAge := 2
	newHeight := 130
	newActivition := true
	newCreationDate := "2025-07-25 02:35:30"
	_, err := (*pointRepo).Store(domain.Usuario{
		Id:             0,
		Nombre:         newName,
		Apellido:       newSecondName,
		Email:          newEmail,
		Edad:           newAge,
		Altura:         newHeight,
		Activo:         newActivition,
		Fecha_creacion: newCreationDate,
	})
	assert.Nil(t, err, "la creacion de margaret no deberia tener problemas")
	usuario := (*pointRepo).GetByName("Margaret")
	assert.Equal(t, newName, usuario.Nombre, "ambos usuarios deberian tener el mismo nombre")
}
