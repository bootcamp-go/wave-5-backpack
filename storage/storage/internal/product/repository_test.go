package product

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DATA-DOG/go-txdb"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"storage/internal/domain"
	"testing"
)

func InitTestDB() (*sql.DB, error) {
	txdb.Register("txdb", "mysql", "root@tcp(localhost:3306)/storage")
	db, err := sql.Open("txdb", uuid.New().String())
	if err != nil || db.Ping() != nil {
		return nil, errors.New("Could not open db connection")
	}
	return db, nil
}

func TestStoreAndGetByName(t *testing.T) {
	db, err := InitTestDB()
	assert.NoError(t, err)
	repo := NewRepository(db)
	ctx := context.TODO()
	prd := domain.Product{
		Name:  "Test",
		Type:  "Tipo",
		Count: 1,
		Price: 50,
	}

	// Test Store
	result, err := repo.Store(ctx, prd)
	assert.NoError(t, err)
	assert.Greater(t, result.ID, 0)

	//Test GetByName
	result, err = repo.GetByName(ctx, "Test")
	assert.NoError(t, err)
	assert.Equal(t, "Test", result.Name)
	assert.Equal(t, "Tipo", result.Type)
	assert.Greater(t, result.ID, 0)
}

func TestStoreAndUpdate(t *testing.T) {
	db, err := InitTestDB()
	assert.NoError(t, err)
	repo := NewRepository(db)
	ctx := context.TODO()
	prd := domain.Product{
		Name:  "Test",
		Type:  "Tipo",
		Count: 1,
		Price: 50,
	}

	// Test Store
	result, err := repo.Store(ctx, prd)
	assert.NoError(t, err)
	assert.Greater(t, result.ID, 0)
	createdId := result.ID

	//Test Update
	updatePrd := domain.Product{
		ID:    createdId,
		Name:  "Test 2",
		Type:  "New Type",
		Count: 2,
		Price: 100,
	}
	result, err = repo.Update(ctx, updatePrd)
	assert.NoError(t, err)

	//Test GetByName
	result, err = repo.GetByName(ctx, "Test 2")
	assert.NoError(t, err)
	assert.Equal(t, "Test 2", result.Name)
	assert.Equal(t, "New Type", result.Type)
	assert.Equal(t, 2, result.Count)
	assert.Equal(t, float64(100), result.Price)
	assert.Equal(t, createdId, result.ID)
}

func TestStoreAndDelete(t *testing.T) {
	db, err := InitTestDB()
	assert.NoError(t, err)
	repo := NewRepository(db)
	ctx := context.TODO()
	prd := domain.Product{
		Name:  "Test",
		Type:  "Tipo",
		Count: 1,
		Price: 50,
	}

	// Test Store
	result, err := repo.Store(ctx, prd)
	assert.NoError(t, err)
	assert.Greater(t, result.ID, 0)
	createdId := result.ID

	//Test Delete
	err = repo.Delete(ctx, createdId)
	assert.NoError(t, err)

	//Test GetByName
	result, err = repo.GetByName(ctx, "Test")
	assert.NoError(t, err)
	assert.Equal(t, 0, result.ID)

	//Test GetAll
	resultList, err := repo.GetAll(ctx)
	hasId := false
	for _, prd := range resultList {
		if prd.ID == createdId {
			hasId = true
		}
	}
	assert.NoError(t, err)
	assert.Equal(t, false, hasId)
}

func TestStoreAndDeleteAsMock(t *testing.T) {
	// Test Store
	prd := domain.Product{
		Name:  "Nombre",
		Type:  "Tipo",
		Count: 1,
		Price: float64(10),
	}
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPrepare("INSERT INTO products")
	mock.ExpectExec("INSERT INTO products").WillReturnResult(sqlmock.NewResult(1, 1))
	repo := NewRepository(db)
	ctx := context.TODO()
	result, err := repo.Store(ctx, prd)
	assert.NoError(t, err)
	assert.Greater(t, result.ID, 0)
	assert.NoError(t, mock.ExpectationsWereMet())

	//Test GetByName
	columns := []string{
		"id", "name", "type", "count", "price",
	}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(1, "Nombre", "Tipo", 1, float64(10))
	mock.ExpectQuery("SELECT id, name, type, count, price FROM products").WithArgs("Nombre").WillReturnRows(rows)
	result, err = repo.GetByName(ctx, "Nombre")
	assert.NoError(t, err)
	assert.Equal(t, 1, result.ID)
	assert.NoError(t, mock.ExpectationsWereMet())

	//Test Delete
	mock.ExpectPrepare("DELETE FROM products")
	mock.ExpectExec("DELETE FROM products").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
	err = repo.Delete(ctx, 1)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())

	//Test GetAll
	newRows := sqlmock.NewRows(columns)
	mock.ExpectQuery("SELECT id, name, type, count, price FROM products").WillReturnRows(newRows)
	resultList, err := repo.GetAll(ctx)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(resultList))
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateFailAsMock(t *testing.T) {
	prd := domain.Product{
		ID:    1,
		Name:  "Nombre",
		Type:  "Tipo",
		Count: 1,
		Price: float64(10),
	}
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPrepare("UPDATE products SET")
	mock.ExpectExec("UPDATE products SET").WithArgs("Nombre", "Tipo", 1, float64(10), 1).WillReturnError(errors.New("id doesn't exists"))

	repo := NewRepository(db)
	ctx := context.TODO()

	result, err := repo.Update(ctx, prd)

	assert.ErrorContains(t, err, "id doesn't exists")
	assert.Empty(t, result)
}
