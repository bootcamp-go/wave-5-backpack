package products

import (
	"bootcamp/wave-5-backpack/storage/internal/domain"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetOneWithContext(t *testing.T) {
	id := 1
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/storage")
	if err != nil {
		t.Fatal(err)
	}
	repo := NewRepo(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = repo.GetOneWithcontext(ctx, id)
	if err != nil {
		t.Errorf("err must be nil, but got %v", err)
	}
}

func TestRepositoryStore(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare("INSERT INTO products")
	//Otra opcion para hacer lo anterior con la sentencia completa
	//mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO products(name, type, count, price, warehouse_id) VALUES( ?, ?, ?, ?, ? )"))

	mock.ExpectExec("INSERT INTO products").WillReturnResult(sqlmock.NewResult(1, 1))

	productId := 1

	repo := NewRepo(db)
	product := domain.Product{
		ID:    productId,
		Name:  "prueba",
		Type:  "test",
		Count: 1,
		Price: 20.2,
	}

	p, err := repo.Store(product)
	assert.NoError(t, err)
	assert.NotZero(t, p)
	assert.Equal(t, product.ID, p.ID)
	assert.NoError(t, mock.ExpectationsWereMet())

}

func TestRepositoryStoreTXDB(t *testing.T) {
	txdb.Register("txdb", "mysql", "root@tcp(localhost:3306)/storage")

	db, err := sql.Open("txdb", uuid.New().String())
	assert.NoError(t, err)

	repo := NewRepo(db)
	product := domain.Product{
		Name:  "prueba",
		Type:  "test",
		Count: 1,
		Price: 20.2,
	}

	p, err := repo.Store(product)
	assert.NoError(t, err)
	assert.NotZero(t, p)
}
