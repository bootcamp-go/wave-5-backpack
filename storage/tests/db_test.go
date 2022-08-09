package tests

import (
	"log"
	"practica/internal/domain"
	"practica/internal/products"
	"practica/pkg/storage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDB(t *testing.T) {
	p := domain.Product{
		Name: "test",
	}
	myRepo := products.NewRepo(storage.StorageDB)
	productResult, err := myRepo.Store(p)
	if err != nil {
		log.Println(err)
	}
	assert.Equal(t, p.Name, productResult.Name)
	productGet, err := myRepo.GetOne(productResult.ID)
	if err != nil {
		log.Println(err)
	}
	assert.Equal(t, productResult, productGet)
}
