package users

import (
	"context"
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
	"github.com/stretchr/testify/assert"
)

type MockDB struct {
	GetWasCalled bool
}

func (s *MockDB) GetAll(ctx context.Context) ([]domain.Users, error) {
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

func (s *MockDB) Store(ctx context.Context, id, age int, name, lastName, email, creationDate string, height float64, active bool) (domain.Users, error) {
	return domain.Users{}, nil
}

func (s *MockDB) Update(ctx context.Context, id, age int, name, lastName, email, creationDate string, height float64, active bool) (domain.Users, error) {
	return domain.Users{}, nil
}

func (s *MockDB) LastID(ctx context.Context) (int, error) {
	return 0, nil
}

func (s *MockDB) GetByName(ctx context.Context, name string) ([]domain.Users, error) {
	return []domain.Users{}, nil
}

func (r *MockDB) GetOne(ctx context.Context, id int) (domain.Users, error) {
	return domain.Users{}, nil
}

func (s *MockDB) UpdateLastNameAndAge(ctx context.Context, id, age int, lastName string) (domain.Users, error) {
	beforeUpdate, _ := s.GetAll(ctx)
	s.GetWasCalled = true

	var user domain.Users
	for i := 0; i < len(beforeUpdate)-1; i++ {
		if beforeUpdate[i].Id == id {
			user = beforeUpdate[i]
			user.Age = age
			user.LastName = lastName
			return user, nil
		}
	}
	return domain.Users{}, nil
}

func (s *MockDB) Delete(ctx context.Context, id int) error {
	return nil
}

func TestUpdateLastNameAndAge(t *testing.T) {
	r := MockDB{}
	s := NewService(&r)

	expected := domain.Users{
		Id:           1,
		Name:         "Lucas",
		LastName:     "Fernandez",
		Email:        "abc",
		Age:          25,
		Height:       1.69,
		Active:       true,
		CreationDate: "12/5/19",
	}

	result, _ := s.UpdateLastNameAndAge(context.TODO(), 1, 25, "Fernandez")

	assert.Equal(t, expected, result)
	assert.True(t, r.GetWasCalled)
}

func TestUpdateLastNameAndAgeNotUpdated(t *testing.T) {
	r := MockDB{}
	s := NewService(&r)

	expected := domain.Users{}

	result, _ := s.UpdateLastNameAndAge(context.TODO(), 5, 25, "Fernandez")

	assert.Equal(t, expected, result)
	assert.True(t, r.GetWasCalled)
}
