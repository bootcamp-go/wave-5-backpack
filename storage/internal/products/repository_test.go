package products

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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

func TestUpdate(t *testing.T) {
	db, err := store.RunTxdb()
	assert.Nil(t, err)

	ctx := context.TODO()
	repository := NewRepository(db)

	productExpected := domain.Product{
		ID:          1,
		Name:        "Product Updated",
		Type:        "negro",
		Price:       3500,
		Count:       90,
		Code:        "AR897656M",
		Public:      1,
		WarehouseID: 1,
	}

	err = repository.Update(ctx, productExpected)
	assert.Nil(t, err)

	product, err := repository.GetProductByName("Product Updated")
	assert.Nil(t, err)
	assert.Equal(t, productExpected, product)
}

func TestDelete(t *testing.T) {
	db, err := store.RunTxdb()
	assert.Nil(t, err)

	repository := NewRepository(db)

	idDelete := 1

	err = repository.Delete(idDelete)
	assert.Nil(t, err)

	product, err := repository.GetProductByName("TV Samsung 60'")
	assert.NotNil(t, err)
	assert.Zero(t, product)

	products, err := repository.GetAll()
	assert.Equal(t, len(products), 2)
	assert.Nil(t, err)
}

func TestGetOneMock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	columns := []string{"id", "name", "type", "price", "count", "code", "public", "warehouseid"}
	rows := sqlmock.NewRows(columns)
	name := "Teclado HP"
	rows.AddRow(4, name, "negro", 5000, 170, "AR98674T", 1, 1)
	mock.ExpectQuery("SELECT .* FROM products").WithArgs(name).WillReturnRows(rows)

	pExpected := domain.Product{
		ID:          4,
		Name:        "Teclado HP",
		Type:        "negro",
		Price:       5000,
		Count:       170,
		Code:        "AR98674T",
		Public:      1,
		WarehouseID: 1,
	}

	repository := NewRepository(db)

	product, err := repository.GetProductByName(name)
	assert.Nil(t, err)
	assert.NotEmpty(t, product)
	assert.Equal(t, pExpected, product)

}

func TestStoreMock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare("INSERT INTO products")
	mock.ExpectExec("INSERT INTO products").WillReturnResult(sqlmock.NewResult(1, 1))

	pExpected := domain.Product{}

	repository := NewRepository(db)

	id, err := repository.Store(pExpected)
	assert.Nil(t, err)
	assert.NotEmpty(t, id)

}

func TestStoreMockQueryFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare("INSERT INTO warehouse")
	mock.ExpectExec("INSERT INTO warehouse").WillReturnResult(sqlmock.NewResult(1, 1))

	pExpected := domain.Product{}

	repository := NewRepository(db)

	id, err := repository.Store(pExpected)
	assert.NotNil(t, err)
	assert.Empty(t, id)

}
