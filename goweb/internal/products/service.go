package products

type Service interface {
	GetAll()
}

// type Service interface {
// 	GetAll() ([]domain.products, error)
// 	GetByID(id int) (models.Products, error)
// 	Store(nombre, color, precio, stock, codigo, publicado, fechaCreacion) (models.Products, error)
// 	Update(id, nombre, color, precio, stock, codigo, publicado, fechaCreacion) (models.Products, error)
// 	UpdatePrecioStock(precio, stock) (models.Products, error)
// 	Delete(id int) (int, error)
// }

// func NewService(r Repository) Service {
// 	return &service{
// 		repository: r,
// 	}
// }

// type service struct {
// 	repository Repository
// }

// func (s service) Store(nombre, color, precio, stock, codigo, publicado, fechaCreacion) (models.Products, error)  {
// 	return s.repository.Store
// }
