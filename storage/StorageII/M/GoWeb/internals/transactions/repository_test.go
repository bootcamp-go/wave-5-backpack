package transactions

import (
	"GoWeb/internals/domain"
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

var StorageDB *sql.DB

func TestGetByIdCtx(t *testing.T) {
	// usamos un Id que exista en la DB
	id := 1

	// definimos un Product cuyo nombre sea igual al registro de la DB
	transaccion := domain.Transanction{
		Receptor: "MePa",
	}
	dataSource := "root@tcp(localhost:3306)/storage"

	StorageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	myRepo := NewRepository(StorageDB)

	// se define un context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := myRepo.GetByIdCtx(ctx, id)
	fmt.Println(err)
	assert.Equal(t, transaccion.Receptor, result.Receptor)

}
