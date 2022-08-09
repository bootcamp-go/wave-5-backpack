package transactions

import (
	"GoWeb/internals/domain"
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var StorageDB *sql.DB

func TestGetByIdCtx(t *testing.T) {
	// usamos un Id que exista en la DB
	id := 1

	// definimos un Product cuyo nombre sea igual al registro de la DB
	transaccion := domain.Transanction{
		Receptor: "MePa",
	}
	dataSource := "root@tcp(localhost:3306)/storage"

	StorageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	myRepo := NewRepository(StorageDB)

	// se define un context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := myRepo.GetByIdCtx(ctx, id)
	fmt.Println(err)
	assert.Equal(t, transaccion.Receptor, result.Receptor)

}

// PUNTO 1
func TestRepositoryStoreTXDB(t *testing.T) {
	txdb.Register("txdb", "mysql", "root@tcp(localhost:3306)/storage") //conexion a la db
	db, err := sql.Open("txdb", uuid.New().String())                   //recibe el driver de conexion
	if err != nil {
		panic(err)
	}
	repo := NewRepository(db)

	tran := domain.Transanction{
		Code:     "AAA",
		Coin:     "USD",
		Amount:   1000,
		Emisor:   "Juan",
		Receptor: "MePa",
		Date:     "2022-08-09",
	}

	res, err := repo.Store(tran.Code, tran.Coin, tran.Amount, tran.Emisor, tran.Receptor, tran.Date)
	assert.NoError(t, err)
	assert.NotZero(t, res)
}

func TestRespositoryGetByIdTXDB(t *testing.T) {
	txdb.Register("txdb", "mysql", "root@tcp(localhost:3306)/storage") //conexion a la db
	db, err := sql.Open("txdb", uuid.New().String())                   //recibe el driver de conexion
	if err != nil {
		panic(err)
	}
	repo := NewRepository(db)

	id := 1
	tran := domain.Transanction{
		Id:       id,
		Code:     "AAA",
		Coin:     "USD",
		Amount:   1000,
		Emisor:   "Juan",
		Receptor: "MePa",
		Date:     "2022-02-02",
	}

	res, err := repo.GetById(tran.Id)
	assert.Nil(t, err)
	assert.Equal(t, tran, res)
}

// PUNTO 2

func TestRespositoryUpdateTXDB(t *testing.T) {
	txdb.Register("txdb", "mysql", "root@tcp(localhost:3306)/storage") //conexion a la db
	db, err := sql.Open("txdb", uuid.New().String())                   //recibe el driver de conexion
	if err != nil {
		panic(err)
	}
	repo := NewRepository(db)

	id := 1
	/*tran := domain.Transanction{
		Id:       id,
		Code:     "AAA",
		Coin:     "USD",
		Amount:   1000,
		Emisor:   "Juan",
		Receptor: "MePa",
		Date:     "2022-02-02",
	}*/
	exp := domain.Transanction{
		Id:       id,
		Code:     "AAA",
		Coin:     "COP",
		Amount:   1000,
		Emisor:   "Juan",
		Receptor: "MePa",
		Date:     "2022-02-02",
	}
	res, err := repo.Update(id, exp.Code, exp.Coin, exp.Amount, exp.Emisor, exp.Receptor, exp.Date)
	assert.Nil(t, err)
	assert.Equal(t, exp, res)
}

func TestRespositoryDeleteTXDB(t *testing.T) {
	txdb.Register("txdb", "mysql", "root@tcp(localhost:3306)/storage") //conexion a la db
	db, err := sql.Open("txdb", uuid.New().String())                   //recibe el driver de conexion
	if err != nil {
		panic(err)
	}
	repo := NewRepository(db)

	id := 1

	exp := domain.Transanction{
		Id:       id,
		Code:     "AAA",
		Coin:     "COP",
		Amount:   1000,
		Emisor:   "Juan",
		Receptor: "MePa",
		Date:     "2022-02-02",
	}
	res, err := repo.Delete(id)
	assert.Nil(t, err)
	assert.NotEqual(t, exp, res)
}

//PUNTO 3
func TestRespositoryUpdateSQLMock(t *testing.T) {

	db, mock, err := sqlmock.New()

	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPrepare("UPDATE transactions SET")
	mock.ExpectExec("UPDATE transactions SET").WillReturnResult(sqlmock.NewResult(1, 1))

	id := 1
	repo := NewRepository(db)
	/*tran := domain.Transanction{
		Id:       id,
		Code:     "AAA",
		Coin:     "USD",
		Amount:   1000,
		Emisor:   "Juan",
		Receptor: "MePa",
		Date:     "2022-02-02",
	}*/
	exp := domain.Transanction{
		Id:       id,
		Code:     "AAA",
		Coin:     "COP",
		Amount:   1000,
		Emisor:   "Juan",
		Receptor: "MePa",
		Date:     "2022-02-02",
	}
	res, err := repo.Update(id, exp.Code, exp.Coin, exp.Amount, exp.Emisor, exp.Receptor, exp.Date)
	assert.Nil(t, err)
	assert.Equal(t, exp, res)
	assert.Equal(t, exp.Id, id)
}

func TestRespositoryDeleteSQLMock(t *testing.T) {

	db, mock, err := sqlmock.New()

	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPrepare("DELETE FROM transactions WHERE")
	mock.ExpectExec("DELETE FROM transactions WHERE").WillReturnResult(sqlmock.NewResult(1, 1))

	id := 1
	repo := NewRepository(db)
	/*tran := domain.Transanction{
		Id:       id,
		Code:     "AAA",
		Coin:     "USD",
		Amount:   1000,
		Emisor:   "Juan",
		Receptor: "MePa",
		Date:     "2022-02-02",
	}*/
	exp := domain.Transanction{
		Id:       id,
		Code:     "AAA",
		Coin:     "COP",
		Amount:   1000,
		Emisor:   "Juan",
		Receptor: "MePa",
		Date:     "2022-02-02",
	}
	res, err := repo.Delete(id)
	assert.Nil(t, err)
	assert.NotEqual(t, exp, res)
}

//PUNTO 4
