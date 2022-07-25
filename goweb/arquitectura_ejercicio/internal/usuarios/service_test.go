package usuarios

import (
	"testing"

	"github.com/anesquivel/wave-5-backpack/goweb/arquitectura_ejercicio/internal/domain"
	"github.com/stretchr/testify/assert"
)

type MockRepository struct {
	hadReadAll bool
	user       domain.Usuario
}

func (r *MockRepository) GetAll() ([]domain.Usuario, error) {
	var localUserList []domain.Usuario
	user1 := domain.Usuario{
		Id:          1,
		Names:       "Andy",
		LastName:    "Esquivel",
		Age:         23,
		Estatura:    1.52,
		Email:       "andy@gmail.com",
		IsActivo:    true,
		DateCreated: "",
	}

	user2 := domain.Usuario{
		Id:          2,
		Names:       "Andy",
		LastName:    "Esquivel",
		Age:         23,
		Estatura:    1.52,
		Email:       "andy@gmail.com",
		IsActivo:    true,
		DateCreated: "",
	}

	localUserList = append(localUserList, user1)
	localUserList = append(localUserList, user2)
	r.hadReadAll = true
	return localUserList, nil
}

func (r *MockRepository) Store(id, age int, names, lastname, email, dateCreated string, estatura float64) (domain.Usuario, error) {

	return domain.Usuario{}, nil
}

func (r *MockRepository) Update(id, age int, names, lastname, email, dateCreated string, estatura float64, activo bool) (domain.Usuario, error) {
	return domain.Usuario{}, nil
}

func (r *MockRepository) UpdateLastNameAndAge(id, age int, lastname string) (domain.Usuario, error) {
	_, err := r.GetAll()

	if err != nil {
		return domain.Usuario{}, err
	}
	r.user.Names = "After Update"
	return r.user, nil
}

func (r *MockRepository) LastID() (int, error) {
	return 1, nil
}

func (r *MockRepository) Delete(id int) error {
	return nil
}

func TestGetAll(t *testing.T) {
	//arrange

	myStubDB := MockRepository{}
	motor := NewService(&myStubDB)
	expectedRes := []domain.Usuario{
		{
			Id:          1,
			Names:       "Andy",
			LastName:    "Esquivel",
			Age:         23,
			Estatura:    1.52,
			Email:       "andy@gmail.com",
			IsActivo:    true,
			DateCreated: "",
		},
		{
			Id:          2,
			Names:       "Andy",
			LastName:    "Esquivel",
			Age:         23,
			Estatura:    1.52,
			Email:       "andy@gmail.com",
			IsActivo:    true,
			DateCreated: "",
		},
	}

	res, err := motor.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, res, expectedRes, res)
}

func TestUpdateNames(t *testing.T) {
	//arrange

	myStubDB := MockRepository{}
	myStubDB.user.Names = "Before Update"
	motor := NewService(&myStubDB)
	expectedRes := "After Update"
	res, err := motor.UpdateLastNameAndAge(1, 23, "After Update")

	assert.Nil(t, err)
	assert.Equal(t, expectedRes, res.Names, "Los nombres deben ser los mismos")
}
