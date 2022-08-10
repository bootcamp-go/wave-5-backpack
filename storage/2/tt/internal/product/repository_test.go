package product

import (
	"context"
	"database/sql"
	"errors"
	"storage/2/tt/internal/domain"
	"storage/2/tt/util"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

// del turno mañana
func TestGetAll(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/storage")
	if err != nil {
		assert.Fail(t, "an error ocurred", err.Error())
	}
	r := NewRepository(db)
	products, err := r.GetAll()
	if err != nil {
		assert.Fail(t, "an error ocurred", err.Error())
	}

	assert.Greater(t, len(products), 0)
}

// del turno tarde
func Test_txdb_Store(t *testing.T) {
	db, err := util.InitDb()
	assert.NoError(t, err)

	r := NewRepository(db)

	product := domain.Product{
		Name:  "test_product",
		Type:  "test_type",
		Count: 1,
		Price: 1.1,
	}

	storeResult, err := r.Store(product)
	assert.NoError(t, err)
	assert.NotZero(t, storeResult)
	assert.Equal(t, product.Name, storeResult.Name)
	assert.Equal(t, product.Type, storeResult.Type)
	assert.Equal(t, product.Count, storeResult.Count)
	assert.Equal(t, product.Price, storeResult.Price)

	getResult, err := r.GetOne(9999999999999)
	assert.NoError(t, err)
	assert.Zero(t, getResult)

	getResult, err = r.GetOne(storeResult.ID)
	assert.NoError(t, err)
	assert.NotZero(t, getResult)
	assert.Equal(t, product.Name, getResult.Name)
	assert.Equal(t, product.Type, getResult.Type)
	assert.Equal(t, product.Count, getResult.Count)
	assert.Equal(t, product.Price, getResult.Price)
}

func Test_txdb_Update(t *testing.T) {
	db, err := util.InitDb()
	assert.NoError(t, err)

	r := NewRepository(db)

	ctx := context.TODO()
	product := domain.Product{
		ID:    7,
		Name:  "test_product",
		Type:  "test_type",
		Count: 1,
		Price: 1.1,
	}

	updateResult, err := r.Update(ctx, product)
	assert.NoError(t, err)
	assert.NotZero(t, updateResult)
	assert.Equal(t, product.ID, updateResult.ID)
	assert.Equal(t, product.Name, updateResult.Name)
	assert.Equal(t, product.Type, updateResult.Type)
	assert.Equal(t, product.Count, updateResult.Count)
	assert.Equal(t, product.Price, updateResult.Price)

	getResult, err := r.GetOne(product.ID)
	assert.NoError(t, err)
	assert.NotZero(t, getResult)
	assert.Equal(t, product.Name, getResult.Name)
	assert.Equal(t, product.Type, getResult.Type)
	assert.Equal(t, product.Count, getResult.Count)
	assert.Equal(t, product.Price, getResult.Price)
}

func Test_txdb_Delete(t *testing.T) {
	db, err := util.InitDb()
	assert.NoError(t, err)

	r := NewRepository(db)

	productID := 7

	err = r.Delete(productID)
	assert.NoError(t, err)

	getResult, err := r.GetOne(productID)
	assert.NoError(t, err)
	assert.Zero(t, getResult)

	getAllResult, err := r.GetAll()
	assert.NoError(t, err)
	for _, item := range getAllResult {
		if item.ID == productID {
			assert.Fail(t, "unexpected id found %d", item.ID)
		}
	}
}

func Test_sqlmock_Store(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare("INSERT INTO products").
		ExpectExec().
		WillReturnResult(sqlmock.NewResult(1, 1))

	columns := []string{"id", "name", "type", "count", "price"}
	rows := sqlmock.NewRows(columns)
	productID := 1
	rows.AddRow(productID, "test_product", "test_type", 1, 1.1)

	mock.ExpectQuery("SELECT .* FROM products .*").
		WithArgs(productID).
		WillReturnRows(rows)

	r := NewRepository(db)
	product := domain.Product{
		ID: productID,
	}

	storeResult, err := r.Store(product)
	assert.NoError(t, err)
	assert.NotZero(t, storeResult)
	assert.Equal(t, productID, storeResult.ID)

	getResult, err := r.GetOne(productID)
	assert.NoError(t, err)
	assert.NotZero(t, storeResult)
	assert.Equal(t, productID, getResult.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_sqlmock_StoreFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	product := domain.Product{}

	mock.ExpectPrepare("INSERT INTO products").
		ExpectExec().
		WithArgs(product.Name, product.Type, product.Count, product.Price).
		WillReturnError(errors.New("error"))

	r := NewRepository(db)

	storeResult, err := r.Store(product)
	assert.Error(t, err)
	assert.Zero(t, storeResult)
	assert.NoError(t, mock.ExpectationsWereMet())
}
