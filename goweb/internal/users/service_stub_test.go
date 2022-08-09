package users

import (
	"fmt"
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
	"github.com/stretchr/testify/assert"
)

type StubDB struct {
	errRead    string
	errWriting string
	errLastID  string
}

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
	if s.errRead != "" {
		return nil, fmt.Errorf("error: %s", s.errRead)
	}
	return users, nil
}

func (s *StubDB) GetByName(name string) ([]domain.Users, error) {
	return []domain.Users{}, nil
}

func (s *StubDB) Store(id, age int, name, lastName, email, creationDate string, height float64, active bool) (domain.Users, error) {
	if s.errRead != "" {
		return domain.Users{}, fmt.Errorf("error: %s", s.errRead)
	}
	user := domain.Users{
		Id:           id,
		Name:         name,
		LastName:     lastName,
		Email:        email,
		Age:          age,
		Height:       height,
		Active:       active,
		CreationDate: creationDate,
	}
	if s.errWriting != "" {
		return domain.Users{}, fmt.Errorf("error: %s", s.errWriting)
	}
	return user, nil
}

func (s *StubDB) Update(id, age int, name, lastName, email, creationDate string, height float64, active bool) (domain.Users, error) {
	return domain.Users{}, nil
}

func (s *StubDB) LastID() (int, error) {
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
	}
	if s.errLastID != "" {
		return 0, fmt.Errorf("error: %s", s.errLastID)
	}
	return len(users), nil
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

func TestGetAllErrRead(t *testing.T) {
	r := StubDB{errRead: "fail"}
	s := NewService(&r)

	_, err := s.GetAll()

	assert.EqualError(t, err, "error: fail")
}

func TestStore(t *testing.T) {
	r := StubDB{}
	s := NewService(&r)

	result, err := s.Store(20, "new", "new", "new", "new", 1.80, true)
	expected := domain.Users{
		Id:           2,
		Age:          20,
		Name:         "new",
		LastName:     "new",
		Email:        "new",
		CreationDate: "new",
		Height:       1.80,
		Active:       true,
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestStoreErrLastId(t *testing.T) {
	r := StubDB{errLastID: "fail"}
	s := NewService(&r)

	_, err := s.Store(20, "new", "new", "new", "new", 1.80, true)
	assert.EqualError(t, err, "error: fail")
}

func TestStoreErrWriting(t *testing.T) {
	r := StubDB{errWriting: "fail"}
	s := NewService(&r)

	_, err := s.Store(20, "new", "new", "new", "new", 1.80, true)
	assert.EqualError(t, err, "error: fail")
}
