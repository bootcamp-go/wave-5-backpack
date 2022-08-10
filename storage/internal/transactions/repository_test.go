package transactions

import (
	"context"
	"github.com/stretchr/testify/assert"
	"goweb/pkg/db"
	"testing"
)

func TestRepository_GetAll(t *testing.T) {
	//Arrange
	db := db.MySQLConnection()

	repository := NewRepository(db)
	ctx := context.TODO()

	//Act
	transactions, err := repository.GetAll(ctx)

	//Assert
	assert.NotNil(t, transactions)
	assert.Nil(t, err)
}
