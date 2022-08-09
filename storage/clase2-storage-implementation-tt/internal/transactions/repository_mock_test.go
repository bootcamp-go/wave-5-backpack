package transactions

import (
	"clase2-storage-implementation-tt/internal/domain"
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var colums = []string{"id", "codeTransaction", "currency", "amount", "transmitter", "receiver", "date"}

func Test_sqlRepository_Ecommerce_Mock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// mock.ExpectPrepare("INSERT INTO transactions")
	mock.ExpectPrepare(regexp.QuoteMeta(InsertTransaction))
	mock.ExpectExec("INSERT INTO transactions").WillReturnResult(sqlmock.NewResult(1, 1))

	transactionID := 1
	repo := NewRepository(db)
	transaction := domain.Transaction{
		ID: transactionID, CodigoTransaccion: "cba", Moneda: "USD", Monto: 233,
		Emisor: "Citigroup Inc", Receptor: "HSBC", Fecha: "2021-02-23",
	}

	tr, err := repo.Ecommerce(transaction)

	// Validation
	assert.NoError(t, err)
	assert.NotZero(t, tr)
	assert.Equal(t, transaction.ID, tr.ID)
}

func Test_sqlRepository_GetAllWithTimeout_Mock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows(colums)
	mock.ExpectQuery("select id, name, type, count, price").WillDelayFor(10 * time.Second).WillReturnRows(rows)
	repo := NewRepository(db)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = repo.GetAll(ctx)

	// Validation
	assert.Error(t, err)
}

func Test_sqlRepository_GetOneWithContextAndTimeout_Mock(t *testing.T) {
	idSelected := 1
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows(colums)
	rows.AddRow(idSelected, "abc", "JPY", 1012.76, "SMFG", "Mitsubishi UFJ", "2018-06-24")
	mock.ExpectQuery("select id, name, type, count, price").WillDelayFor(10 * time.Second).WillReturnRows(rows)
	repo := NewRepository(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = repo.GetOneWithContext(ctx, idSelected)

	// Validation
	assert.Error(t, err)

}
