package products

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	cnn "github.com/bootcamp-go/storage/db"
	"github.com/bootcamp-go/storage/internal/domains"
	"github.com/bootcamp-go/storage/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/* Ejercicio 1 - Testear Store() y GetOne()
Suponiendo la correcta implementación de los métodos Store y GetOne. Escribir un test utilizando go-txdb en donde se guarde un User y luego se obtenga, en el caso de que no exista GetOne debe devolver un resultado nulo. No hay límite de cantidad de tests ni casos, crear los que considere necesarios para asegurar el correcto funcionamiento de las distintas variaciones. */

func Test_txtdbRepositoryStore(t *testing.T) {
	db, err := tests.InitDb()
	assert.NoError(t, err)
	defer db.Close()

	p := &domains.Product{
		Name:        "name test",
		Type:        "type test",
		Count:       9,
		Price:       9.9,
		WarehouseId: 1,
	}

	repo := NewRepository(db)
	ctx := context.TODO()
	id, err := repo.Store(ctx, *p)

	p.ID = id
	assert.NoError(t, err)

	product, err := repo.Get(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, p, &product)
}

/* Ejercicio 2 - Testear Update() y Delete()
Generar tests para update, en donde se verifique que luego de modificarse un modelo, al obtenerlo el mismo posea los cambios realizados.
Generar tests para delete para verificar que un registro fue borrado correctamente y ya no se puede obtener ni utilizando GetOne ni al llamar a GetAll. */

func Test_txtdbRepositoryUpdate(t *testing.T) {
	db, err := tests.InitDb()
	assert.NoError(t, err)
	defer db.Close()

	p := &domains.Product{
		Name:        "no actualizado",
		Type:        "no actualizado",
		Count:       100,
		Price:       99.9,
		WarehouseId: 1,
	}

	ctx := context.TODO()
	repo := NewRepository(db)
	id, err := repo.Store(ctx, *p)
	require.NoError(t, err)

	// Update
	p.ID = id
	p.Name = "actualizando name"
	err = repo.Update(ctx, *p)
	assert.NoError(t, err)
	assert.Nil(t, err)

	// Get and compare
	pUpdated, err := repo.Get(ctx, id)
	assert.NoError(t, err)
	assert.Nil(t, err)

	t.Log(p.Name, pUpdated.Name)
	assert.Equal(t, p, &pUpdated)
}

func Test_txtdbRepositoryDelete(t *testing.T) {
	db, err := tests.InitDb()
	assert.NoError(t, err)
	defer db.Close()

	p := &domains.Product{
		Name:        "nombre test",
		Type:        "tipo test",
		Count:       100,
		Price:       99.9,
		WarehouseId: 1,
	}

	ctx := context.TODO()
	repo := NewRepository(db)
	id, err := repo.Store(ctx, *p)
	require.NoError(t, err)

	// Update
	err = repo.Delete(ctx, id)
	assert.NoError(t, err)
	assert.Nil(t, err)

	// Get and compare
	exists, err := repo.Get(ctx, id)
	assert.NotNil(t, err)
	assert.Empty(t, exists)

	existsGetAll, err := repo.GetAll(ctx, id)
	assert.Nil(t, err)
	assert.Empty(t, existsGetAll)
}

/* Ejercicio 3 - Replicar tests anteriores utilizando mocks
Tomar alguno de los tests realizados en los ejercicios anteriores (o todos) y replicarlos utilizando go-sqlmock. */
func Test_sqlRepositoryStore(t *testing.T) {
	db, mock, err := tests.NewDBMock(t)
	assert.NoError(t, err)
	defer db.Close()

	p := &domains.Product{
		ID:          1,
		Name:        "name test",
		Type:        "type test",
		Count:       9,
		Price:       9.9,
		WarehouseId: 1,
	}

	ctx := context.TODO()
	repo := NewRepository(db)

	mock.ExpectPrepare(regexp.QuoteMeta(INSERT_QUERY)).
		ExpectExec().
		WithArgs(p.Name, p.Type, p.Count, p.Price, p.WarehouseId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Store
	id, err := repo.Store(ctx, *p)
	assert.NoError(t, err)
	assert.Equal(t, p.ID, id)

	columns := []string{"id", "name", "type", "count", "price", "id_warehouse"}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(p.ID, p.Name, p.Type, p.Count, p.Price, p.WarehouseId)
	mock.ExpectPrepare(regexp.QuoteMeta(GET_BY_ID_QUERY)).
		ExpectQuery().
		WithArgs(p.ID).
		WillReturnRows(rows)

	// Get
	product, err := repo.Get(ctx, p.ID)
	assert.NoError(t, err)
	assert.Equal(t, p, &product)
}

/*
Ejercicio 4 - Testear que sucede en el caso de que falle una query
Seleccione alguno de los métodos (Store, GetOne, GetAll, Update o Delete) y simule un error al ejecutar la consulta. El test debe pasar si ocurre un error. (assert.Error). El objetivo de este ejercicio es conocer la forma de ver como reacciona nuestra implementación frente a una falla.
*/
func Test_sqlRpositoryStoreError(t *testing.T) {
	db, mock, err := tests.NewDBMock(t)
	assert.NoError(t, err)
	defer db.Close()

	p := &domains.Product{}

	ctx := context.TODO()
	repo := NewRepository(db)

	mock.ExpectPrepare(regexp.QuoteMeta(INSERT_QUERY)).
		ExpectExec().
		WithArgs(p.Name, p.Type, p.Count, p.Price, p.WarehouseId).
		WillReturnError(errors.New("you have not provided the necessary fields to insert"))

	// Store failed
	id, err := repo.Store(ctx, *p)
	assert.Equal(t, 0, id)
	assert.NotNil(t, err)
	assert.Error(t, err)
}

func TestGetAll(t *testing.T) {
	db := cnn.MySQLConnection()
	repo := NewRepository(db)

	// recordar tener el producto con id 1 y tener un warehouse asociado a este producto
	products, err := repo.GetAll(context.Background(), 1)

	assert.NoError(t, err)
	assert.True(t, len(products) > 0)
}
