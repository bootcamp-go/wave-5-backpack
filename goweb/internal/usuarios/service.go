/*Servicio, debe contener la lógica de nuestra aplicación.
OK Se debe crear el archivo service.go.
OK Se debe generar la interface Service con todos sus métodos.
OK Se debe generar la estructura service que contenga el repositorio.
OK Se debe generar una función que devuelva el Servicio.
OK Se deben implementar todos los métodos correspondientes a las operaciones a realizar (GetAll, Store, etc..).
*/
package usuarios

import "github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"

type Service interface {
	GetAll() ([]domain.Usuarios, error)
	Guardar(nombre string, apellido string, email string, edad int, altura float64, actico bool, fecha string) (domain.Usuarios, error)
	Update(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fecha string) (domain.Usuarios, error)
	Delete(id int) error
	UpdateNameAndLastName(id int, name string, apellido string) (domain.Usuarios, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) UpdateNameAndLastName(id int, nombre, apellido string) (domain.Usuarios, error) {
	return s.repository.UpdateNameAndLastName(id, nombre, apellido)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) Update(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fecha string) (domain.Usuarios, error) {
	usuario, error := s.repository.Update(id, nombre, apellido, email, edad, altura, activo, fecha)
	return usuario, error
}

func (s *service) GetAll() ([]domain.Usuarios, error) {
	us, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return us, nil
}
func (s *service) Guardar(nombre string, apellido string, email string, edad int, altura float64, actico bool, fecha string) (domain.Usuarios, error) {
	lasNro, erro := s.repository.lastId()
	if erro != nil {
		return domain.Usuarios{}, erro
	}
	lasNro++
	usuario, error := s.repository.Guardar(lasNro, nombre, apellido, email, edad, altura, actico, fecha)
	if error != nil {
		return domain.Usuarios{}, error
	}
	return usuario, nil
}
