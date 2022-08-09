package product

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DATA-DOG/go-txdb"
	"github.com/google/uuid"

	"github.com/bootcamp-go/wave-5-backpack/tree/ospina_christian/storage/storageprueba/internal/domain"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

var (
	StorageDB *sql.DB
	err       error
	path      string = "meli_sprint_user:Meli_Sprint#123@/storage"
)

func TestStore(t *testing.T) {
	producto := domain.Productos{
		Nombre:        "Frijoles",
		Color:         "Rojos",
		Precio:        2500,
		Stock:         15,
		Codigo:        "1f",
		Publicado:     true,
		FechaCreación: "2021-03-15",
		IdWarehouse:   1,
	}
	StorageDB, err = sql.Open("mysql", path)

	assert.Nil(t, err)

	myrepo := NewRepo(StorageDB)
	productResult, err := myrepo.Store(producto)
	if err != nil {
		log.Println(err)
	}
	assert.Equal(t, producto.Nombre, productResult.Nombre)
}

func TestGetOne(t *testing.T) {
	producto := domain.Productos{
		Nombre:        "Mandarinas Francesas",
		Color:         "Naranjas",
		Precio:        1500,
		Stock:         50,
		Codigo:        "12mand",
		Publicado:     true,
		FechaCreación: "2021-04-06",
		IdWarehouse:   1,
	}

	StorageDB, err = sql.Open("mysql", path)
	assert.Nil(t, err)

	myrepo := NewRepo(StorageDB)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	productResult, err := myrepo.GetOne(ctx, 2)
	if err != nil {
		log.Println(err)
	}
	assert.Equal(t, producto.Nombre, productResult.Nombre)
}

func TestUpdate(t *testing.T) {
	//arrange
	StorageDB, err = sql.Open("mysql", path)
	assert.Nil(t, err)
	myrepo := NewRepo(StorageDB)

	producto := domain.Productos{
		Nombre:        "Mandarinas Francesas",
		Color:         "Naranjas",
		Precio:        1500,
		Stock:         50,
		Codigo:        "12mand",
		Publicado:     true,
		FechaCreación: "2021-04-06",
		IdWarehouse:   1,
	}

	//act
	idtest := 2
	producto.Id = idtest
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	productResult, err := myrepo.Update(ctx, idtest, producto.Nombre, producto.Color, producto.Precio, producto.Stock, producto.Codigo, producto.Publicado, producto.FechaCreación, producto.IdWarehouse)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, producto, productResult)
}

func TestGetAll(t *testing.T) {
	//arrange
	StorageDB, err = sql.Open("mysql", path)
	assert.Nil(t, err)
	myrepo := NewRepo(StorageDB)

	producto := domain.Productos{
		Nombre:        "Mandarinas Francesas",
		Color:         "Naranjas",
		Precio:        1500,
		Stock:         50,
		Codigo:        "12mand",
		Publicado:     true,
		FechaCreación: "2021-04-06",
		IdWarehouse:   1,
	}

	//act
	idtest := 2
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	producto.Id = idtest
	productResult, err := myrepo.GetAll(ctx)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, producto, productResult[idtest-1])
}

func TestDelete(t *testing.T) {
	//arrange
	StorageDB, err = sql.Open("mysql", path)
	assert.Nil(t, err)
	myrepo := NewRepo(StorageDB)

	//act
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	idtest := 5
	err := myrepo.Delete(ctx, idtest)

	//assert
	assert.Nil(t, err)
}

func TestGetAllProdWare(t *testing.T) {
	//arrange
	StorageDB, err = sql.Open("mysql", path)
	assert.Nil(t, err)
	myrepo := NewRepo(StorageDB)

	//acts
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	idtest := 1
	prodware := myrepo.GetAllProdWare(ctx, idtest)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, domain.Productos{}, prodware)
}

//TEST go-txdb

//Store
func TestStoretxdb(t *testing.T) {
	//arrange
	txdb.Register("txdb", "mysql", path)
	db, err := sql.Open("txdb", uuid.New().String())
	assert.Nil(t, err)
	repo := NewRepo(db)

	producto := domain.Productos{
		Nombre:        "Frijoles",
		Color:         "Rojos",
		Precio:        2500,
		Stock:         15,
		Codigo:        "1f",
		Publicado:     true,
		FechaCreación: "2021-03-15",
		IdWarehouse:   1,
	}

	//act
	productResult, err := repo.Store(producto)
	producto.Id = productResult.Id

	//assert
	assert.NoError(t, err)
	assert.NotZero(t, productResult)
	assert.Equal(t, producto, productResult)
}

//GetOne
//GetOne caso de existo, se valida que no haya error, que el contenido sea el mismo.
func TestGetOnetxdb(t *testing.T) {
	//arrange
	txdb.Register("txdb", "mysql", path)
	db, err := sql.Open("txdb", uuid.New().String())
	assert.Nil(t, err)

	repo := NewRepo(db)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	producto := domain.Productos{
		Nombre:        "Mandarinas Francesas",
		Color:         "Naranjas",
		Precio:        1500,
		Stock:         50,
		Codigo:        "12mand",
		Publicado:     true,
		FechaCreación: "2021-04-06",
		IdWarehouse:   1,
	}

	//act
	productResult, err := repo.GetOne(ctx, 2)

	//assert
	assert.NoError(t, err)
	assert.Equal(t, producto.Nombre, productResult.Nombre)
}

//ERROR getOne al intentar obtener un producto que no existe.
func TestGetOneFailtxdb(t *testing.T) {
	//arrange
	txdb.Register("txdb", "mysql", path)
	db, err := sql.Open("txdb", uuid.New().String())
	assert.Nil(t, err)
	repo := NewRepo(db)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//act
	productResult, err := repo.GetOne(ctx, 140)

	//assert
	assert.NoError(t, err)
	assert.Zero(t, productResult)
}

//Update
func TestUpdatetxdb(t *testing.T) {
	//arrange
	txdb.Register("txdb", "mysql", path)
	db, err := sql.Open("txdb", uuid.New().String())
	assert.Nil(t, err)

	repo := NewRepo(db)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	producto := domain.Productos{
		Nombre:        "Mandarinas Francesas",
		Color:         "Naranjas",
		Precio:        1500,
		Stock:         76,
		Codigo:        "12mand",
		Publicado:     true,
		FechaCreación: "2021-04-06",
		IdWarehouse:   1,
	}

	//act
	idtest := 2
	productResult, err := repo.Update(ctx, idtest, producto.Nombre, producto.Color, producto.Precio, producto.Stock, producto.Codigo, producto.Publicado, producto.FechaCreación, producto.IdWarehouse)
	producto.Id = productResult.Id

	//assert
	assert.Nil(t, err)
	assert.Equal(t, producto.Precio, productResult.Precio)
}

//Delete
func TestDeletetxdb(t *testing.T) {
	//arrange
	txdb.Register("txdb", "mysql", path)
	db, err := sql.Open("txdb", uuid.New().String())
	assert.Nil(t, err)

	repo := NewRepo(db)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//act
	idtest := 1
	err = repo.Delete(ctx, idtest)

	//assert
	assert.Error(t, err)
}

//Usando MOCKS sqlmock

//Store
func TestStoresqlmock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPrepare("INSERT INTO products")
	mock.ExpectExec("INSERT INTO products").WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewRepo(db)

	producto := domain.Productos{
		Nombre:        "Mandarinas Francesas",
		Color:         "Naranjas",
		Precio:        1500,
		Stock:         76,
		Codigo:        "12mand",
		Publicado:     true,
		FechaCreación: "2021-04-06",
		IdWarehouse:   1,
	}
	productResult, err := repo.Store(producto)
	producto.Id = productResult.Id

	//assert
	assert.NoError(t, err)
	assert.NotZero(t, productResult)
	assert.Equal(t, producto, productResult)
	assert.NoError(t, mock.ExpectationsWereMet())
}

//GetOne

func TestGetOnesqlmock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	productId := 1
	columns := []string{"id", "nombre", "color", "precio", "stock", "codigo", "publicado", "fecha_creacion", "id_warehouse"}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(productId, "Mandarinas Francesas", "Naranjas", 1500, 76, "12mand", true, "2021-04-06", 1)
	mock.ExpectQuery("SELECT id, nombre, color, precio, stock, codigo, publicado, fechaCreacion, id_warehouse").WillDelayFor(30 * time.Second).WillReturnRows(rows)
	repo := NewRepo(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = repo.GetOne(ctx, productId)

	assert.Error(t, err)
}
