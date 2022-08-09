package product

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/storage")
	if err != nil {
		assert.Fail(t, "an error ocurred", err.Error())
	}
	r := NewRepository(db)
	products, err := r.GetAll()
	if err != nil {
		assert.Fail(t, "an error ocurred", err.Error())
	}

	assert.Greater(t, len(products), 0)
}
