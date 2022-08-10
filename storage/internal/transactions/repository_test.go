package transactions

import (
	"context"
	"database/sql"
	"goweb/internal/domain"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {

	txdb.Register("txdb", "mysql", "meli_sprint_user:Meli_Sprint#123@/transactions?parseTime=true")
	db, _ := sql.Open("txdb", uuid.NewString())
	repo := NewRepository(db)
	transactions, err := repo.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, transactions)

}

func TestGetAllMock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	expectedTime, err := time.Parse("2006-01-02T15:04:05Z07:00", "2022-08-08T20:37:32Z")
	assert.NoError(t, err)
	repo := NewRepository(db)
	rows := mock.NewRows([]string{"id", "transaction_code", "currency", "amount", "sender", "reciever", "transaction_date"}).
		AddRow(1, "dsad22313asd", "CLP", 20000, "Adriana Sepulveda", "Jose Luis", expectedTime).
		AddRow(2, "dsa212dasd", "CLP", 20000, "Claudio Figueroa", "Luz Carime", expectedTime)

	expectedTransactions := []domain.Transaction{
		{Id: 1, TransactionCode: "dsad22313asd", Currency: "CLP", Amount: 20000, Sender: "Adriana Sepulveda", Reciever: "Jose Luis", TransactionDate: expectedTime},
		{Id: 2, TransactionCode: "dsa212dasd", Currency: "CLP", Amount: 20000, Sender: "Claudio Figueroa", Reciever: "Luz Carime", TransactionDate: expectedTime},
	}

	mock.ExpectQuery("SELECT id, transaction_code, currency, amount, sender, reciever, transaction_date FROM transactions").WillReturnRows(rows)
	resultTransactions, err := repo.GetAll()

	assert.NoError(t, err)
	assert.NotEmpty(t, resultTransactions)
	assert.Equal(t, expectedTransactions, resultTransactions)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetByIdMock(t *testing.T) {
	db, mock, err := sqlmock.New()
	searchId := 1
	assert.NoError(t, err)
	expectedTime, err := time.Parse("2006-01-02T15:04:05Z07:00", "2022-08-08T20:37:32Z")
	assert.NoError(t, err)
	expectedTransaction := domain.Transaction{Id: 1, TransactionCode: "dsad22313asd", Currency: "CLP", Amount: 20000, Sender: "Adriana Sepulveda", Reciever: "Jose Luis", TransactionDate: expectedTime}
	repo := NewRepository(db)
	rows := mock.NewRows([]string{"id", "transaction_code", "currency", "amount", "sender", "reciever", "transaction_date"}).
		AddRow(1, "dsad22313asd", "CLP", 20000, "Adriana Sepulveda", "Jose Luis", expectedTime)

	mock.ExpectQuery("SELECT id,transaction_code,currency, amount,sender,reciever,transaction_date FROM transactions WHERE id = ?").WithArgs(searchId).WillReturnRows(rows)
	resultTransaction, err := repo.GetById(searchId)

	assert.NoError(t, err)
	assert.NotEmpty(t, resultTransaction)
	assert.Equal(t, expectedTransaction, resultTransaction)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetById(t *testing.T) {

	txdb.Register("txdb", "mysql", "meli_sprint_user:Meli_Sprint#123@/transactions?parseTime=true")
	db, _ := sql.Open("txdb", uuid.NewString())
	repo := NewRepository(db)
	expectedTime, err := time.Parse("2006-01-02T15:04:05Z07:00", "2022-08-08T20:37:32Z")
	assert.NoError(t, err)
	transactionExpected := domain.Transaction{
		Id:              1,
		TransactionCode: "7401fd17d38171b2686602781b2473",
		Currency:        "CLP",
		Amount:          120000,
		Reciever:        "Luz Lucumi",
		Sender:          "Carlos Desidero",
		TransactionDate: expectedTime,
	}

	transaction, err := repo.GetById(transactionExpected.Id)

	assert.NoError(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, transactionExpected, transaction)

}

func TestStoreAndGetOne(t *testing.T) {
	txdb.Register("txdb", "mysql", "meli_sprint_user:Meli_Sprint#123@/transactions?parseTime=true")
	db, _ := sql.Open("txdb", uuid.NewString())
	repo := NewRepository(db)
	expectedTime, errTime := time.Parse("2006-01-02T15:04:05Z07:00", "2022-08-08T20:37:32Z")

	expectedTransaction := domain.Transaction{
		TransactionCode: "7401fd17d38171bsasd112dsx",
		Currency:        "CLP",
		Amount:          120000,
		Reciever:        "Luz Lucumi",
		Sender:          "Carlos Desidero",
		TransactionDate: expectedTime,
	}
	storeTransaction, errStore := repo.Store(expectedTransaction.TransactionCode, expectedTransaction.Currency, expectedTransaction.Amount, expectedTransaction.Sender, expectedTransaction.Reciever, expectedTransaction.TransactionDate)
	expectedTransaction.Id = storeTransaction.Id
	gettedTransaction, errGet := repo.GetById(storeTransaction.Id)

	assert.NoError(t, errTime)
	assert.NoError(t, errStore)
	assert.NoError(t, errGet)
	assert.Equal(t, expectedTransaction, storeTransaction)
	assert.Equal(t, storeTransaction, gettedTransaction)
}

func TestUpdate(t *testing.T) {
	txdb.Register("txdb", "mysql", "meli_sprint_user:Meli_Sprint#123@/transactions?parseTime=true")
	db, _ := sql.Open("txdb", uuid.NewString())
	repo := NewRepository(db)
	dbTime, errTime := time.Parse("2006-01-02T15:04:05Z07:00", "2022-08-08T20:37:32Z")

	initTransaction := domain.Transaction{
		TransactionCode: "7401fd17d38171bsasd112dsx",
		Currency:        "CLP",
		Amount:          120000,
		Reciever:        "Luz Lucumi",
		Sender:          "Carlos Desidero",
		TransactionDate: dbTime,
	}
	expectedTransaction := domain.Transaction{
		TransactionCode: "7401fd17d38171bsasd112dsx",
		Currency:        "CLP",
		Amount:          120000,
		Reciever:        "Luz Lucumi",
		Sender:          "Carlos Desidero",
		TransactionDate: dbTime,
	}
	storeTransaction, errStore := repo.Store(initTransaction.TransactionCode, initTransaction.Currency, initTransaction.Amount, initTransaction.Sender, initTransaction.Reciever, initTransaction.TransactionDate)
	initTransaction.Id = storeTransaction.Id
	expectedTransaction.Id = storeTransaction.Id
	_, updateErr := repo.Update(storeTransaction.Id, expectedTransaction.Currency, expectedTransaction.Amount, expectedTransaction.Sender, expectedTransaction.Reciever)
	gettedProduct, errGet := repo.GetById(expectedTransaction.Id)

	assert.NoError(t, errTime)
	assert.NoError(t, errStore)
	assert.NoError(t, updateErr)
	assert.NoError(t, errGet)
	assert.Equal(t, initTransaction, storeTransaction)
	assert.Equal(t, expectedTransaction, gettedProduct)

}

func TestDelete(t *testing.T) {
	txdb.Register("txdb", "mysql", "meli_sprint_user:Meli_Sprint#123@/transactions?parseTime=true")
	db, _ := sql.Open("txdb", uuid.NewString())
	repo := NewRepository(db)
	dbTime, errTime := time.Parse("2006-01-02T15:04:05Z07:00", "2022-08-08T20:37:32Z")

	initTransaction := domain.Transaction{
		TransactionCode: "7401fd17d38171bsasd112dsx",
		Currency:        "CLP",
		Amount:          120000,
		Reciever:        "Luz Lucumi",
		Sender:          "Carlos Desidero",
		TransactionDate: dbTime,
	}

	storeProduct, errStore := repo.Store(initTransaction.TransactionCode, initTransaction.Currency, initTransaction.Amount, initTransaction.Sender, initTransaction.Reciever, initTransaction.TransactionDate)
	initTransaction.Id = storeProduct.Id
	beforeDeleteProduct, errBeforeGet := repo.GetById(storeProduct.Id)
	errDelete := repo.Delete(initTransaction.Id)
	gettedProduct, errGet := repo.GetById(storeProduct.Id)
	getAllProducts, errGetAll := repo.GetAll()

	assert.NoError(t, errTime)
	assert.NoError(t, errStore)
	assert.NoError(t, errBeforeGet)
	assert.NoError(t, errDelete)
	assert.Error(t, errGet)
	assert.NoError(t, errGetAll)

	assert.Equal(t, initTransaction, beforeDeleteProduct)
	assert.Empty(t, gettedProduct)
	assert.NotContains(t, getAllProducts, errGetAll)
}

func TestUpdateWithContext(t *testing.T) {
	txdb.Register("txdb", "mysql", "meli_sprint_user:Meli_Sprint#123@/transactions?parseTime=true")
	db, _ := sql.Open("txdb", uuid.NewString())
	repo := NewRepository(db)
	transactionId := 2
	transactionUpdate := domain.Transaction{
		Currency: "CLP",
		Amount:   50000,
		Reciever: "Jorge Figueroa",
		Sender:   "Luz Carime",
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	transaction, err := repo.UpdateWithContext(ctx, transactionId, transactionUpdate.Currency, transactionUpdate.Amount, transactionUpdate.Sender, transactionUpdate.Reciever)

	assert.NoError(t, err)
	assert.Equal(t, transactionUpdate, transaction)
}
