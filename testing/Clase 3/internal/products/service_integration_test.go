package products

import (
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/internal/domain"
	"github.com/stretchr/testify/assert"
)

type MockFileService struct {
	mockData []domain.Product
}

func (s *MockFileService) Read(data interface{}) error {
	a := data.(*[]domain.Product)
	*a = s.mockData
	return nil

}

func (s *MockFileService) Write(data interface{}) error {
	a := data.(*[]domain.Product)
	s.mockData = *a
	return nil
}

func (s MockFileService) Ping() error {
	return nil
}

func TestUpdate(t *testing.T) {
	expected := []domain.Product{
		{
			Id:         1,
			Name:       "Laptop mod",
			Color:      "black",
			Price:      999.99,
			Stock:      100,
			Code:       "SJD23RFG",
			Published:  false,
			Created_at: "2022-06-30",
		},
		{
			Id:         2,
			Name:       "3090",
			Color:      "green",
			Price:      2000,
			Stock:      10,
			Code:       "QBH76BC",
			Published:  true,
			Created_at: "2022-01-24",
		},
	}
	mock := MockFileService{
		mockData: []domain.Product{
			{
				Id:         1,
				Name:       "Laptop",
				Color:      "black",
				Price:      999.99,
				Stock:      100,
				Code:       "SJD23RFG",
				Published:  false,
				Created_at: "2022-06-30",
			},
			{
				Id:         2,
				Name:       "3090",
				Color:      "green",
				Price:      2000,
				Stock:      10,
				Code:       "QBH76BC",
				Published:  true,
				Created_at: "2022-01-24",
			},
		},
	}
	r := NewRepository(&mock)
	s := NewService(r)
	_, errUpdate := s.UpdatePartial(1, "Laptop mod", "", 0, 100, "", false, "")
	newList, errGetAll := s.GetAll()
	assert.Nil(t, errUpdate)
	assert.Nil(t, errGetAll)
	assert.Equal(t, expected, newList)
}

func TestDelete(t *testing.T) {
	expected := []domain.Product{
		{
			Id:         2,
			Name:       "3090",
			Color:      "green",
			Price:      2000,
			Stock:      10,
			Code:       "QBH76BC",
			Published:  true,
			Created_at: "2022-01-24",
		},
	}
	mock := MockFileService{
		mockData: []domain.Product{
			{
				Id:         1,
				Name:       "Laptop",
				Color:      "black",
				Price:      999.99,
				Stock:      100,
				Code:       "SJD23RFG",
				Published:  false,
				Created_at: "2022-06-30",
			},
			{
				Id:         2,
				Name:       "3090",
				Color:      "green",
				Price:      2000,
				Stock:      10,
				Code:       "QBH76BC",
				Published:  true,
				Created_at: "2022-01-24",
			},
		},
	}
	r := NewRepository(&mock)
	s := NewService(r)
	_, errUpdate := s.Delete(1)
	newList, errGetAll := s.GetAll()
	assert.Nil(t, errUpdate)
	assert.Nil(t, errGetAll)
	assert.Equal(t, expected, newList)
}
