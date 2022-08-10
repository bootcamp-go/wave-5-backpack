package products

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nictes1/storage-implementation/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestSqlRepositoryStoreMock(t *testing.T)  {
  //Arrange
  db, mock, err := sqlmock.New()
  assert.NoError(t, err)
  
  repo := NewRepo(db)

  product := domain.Product{
    ID: 1,
    Name: "destornillador",
    Type: "ferreteria",
    Count: 100,
    Price: 1000,
    Warehouse: "abc123",
    WarehouseAdress: "Calle Falsa 123",
  }

  mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO products(name, type, count, price) VALUES( ?, ?, ?, ? )"))
  mock.ExpectExec("").WithArgs(product.Name, product.Type, product.Count, product.Price).WillReturnResult(sqlmock.NewResult(1,1)) 
  //Act
  p, err := repo.Store(product)

  //Assert
  assert.NoError(t, err)
  assert.Equal(t, product, p)
  assert.Equal(t, product.ID, p.ID)
}
