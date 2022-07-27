package usuarios

import (
	"Clase3-1/internal/domain"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceIntegrationGetAll(t *testing.T) {
	//arrange
	database := []domain.Usuario{
		{Id: 1, Nombre: "prueba1", Apellido: "BeforeUpdate", Email: "prueba1Email", Edad: 25, Altura: 180, Activo: true, FechaDeCreacion: "29/10/2004"},
		{Id: 2, Nombre: "prueba3", Apellido: "BeforeUpdate", Email: "prueba1Email", Edad: 25, Altura: 180, Activo: true, FechaDeCreacion: "29/10/2004"},
	}
	myMockStore := MockStore{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}
	repo := NewRepository(&myMockStore)
	servicio := NewService(repo)

	usuariosEsperados := database

	//act
	resultado, err := servicio.GetAll()

	//assert
	assert.True(t, myMockStore.ReadWasCalled)
	assert.Nil(t, err)
	assert.Equal(t, resultado, usuariosEsperados)
}

func TestServiceIntegrationGetAllFail(t *testing.T) {
	// arrange
	database := []domain.Usuario{
		{Id: 1, Nombre: "prueba1", Apellido: "BeforeUpdate", Email: "prueba1Email", Edad: 25, Altura: 180, Activo: true, FechaDeCreacion: "29/10/2004"},
		//	{Id: 2, Nombre: "prueba2", Apellido: "BeforeUpdate", Email: "prueba1Email", Edad: 25, Altura: 180, Activo: true, FechaDeCreacion: "29/10/2004"},
	}
	expectedError := fmt.Errorf("cant read database")
	mockStorage := MockStore{
		dataMock: database,
		errWrite: "",
		errRead:  "esto es un error del store",
	}
	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	_, err := service.GetAll()
	// assert
	assert.ErrorContains(t, err, expectedError.Error())
}
func TestServiceIntegrationUpdate(t *testing.T) {
	//arrange
	database := []domain.Usuario{
		{Id: 1, Nombre: "prueba1", Apellido: "BeforeUpdate", Email: "prueba1Email", Edad: 25, Altura: 180, Activo: true, FechaDeCreacion: "29/10/2004"},
		{Id: 2, Nombre: "prueba3", Apellido: "BeforeUpdate", Email: "prueba1Email", Edad: 25, Altura: 180, Activo: true, FechaDeCreacion: "29/10/2004"},
	}
	myMockStore := MockStore{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}
	repo := NewRepository(&myMockStore)
	servicio := NewService(repo)
	id := 1
	name := "Jose"
	surname := "Olivera"
	mail := "joli@gmail.com"
	edad := 15
	alt := 195
	activo := false
	fechaCreacion := "29/10/2004"

	usuarioEsperado := domain.Usuario{
		Id:              id,
		Nombre:          name,
		Apellido:        surname,
		Email:           mail,
		Edad:            edad,
		Altura:          alt,
		Activo:          activo,
		FechaDeCreacion: fechaCreacion,
	}

	//act
	resultado, err := servicio.Update(id, name, surname, mail, edad, alt, activo, fechaCreacion)

	//assert
	assert.True(t, myMockStore.ReadWasCalled)
	assert.Nil(t, err)
	assert.Equal(t, resultado, usuarioEsperado)
	assert.Equal(t, resultado, myMockStore.dataMock[0])
}

func TestServiceIntegrationDelete(t *testing.T) {
	//arrange
	database := []domain.Usuario{
		{Id: 1, Nombre: "prueba1", Apellido: "BeforeUpdate", Email: "prueba1Email", Edad: 25, Altura: 180, Activo: true, FechaDeCreacion: "29/10/2004"},
		//	{Id: 2, Nombre: "prueba2", Apellido: "BeforeUpdate", Email: "prueba1Email", Edad: 25, Altura: 180, Activo: true, FechaDeCreacion: "29/10/2004"},
	}
	myMockStore := MockStore{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}
	repo := NewRepository(&myMockStore)
	servicio := NewService(repo)

	//act
	err := servicio.Delete(1)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, len(myMockStore.dataMock), len(database)-1)
	//fmt.Printf("%v", myMockStore.dataMock)
}

func TestServiceIntegrationDeleteFail(t *testing.T) {
	// arrange
	database := []domain.Usuario{
		{Id: 1, Nombre: "prueba1", Apellido: "BeforeUpdate", Email: "prueba1Email", Edad: 25, Altura: 180, Activo: true, FechaDeCreacion: "29/10/2004"},
		//	{Id: 2, Nombre: "prueba2", Apellido: "BeforeUpdate", Email: "prueba1Email", Edad: 25, Altura: 180, Activo: true, FechaDeCreacion: "29/10/2004"},
	}
	idInexistente := 2
	expectedError := fmt.Errorf("user %v not found", idInexistente)
	mockStorage := MockStore{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}
	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	err := service.Delete(idInexistente)
	// assert
	assert.ErrorContains(t, err, expectedError.Error())
}

func TestServiceIntegrationStore(t *testing.T) {
	//arrange
	listaUs := []domain.Usuario{
		{Id: 1, Nombre: "prueba1", Apellido: "BeforeUpdate", Email: "prueba1Email", Edad: 25, Altura: 180, Activo: true, FechaDeCreacion: "29/10/2004"},
		{Id: 2, Nombre: "prueba3", Apellido: "BeforeUpdate", Email: "prueba1Email", Edad: 25, Altura: 180, Activo: true, FechaDeCreacion: "29/10/2004"},
	}
	database := listaUs
	myMockStore := MockStore{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}
	repo := NewRepository(&myMockStore)
	servicio := NewService(repo)
	name := "Jose"
	surname := "Olivera"
	mail := "joli@gmail.com"
	edad := 15
	alt := 195
	activo := false
	fechaCreacion := "29/10/2004"

	usuarioEsperado := domain.Usuario{
		Id:              3,
		Nombre:          name,
		Apellido:        surname,
		Email:           mail,
		Edad:            edad,
		Altura:          alt,
		Activo:          activo,
		FechaDeCreacion: fechaCreacion,
	}
	//act
	resultado, err := servicio.Store(name, surname, mail, edad, alt, activo, fechaCreacion)

	//assert
	assert.True(t, myMockStore.ReadWasCalled)
	assert.Nil(t, err)
	assert.Equal(t, resultado, usuarioEsperado)
	assert.Equal(t, len(listaUs)+1, len(myMockStore.dataMock))
}

func TestServiceIntegrationStoreFail(t *testing.T) {
	// arrange
	database := []domain.Usuario{
		{Id: 1, Nombre: "prueba1", Apellido: "BeforeUpdate", Email: "prueba1Email", Edad: 25, Altura: 180, Activo: true, FechaDeCreacion: "29/10/2004"},
		//	{Id: 2, Nombre: "prueba2", Apellido: "BeforeUpdate", Email: "prueba1Email", Edad: 25, Altura: 180, Activo: true, FechaDeCreacion: "29/10/2004"},
	}
	expectedError := fmt.Errorf("cant write database")
	mockStorage := MockStore{
		dataMock: database,
		errWrite: "esto es un error del store",
		errRead:  "",
	}
	// act
	name := "Jose"
	surname := "Olivera"
	mail := "joli@gmail.com"
	edad := 15
	alt := 195
	activo := false
	fechaCreacion := "29/10/2004"
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	_, err := service.Store(name, surname, mail, edad, alt, activo, fechaCreacion)
	// assert
	assert.ErrorContains(t, err, expectedError.Error())
}
