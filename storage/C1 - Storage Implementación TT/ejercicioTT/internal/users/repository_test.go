package users

import (
	"database/sql"
	"ejercicioTT/internal/domain"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

var dataSource = "root@tcp(localhost:3306)/storage?parseTime=true"

func TestGetByName(t *testing.T) {
	//Arrange
	StorageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	repositorio := NewRepository(StorageDB)

	//Act
	usuario, err := repositorio.GetByName("Luz")

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, "Luz", usuario.Nombre)
}

func TestRepositoryStore(t *testing.T) {
	//Arrange
	usuario := domain.Usuarios{
		Nombre:   "Martha",
		Apellido: "Hernandez",
		Email:    "marthahg@gmail.es",
		Edad:     61,
		Altura:   1.65,
	}
	StorageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	repositorio := NewRepository(StorageDB)

	//Act
	resultUsuario, err := repositorio.Store(usuario)
	if err != nil {
		t.Errorf("Error al guardar el usuario: %v", err)
	}

	//Assert
	assert.Equal(t, usuario.Nombre, resultUsuario.Nombre)
}

func TestGetOne(t *testing.T) {
	//Arrange
	StorageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	repositorio := NewRepository(StorageDB)

	//Act
	usuario, err := repositorio.GetOne(4)

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, "Luz", usuario.Nombre)
	assert.Equal(t, "Lucumi Hernandez", usuario.Apellido)
}

func TestUpdate(t *testing.T) {
	//Arrange
	usuario := domain.Usuarios{
		Id:       39,
		Nombre:   "Martha Luisa Luisa Luz",
		Apellido: "Hernandez Gil",
		Email:    "marthahg@gmail.es",
		Edad:     61,
		Altura:   1.65,
	}
	StorageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	repositorio := NewRepository(StorageDB)

	//Act
	resultUsuario, err := repositorio.Update(usuario)
	if err != nil {
		t.Errorf("Error al actualizar el usuario: %v", err)
	}

	//Assert
	assert.Equal(t, usuario.Nombre, resultUsuario.Nombre)
}
