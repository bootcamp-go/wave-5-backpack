package users

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/internal/domain"
	"github.com/stretchr/testify/assert"
)

type MockStorage struct {
	dataMock []domain.User
	errWrite string
	errRead  string
	ReadOk   bool
}

func (m *MockStorage) Ping() error {
	return nil
}
func (m *MockStorage) Read(data interface{}) error {
	if m.errRead != "" {
		return fmt.Errorf(m.errRead)
	}
	a := data.(*[]domain.User)
	*a = m.dataMock
	m.ReadOk = true
	return nil
}

func (m *MockStorage) Write(data interface{}) error {
	if m.errWrite != "" {
		return fmt.Errorf(m.errWrite)
	}
	a := data.([]domain.User)
	m.dataMock = append(m.dataMock, a[len(a)-1])
	return nil
}
func TestDeleteIntegration(t *testing.T) {
	//arrange
	db := []domain.User{
		{
			ID:         1,
			Name:       "name1",
			Lastname:   "lastname1",
			Email:      "1@mail.com",
			Age:        22,
			Height:     1.55,
			Active:     true,
			DoCreation: "02-03-2020",
		}, {
			ID:         2,
			Name:       "name2",
			Lastname:   "lastname2",
			Email:      "2@mail.com",
			Age:        23,
			Height:     1.56,
			Active:     true,
			DoCreation: "02-03-2021",
		},
	}
	mock := MockStorage{
		dataMock: db,
		errWrite: "",
		errRead:  "",
	}
	id := 1
	//act
	repo := NewRepository(&mock)
	serv := NewService(repo)

	err := serv.DeleteUser(id)

	//assert
	assert.Nil(t, err)
}
func TestUpdateIntegration(t *testing.T) {
	//arrange
	db := []domain.User{
		{
			ID:         1,
			Name:       "name1",
			Lastname:   "lastname1",
			Email:      "1@mail.com",
			Age:        22,
			Height:     1.55,
			Active:     true,
			DoCreation: "02-03-2020",
		}, {
			ID:         2,
			Name:       "name2",
			Lastname:   "lastname2",
			Email:      "2@mail.com",
			Age:        23,
			Height:     1.56,
			Active:     true,
			DoCreation: "02-03-2021",
		},
	}
	new := domain.User{
		ID:         1,
		Name:       "new name 1",
		Lastname:   "new lastname 1",
		Email:      "new@mail.com",
		Age:        44,
		Height:     1.44,
		Active:     true,
		DoCreation: "04-04-2014",
	}
	mock := MockStorage{
		dataMock: db,
		errWrite: "",
		errRead:  "",
	}

	//act
	repo := NewRepository(&mock)
	serv := NewService(repo)

	response, err := serv.UpdateUser(new.ID, new.Name, new.Lastname, new.Email, new.Age, new.Height, new.Active, new.DoCreation)

	//assert
	assert.Equal(t, new.ID, response.ID)
	assert.Nil(t, err)
	assert.Equal(t, new, response)
	assert.True(t, mock.ReadOk)
}
func TestIntegrationFailRead(t *testing.T) {
	// arrange
	e := errors.New("fail to read database")

	mock := MockStorage{
		dataMock: nil,
		errWrite: "",
		errRead:  "fail to read database",
	}
	// act
	repo := NewRepository(&mock)
	service := NewService(repo)
	response, err := service.UpdateUser(1, "", "", "", 0, 0, false, "")
	// accert
	assert.Equal(t, e, err)
	assert.Equal(t, response.ID, 0)
}

func TestIntegrationFailWrite(t *testing.T) {
	// arrange
	e := errors.New("fail to write database")
	db := []domain.User{
		{
			ID:         1,
			Name:       "name1",
			Lastname:   "lastname1",
			Email:      "1@mail.com",
			Age:        22,
			Height:     1.55,
			Active:     true,
			DoCreation: "02-03-2020",
		}, {
			ID:         2,
			Name:       "name2",
			Lastname:   "lastname2",
			Email:      "2@mail.com",
			Age:        23,
			Height:     1.56,
			Active:     true,
			DoCreation: "02-03-2021",
		},
	}
	mock := MockStorage{
		dataMock: db,
		errRead:  "",
		errWrite: "fail to write database",
	}
	// act
	repo := NewRepository(&mock)
	service := NewService(repo)
	response, err := service.UpdateUser(1, "", "", "", 0, 0, false, "")
	// accert
	assert.Equal(t, e, err)
	assert.Equal(t, response.ID, 0)
}
func TestGetAll(t *testing.T) {
	//arrange
	db := []domain.User{
		{
			ID:         1,
			Name:       "name1",
			Lastname:   "lastname1",
			Email:      "1@mail.com",
			Age:        22,
			Height:     1.55,
			Active:     true,
			DoCreation: "02-03-2020",
		}, {
			ID:         2,
			Name:       "name2",
			Lastname:   "lastname2",
			Email:      "2@mail.com",
			Age:        23,
			Height:     1.56,
			Active:     true,
			DoCreation: "02-03-2021",
		},
	}
	mockStorage := MockStorage{
		dataMock: db,
		errWrite: "",
		errRead:  "",
	}

	//act
	repo := NewRepository(&mockStorage)
	response, err := repo.GetAll(context.TODO())

	//assert
	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataMock, response)

}

//unit testing for services
type StubDataBase struct{}

func (st *StubDataBase) GetById(ctx context.Context,id int) (domain.User, error) {
	return domain.User{}, nil
}
func (st *StubDataBase) LastId() (int, error) {
	return 0, nil
}
func (st *StubDataBase) StoreUser(id int, name, lastname, email string, age int, height float32, active bool, doCreation string) (domain.User, error) {
	return domain.User{}, nil
}
func (st *StubDataBase) UpdateUser(id int, name, lastname, email string, age int, height float32, active bool, doCreation string) (domain.User, error) {
	return domain.User{}, nil
}
func (st *StubDataBase) DeleteUser(id int) error {
	return nil
}
func (st *StubDataBase) GetByName(ctx context.Context, name string) ([]domain.User, error) {
	return []domain.User{}, nil
}
func (st *StubDataBase) UpdateLastnameAndAge(id int, lastname string, age int) (*domain.User, error) {
	return &domain.User{}, nil
}

func (st *StubDataBase) GetAll(ctx context.Context) ([]domain.User, error) {
	allUsers := []domain.User{{
		ID:         1,
		Name:       "name1",
		Lastname:   "lastname1",
		Email:      "1@mail.com",
		Age:        22,
		Height:     1.55,
		Active:     true,
		DoCreation: "02-03-2020",
	}, {
		ID:         2,
		Name:       "name2",
		Lastname:   "lastname2",
		Email:      "2@mail.com",
		Age:        23,
		Height:     1.56,
		Active:     true,
		DoCreation: "02-03-2021",
	}}
	return allUsers, nil
}

func TestGetAllUser(t *testing.T) {
	stub := StubDataBase{}
	serv := NewService(&stub)
	expectedResponse := []domain.User{{
		ID:         1,
		Name:       "name1",
		Lastname:   "lastname1",
		Email:      "1@mail.com",
		Age:        22,
		Height:     1.55,
		Active:     true,
		DoCreation: "02-03-2020",
	}, {
		ID:         2,
		Name:       "name2",
		Lastname:   "lastname2",
		Email:      "2@mail.com",
		Age:        23,
		Height:     1.56,
		Active:     true,
		DoCreation: "02-03-2021",
	}}
	response, err := serv.GetAll(context.TODO())
	assert.Equal(t, expectedResponse, response)
	assert.Nil(t, err)
}
