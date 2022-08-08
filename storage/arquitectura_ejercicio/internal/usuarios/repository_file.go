/*package usuarios

import (
	"errors"

	"github.com/anesquivel/wave-5-backpack/storage/arquitectura_ejercicio/internal/domain"
	"github.com/anesquivel/wave-5-backpack/storage/arquitectura_ejercicio/pkg/store"
)

const (
	STRING_HUBO       = "Hubo un error ... "
	ERROR_READING     = "Hubo un error al leer los datos de la BD."
	ERR_WRITING       = "Hubo un error al guardar los datos en la BD."
	ERR_UPDATING_USER = ".. al no encontrar el usuario a actualizar"
)

type Repository interface {
	GetAll() ([]domain.Usuario, error)
	Store(id, age int, names, lastname, email, dateCreated string, estatura float64) (domain.Usuario, error)
	LastID() (int, error)
	Update(id, age int, names, lastname, email, dateCreated string, estatura float64, activo bool) (domain.Usuario, error)
	UpdateLastNameAndAge(id, age int, lastname string) (domain.Usuario, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func (r *repository) GetAll() ([]domain.Usuario, error) {
	var localUserList []domain.Usuario

	if err := r.db.Read(&localUserList); err != nil {
		return []domain.Usuario{}, errors.New(ERROR_READING)
	}

	return localUserList, nil
}

func (r *repository) Store(id, age int, names, lastname, email, dateCreated string, estatura float64) (domain.Usuario, error) {
	var localUserList []domain.Usuario

	if err := r.db.Read(&localUserList); err != nil {
		return domain.Usuario{}, errors.New(ERROR_READING)
	}

	nwUsuario := domain.Usuario{
		Id:          id,
		Names:       names,
		LastName:    lastname,
		Age:         age,
		DateCreated: dateCreated,
		Estatura:    estatura,
		Email:       email,
		IsActivo:    true,
	}
	localUserList = append(localUserList, nwUsuario)

	if err := r.db.Write(localUserList); err != nil {
		return domain.Usuario{}, errors.New(ERR_WRITING)
	}

	return nwUsuario, nil
}

func (r *repository) Update(id, age int, names, lastname, email, dateCreated string, estatura float64, activo bool) (domain.Usuario, error) {

	var localUserList []domain.Usuario

	if err := r.db.Read(&localUserList); err != nil {
		return domain.Usuario{}, errors.New(ERROR_READING)
	}

	upUsuario := domain.Usuario{
		Id:          id,
		Names:       names,
		LastName:    lastname,
		Age:         age,
		DateCreated: dateCreated,
		Estatura:    estatura,
		Email:       email,
		IsActivo:    activo,
	}

	update := false

	for i := range localUserList {
		if localUserList[i].Id == id {
			update = true
			localUserList[i] = upUsuario
		}
	}

	if !update {
		return domain.Usuario{}, errors.New(STRING_HUBO + ERR_UPDATING_USER)
	}
	if err := r.db.Write(localUserList); err != nil {
		return domain.Usuario{}, errors.New(ERR_WRITING)
	}
	return upUsuario, nil
}

func (r *repository) UpdateLastNameAndAge(id, age int, lastname string) (domain.Usuario, error) {
	var usersList []domain.Usuario

	if err := r.db.Read(&usersList); err != nil {
		return domain.Usuario{}, errors.New(ERROR_READING)
	}

	upUsuario := domain.Usuario{}
	update := false

	for i := range usersList {
		if usersList[i].Id == id {
			update = true

			usersList[i].Age = age
			usersList[i].LastName = lastname
			upUsuario = usersList[i]
		}
	}

	if !update {
		return domain.Usuario{}, errors.New("No se encontró el usuario a actualizar.")
	}
	return upUsuario, nil
}

func (r *repository) LastID() (int, error) {
	var localUserList []domain.Usuario
	if err := r.db.Read(&localUserList); err != nil {
		return 0, errors.New(ERROR_READING)
	}
	if len(localUserList) == 0 {
		return 0, nil
	}

	return localUserList[len(localUserList)-1].Id, nil
}

func (r *repository) Delete(id int) error {
	var usersList []domain.Usuario

	if err := r.db.Read(&usersList); err != nil {
		return errors.New(ERROR_READING)
	}

	deleted := false
	var indexAux int

	for i := range usersList {
		if usersList[i].Id == id {
			deleted = true
			indexAux = i
		}
	}

	if !deleted {
		return errors.New("No se encontró el usuario a eliminar.")
	}

	usersList = append(usersList[:indexAux], usersList[indexAux+1:]...)

	if err := r.db.Write(usersList); err != nil {
		return errors.New(ERR_WRITING)
	}

	return nil
}
func NewRepository(db store.Store) Repository {
	return &repository{db: db}
}
*/