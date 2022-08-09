package transactions

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/storage/internal/models"
	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/storage/pkg/db"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestRepositoryGetByID(t *testing.T) {
	//Arrange
	db, err := db.NewConnection()
	assert.Nil(t, err)

	repo := NewRepository(db)
	ctx := context.TODO()

	//Act
	transaction, err := repo.GetByID(ctx, 1)

	//Assert
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
}

func TestRepositoryGetAll(t *testing.T) {
	//Arrange
	db, err := db.NewConnection()
	assert.Nil(t, err)

	repo := NewRepository(db)
	ctx := context.TODO()

	//Act
	transactions, err := repo.GetAll(ctx)

	//Assert
	assert.NotNil(t, transactions)
	assert.Nil(t, err)
}

func TestRepositoryStore(t *testing.T) {
	//Arrange
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()

	tr := models.Transaction{
		ID:       1,
		Monto:    1000,
		Cod:      "abc1234abc",
		Moneda:   "ARS",
		Emisor:   "BBVA",
		Receptor: "Mercado Pago",
		Fecha:    time.Now().Format("2006-01-01"),
	}

	ctx := context.TODO()

	mock.ExpectPrepare(regexp.QuoteMeta(queryStore))
	mock.ExpectExec(regexp.QuoteMeta(queryStore)).WithArgs(tr.Monto, tr.Cod, tr.Moneda, tr.Emisor, tr.Receptor, tr.Fecha).WillReturnResult(sqlmock.NewResult(1, 1))

	idExpected := 1

	repo := NewRepository(db)

	//Act
	transaction, err := repo.Store(ctx, tr.Monto, tr.Cod, tr.Moneda, tr.Emisor, tr.Receptor)

	//Assert
	assert.NoError(t, err)
	assert.NotZero(t, t)
	assert.Equal(t, idExpected, transaction.ID)
	assert.Equal(t, tr, transaction)
}
