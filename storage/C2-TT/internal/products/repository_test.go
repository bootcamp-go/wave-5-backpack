package products

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/nictes1/storage-implementation/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryStore(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err) //Verificamos poder crear el mock de base de datos
	defer db.Close()       // Cerramos la db (*sql.DB)
	//mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO products(name, type, count, price) VALUES( ?, ?, ?, ? )")) //para escapar de los metadatos "? ? ? ?"
	mock.ExpectPrepare("INSERT INTO products")                                        //Esperamos que se realice un prepare.
	mock.ExpectExec("INSERT INTO products").WillReturnResult(sqlmock.NewResult(1, 1)) //Esperamos que se realice un execute.

	productId := 1
	repo := NewRepo(db)
	user := domain.Product{
		ID:    productId,
		Name:  "Iphone",
		Type:  "Tecnologia",
		Count: 987,
		Price: 1200,
	}
	p, err := repo.Store(user)
	assert.NoError(t, err)
	assert.NotZero(t, p)
	assert.Equal(t, user.ID, p.ID)
	assert.NoError(t, mock.ExpectationsWereMet()) //Verificamos que los mock.Expect se ejecutaron y en orden

}

func TestRepositoryStoreTXDB(t *testing.T) { //Realizamos mock sobre la transaccion de la base de datos
	txdb.Register("txdb", "mysql", "root@tcp(localhost:3306)/storage") //nos conectamos a nuestra base de datos
	db, err := sql.Open("txdb", uuid.New().String())
	assert.NoError(t, err)
	//sql.Open recibe el driver de base de datos y un string de conexion
	repo := NewRepo(db) //Generamos nuestro repository
	user := domain.Product{
		Name:  "Iphone",
		Type:  "Tecnologia",
		Count: 987,
		Price: 1200,
	}
	p, err := repo.Store(user) //consulta el repo.
	assert.NoError(t, err)
	assert.NotZero(t, p)
}

func TestUpdateTXDB(t *testing.T) {
	txdb.Register("txdb", "mysql", "root@tcp(localhost:3306)/storage") //nos conectamos a nuestra base de datos
	db, err := sql.Open("txdb", uuid.New().String())                   //sql.Open recibe el driver de base de datos y un string de conexion
	assert.NoError(t, err)

	repo := NewRepo(db) //Generamos nuestro repository
	product := domain.Product{
		ID:    120,
		Name:  "Iphone",
		Type:  "Tecnologia",
		Count: 987,
		Price: 1200,
	}
	p, err := repo.Update(product.ID, product.Name, product.Type, product.Count, product.Price) //consulta el repo.
	assert.NoError(t, err)
	assert.Equal(t, nil, err)
	assert.NotZero(t, p)
}
func TestDeleteTXDB(t *testing.T) {
	txdb.Register("txdb", "mysql", "root@tcp(localhost:3306)/storage") //nos conectamos a nuestra base de datos
	db, err := sql.Open("txdb", uuid.New().String())                   //sql.Open recibe el driver de base de datos y un string de conexion
	assert.NoError(t, err)

	repo := NewRepo(db) //Generamos nuestro repository
	product := domain.Product{
		ID:    120,
		Name:  "Iphone",
		Type:  "Tecnologia",
		Count: 987,
		Price: 1200,
	}
	err = repo.Delete(product.ID) //consulta el repo.
	assert.NoError(t, err)
}
func TestRepositoryGetWithTimeout(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	productId := 1
	columns := []string{"id", "name", "type", "count", "price"}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(productId, "remera", "indumentaria", 3, 1500)
	mock.ExpectQuery("select id, name, type, count, price").WillDelayFor(30 * time.Second).WillReturnRows(rows)
	repo := NewRepo(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = repo.GetOneWithContext(ctx, productId)

	assert.Error(t, err)
}
