package db_test

import (
	"C1-TT/internal/domain"
	"C1-TT/internal/repositorio"
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	u := domain.User{
		FirstName: "testify",
		LastName:  "testing",
		Email:     "roko@dogmail.com",
		Age:       4,
		Height:    80.2,
		Activo:    true,
		CreatedAt: "11/05/2018",
	}
	myRepo := repositorio.NewRepository()
	userResult, err := myRepo.Store(u)
	if err != nil {
		log.Println(err)
	}
	assert.Equal(t, u.FirstName, userResult.FirstName)

}

func TestGetOne(t *testing.T) {
	expectedUser := domain.User{
		Id:        2,
		FirstName: "test",
		LastName:  "testing",
		Email:     "roko@dogmail.com",
		Age:       4,
		Height:    80.2,
		Activo:    true,
		CreatedAt: "11/05/2018",
	}
	myRepo := repositorio.NewRepository()
	actualUser, err := myRepo.GetOne(expectedUser.Id)
	if err != nil {
		log.Println(err)
	}
	assert.Nil(t, err)
	assert.Equal(t, expectedUser, actualUser)

}

func TestGetByName(t *testing.T) {
	expectedUser := domain.User{
		Id:        4,
		FirstName: "testify",
		LastName:  "testing",
		Email:     "roko@dogmail.com",
		Age:       4,
		Height:    80.2,
		Activo:    true,
		CreatedAt: "11/05/2018",
	}
	myRepo := repositorio.NewRepository()
	actualUser, err := myRepo.GetByName(expectedUser.FirstName)
	if err != nil {
		log.Println(err)
	}
	assert.Nil(t, err)
	assert.Equal(t, expectedUser, actualUser)

}

func TestGetAll(t *testing.T) {
	expectedLenUsers := 4
	myRepo := repositorio.NewRepository()
	actualUsers, err := myRepo.GetAll()
	if err != nil {
		log.Println(err)
	}
	actualLenUsers := len(actualUsers)
	assert.Nil(t, err)
	assert.Equal(t, expectedLenUsers, actualLenUsers)
}
func TestUpdate(t *testing.T) {
	expectedUser := domain.User{
		Id:        1,
		FirstName: "testifyUpdatex",
		LastName:  "testingUpdateX",
		Email:     "roko@dogmail.com",
		Age:       4,
		Height:    80.2,
		Activo:    true,
		CreatedAt: "11/05/2018",
	}
	myRepo := repositorio.NewRepository()

	actualUser, err := myRepo.Update(expectedUser)
	fmt.Println(err)

	assert.Nil(t, err)
	assert.Equal(t, expectedUser, actualUser)

}
func TestDelete(t *testing.T) {
	var userId2Delete = 2
	myRepo := repositorio.NewRepository()
	err := myRepo.Delete(userId2Delete)
	assert.Nil(t, err)
}

func TestGetOneWithContext(t *testing.T) {
	expectedUser := domain.User{
		Id:        4,
		FirstName: "testify",
		LastName:  "testing",
		Email:     "roko@dogmail.com",
		Age:       4,
		Height:    80.2,
		Activo:    true,
		CreatedAt: "11/05/2018",
	}
	myRepo := repositorio.NewRepository()
	//Definimos un context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	actualUser, err := myRepo.GetOneWithContext(ctx, expectedUser.Id)
	fmt.Println(err)

	assert.Nil(t, err)
	assert.Equal(t, expectedUser, actualUser)

}

func TestUpdateWithContext(t *testing.T) {
	expectedUser := domain.User{
		Id:        1,
		FirstName: "testifyUpdate",
		LastName:  "testingUpdate",
		Email:     "roko@dogmail.com",
		Age:       4,
		Height:    80.2,
		Activo:    true,
		CreatedAt: "11/05/2018",
	}
	myRepo := repositorio.NewRepository()
	//Definimos un context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	actualUser, err := myRepo.UpdateWithContext(ctx, expectedUser)
	fmt.Println(err)

	assert.Nil(t, err)
	assert.Equal(t, expectedUser, actualUser)

}
