package products

import (
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/storage/pkg/store"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	repository := NewRepository(store.DBConnection())

	products, err := repository.GetAll()

	assert.Nil(t, err)
	assert.NotEmpty(t, products)
}
