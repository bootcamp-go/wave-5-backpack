package transactions

import (
	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/storage/internal/models"
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetAll(ctx *gin.Context) ([]models.Transaction, error)
	GetByID(ctx *gin.Context, id int) (models.Transaction, error)
	Store(ctx *gin.Context, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error)
	Update(ctx *gin.Context, id int, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error)
	Patch(ctx *gin.Context, id int, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error)
	Delete(ctx *gin.Context, id int) (int, error)
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

type service struct {
	repository Repository
}

func (s service) Store(ctx *gin.Context, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error) {
	return s.repository.Store(ctx, monto, cod, moneda, emisor, receptor)
}

func (s service) GetAll(ctx *gin.Context) ([]models.Transaction, error) {
	return s.repository.GetAll(ctx)
}

func (s service) GetByID(ctx *gin.Context, id int) (models.Transaction, error) {
	return s.repository.GetByID(ctx, id)
}

func (s service) Update(ctx *gin.Context, id int, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error) {
	return s.repository.Update(ctx, id, monto, cod, moneda, emisor, receptor)
}

func (s service) Patch(ctx *gin.Context, id int, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error) {
	return models.Transaction{}, nil
}

func (s service) Delete(ctx *gin.Context, id int) (int, error) {
	return 0, nil
}
