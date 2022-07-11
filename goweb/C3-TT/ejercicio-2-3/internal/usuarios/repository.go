package usuarios

import (
	"ejercicio-2-3/pkg/registro"
	"fmt"
)

type Usuario struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre" binding:"required"`
	Apellido      string `json:"apellido" binding:"required"`
	Email         string `json:"email" binding:"required"`
	Edad          int    `json:"edad" binding:"required"`
	Altura        int    `json:"altura" binding:"required"`
	Activo        bool   `json:"activo" binding:"required"`
	FechaCreacion string `json:"fecha_creacion" binding:"required"`
}

type repository struct {
	db registro.Registro
}

type Repository interface {
	GetAll() ([]*Usuario, error)
	Registrar(nombre, apellido, email string, edad, altura int, activo bool, fecha_creacion string) (*Usuario, error)
	Modificar(id int, nombre, apellido, email string, edad, altura int, activo bool, fecha_creacion string) (*Usuario, error)
	Eliminar(id int) error
	ModificarAE(id int, apellido string, edad int) (*Usuario, error)
}

func NewRepository(db registro.Registro) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll() ([]*Usuario, error) {
	var users []*Usuario
	if err := r.db.Read(&users); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repository) Registrar(nombre, apellido, email string, edad, altura int, activo bool, fecha_creacion string) (*Usuario, error) {
	var users []Usuario
	r.db.Read(&users)
	lastID := 0

	for _, val := range users {
		if val.ID > lastID {
			lastID = val.ID
		}
	}

	lastID++
	usuario := Usuario{
		ID:            lastID,
		Nombre:        nombre,
		Apellido:      apellido,
		Email:         email,
		Edad:          edad,
		Altura:        altura,
		Activo:        activo,
		FechaCreacion: fecha_creacion,
	}

	users = append(users, usuario)
	if err := r.db.Write(&users); err != nil {
		return nil, err
	}

	return &usuario, nil
}

func (r *repository) Modificar(id int, nombre, apellido, email string, edad, altura int, activo bool, fecha_creacion string) (*Usuario, error) {
	user := Usuario{Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura, Activo: activo, FechaCreacion: fecha_creacion}
	updated := false

	var users []*Usuario
	r.db.Read(&users)

	for i := range users {
		if users[i].ID == id {
			user.ID = id
			users[i] = &user
			updated = true
		}
	}

	if !updated {
		return &Usuario{}, fmt.Errorf("Usuario %d no encontrado", id)
	}

	if err := r.db.Write(&users); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) Eliminar(id int) error {
	var users []*Usuario
	if err := r.db.Read(&users); err != nil {
		return err
	}

	eliminado := false
	var index int
	for i, val := range users {
		if val.ID == id {
			index = i
			eliminado = true
		}
	}
	if !eliminado {
		return fmt.Errorf("Producto %d no encontrado", id)
	}

	users = append(users[:index], users[index+1:]...)

	if err := r.db.Write(&users); err != nil {
		return err
	}

	return nil
}

func (r *repository) ModificarAE(id int, apellido string, edad int) (*Usuario, error) {
	var users []*Usuario
	if err := r.db.Read(&users); err != nil {
		return nil, err
	}

	user := Usuario{}
	updated := false

	for i, val := range users {
		if val.ID == id {
			user = Usuario{Nombre: val.Nombre, Apellido: apellido, Email: val.Email, Edad: edad, Altura: val.Altura, Activo: val.Activo, FechaCreacion: val.FechaCreacion}
			users[i] = &user
			updated = true
		}
	}

	if !updated {
		return &Usuario{}, fmt.Errorf("Usuario %d no encontrado", id)
	}

	if err := r.db.Write(&users); err != nil {
		return nil, err
	}

	return &user, nil
}
