package products

import (
	"fmt"
	"goweb/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllIntegration(t *testing.T) {

	database := []domain.Products{

		{
			Id:            1,
			Nombre:        "aguacate",
			Color:         "verde",
			Precio:        30000,
			Stock:         5,
			Codigo:        "23fe2",
			Publicado:     true,
			FechaCreacion: "23/10/2022",
		},

		{
			Id:            2,
			Nombre:        "Banana",
			Color:         "Amarillo",
			Precio:        60000,
			Stock:         13,
			Codigo:        "d7fe2",
			Publicado:     true,
			FechaCreacion: "30/11/2022",
		},
	}

	mockStorage := MockStorage{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}

	repo := InitRepository(&mockStorage)
	service := InitService(repo)
	result, err := service.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataMock, result)
}

func TestUpdateIntegration(t *testing.T) {

	updateData := []domain.Products{
		{Id: 1,
			Nombre:        "Sandia",
			Color:         "verde",
			Precio:        12332,
			Stock:         13,
			Codigo:        "asd7sd",
			Publicado:     true,
			FechaCreacion: "12/02/2022"},
	}

	mockUpdate := MockStorage{
		dataMock: updateData,
		errWrite: "",
		errRead:  "",
	}

	repo := InitRepository(&mockUpdate)
	service := InitService(repo)
	result, err := service.Update(updateData[0].Id, updateData[0].Nombre, updateData[0].Color, updateData[0].Precio, updateData[0].Stock, updateData[0].Codigo, updateData[0].Publicado, updateData[0].FechaCreacion)
	assert.Nil(t, err)
	assert.Equal(t, mockUpdate.dataMock[0], result)

}

func TestDeleteIntegration(t *testing.T) {
	deleteData := []domain.Products{
		{Id: 1,
			Nombre:        "Sandia",
			Color:         "verde",
			Precio:        12332,
			Stock:         13,
			Codigo:        "asd7sd",
			Publicado:     true,
			FechaCreacion: "12/02/2022"},
	}
	mockUpdate := MockStorage{
		dataMock: deleteData,
		errWrite: "",
		errRead:  "",
	}

	repo := InitRepository(&mockUpdate)
	service := InitService(repo)
	result := service.Delete(1)
	assert.Nil(t, result)
	assert.Equal(t, nil, result)

}

func TestDeleteFailIntegration(t *testing.T) {
	deleteData := []domain.Products{
		{Id: 2,
			Nombre:        "Sandia",
			Color:         "verde",
			Precio:        12332,
			Stock:         13,
			Codigo:        "asd7sd",
			Publicado:     true,
			FechaCreacion: "12/02/2022"},
	}
	mockUpdate := MockStorage{
		dataMock: deleteData,
		errWrite: "",
		errRead:  "",
	}

	errTest := fmt.Errorf("Producto %d no encontrado", 1)

	repo := InitRepository(&mockUpdate)
	service := InitService(repo)
	err := service.Delete(1)
	assert.ErrorContains(t, err, err.Error())
	assert.Equal(t, errTest, err)

}

func TestUpadeOneIntegration(t *testing.T) {

	updateData := []domain.Products{
		{
			Id:            1,
			Nombre:        "Papaya",
			Color:         "verde",
			Precio:        231522,
			Stock:         13,
			Codigo:        "asd7sd",
			Publicado:     true,
			FechaCreacion: "12/02/2022",
		},
	}

	mockUpdate := MockStorage{
		dataMock: updateData,
		errWrite: "",
		errRead:  "",
	}

	repo := InitRepository(&mockUpdate)
	service := InitService(repo)
	result, err := service.UpdateOne(updateData[0].Id, updateData[0].Nombre, updateData[0].Precio)
	assert.Nil(t, err)
	assert.Equal(t, mockUpdate.dataMock[0], result)

}
