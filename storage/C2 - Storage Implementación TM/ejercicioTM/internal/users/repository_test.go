package users

import (
	"context"
	"database/sql"
	"ejercicioTM/internal/domain"
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
	usuario, err := repositorio.GetByName("Luber")

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, "Luber", usuario.Nombre)
}

func TestRepositoryStore(t *testing.T) {
	//Arrange
	usuario := domain.Usuarios{
		Nombre:   "Angela",
		Apellido: "Lucumi",
		Email:    "angela@gmail.es",
		Edad:     34,
		Altura:   1.60,
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
	assert.Equal(t, "Paquita", usuario.Nombre)
	assert.Equal(t, "Lucumi Hernandez", usuario.Apellido)
}

func TestUpdate(t *testing.T) {
	//Arrange
	usuario := domain.Usuarios{
		Id:       1,
		Nombre:   "Martha Luisa",
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
	resultUsuario, err := repositorio.Update(context.TODO(), usuario)
	if err != nil {
		t.Errorf("Error al actualizar el usuario: %v", err)
	}

	//Assert
	assert.Equal(t, usuario.Nombre, resultUsuario.Nombre)
}

func TestGetAll(t *testing.T) {
	//Arrange
	StorageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	//Act
	repositorio := NewRepository(StorageDB)
	usuarios, err := repositorio.GetAll()

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, 5, len(usuarios))
}
