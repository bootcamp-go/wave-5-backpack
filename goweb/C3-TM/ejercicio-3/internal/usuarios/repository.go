package usuarios

import (
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

type Repository interface {
	GetAll() ([]*Usuario, error)
	Registrar(nombre, apellido, email string, edad, altura int, activo bool, fecha_creacion string) (*Usuario, error)
	Modificar(id int, nombre, apellido, email string, edad, altura int, activo bool, fecha_creacion string) (*Usuario, error)
	Eliminar(id int) error
	ModificarAE(id int, apellido string, edad int) (*Usuario, error)
}

var usuarios []*Usuario
var lastID int

func NewRepository() Repository {
	usuarios = append(usuarios, &Usuario{
		Nombre:        "Matias",
		Apellido:      "Vince",
		Email:         "matiasvince9@gmail.com",
		Edad:          22,
		Altura:        175,
		Activo:        true,
		FechaCreacion: "09121999",
	})
	return &Usuario{}
}

func (u *Usuario) GetAll() ([]*Usuario, error) {
	return usuarios, nil
}

func (u *Usuario) Registrar(nombre, apellido, email string, edad, altura int, activo bool, fecha_creacion string) (*Usuario, error) {
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

	usuarios = append(usuarios, &usuario)

	return &usuario, nil
}

func (u *Usuario) Modificar(id int, nombre, apellido, email string, edad, altura int, activo bool, fecha_creacion string) (*Usuario, error) {
	user := Usuario{Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura, Activo: activo, FechaCreacion: fecha_creacion}
	updated := false

	for i := range usuarios {
		if usuarios[i].ID == id {
			user.ID = id
			usuarios[i] = &user
			updated = true
		}
	}

	if !updated {
		return &Usuario{}, fmt.Errorf("Usuario %d no encontrado", id)
	}

	return &user, nil
}

func (u *Usuario) Eliminar(id int) error {
	eliminado := false
	var index int
	for i, val := range usuarios {
		if val.ID == id {
			index = i
			eliminado = true
		}
	}
	if !eliminado {
		return fmt.Errorf("Producto %d no encontrado", id)
	}

	usuarios = append(usuarios[:index], usuarios[index+1:]...)

	return nil
}

func (u *Usuario) ModificarAE(id int, apellido string, edad int) (*Usuario, error) {
	user := Usuario{}
	updated := false

	for i, val := range usuarios {
		if val.ID == id {
			user = Usuario{Nombre: val.Nombre, Apellido: apellido, Email: val.Email, Edad: edad, Altura: val.Altura, Activo: val.Activo, FechaCreacion: val.FechaCreacion}
			usuarios[i] = &user
			updated = true
		}
	}

	if !updated {
		return &Usuario{}, fmt.Errorf("Usuario %d no encontrado", id)
	}

	return &user, nil
}
