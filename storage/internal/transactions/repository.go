package transactions

import (
	"database/sql"
	"wave-5-backpack/storage/internal/domain"
)

type Repository interface {
	GetByName(transactionCode string) (domain.Transactions, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r repository) GetByName(transactionCode string) (domain.Transactions, error) {
	//TODO implement me
	panic("implement me")
}
