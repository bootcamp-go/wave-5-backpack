package transactions

import (
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/storage/cmd/db"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryGetAll(t *testing.T) {
	//Arrange
	db, err := db.NewConnection()
	assert.Nil(t, err)

	repo := NewRepository(db)

	//Act
	transactions, err := repo.GetAll(&gin.Context{})

	//Assert
	assert.NotNil(t, transactions)
	assert.Nil(t, err)
}
