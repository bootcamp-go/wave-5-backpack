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

func TestSqlRepositoryGetOneMock(t *testing.T) {
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
	}

	columns := []string{"id", "name", "type", "count", "price"}
	rows := mock.NewRows(columns)
	rows.AddRow(product.ID, product.Name, product.Type, product.Count, product.Price)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, type, count, price FROM products WHERE id = ?")).WithArgs(product.ID).WillReturnRows(rows)

	//Act
	p, err := repo.GetOne(1)

	//Assert
	assert.NoError(t, err)
	assert.Equal(t, product, p)
}

func TestSqlRepositoryUpdateMock(t *testing.T) {
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
	}

	mock.ExpectPrepare(regexp.QuoteMeta("UPDATE products SET name = ?, type = ?, count = ?, price = ? WHERE id = ?"))
  mock.ExpectExec("").WithArgs(product.Name, product.Type, product.Count, product.Price, product.ID).WillReturnResult(sqlmock.NewResult(1,1)) 

	//Act
	p, err := repo.Update(product.ID, product.Name, product.Type, product.Count,product.Price)

	//Assert
	assert.NoError(t, err)
	assert.Equal(t, product, p)
}

func TestSqlRepositoryDeleteMock(t *testing.T) {
	//Arrange
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	repo := NewRepo(db)

	productID := 1

	mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM products WHERE id = ?"))
  mock.ExpectExec("").WithArgs(productID).WillReturnResult(sqlmock.NewResult(1,1)) 

	//Act
	err = repo.Delete(productID)

	//Assert
	assert.NoError(t, err)
}
