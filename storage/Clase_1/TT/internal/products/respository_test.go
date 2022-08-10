package products

import (
	"context"
	"database/sql"
	"goweb/internal/domain"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {

	products := []domain.Product{
		{
			Id:           4,
			Name:         "televisor",
			Type:         "electro",
			Count:        3,
			Price:        399.99,
			Id_warehouse: 1,
		},
		{
			Id:           5,
			Name:         "laptop",
			Type:         "electro",
			Count:        10,
			Price:        299.99,
			Id_warehouse: 1,
		},
		{
			Id:           6,
			Name:         "Lapiz",
			Type:         "Libreria",
			Count:        50,
			Price:        990,
			Id_warehouse: 1,
		},
		{
			Id:           7,
			Name:         "Lapiz",
			Type:         "Libreria",
			Count:        50,
			Price:        990,
			Id_warehouse: 1,
		},
		{
			Id:           8,
			Name:         "Ipad",
			Type:         "Electronica",
			Count:        10,
			Price:        199.99,
			Id_warehouse: 1,
		},
		{
			Id:           9,
			Name:         "Botella",
			Type:         "Otros",
			Count:        10,
			Price:        1.99,
			Id_warehouse: 1,
		},
		{
			Id:           10,
			Name:         "Cuchara",
			Type:         "Hogar",
			Count:        100,
			Price:        290,
			Id_warehouse: 1,
		},
	}

	dataSource := "root:@tcp(localhost:3306)/storage"
	StorageDB, err := sql.Open("mysql", dataSource)
	assert.Nil(t, err)
	myRepo := NewRepository(StorageDB)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	productsResult, err := myRepo.GetAll(ctx)
	assert.Nil(t, err)
	assert.Equal(t, products, productsResult)

}

func TestStoreAndGetOneTXDB(t *testing.T) {

	txdb.Register("textdb", "mysql", "root:@tcp(localhost:3306)/storage")
	db, err := sql.Open("textdb", uuid.New().String())
	assert.NoError(t, err)

	repo := NewRepository(db)
	ctx := context.TODO()
	product := domain.Product{
		Name:         "Iphone",
		Type:         "Electro",
		Count:        253,
		Price:        599.990,
		Id_warehouse: 1,
	}

	p, err := repo.Store(product)
	product.Id = p.Id

	// Crear nuevo producto en el Store
	assert.NoError(t, err)
	assert.NotZero(t, p)

	// Verificando que el producto obtenido corresponda a lo esperado
	getProductExist, err := repo.GetOne(ctx, p.Id)
	assert.NoError(t, err)
	assert.Equal(t, product, getProductExist)

	// Se verifica que si el producto no existe se obtenga producto vacío
	getProductNonExist, err := repo.GetOne(ctx, 100)
	assert.NoError(t, err)
	assert.Zero(t, getProductNonExist)
}

func TestUpdateTXDB(t *testing.T) {

	txdb.Register("textdb", "mysql", "root:@tcp(localhost:3306)/storage")
	db, err := sql.Open("textdb", uuid.New().String())
	assert.NoError(t, err)

	repo := NewRepository(db)
	ctx := context.TODO()
	product := domain.Product{
		Id:           9,
		Name:         "PS5",
		Type:         "Electro",
		Count:        20,
		Price:        799.990,
		Id_warehouse: 1,
	}

	p, err := repo.Update(ctx, product)

	// Verificar producto actualizado en el Store
	assert.NoError(t, err)
	assert.NotZero(t, p)

	// Verificando que el producto obtenido corresponda a lo esperado
	assert.Equal(t, product, p)
}

func TestDeleteTXDB(t *testing.T) {

	txdb.Register("txdb", "mysql", "root:@tcp(localhost:3306)/storage")
	db, err := sql.Open("txdb", uuid.New().String())
	assert.NoError(t, err)

	repo := NewRepository(db)
	ctx := context.TODO()
	id := 5

	err = repo.Delete(ctx, id)

	// Verificar producto actualizado en el Store
	assert.NoError(t, err)

	// Verificando que el producto eliminado no se encuentre
	getProductNonExist, err := repo.GetOne(ctx, id)
	assert.NoError(t, err)
	assert.Zero(t, getProductNonExist)

}

func TestStoreAndGetOneSQLMOCK(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO products (name,type,count,price,id_warehouse) VALUES(?,?,?,?,?)"))
	mock.ExpectExec("INSERT INTO products").WillReturnResult(sqlmock.NewResult(1, 1))

	columns := []string{"id", "name", "type", "count", "price", "id_warehouse"}
	rows := sqlmock.NewRows(columns)
	newId := 1
	rows.AddRow(newId, "Iphone", "Electro", "2", "599.990", "1")
	mock.ExpectQuery("SELECT id,name,type,count,price,id_warehouse FROM products").WithArgs(newId).WillReturnRows(rows)

	repo := NewRepository(db)

	ctx := context.TODO()
	product := domain.Product{
		Name:         "Iphone",
		Type:         "Electro",
		Count:        2,
		Price:        599.990,
		Id_warehouse: 1,
	}

	product.Id = 1
	pObtenido, err := repo.GetOne(ctx, 1)

	//Obtener producto no creado y verificar que indique error y contenido igual a 0.
	assert.Error(t, err)
	assert.Zero(t, pObtenido)

	p, err := repo.Store(product)

	// Crear nuevo producto en el Store.
	assert.NoError(t, err)
	assert.NotZero(t, p)
	assert.Equal(t, product.Id, p.Id)

	pObtenido, err = repo.GetOne(ctx, 1)

	// Obtener producto creado.
	assert.NoError(t, err)
	assert.NotNil(t, pObtenido)
	assert.Equal(t, product.Id, pObtenido.Id)
	assert.Equal(t, product, pObtenido)

}

func TestGetAllConflic(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT id,name,type,count,price,id_warehouse FROM products").WillReturnError(sql.ErrConnDone)

	repo := NewRepository(db)

	ctx := context.TODO()
	result, err := repo.GetAll(ctx)

	assert.NoError(t, mock.ExpectationsWereMet())
	assert.Equal(t, sql.ErrConnDone, err)
	assert.Nil(t, result)

}
