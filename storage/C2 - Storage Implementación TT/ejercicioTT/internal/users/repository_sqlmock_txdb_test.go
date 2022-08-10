package users

import (
	"context"
	"database/sql"
	"ejercicioTT/internal/domain"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var conn = "root@tcp(localhost:3306)/storage?parseTime=true"

//Ejercicio en clase
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

//Ejercicio 1 - Testear Store y GetOne()
//go-txdb
func TestRepositoryStoreTXDB(t *testing.T) {
	//Mockeando la transacción en la base de datos
	txdb.Register("txdb", "mysql", conn)

	//sql.Open recibe el driver de base de datos y un string de conexión
	//uuid nos da un identificador único
	db, err := sql.Open("txdb", uuid.New().String())
	//Se comprueba que no haya error
	assert.NoError(t, err)

	//Generando un repository
	repo := NewRepository(db)
	usuario := domain.Usuarios{
		Nombre:   "Angela",
		Apellido: "Lucumi",
		Email:    "angela@hotmail.es",
		Edad:     34,
		Altura:   1.60,
	}

	//Consulta el repo
	user, err := repo.Store(usuario)
	usuario.Id = user.Id

	assert.NoError(t, err)
	//Verifico que esté completa la estructura
	assert.NotZero(t, user)

	result, err := repo.GetOne(user.Id)
	assert.NoError(t, err)
	assert.Equal(t, user.Nombre, result.Nombre)
	assert.Equal(t, user.Id, result.Id)
}

//Ejercicio 2 - Testear Update() y Delete()
//Test Update y GetOne para verificar modificación
func TestRepositoryUpdateTXDB(t *testing.T) {
	txdb.Register("txdb", "mysql", conn)

	db, err := sql.Open("txdb", uuid.New().String())
	if err != nil {
		panic(err)
	}

	repo := NewRepository(db)
	ctx := context.TODO()

	usuarioId := 1
	usuario := domain.Usuarios{
		Id:       usuarioId,
		Nombre:   "Angela",
		Apellido: "Lucumi",
		Email:    "angela@gmail.es",
		Edad:     34,
		Altura:   1.60,
	}

	user, err := repo.Update(ctx, usuario)
	assert.NoError(t, err)
	assert.NotZero(t, user)

	result, err := repo.GetOne(user.Id)
	assert.NoError(t, err)
	assert.Equal(t, usuario.Nombre, result.Nombre)
	assert.Equal(t, usuario.Id, result.Id)
}

//Test Delete: GetOne y GetAll para obtener el Zero value al eliminarlo
func TestRepositoryDeleteTXDB(t *testing.T) {
	txdb.Register("txdb", "mysql", conn)
	db, err := sql.Open("txdb", uuid.New().String())
	if err != nil {
		panic(err)
	}
	repo := NewRepository(db)
	idUsuario := 1

	err = repo.Delete(idUsuario)
	assert.NoError(t, err)

	resultOne, err := repo.GetOne(idUsuario)
	assert.NoError(t, err)
	assert.Zero(t, resultOne)

	resultAll, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Zero(t, resultAll)
}

//Ejercicio 3 - Replicar tests anteriors con mocks
//go-sqlmock para Update y Delete
func TestRepositoryUpdateMock(t *testing.T) {
	//Inicializando el sqlmock
	db, mock, err := sqlmock.New()
	//Verifico que no suceda el error
	assert.NoError(t, err)
	//Cierro conexión
	defer db.Close()

	//Si quiero usar toda la sentencia, utilizo regexp.QuoteMeta
	mock.ExpectPrepare("UPDATE usuarios SET")
	//inserta y afecta a una columna
	//id 1, 1 row afectado
	mock.ExpectExec("UPDATE usuarios SET").WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewRepository(db)

	usuario := domain.Usuarios{
		Id:       1,
		Nombre:   "Angela",
		Apellido: "Lucumi",
		Email:    "angela@gmail.es",
		Edad:     34,
		Altura:   1.60,
	}
	//Actualizo usuario
	user, err := repo.Update(context.TODO(), usuario)
	//Asserts
	assert.NoError(t, err)
	assert.Equal(t, usuario.Id, user.Id)
	assert.Equal(t, usuario, user)
}

func TestRepositoryDeleteMock(t *testing.T) {
	//Inicializando el sqlmock
	db, mock, err := sqlmock.New()
	//Verifico que no suceda el error
	assert.NoError(t, err)
	//Cierro conexión
	defer db.Close()

	mock.ExpectPrepare("DELETE FROM usuarios WHERE")
	//inserta y afecta a una columna
	//id 1, 1 row afectado
	mock.ExpectExec("DELETE FROM usuarios WHERE").WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewRepository(db)
	//Elimino usuario con id 1
	err = repo.Delete(1)
	//Asserts
	assert.NoError(t, err)
}

//Ejercicio 4 - Testear en caso de fallo de un query
//StoreMockError
func TestStoreMockError(t *testing.T) {
	//Inicializando el sqlmock
	db, mock, err := sqlmock.New()
	//Verifico que no suceda el error
	assert.NoError(t, err)
	//Cierro conexión
	defer db.Close()

	mock.ExpectPrepare("INSERT INTO usuarios")
	//inserta y afecta a una columna
	//id 1, 1 row afectado
	mock.ExpectExec("INSERT INTO usuarios").WillReturnError(fmt.Errorf("error creating the user"))

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
	expectedError := fmt.Errorf("error creating the user")
	//Paso el usuario
	usuarioStore, errStore := repo.Store(usuario)
	//Asserts
	assert.Error(t, expectedError, errStore)
	assert.NotEqual(t, usuario.Id, usuarioStore.Id)
}
