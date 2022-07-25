package users

import (
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
	"github.com/stretchr/testify/assert"
)

type StubDB struct{}

func (s *StubDB) GetAll() ([]domain.Users, error) {
	users := []domain.Users{
		{
			Id:           1,
			Name:         "Lucas",
			LastName:     "Hernandez",
			Email:        "abc",
			Age:          15,
			Height:       1.69,
			Active:       true,
			CreationDate: "12/5/19",
		},
		{
			Id:           2,
			Name:         "Diana",
			LastName:     "Hernandez",
			Email:        "abc",
			Age:          18,
			Height:       1.78,
			Active:       true,
			CreationDate: "12/5/17",
		},
	}
	return users, nil
}

func (s *StubDB) Store(id, age int, name, lastName, email, creationDate string, height float64, active bool) (domain.Users, error) {
	return domain.Users{}, nil
}

func (s *StubDB) Update(id, age int, name, lastName, email, creationDate string, height float64, active bool) (domain.Users, error) {
	return domain.Users{}, nil
}

func (s *StubDB) LastID() (int, error) {
	return 0, nil
}

func (s *StubDB) UpdateLastNameAndAge(id, age int, lastName string) (domain.Users, error) {
	return domain.Users{}, nil
}

func (s *StubDB) Delete(id int) error {
	return nil
}

func TestGetAll(t *testing.T) {
	r := StubDB{}
	s := NewService(&r)

	expected := []domain.Users{
		{
			Id:           1,
			Name:         "Lucas",
			LastName:     "Hernandez",
			Email:        "abc",
			Age:          15,
			Height:       1.69,
			Active:       true,
			CreationDate: "12/5/19",
		},
		{
			Id:           2,
			Name:         "Diana",
			LastName:     "Hernandez",
			Email:        "abc",
			Age:          18,
			Height:       1.78,
			Active:       true,
			CreationDate: "12/5/17",
		},
	}

	result, _ := s.GetAll()

	assert.Equal(t, expected, result)
}
