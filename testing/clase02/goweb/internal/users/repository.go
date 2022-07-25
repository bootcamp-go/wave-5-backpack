package users

import (
	"fmt"
	"goweb/internal/domain"
	"goweb/pkg/store"
)

// variables globales----------------------------------------------------------------------------------------------------------------------
//var users []domain.User
//var user domain.User
var lastID int
//-----------------------------------------------------------------------------------------------------------------------------------------

const (
	UserNotFound = "user %d not found"
	FailReading     = "cant read database"
	FailWriting     = "cant write database, error: %w"
)



type Repository interface {
	GetAllUsers() ([]domain.User, error)
	GetUserById(id int) (domain.User, error)
	StoreUser(id int, name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error)
	LastID() (int,error)
	UpdateTotal(id int, name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error)
	UpdatePartial(id int, lastname string, age int) (domain.User, error)
	Delete(id int) error
}

// agregamos un store al repo
type repository struct{
	db store.Store
}

// la db tiene que ir como parametro al inicializar el repository
func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

// VERSION ANTERIOR: ALMACENADO EN MEMORIA
/* func (r *repository) GetAllUsers() ([]domain.User, error) {
	return users, nil
} */

// ALMACENADO EN JSON
func (r *repository) GetAllUsers() ([]domain.User, error) {
	var users []domain.User
	if err := r.db.Read(&users); err != nil {
		return nil, fmt.Errorf(FailReading)
	}
	return users, nil
}

// VERSION ANTERIOR: ALMACENADO EN MEMORIA
/* func (r *repository) GetUserById(id int) (domain.User, error) {
	var userFounded domain.User
	find := false
	for _,u :=range users{
		if u.Id == id{
			userFounded = u
			find = true
			break
		}
	}
	if !find{
		return domain.User{}, fmt.Errorf("No existe el usuario con id %d", id)
	}
	return userFounded, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}
 */

// ALMACENADO EN JSON
func (r *repository) GetUserById(id int) (domain.User, error) {
	// 1ero leo los datos
	var users []domain.User
	if err := r.db.Read(&users); err != nil {
		return domain.User{}, fmt.Errorf(FailReading)
	}

	// 2do busco el Id
	var userFounded domain.User
	find := false
	for _,u :=range users{
		if u.Id == id{
			userFounded = u
			find = true
			break
		}
	}
	if !find{
		return domain.User{}, fmt.Errorf("No existe el usuario con id %d", id)
	}
	return userFounded, nil
}


// VERSION ANTERIOR. ALMACENADO EN MEMORIA
/* func (r *repository) LastID() (int, error) {
	return lastID, nil
} */

// ALMACENADO EN JSON
func (r *repository) LastID() (int, error) {
	var users []domain.User
	if err := r.db.Read(&users); err != nil {
		return 0, fmt.Errorf(FailReading)
	}
	if len(users) == 0 {
		return 0, nil
	}

	return users[len(users)-1].Id, nil
}

// VERSION ANTERIOR. ALMACENADO EN MEMORIA
/* func (r *repository) StoreUser(id int, name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error){
	user = domain.User{Id: id, Name:name, LastName: lastname, Email: email, Age: age, Height: height, Active: active, CreatedAt: createdat}

	users = append(users, user)
	lastID= user.Id

	return user, nil
}
 */

// ALMACENADO EN JSON
func (r *repository) StoreUser(id int, name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error) {
	var users []domain.User

	// si falla la lectura, devolvemos un producto vacío y un msj que definimos previamente en una variable. Para la lectura le paso users como referencia.
	if err := r.db.Read(&users); err != nil {
		return domain.User{}, fmt.Errorf(FailReading)
	}

	// si no falla la lectura, armo el user que voy a agregar a la lista
	user := domain.User{Id: id, Name:name, LastName: lastname, Email: email, Age: age, Height: height, Active: active, CreatedAt: createdat}
	users = append(users, user)

	// ahora para escribir, hago el paso por valor. Tmb incluyo el manejo de errores
	if err := r.db.Write(users); err != nil {
		return domain.User{}, fmt.Errorf(FailWriting, err)
	}

	return user, nil
}

//ALMACENADO EN MEMORIA
/* func(r *repository) UpdateTotal(id int, name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error) {
	userToUpdate := domain.User{Name: name, LastName: lastname, Email: email, Age: age, Height: height, Active: active, CreatedAt: createdat}
	updated := false
	// busco el id del usuario que quiero modificar y le cambio los valores
	for i:= range users {
		if users[i].Id == id{
			userToUpdate.Id = id
			users[i] = userToUpdate
			updated = true
			break
		}
	}
	if !updated{
		return domain.User{}, fmt.Errorf(UserNotFound, id)
	}

	return userToUpdate, nil
}
 */

// ALMACENADO JSON
func(r *repository) UpdateTotal(id int, name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error) {
	// como siempre, 1ero intento leer
	var users []domain.User
	if err := r.db.Read(&users); err != nil {
		return domain.User{}, fmt.Errorf(FailReading)
	}
	
	userToUpdate := domain.User{Name: name, LastName: lastname, Email: email, Age: age, Height: height, Active: active, CreatedAt: createdat}
	updated := false
	
	// busco el id del usuario que quiero modificar y le cambio los valores
	for i:= range users {
		if users[i].Id == id{
			userToUpdate.Id = id
			users[i] = userToUpdate
			updated = true
			break
		}
	}
	if !updated{
		return domain.User{}, fmt.Errorf(UserNotFound, id)
	}

	// ahora sobrescribo todo el archivo
	if err := r.db.Write(users); err != nil {
		return domain.User{}, fmt.Errorf(FailWriting, err)
	}

	return userToUpdate, nil
}

//ALMACENADO EN MEMORIA
/* func(r *repository) UpdatePartial(id int, lastname string, age int) (domain.User, error) {
	// busco por Id
	updated := false
	var userUpdated domain.User 
	for i:= range users {
		if users[i].Id == id{
			users[i].LastName = lastname
			users[i].Age = age
			updated = true
			userUpdated = users[i]
			break
		}
	}
	if !updated{
		return domain.User{}, fmt.Errorf(UserNotFound, id)
	}

	return userUpdated, nil
} */

// ALMACENADO JSON
func(r *repository) UpdatePartial(id int, lastname string, age int) (domain.User, error) {
	// como siempre, 1ero intento leer
	var users []domain.User
	if err := r.db.Read(&users); err != nil {
		return domain.User{}, fmt.Errorf(FailReading)
	}
	// busco por Id
	updated := false
	var userUpdated domain.User 
	for i:= range users {
		if users[i].Id == id{
			users[i].LastName = lastname
			users[i].Age = age
			updated = true
			userUpdated = users[i]
			break
		}
	}
	if !updated{
		return domain.User{}, fmt.Errorf(UserNotFound, id)
	}

	// ahora sobrescribo todo el archivo
	if err := r.db.Write(users); err != nil {
		return domain.User{}, fmt.Errorf(FailWriting, err)
	}

	return userUpdated, nil
}

// ALMACENADO EN MEMORIA
/* func (r *repository) Delete(id int) error {
	
	var indexToDelete int
	find := false
	for i :=range users{
		if users[i].Id == id{
			indexToDelete = i
			find = true
			break
		}
	}
	if !find{
		return fmt.Errorf("No existe el usuario con id %d", id)
	}
	users = append(users[:indexToDelete],users[indexToDelete+1:]... )
	return nil
} */

//ALMACENADO JSON
func (r *repository) Delete(id int) error {
    // como siempre, 1ero intento leer
    var users []domain.User
    if err := r.db.Read(&users); err != nil {
        return fmt.Errorf(FailReading)
    }

    var indexToDelete int
    find := false
    for i :=range users{
        if users[i].Id == id{
            indexToDelete = i
            find = true
            break
        }
    }
    if !find{
        return fmt.Errorf(UserNotFound, id)
    }
    users = append(users[:indexToDelete],users[indexToDelete+1:]... )

    if err := r.db.Write(users); err != nil {
        return fmt.Errorf(FailWriting, err)
    }
    
    return nil
}