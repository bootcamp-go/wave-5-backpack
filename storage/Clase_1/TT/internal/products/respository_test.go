package products

import (
	"context"
	"database/sql"
	"goweb/internal/domain"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
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
