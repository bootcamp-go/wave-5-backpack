package usuarios

import (
	"fmt"

	"github.com/del_rio/web-server/internal/domain"
	"github.com/del_rio/web-server/pkg/store"
)

type Repository interface {
	GetAll() ([]domain.Usuario, error)
	Save(Nombre, Apellido, Email, Fecha_creacion string, Id, Edad, Altura int, Activo bool) (domain.Usuario, error)
	LastId() (int, error)
	// UpdateUsuario(Nombre, Apellido, Email, Fecha_creacion string, Id, Edad, Altura int, Activo bool) (domain.Usuario, error)
	// UpdateAtributos(Nombre, Apellido, Email, Fecha_creacion string, Id, Edad, Altura int, Activo *bool) (domain.Usuario, error)
	// DeleteUsuario(Id int) error
}

type repository struct {
	db store.Store
}

func (r *repository) GetAll() ([]domain.Usuario, error) {
	var ListUsuarios []domain.Usuario
	if err := r.db.Read(&ListUsuarios); err != nil {
		// return ListUsuarios, err
	}
	if len(ListUsuarios) == 0 {
		return ListUsuarios, fmt.Errorf("lista vacia perdone usted")
	}
	return ListUsuarios, nil
}

func (r *repository) Save(Nombre, Apellido, Email, Fecha_creacion string, Id, Edad, Altura int, Activo bool) (domain.Usuario, error) {
	var ListUsuarios []domain.Usuario
	if err := r.db.Read(&ListUsuarios); err != nil {
		// usuarioNulo := domain.Usuario{}
		// fmt.Println(err)
		// return usuarioNulo, err
	}
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
	ListUsuarios = append(ListUsuarios, usuario)
	if err := r.db.Write(ListUsuarios); err != nil {
		usuarioNulo := domain.Usuario{}
		return usuarioNulo, err
	}
	return usuario, nil
}
func (r *repository) LastId() (int, error) {
	//aqui podria pasar algo
	var ListUsuarios []domain.Usuario
	if err := r.db.Read(&ListUsuarios); err != nil {
		return 0, err
	}
	if len(ListUsuarios) == 0 {
		return 0, nil
	}
	lastId := ListUsuarios[len(ListUsuarios)-1].Id
	return lastId, nil
}

func NewRepository(db store.Store) Repository {
	return &repository{db}
}

// func (r *repository) UpdateUsuario(Nombre, Apellido, Email, Fecha_creacion string, Id, Edad, Altura int, Activo bool) (domain.Usuario, error) {
// 	usuario := domain.Usuario{
// 		Nombre:         Nombre,
// 		Apellido:       Apellido,
// 		Email:          Email,
// 		Edad:           Edad,
// 		Altura:         Altura,
// 		Activo:         Activo,
// 		Fecha_creacion: Fecha_creacion,
// 	}
// 	for i := 0; i < len(listUsuarios); i++ {
// 		if Id == listUsuarios[i].Id {
// 			usuario.Id = Id
// 			listUsuarios[i] = usuario
// 			return usuario, nil
// 		}
// 	}
// 	usuarioNulo := domain.Usuario{}
// 	return usuarioNulo, fmt.Errorf("ese Id pareciera no existir dentro de nuestra BD")
// }
// func (r *repository) UpdateAtributos(Nombre, Apellido, Email, Fecha_creacion string, Id, Edad, Altura int, Activo *bool) (domain.Usuario, error) {

// 	for i := 0; i < len(listUsuarios); i++ {
// 		if Id == listUsuarios[i].Id {
// 			r.changeNombre(i, Nombre)
// 			r.changeApellido(i, Apellido)
// 			r.changeEmail(i, Email)
// 			r.changeFecha_creacion(i, Fecha_creacion)
// 			r.changeEdad(i, Edad)
// 			r.changeAltura(i, Altura)
// 			r.changeActivo(i, Activo)
// 			return listUsuarios[i], nil
// 		}
// 	}
// 	usuarioNulo := domain.Usuario{}
// 	return usuarioNulo, fmt.Errorf("ese Id pareciera no existir dentro de nuestra BD")
// }
// func (r *repository) DeleteUsuario(Id int) error {
// 	for i := 0; i < len(listUsuarios); i++ {
// 		if listUsuarios[i].Id == Id {
// 			listUsuarios = append(listUsuarios[:i], listUsuarios[i+1:]...)
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf("no se encontro esa direccion id")
// }

//Nombre, Apellido, Email, Fecha_creacion string
// func (r *repository) changeNombre(position int, newNombre string) {
// 	if newNombre != "" {
// 		listUsuarios[position].Nombre = newNombre
// 	}
// }

// func (r *repository) changeApellido(position int, newApellido string) {
// 	if newApellido != "" {
// 		listUsuarios[position].Apellido = newApellido
// 	}
// }
// func (r *repository) changeEmail(position int, newEmail string) {
// 	if newEmail != "" {
// 		listUsuarios[position].Email = newEmail
// 	}
// }

// func (r *repository) changeFecha_creacion(position int, newFecha_creacion string) {
// 	if newFecha_creacion != "" {
// 		listUsuarios[position].Fecha_creacion = newFecha_creacion
// 	}
// }

// // Edad, Altura int
// func (r *repository) changeEdad(position int, newEdad int) {
// 	if newEdad != 0 {
// 		listUsuarios[position].Edad = newEdad
// 	}
// }

// func (r *repository) changeAltura(position int, newAltura int) {
// 	if newAltura != 0 {
// 		listUsuarios[position].Altura = newAltura
// 	}
// }

// //Activo *bool
// func (r *repository) changeActivo(position int, newActivo *bool) {
// 	if newActivo != nil {
// 		listUsuarios[position].Activo = *newActivo
// 	}
// }
