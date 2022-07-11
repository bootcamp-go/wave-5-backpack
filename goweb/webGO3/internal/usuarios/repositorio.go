package usuarios

import (
	"fmt"

	"github.com/del_rio/web-server/internal/domain"
)

var lastId int
var listUsuarios []domain.Usuario

type Repository interface {
	GetAll() ([]domain.Usuario, error)
	Save(Nombre, Apellido, Email, Fecha_creacion string, Id, Edad, Altura int, Activo bool) (domain.Usuario, error)
	LastId() (int, error)
	UpdateUsuario(Nombre, Apellido, Email, Fecha_creacion string, Id, Edad, Altura int, Activo bool) (domain.Usuario, error)
	UpdateAtributos(Nombre, Apellido, Email, Fecha_creacion string, Id, Edad, Altura int, Activo *bool) (domain.Usuario, error)
	DeleteUsuario(Id int) error
}

type repository struct{}

func (r *repository) GetAll() ([]domain.Usuario, error) {
	if len(listUsuarios) == 0 {
		return listUsuarios, fmt.Errorf("lista vacia perdone usted")
	}
	return listUsuarios, nil
}

func (r *repository) Save(Nombre, Apellido, Email, Fecha_creacion string, Id, Edad, Altura int, Activo bool) (domain.Usuario, error) {
	if Nombre == "" || Apellido == "" || Email == "" || Fecha_creacion == "" || Id <= 0 || Edad < 0 || Altura <= 0 {
		usuarioNulo := domain.Usuario{}
		return usuarioNulo, fmt.Errorf("campo invalido o campos invalidos")
	}
	usuario := domain.Usuario{
		Id:             Id,
		Nombre:         Nombre,
		Apellido:       Apellido,
		Email:          Email,
		Edad:           Edad,
		Altura:         Altura,
		Activo:         Activo,
		Fecha_creacion: Fecha_creacion,
	}
	listUsuarios = append(listUsuarios, usuario)
	lastId = usuario.Id
	return usuario, nil
}
func (r *repository) LastId() (int, error) {
	//aqui podria pasar algo
	return lastId, nil
}
func (r *repository) UpdateUsuario(Nombre, Apellido, Email, Fecha_creacion string, Id, Edad, Altura int, Activo bool) (domain.Usuario, error) {
	usuario := domain.Usuario{
		Nombre:         Nombre,
		Apellido:       Apellido,
		Email:          Email,
		Edad:           Edad,
		Altura:         Altura,
		Activo:         Activo,
		Fecha_creacion: Fecha_creacion,
	}
	for i := 0; i < len(listUsuarios); i++ {
		if Id == listUsuarios[i].Id {
			usuario.Id = Id
			listUsuarios[i] = usuario
			return usuario, nil
		}
	}
	usuarioNulo := domain.Usuario{}
	return usuarioNulo, fmt.Errorf("ese Id pareciera no existir dentro de nuestra BD")
}
func (r *repository) UpdateAtributos(Nombre, Apellido, Email, Fecha_creacion string, Id, Edad, Altura int, Activo *bool) (domain.Usuario, error) {

	for i := 0; i < len(listUsuarios); i++ {
		if Id == listUsuarios[i].Id {
			r.changeNombre(i, Nombre)
			r.changeApellido(i, Apellido)
			r.changeEmail(i, Email)
			r.changeFecha_creacion(i, Fecha_creacion)
			r.changeEdad(i, Edad)
			r.changeAltura(i, Altura)
			r.changeActivo(i, Activo)
			return listUsuarios[i], nil
		}
	}
	usuarioNulo := domain.Usuario{}
	return usuarioNulo, fmt.Errorf("ese Id pareciera no existir dentro de nuestra BD")
}
func (r *repository) DeleteUsuario(Id int) error {
	for i := 0; i < len(listUsuarios); i++ {
		if listUsuarios[i].Id == Id {
			listUsuarios = append(listUsuarios[:i], listUsuarios[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("no se encontro esa direccion id")
}

func NewRepository() Repository {
	return &repository{}
}

//Nombre, Apellido, Email, Fecha_creacion string
func (r *repository) changeNombre(position int, newNombre string) {
	if newNombre != "" {
		listUsuarios[position].Nombre = newNombre
	}
}

func (r *repository) changeApellido(position int, newApellido string) {
	if newApellido != "" {
		listUsuarios[position].Apellido = newApellido
	}
}
func (r *repository) changeEmail(position int, newEmail string) {
	if newEmail != "" {
		listUsuarios[position].Email = newEmail
	}
}

func (r *repository) changeFecha_creacion(position int, newFecha_creacion string) {
	if newFecha_creacion != "" {
		listUsuarios[position].Fecha_creacion = newFecha_creacion
	}
}

// Edad, Altura int
func (r *repository) changeEdad(position int, newEdad int) {
	if newEdad != 0 {
		listUsuarios[position].Edad = newEdad
	}
}

func (r *repository) changeAltura(position int, newAltura int) {
	if newAltura != 0 {
		listUsuarios[position].Altura = newAltura
	}
}

//Activo *bool
func (r *repository) changeActivo(position int, newActivo *bool) {
	if newActivo != nil {
		listUsuarios[position].Activo = *newActivo
	}
}
