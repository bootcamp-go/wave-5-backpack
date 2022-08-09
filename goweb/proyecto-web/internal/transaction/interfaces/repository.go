package interfaces

import (
	"context"
	"proyecto-web/internal/domain"
)

type IRepository interface {
	GetAll() ([]domain.Transaction, error)
	Create(codigoTransaccion string, moneda string, monto float64, emisor string, receptor string, fecha string) (domain.Transaction, error)
	GetById(id int) (domain.Transaction, error)
	Update(ctx context.Context, id int, codigoTransaccion string, moneda string, monto float64, emisor string, receptor string, fecha string) (domain.Transaction, error)
	UpdateParcial(id int, codigoTransaccion string, monto float64) (domain.Transaction, error)
	Delete(id int) error
	GetByCodigoTransaccion(name string) (domain.Transaction, error)
}
