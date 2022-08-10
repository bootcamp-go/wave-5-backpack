package users

import (
	"context"
	"database/sql"
	"ejercicioTT/internal/domain"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

//Mock Store
func TestRepositoryStoreMock(t *testing.T) {
	//Inicializando el sqlmock
	db, mock, err := sqlmock.New()
	//Verifico que no suceda el error
	assert.NoError(t, err)
	//Cierro conexión
	defer db.Close()

	//Nos pide un string que es la query, se va a matchear por strings, por lo que la consulta resulta exitosa de esta manera porque no entiende la expresión ?,?,...

	//Si quiero usar toda la sentencia, utilizo regexp.QuoteMeta
	mock.ExpectPrepare("INSERT INTO usuarios")
	//inserta y afecta a una columna
	//id 1, 1 row afectado
	mock.ExpectExec("INSERT INTO usuarios").WillReturnResult(sqlmock.NewResult(1, 1))

	//Fase de creación para comparar expect y ejecución
	usuarioId := 1
	repo := NewRepository(db)
	usuario := domain.Usuarios{
		Id:       usuarioId,
		Nombre:   "Angela",
		Apellido: "Lucumi",
		Email:    "angela@gmail.es",
		Edad:     34,
		Altura:   1.60,
	}
	//Paso el usuario
	usuarioStore, err := repo.Store(usuario)
	//Asserts
	assert.NoError(t, err)
	assert.NotZero(t, usuarioStore)
	assert.Equal(t, usuario.Id, usuarioStore.Id)
}

func TestStoreTxdb(t *testing.T) {
	txdb.Register("txdb", "mysql", dataSource)
	db, err := sql.Open("txdb", uuid.New().String())
	assert.NoError(t, err)
	repo := NewRepository(db)

	ctx := context.TODO()
	usuario := domain.Usuarios{
		Nombre:   "Angela",
		Apellido: "Lucumi",
		Email:    "angela@gmail.es",
		Edad:     34,
		Altura:   1.60,
	}

	p, err := repo.Store(ctx, product)
	product.ID = p.ID

	assert.NoError(t, err)
	assert.NotZero(t, p)
	getResult, err := repo.GetOne(ctx, p.ID)
	assert.NoError(t, err)
	assert.Equal(t, product.Name, getResult.Name)
	assert.Equal(t, product.ID, getResult.ID)
}
