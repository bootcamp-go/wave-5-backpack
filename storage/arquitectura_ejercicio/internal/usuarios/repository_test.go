package usuarios

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/anesquivel/wave-5-backpack/storage/arquitectura_ejercicio/db"
	"github.com/anesquivel/wave-5-backpack/storage/arquitectura_ejercicio/internal/domain"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

type MockStorage struct {
	dataMock       []domain.Usuario
	errRead        string
	errWrite       string
	hadCalledRead  bool
	hadCalledWrite bool
}

func (m *MockStorage) Read(data interface{}) error {
	if m.errRead != "" {
		return fmt.Errorf(m.errRead)
	}

	user := data.(*[]domain.Usuario)
	*user = m.dataMock

	m.hadCalledRead = true
	return nil
}

func (m *MockStorage) Write(data interface{}) error {

	if m.errWrite != "" {
		return fmt.Errorf(m.errWrite)
	}

	a := data.([]domain.Usuario)
	m.dataMock = append(m.dataMock, a...)

	m.hadCalledWrite = true
	return nil

}

func (m *MockStorage) Ping() error {
	if m.errRead != "" {
		return fmt.Errorf(m.errRead)
	}
	return nil
}

/*func TestGetall(t *testing.T) {
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
	res, err := repo.GetAll()

	// assert

	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataMock, res)
}*/

func TestStore(t *testing.T) {
	db.Init()
	repo := NewRepository(db.StorageDB)
	newUser := domain.Usuario{
		Names:       "Ashton",
		LastName:    "Brooke",
		Email:       "ash2@gmail.com",
		Estatura:    1.80,
		IsActivo:    true,
		DateCreated: "2022-08-08",
		Age:         28,
	}

	userResult, err := repo.Store(newUser)

	if err != nil {
		log.Println("----- ERROR- TEST:", err.Error())
	}

	assert.Equal(t, 1, userResult.Id)
	assert.Equal(t, newUser.Names, userResult.Names)
}

func TestByName(t *testing.T) {
	db.Init()
	repo := NewRepository(db.StorageDB)
	name := "Ashton"
	userResult, err := repo.GetByName(name)

	if err != nil {
		log.Println("----- ERROR- TEST:", err.Error())
	}

	assert.Equal(t, 1, userResult.Id)
}

func TestGetAll(t *testing.T) {
	db.Init()
	repo := NewRepository(db.StorageDB)
	totalOfUsers := 4
	userResult, err := repo.GetAll()

	if err != nil {
		log.Println("----- ERROR- TEST:", err.Error())
	}

	assert.Equal(t, totalOfUsers, len(userResult))
}

func TestUpdateLASTAGE(t *testing.T) {
	db.Init()
	repo := NewRepository(db.StorageDB)
	id, lastName, age := 1, "Irwin", 26
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	userResult, err := repo.UpdateLastNameAndAge(ctx, id, age, lastName)

	if err != nil {
		log.Println("----- ERROR- TEST:", err.Error())
	}

	assert.Equal(t, age, userResult.Age)
	assert.Equal(t, lastName, userResult.LastName)

}
