package users

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
	"github.com/bootcamp-go/wave-5-backpack/goweb/pkg/store"
)

//var usersLocal []domain.ModelUser
//var lastId int
const (
	errorLectura   = "no se puede leer la db, error: %s"
	errorEscritura = "no se puede escribir en la db, error: %s"
)

type Repository interface {
	GetAll() ([]domain.ModelUser, error)
	GetById(id int) (domain.ModelUser, error)
	Store(nombre string, apellido string, email string, edad int, altura float64) (domain.ModelUser, error)
	Update(id int, nombre string, apellido string, email string, edad int, altura float64) (domain.ModelUser, error)
	UpdateApellidoEdad(id int, nombre string, edad int) (domain.ModelUser, error)
	Delete(id int) error
	SearchUser(nombreQuery string, apellidoQuery string, emailQuery string, edadQuery string, alturaQuery string, activoQuery string, fechaCreacionQuery string) ([]domain.ModelUser, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

// Función para obtener todas las entidades
func (r *repository) GetAll() ([]domain.ModelUser, error) {
	// Leemos los usuarios del JSON
	var users []domain.ModelUser
	if err := r.db.Read(&users); err != nil {
		return nil, fmt.Errorf(errorLectura, err)
	}

	// Obtenemos los usuarios que no están borrados
	var filtro []domain.ModelUser
	for i := range users {
		if !users[i].Borrado {
			filtro = append(filtro, users[i])
		}
	}

	// Devolvemos todos los usuarios del JSON
	return filtro, nil
}

// Función para devolver una entidad por id
func (r *repository) GetById(id int) (domain.ModelUser, error) {
	var user domain.ModelUser

	// Leemos los usuarios del JSON
	var users []domain.ModelUser
	if err := r.db.Read(&users); err != nil {
		return domain.ModelUser{}, fmt.Errorf(errorLectura, err)
	}

	// Buscamos el usuario por id
	found := false
	for i := range users {
		if users[i].Id == id && !users[i].Borrado && !found {
			user = users[i]
			found = true
		}
	}

	if !found {
		return domain.ModelUser{}, fmt.Errorf("usuario %d no encontrado", id)
	}

	// Devolvemos el usuario por id
	return user, nil
}

// Función para guardar una entidad
func (r *repository) Store(nombre string, apellido string, email string, edad int, altura float64) (domain.ModelUser, error) {
	user := domain.ModelUser{Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura}
	lastId := 0

	// Leemos los usuarios del JSON
	var users []domain.ModelUser
	if err := r.db.Read(&users); err != nil {
		return domain.ModelUser{}, fmt.Errorf(errorLectura, err)
	}

	// Se Calcula el id siguiente
	for i := range users {
		if users[i].Id > lastId {
			lastId = users[i].Id
		}
	}

	// Se genera el id consecutivo y la fecha de creación
	user.Id = lastId + 1
	user.FechaCreacion = time.Now()
	user.Borrado = false
	user.Activo = true

	// Se adiciona el usuario al slice de usuarios
	users = append(users, user)

	// Se guarda la información y se verifica que no haya ocurrido un error
	if err := r.db.Write(&users); err != nil {
		return domain.ModelUser{}, fmt.Errorf(errorEscritura, err)
	}

	// Devolvemos el nuevo usuario
	return user, nil
}

// Función para actualizar una entidad completa
func (r *repository) Update(id int, nombre string, apellido string, email string, edad int, altura float64) (domain.ModelUser, error) {
	// Leemos los usuarios del JSON
	var users []domain.ModelUser
	if err := r.db.Read(&users); err != nil {
		return domain.ModelUser{}, fmt.Errorf(errorLectura, err)
	}

	user := domain.ModelUser{Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura}
	found := false
	for i := range users {
		if users[i].Id == id && !users[i].Borrado && !found {
			user.Id = id
			user.FechaCreacion = users[i].FechaCreacion
			user.Activo = users[i].Activo
			users[i] = user
			found = true
		}
	}

	// Verificamos que haya existido el usuario a actualizar
	if !found {
		return domain.ModelUser{}, fmt.Errorf("usuario %d no encontrado", id)
	}

	// Se guarda la información y se verifica que no haya ocurrido un error
	if err := r.db.Write(&users); err != nil {
		return domain.ModelUser{}, fmt.Errorf(errorEscritura, err)
	}

	// Se devuelve el usuario actualizado
	return user, nil
}

// Función para actualizar 2 campos de una entidad
func (r *repository) UpdateApellidoEdad(id int, apellido string, edad int) (domain.ModelUser, error) {
	// Leemos los usuarios del JSON
	var users []domain.ModelUser
	if err := r.db.Read(&users); err != nil {
		return domain.ModelUser{}, fmt.Errorf(errorLectura, err)
	}

	var user domain.ModelUser
	found := false
	for i := range users {
		if users[i].Id == id && !users[i].Borrado && !found {
			users[i].Apellido = apellido
			users[i].Edad = edad
			user = users[i]
			found = true
		}
	}

	// Verificamos que haya existido el usuario a actualizar su Apellido y Edad
	if !found {
		return domain.ModelUser{}, fmt.Errorf("usuario %d no encontrado", id)
	}

	// Se guarda la información y se verifica que no haya ocurrido un error
	if err := r.db.Write(&users); err != nil {
		return domain.ModelUser{}, fmt.Errorf(errorEscritura, err)
	}

	// Se devuelve el usuario actualizado del Apellido y Edad
	return user, nil
}

// Función para borrar una entidad
func (r *repository) Delete(id int) error {
	// Leemos los usuarios del JSON
	var users []domain.ModelUser
	if err := r.db.Read(&users); err != nil {
		return fmt.Errorf(errorLectura, err)
	}

	found := false
	for i := range users {
		if users[i].Id == id && !users[i].Borrado && !found {
			users[i].Borrado = true
			found = true
		}
	}

	// Verificamos que haya existido el usuario a borrar
	if !found {
		return fmt.Errorf("usuario %d no econtrado", id)
	}

	// Se guarda la información y se verifica que no haya ocurrido un error
	if err := r.db.Write(&users); err != nil {
		return fmt.Errorf(errorEscritura, err)
	}

	return nil
}

func (r *repository) SearchUser(nombreQuery string, apellidoQuery string, emailQuery string, edadQuery string, alturaQuery string, activoQuery string, fechaCreacionQuery string) ([]domain.ModelUser, error) {
	// Leemos los usuarios del JSON
	var users []domain.ModelUser
	if err := r.db.Read(&users); err != nil {
		return []domain.ModelUser{}, fmt.Errorf(errorLectura, err)
	}

	var filtro []domain.ModelUser
	var temporal []domain.ModelUser

	// Buscamos por nombre
	if nombreQuery != "" {
		for _, u := range users {
			if strings.Contains(strings.ToUpper(u.Nombre), strings.ToUpper(nombreQuery)) && !u.Borrado {
				filtro = append(filtro, u)
			}
		}
	}

	// Buscamos por apellido
	if apellidoQuery != "" {
		if len(filtro) > 0 {
			temporal = nil
			for _, u := range filtro {
				if strings.Contains(strings.ToUpper(u.Apellido), strings.ToUpper(apellidoQuery)) && !u.Borrado {
					temporal = append(temporal, u)
				}
			}
			filtro = temporal
		} else {
			for _, u := range users {
				if strings.Contains(strings.ToUpper(u.Apellido), strings.ToUpper(apellidoQuery)) && !u.Borrado {
					filtro = append(filtro, u)
				}
			}
		}
	}

	// Buscamos por Email
	if emailQuery != "" {
		if len(filtro) > 0 {
			temporal = nil
			for _, u := range filtro {
				if strings.Contains(strings.ToUpper(u.Email), strings.ToUpper(emailQuery)) && !u.Borrado {
					temporal = append(temporal, u)
				}
			}
			filtro = temporal
		} else {
			for _, u := range users {
				if strings.Contains(strings.ToUpper(u.Email), strings.ToUpper(emailQuery)) && !u.Borrado {
					filtro = append(filtro, u)
				}
			}
		}
	}

	// Buscamos por edad
	if edadQuery != "" {
		edadInt, err := strconv.Atoi(edadQuery)
		if err == nil {
			if len(filtro) > 0 {
				temporal = nil
				for _, u := range filtro {
					if edadInt == u.Edad && !u.Borrado {
						temporal = append(temporal, u)
					}
				}
				filtro = temporal
			} else {
				for _, u := range users {
					if edadInt == u.Edad && !u.Borrado {
						filtro = append(filtro, u)
					}
				}
			}
		}
	}

	// Buscamos por altura
	if alturaQuery != "" {
		alturaFloat64, err := strconv.ParseFloat(alturaQuery, 64)
		if err == nil {
			if len(filtro) > 0 {
				temporal = nil
				for _, u := range filtro {
					if alturaFloat64 == u.Altura && !u.Borrado {
						temporal = append(temporal, u)
					}
				}
				filtro = temporal
			} else {
				for _, u := range users {
					if alturaFloat64 == u.Altura && !u.Borrado {
						filtro = append(filtro, u)
					}
				}
			}
		}
	}

	// Buscamos por activo
	if activoQuery != "" {
		activoBool, err := strconv.ParseBool(activoQuery)
		if err == nil {
			if len(filtro) > 0 {
				temporal = nil
				for _, u := range filtro {
					if activoBool == u.Activo && !u.Borrado {
						temporal = append(temporal, u)
					}
				}
				filtro = temporal
			} else {
				for _, u := range users {
					if activoBool == u.Activo && !u.Borrado {
						filtro = append(filtro, u)
					}
				}
			}
		}
	}

	// Buscamos por fecha
	if fechaCreacionQuery != "" {
		fechaCreacionDate, err := time.Parse("2006-01-02", fechaCreacionQuery)
		if err == nil {
			if len(filtro) > 0 {
				temporal = nil
				for _, u := range filtro {
					if fechaCreacionDate.Format("2006-01-02") == u.FechaCreacion.Format("2006-01-02") && !u.Borrado {
						temporal = append(temporal, u)
					}
				}
				filtro = temporal
			} else {
				for _, u := range users {
					if fechaCreacionDate.Format("2006-01-02") == u.FechaCreacion.Format("2006-01-02") && !u.Borrado {
						filtro = append(filtro, u)
					}
				}
			}
		}
	}

	if len(filtro) > 0 {
		return filtro, nil
	}
	return []domain.ModelUser{}, nil
}
