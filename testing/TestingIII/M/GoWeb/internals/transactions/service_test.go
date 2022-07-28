package transactions

import (
	"GoWeb/internals/domain"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationUpdate(t *testing.T) {
	//arrange

	data := []domain.Transanction{
		{Id: 1, Code: "QWE123", Coin: "COP", Amount: 1000, Emisor: "Juan", Receptor: "MeLi", Date: "07-27-2022"},
		{Id: 2, Code: "QWE123", Coin: "COP", Amount: 2000, Emisor: "David", Receptor: "MeLi", Date: "07-27-2022"},
	}

	mockStore := MockStore{
		dataMock:      data,
		errRead:       "",
		errWrite:      "",
		ReadWasCalled: false,
	}

	//act
	repo := NewRepository(&mockStore)
	service := NewService(repo)
	res, err := service.Update(1, "BBB", "USD", 4000, "NewName", "MeLi", "07-28-2022")
	//assert

	assert.Nil(t, err)
	assert.True(t, mockStore.ReadWasCalled)
	assert.Equal(t, mockStore.dataMock[0], res)
	assert.Equal(t, mockStore.dataMock[0].Id, 1)
}

func TestIntegrationUpdateFailRead(t *testing.T) {
	//arrange
	mockStore := MockStore{
		dataMock:      nil,
		errRead:       "cant read database",
		errWrite:      "",
		ReadWasCalled: false,
	}

	readErr := fmt.Errorf("cant read database")
	expectedError := fmt.Errorf("%w", readErr)

	//act
	repo := NewRepository(&mockStore)
	service := NewService(repo)
	res, err := service.Update(1, "BBB", "USD", 4000, "NewName", "MeLi", "07-28-2022")

	//assert
	assert.Equal(t, domain.Transanction{}, res)
	assert.ErrorContains(t, err, expectedError.Error())
	assert.False(t, mockStore.ReadWasCalled)
}

func TestIntegrationUpdateFailWrite(t *testing.T) {
	//arrange
	data := []domain.Transanction{
		{Id: 1, Code: "QWE123", Coin: "COP", Amount: 1000, Emisor: "Juan", Receptor: "MeLi", Date: "07-27-2022"},
		{Id: 2, Code: "QWE123", Coin: "COP", Amount: 2000, Emisor: "David", Receptor: "MeLi", Date: "07-27-2022"},
	}

	mock := MockStore{
		dataMock:      data,
		errRead:       "",
		errWrite:      "cant write database",
		ReadWasCalled: false,
	}

	writeErr := fmt.Errorf("cant write database")
	expectedError := fmt.Errorf("%w", writeErr)

	//act
	repo := NewRepository(&mock)
	service := NewService(repo)
	res, err := service.Update(1, "BBB", "USD", 4000, "NewName", "MeLi", "07-28-2022")

	//assert
	assert.True(t, mock.ReadWasCalled)
	assert.ErrorContains(t, err, expectedError.Error())
	assert.Equal(t, domain.Transanction{}, res)

}

func TestIntegrationDelete(t *testing.T) {
	//arrange
	data := []domain.Transanction{
		{Id: 1, Code: "QWE123", Coin: "COP", Amount: 1000, Emisor: "Juan", Receptor: "MeLi", Date: "07-27-2022"},
		{Id: 2, Code: "QWE123", Coin: "COP", Amount: 2000, Emisor: "David", Receptor: "MeLi", Date: "07-27-2022"},
	}

	mock := MockStore{
		dataMock:      data,
		errRead:       "",
		errWrite:      "",
		ReadWasCalled: false,
	}

	//act
	repo := NewRepository(&mock)
	service := NewService(repo)
	err := service.Delete(1)

	// assert

	assert.Nil(t, err)
	assert.True(t, mock.ReadWasCalled)
	assert.NotEqual(t, mock.dataMock[0].Id, 1)
}

func TestIntegrationDeleteFail(t *testing.T) {
	//arrange
	data := []domain.Transanction{
		{Id: 1, Code: "QWE123", Coin: "COP", Amount: 1000, Emisor: "Juan", Receptor: "MeLi", Date: "07-27-2022"},
		{Id: 2, Code: "QWE123", Coin: "COP", Amount: 2000, Emisor: "David", Receptor: "MeLi", Date: "07-27-2022"},
	}

	mock := MockStore{
		dataMock:      data,
		errRead:       "",
		errWrite:      "",
		ReadWasCalled: false,
	}

	//act
	repo := NewRepository(&mock)
	service := NewService(repo)
	_ = service.Delete(4)

	// assert

	assert.True(t, mock.ReadWasCalled)
}
