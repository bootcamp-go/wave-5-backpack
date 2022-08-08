package usuarios

import (
	"errors"
	"fmt"
	"testing"

	"github.com/anesquivel/wave-5-backpack/storage/arquitectura_ejercicio/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestIntegracionGetAll(t *testing.T) {
	// arrange
	database := []domain.Usuario{
		{
			Id:          1,
			Names:       "Andy",
			LastName:    "Esquivel",
			Age:         23,
			Estatura:    1.52,
			Email:       "andy@gmail.com",
			IsActivo:    true,
			DateCreated: "25-07-2022",
		},
		{
			Id:          2,
			Names:       "Gabriela",
			LastName:    "Rueda",
			Age:         23,
			Estatura:    1.52,
			Email:       "gr@gmail.com",
			IsActivo:    true,
			DateCreated: "26-07-2022",
		},
	}

	mockStorage := MockStorage{
		dataMock: database,
		errRead:  "",
		errWrite: "",
	}

	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	res, err := service.GetAll()

	// assert

	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataMock, res)
}

func TestIntegracionGetAllFail(t *testing.T) {
	// arrange
	database := []domain.Usuario{
		{
			Id:          1,
			Names:       "Andy",
			LastName:    "Esquivel",
			Age:         23,
			Estatura:    1.52,
			Email:       "andy@gmail.com",
			IsActivo:    true,
			DateCreated: "25-07-2022",
		},
		{
			Id:          2,
			Names:       "Gabriela",
			LastName:    "Rueda",
			Age:         23,
			Estatura:    1.52,
			Email:       "gr@gmail.com",
			IsActivo:    true,
			DateCreated: "26-07-2022",
		},
	}

	mockStorage := MockStorage{
		dataMock: database,
		errRead:  "",
		errWrite: "",
	}

	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	res, err := service.GetAll()
	//errExpected := fmt.Errorf("can't read database")
	// assert

	assert.Nil(t, err)
	//assert.Equal(t, errExpected, err)
	assert.Equal(t, mockStorage.dataMock, res)
}

func TestIntegracionStore(t *testing.T) {
	newUser := domain.Usuario{
		Id:          0,
		Names:       "Ashton",
		LastName:    "Brooke",
		Email:       "ash2@gmail.com",
		Estatura:    1.80,
		IsActivo:    true,
		DateCreated: "26-07-2022",
		Age:         28,
	}

	mockStorage := MockStorage{
		dataMock: []domain.Usuario{},
		errRead:  "",
		errWrite: "",
	}

	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	res, err := service.Store(newUser.Age, newUser.Names, newUser.LastName,
		newUser.Email, newUser.Estatura)

	// assert

	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataMock[0], res)
	assert.Equal(t, mockStorage.dataMock[0].Id, res.Id)

}

func TestIntegracionStoreFail(t *testing.T) {
	newUser := domain.Usuario{
		Id:          0,
		Names:       "Ashton",
		LastName:    "Brooke",
		Email:       "ash2@gmail.com",
		Estatura:    1.80,
		IsActivo:    true,
		DateCreated: "26-07-2022",
		Age:         28,
	}

	mockStorage := MockStorage{
		dataMock: []domain.Usuario{},
		errRead:  "",
		errWrite: "can't write to the database",
	}

	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	errExpected := fmt.Errorf("can't write to the databse")
	_, err := service.Store(newUser.Age, newUser.Names, newUser.LastName,
		newUser.Email, newUser.Estatura)

	// assert
	assert.NotNil(t, err)
	assert.Equal(t, errExpected, err)

}

func TestIntegracionUpdate(t *testing.T) {
	newUser := domain.Usuario{
		Id:          0,
		Names:       "ASH",
		Email:       "ash2@gmail.com",
		Estatura:    1.83,
		DateCreated: "26-07-2022",
	}

	mockStorage := MockStorage{
		dataMock: []domain.Usuario{newUser},
		errRead:  "",
		errWrite: "",
	}

	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	res, err := service.Update(newUser.Id, newUser.Age, newUser.Names, newUser.LastName,
		newUser.Email, newUser.DateCreated, newUser.Estatura, newUser.IsActivo)

	// assert

	assert.Nil(t, err)
	assert.Equal(t, true, mockStorage.hadCalledRead)
	assert.Equal(t, mockStorage.dataMock[0], res)
	assert.Equal(t, mockStorage.dataMock[0].Id, res.Id)

}

func TestIntegracionUpdateFail(t *testing.T) {
	newUser := domain.Usuario{
		Id:    0,
		Names: "ASH",
	}

	mockStorage := MockStorage{
		dataMock: []domain.Usuario{newUser},
		errRead:  "",
		errWrite: "",
	}

	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	res, err := service.Update(5, newUser.Age, newUser.Names, newUser.LastName,
		newUser.Email, newUser.DateCreated, newUser.Estatura, newUser.IsActivo)

	// assert

	assert.Nil(t, err)
	assert.Equal(t, true, mockStorage.hadCalledRead)
	assert.Equal(t, mockStorage.dataMock[0], res)
	assert.Equal(t, mockStorage.dataMock[0].Id, res.Id)

}

func TestIntegracionDelete(t *testing.T) {
	newUser := domain.Usuario{
		Id:          0,
		Names:       "ASH",
		Email:       "ash2@gmail.com",
		Estatura:    1.83,
		DateCreated: "26-07-2022",
	}

	mockStorage := MockStorage{
		dataMock: []domain.Usuario{newUser},
		errRead:  "",
		errWrite: "",
	}

	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	err := service.Delete(newUser.Id)

	// assert
	assert.Equal(t, true, mockStorage.hadCalledRead)
	assert.Nil(t, err)
}

func TestIntegracionDeleteFails(t *testing.T) {
	newUser := domain.Usuario{
		Id:          0,
		Names:       "ASH",
		Email:       "ash2@gmail.com",
		Estatura:    1.83,
		DateCreated: "26-07-2022",
	}

	mockStorage := MockStorage{
		dataMock: []domain.Usuario{newUser},
		errRead:  "",
		errWrite: "",
	}

	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	errExpected := errors.New("No se encontró el usuario a eliminar.")
	err := service.Delete(100)

	// assert
	assert.Equal(t, true, mockStorage.hadCalledRead)
	assert.NotNil(t, err)
	assert.Equal(t, errExpected, err)
}

func TestIntegracionPatch(t *testing.T) {
	user := domain.Usuario{
		Id:          0,
		Names:       "Ashton",
		LastName:    "Irwin",
		Email:       "ash2@gmail.com",
		Age:         34,
		Estatura:    1.88,
		DateCreated: "26-07-2022",
		IsActivo:    true,
	}

	dataUserToUpd := domain.Usuario{
		Id:       0,
		LastName: "Brooke",
		Age:      34,
	}
	mockStorage := MockStorage{
		dataMock: []domain.Usuario{user},
		errRead:  "",
		errWrite: "",
	}

	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	res, err := service.UpdateLastNameAndAge(dataUserToUpd.Id, dataUserToUpd.Age, dataUserToUpd.LastName)
	// assert

	assert.Nil(t, err)
	assert.Equal(t, true, mockStorage.hadCalledRead)
	assert.Equal(t, mockStorage.dataMock[0].Age, res.Age)
	assert.Equal(t, mockStorage.dataMock[0].LastName, res.LastName)
	assert.Equal(t, mockStorage.dataMock[0].Id, res.Id)

}

func TestIntegracionPatchFails(t *testing.T) {
	user := domain.Usuario{
		Id:          0,
		Names:       "Ashton",
		LastName:    "Irwin",
		Email:       "ash2@gmail.com",
		Age:         34,
		Estatura:    1.88,
		DateCreated: "26-07-2022",
		IsActivo:    true,
	}

	dataUserToUpd := domain.Usuario{
		Id:       0,
		LastName: "Brooke",
		Age:      34,
	}
	mockStorage := MockStorage{
		dataMock: []domain.Usuario{user},
		errRead:  "",
		errWrite: "",
	}

	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	errExpected := errors.New("No se encontró el usuario a actualizar.")
	_, err := service.UpdateLastNameAndAge(600, dataUserToUpd.Age, dataUserToUpd.LastName)

	// assert
	assert.Equal(t, true, mockStorage.hadCalledRead)
	assert.NotNil(t, err)
	assert.Equal(t, errExpected, err)

}
