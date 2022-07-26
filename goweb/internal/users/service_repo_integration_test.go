package users

import (
	"fmt"
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
	"github.com/stretchr/testify/assert"
)

type MockIntegration struct {
	dataMock      []domain.Users
	errWrite      string
	errRead       string
	readWasCalled bool
}

func (ms *MockIntegration) Ping() error {
	return nil
}
func (ms *MockIntegration) Write(data interface{}) error {
	if ms.errWrite != "" {
		return fmt.Errorf(ms.errWrite)
	}
	a := data.([]domain.Users)
	ms.dataMock = append(ms.dataMock, a...)
	return nil
}
func (ms *MockIntegration) Read(data interface{}) error {
	if ms.errRead != "" {
		return fmt.Errorf(ms.errRead)
	}
	user := data.(*[]domain.Users)
	*user = ms.dataMock
	ms.readWasCalled = true
	return nil
}
func TestIntegrationUpdate(t *testing.T) {
	// arrange
	database := []domain.Users{
		{Id: 1, Name: "Juan", LastName: "Perez", Height: 1.82, CreationDate: "1992"},
		{Id: 2, Name: "Simon", LastName: "Fernandez", Height: 1.65, CreationDate: "1232"},
	}
	updateUser := domain.Users{
		Id: 1, Name: "New", LastName: "New", Email: "New", Age: 10, Height: 1.91, Active: true, CreationDate: "New",
	}
	mockInte := MockIntegration{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}
	// act
	repo := NewRepository(&mockInte)
	service := NewService(repo)
	user, err := service.Update(
		updateUser.Id,
		updateUser.Age,
		updateUser.Name,
		updateUser.LastName,
		updateUser.Email,
		updateUser.CreationDate,
		updateUser.Height,
		updateUser.Active,
	)
	// assert
	assert.Nil(t, err)
	assert.Equal(t, mockInte.dataMock[0], user)
	assert.Equal(t, mockInte.dataMock[0].Id, 1)
	assert.True(t, mockInte.readWasCalled)
}

func TestIntegrationUpdateFailReading(t *testing.T) {
	// arrange
	database := []domain.Users{
		{Id: 1, Name: "Juan", LastName: "Perez", Height: 1.82, CreationDate: "1992"},
		{Id: 2, Name: "Simon", LastName: "Fernandez", Height: 1.65, CreationDate: "1232"},
	}
	updateUser := domain.Users{
		Id: 1, Name: "New", LastName: "New", Email: "New", Age: 10, Height: 1.91, Active: true, CreationDate: "New",
	}
	mockInte := MockIntegration{
		dataMock: database,
		errWrite: "",
		errRead:  "fail",
	}
	// act
	repo := NewRepository(&mockInte)
	service := NewService(repo)
	_, err := service.Update(
		updateUser.Id,
		updateUser.Age,
		updateUser.Name,
		updateUser.LastName,
		updateUser.Email,
		updateUser.CreationDate,
		updateUser.Height,
		updateUser.Active,
	)
	// assert
	assert.EqualError(t, err, "cant read database")
}

func TestIntegrationUpdateFailWriting(t *testing.T) {
	// arrange
	database := []domain.Users{
		{Id: 1, Name: "Juan", LastName: "Perez", Height: 1.82, CreationDate: "1992"},
		{Id: 2, Name: "Simon", LastName: "Fernandez", Height: 1.65, CreationDate: "1232"},
	}
	updateUser := domain.Users{
		Id: 1, Name: "New", LastName: "New", Email: "New", Age: 10, Height: 1.91, Active: true, CreationDate: "New",
	}
	mockInte := MockIntegration{
		dataMock: database,
		errWrite: "fail",
		errRead:  "",
	}
	// act
	repo := NewRepository(&mockInte)
	service := NewService(repo)
	_, err := service.Update(
		updateUser.Id,
		updateUser.Age,
		updateUser.Name,
		updateUser.LastName,
		updateUser.Email,
		updateUser.CreationDate,
		updateUser.Height,
		updateUser.Active,
	)
	// assert
	assert.EqualError(t, err, "cant write database, error: fail")
}

func TestIntegrationDelete(t *testing.T) {
	// arrange
	database := []domain.Users{
		{Id: 1, Name: "Juan", LastName: "Perez", Height: 1.82, CreationDate: "1992"},
		{Id: 2, Name: "Simon", LastName: "Fernandez", Height: 1.65, CreationDate: "1232"},
	}
	mockInte := MockIntegration{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}
	// act
	repo := NewRepository(&mockInte)
	service := NewService(repo)
	err := service.Delete(1)
	// assert
	assert.Nil(t, err)
}

func TestIntegrationDeleteNotFound(t *testing.T) {
	// arrange
	database := []domain.Users{
		{Id: 1, Name: "Juan", LastName: "Perez", Height: 1.82, CreationDate: "1992"},
		{Id: 2, Name: "Simon", LastName: "Fernandez", Height: 1.65, CreationDate: "1232"},
	}
	mockInte := MockIntegration{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}
	writeErr := fmt.Errorf("user 30 not found")
	// act
	repo := NewRepository(&mockInte)
	service := NewService(repo)
	errNotFound := service.Delete(30)
	// assert
	assert.ErrorContains(t, errNotFound, writeErr.Error())
}
