package products

import (
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/storage/internal/domain"
	"github.com/bootcamp-go/wave-5-backpack/storage/pkg/store"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	db, err := store.RunTxdb()
	assert.Nil(t, err)

	repository := NewRepository(db)

	products, err := repository.GetAll()

	assert.Nil(t, err)
	assert.NotEmpty(t, products)
}

func TestGetOne(t *testing.T) {
	db, err := store.RunTxdb()
	assert.Nil(t, err)

	repository := NewRepository(db)

	pExpected := domain.Product{
		ID:     4,
		Name:   "Teclado HP",
		Type:   "negro",
		Price:  5000,
		Count:  170,
		Code:   "AR98674T",
		Public: 1,
	}

	product, err := repository.GetProductByName(pExpected.Name)
	assert.Nil(t, err)
	assert.Equal(t, pExpected, product)

	product, err = repository.GetProductByName("nombreIncorrecto")
	assert.NotNil(t, err)
	assert.Zero(t, product)
}

func TestStore(t *testing.T) {
	db, err := store.RunTxdb()
	assert.Nil(t, err)

	repository := NewRepository(db)

	product := domain.Product{
		Name:        "Mouse HP",
		Type:        "negro",
		Price:       3500,
		Count:       90,
		Code:        "AR897656M",
		Public:      1,
		WarehouseID: 1,
	}

	pID, err := repository.Store(product)
	assert.Nil(t, err)
	assert.NotEmpty(t, pID)

	p, err := repository.GetProductByName("Mouse HP")
	product.ID = pID
	assert.Nil(t, err)
	assert.Equal(t, product, p)
}
