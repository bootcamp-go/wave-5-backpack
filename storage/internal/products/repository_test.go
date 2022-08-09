package products

import (
	"context"
	"database/sql"
	"storage/internal/domain"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
	_ "github.com/go-sql-driver/mysql"
)

func TestUpdateWithContext(t *testing.T) {

	productTest := domain.Products{
		Nombre: "Sandia",
		Color:  "verdoso",
		Precio: 2132222,
		Stock:  2,
	}

	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/storage")
	if err != nil {
		t.Fatal(err)
	}
	repo := InitRepository(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = repo.Update(ctx, productTest)
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func TestGetAll(t *testing.T) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/storage")
	if err != nil {
		t.Fatal(err)
	}

	expect := []domain.Products{
		{
			Id:            1,
			Nombre:        "Sandia",
			Color:         "verdoso",
			Precio:        2132220,
			Stock:         2,
			Codigo:        "21fe2",
			Publicado:     true,
			FechaCreacion: "23/10/2022",
		},
		{
			Id:            2,
			Nombre:        "Guayaba",
			Color:         "Amarillo",
			Precio:        323123,
			Stock:         24,
			Codigo:        "3efe2",
			Publicado:     true,
			FechaCreacion: "30/10/2022",
		},
		{
			Id:            3,
			Nombre:        "Pera",
			Color:         "Verde",
			Precio:        3213430,
			Stock:         43,
			Codigo:        "y2h76",
			Publicado:     true,
			FechaCreacion: "14/11/2022",
		},
	}

	repo := InitRepository(db)

	r, err := repo.GetAll()

	assert.Equal(t, err, nil)
	assert.Equal(t, expect, r)
}
