package users

import (
	"goweb/clase4_parte2/internal/domain"
	"goweb/clase4_parte2/pkg/store"
	"fmt"
	"time"
)

const (
	UserNotFound = "Usuario %d no encontrado"
	ReadingError = "No se puede leer la BD"
	WritingError = "No se puede escribir en la BD, error: %w"
)


// Se genera la interface Repository con todos sus métodos
type Repository interface {
	GetAll() ([]domain.User, error)
	Store(nombre, apellido, email string, edad int, altura float64, activo *bool) (domain.User, error)
	Update(id int, nombre, apellido, email string, edad int, altura float64, activo *bool) (domain.User, error)
	UpdateLastNameAndAge(id int, apellido string, edad int) (domain.User, error)
	Delete(id int) error
}

/* Se genera la estructura repository y se declara el campo de tipo Store que se 
importará del paquete store */
type repository struct {
	db store.Store
}

// Se genera una función que devuelve el Repositorio
func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

/* Se implementan todos los métodos correspondientes a las operaciones a realizar (GetAll, Store, etc..) */

func (r *repository) GetAll() ([]domain.User, error) {
	
	/* Para obtener los usuarios se declara un slice de Usuarios y se le pasa la variable al método
	Read que es responsable de cargarla con la información del archivo, la cual finalmente
	será retornada por el método */
	var users []domain.User
	if err := r.db.Read(&users); err != nil {
		return nil, fmt.Errorf(ReadingError)
	}

	if len(users) == 0 {
		return users, fmt.Errorf("No hay usuarios registrados aún")
	}

	return users, nil
}

func (r *repository) Store(nombre, apellido, email string, edad int, altura float64, activo *bool) (domain.User, error) {
	
	var users []domain.User

	/* Se obtienen los usuarios desde el archivo json y le pasamos la referencia de la variable User
	para que al momento de ser modificada dentro del método Read también esa modificación se vea 
	reflejada fuera de él */
	if err := r.db.Read(&users); err != nil {
		return domain.User{}, fmt.Errorf(ReadingError)
	}
	
	var lastID int
	if len(users) > 0 {
		lastID = users[len(users)-1].ID
	} 

	lastID++
	fechaCreacion := time.Now().Format("02/01/2006")
	
	u := domain.User{lastID, nombre, apellido, email, edad, altura, activo, fechaCreacion}
	users = append(users, u)
	
	if err := r.db.Write(users); err != nil {
		return domain.User{}, fmt.Errorf(WritingError, err)
	}
	return u, nil	
}

/* Se implementa la funcionalidad para actualizar el usuario en memoria, en caso que coincida
con el ID enviado; en caso contrario, retorna un error */
func (r *repository) Update(id int, nombre, apellido, email string, edad int, altura float64, activo *bool) (domain.User, error) {
	
	var users []domain.User
	if err := r.db.Read(&users); err != nil {
		return domain.User{}, fmt.Errorf(ReadingError)
	}
	
	u := domain.User{Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura, Activo: activo}
	updated := false
	for i := range users {
		if users[i].ID == id {
			u.ID = id
			u.FechaCreacion = users[i].FechaCreacion
			users[i] = u
			updated = true
		}
	}

	if !updated {
		return domain.User{}, fmt.Errorf(UserNotFound, id)
	}

	if err := r.db.Write(&users); err != nil {
		return domain.User{}, fmt.Errorf(WritingError, err)
	}

	return u, nil
}

/* Se implementa la funcionalidad para actualizar el nombre y la edad del usuario en memoria,
en caso que coincida con el ID enviado; en caso contrario, se retorna un error */
func (r *repository) UpdateLastNameAndAge(id int, apellido string, edad int) (domain.User, error) {
	
	var users []domain.User
	if err := r.db.Read(&users); err != nil {
		return domain.User{}, fmt.Errorf(ReadingError)
	}
	
	var u domain.User
	updated := false
	for i := range users {
		if users[i].ID == id {
			users[i].Apellido = apellido
			users[i].Edad = edad
			updated = true
			u = users[i]
		}
	}
	if !updated {
		return domain.User{}, fmt.Errorf(UserNotFound, id)
	}

	if err := r.db.Write(&users); err != nil {
		return domain.User{}, fmt.Errorf(WritingError, err)
	}

	return u, nil
}

/* Se implementa la funcionalidad para eliminar el usuario en memoria, en caso que 
coincida con el ID enviado; en caso contrario, retorna un error */
func (r *repository) Delete(id int) error {

	var users []domain.User
	if err := r.db.Read(&users); err != nil {
		return fmt.Errorf(ReadingError)
	}

	if len(users) == 0 {
		return fmt.Errorf("No hay usuarios registrados aún")
	}

	deleted := false
	var index int
	for i := range users {
		if users[i].ID == id {
			index = i
			deleted = true
		} 
	}
	if !deleted {
		return fmt.Errorf(UserNotFound, id)
	}

	users = append(users[:index], users[index+1:]...)
	if err := r.db.Write(&users); err != nil {
		return fmt.Errorf(WritingError, err)
	}

	return nil
}



