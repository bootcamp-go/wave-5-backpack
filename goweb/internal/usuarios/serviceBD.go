package usuarios

import (
	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
)

type ServiceBD interface {
	GetByName(name string) (domain.Usuarios, error)
}
type serviceBD struct {
	repoBD RepositoryBD
}

func NewServiceBD(rB RepositoryBD) ServiceBD {
	return &serviceBD{
		repoBD: rB,
	}
}
func (s *serviceBD) GetByName(name string) (domain.Usuarios, error) {
	us, err := s.repoBD.GetByName(name)
	if err != nil {
		return domain.Usuarios{}, err
	}
	return us, nil
}
