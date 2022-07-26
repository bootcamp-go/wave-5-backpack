package products

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockStore struct {
	infodb interface{}
	read   bool
}

func (ms *MockStore) Read(data interface{}) error {
	r := reflect.ValueOf(data)
	r = reflect.Indirect(r)
	r.Set(reflect.ValueOf(ms.infodb))
	ms.read = true
	return nil
}

func (ms *MockStore) Write(data interface{}) error {
	ms.infodb = data
	return nil
}

func (ms *MockStore) Ping() error {
	return nil
}

func TestUpdate(t *testing.T) {
	testUpdateUsersNow := []Product{
		{ID: 1, Name: "Cama", Type: "Madera", Count: 2, Price: 270000.0},
		{ID: 2, Name: "Silla", Type: "Madera", Count: 2, Price: 270000.0},
	}

	testUpdateUsersThen := []Product{
		{ID: 1, Name: "Escritorio", Type: "Madera", Count: 2, Price: 130000.0},
		{ID: 2, Name: "Silla", Type: "Madera", Count: 2, Price: 270000.0},
	}

	info_update := &MockStore{
		infodb: testUpdateUsersNow,
		read:   false,
	}

	repository := NewRepository(info_update)

	result, _ := repository.Patch(1, "Escritorio", 130000.0)
	assert.Equal(t, testUpdateUsersThen, info_update.infodb)
	assert.Equal(t, testUpdateUsersThen[0], result)
	assert.True(t, info_update.read)
}
