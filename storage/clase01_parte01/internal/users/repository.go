package users

import (
	"database/sql"
	"fmt"
	//"fmt"
	"goweb/internal/domain"
	"log"
)

//var lastID int
const (
	UserNotFound = "user %d not found"
	FailReading     = "cant read database"
	FailWriting     = "cant write database"
)


type Repository interface {
	// métodos viejos
	GetAllUsers() ([]domain.User, error)
	GetUserById(id int) (domain.User, error)
	UpdateTotal(id int, name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error)
	UpdatePartial(id int, lastname string, age int) (domain.User, error)
	Delete(id int) error
	
	// métodos nuevos o modificados
	GetUserByName(name string) (domain.User, error)
	StoreUser(name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error)
}

//modifico el tipo de db porque ya no es más un json
type repository struct{
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

const (
    GetUserByName    = "SELECT * FROM users WHERE name = ?"
    GetUserById    = "SELECT * FROM users WHERE id = ?"
)

func (r *repository) GetUserByName(name string) (domain.User, error) {
	
	var user domain.User
    rows, err := r.db.Query(GetUserByName, name)
    if err != nil {
        log.Println(err)
        return domain.User{},err
    }
    for rows.Next() {
        err := rows.Scan(&user.Id, &user.Name, &user.LastName, &user.Email, &user.Age, &user.Height, &user.Active, &user.CreatedAt)
        if err != nil {
            log.Fatal(err)
            return domain.User{},err
        }
	}

    return user, nil
}


func (r *repository) GetUserById(id int) (domain.User, error) {
	var user domain.User
    rows, err := r.db.Query(GetUserById, id)
    if err != nil {
        log.Println(err)
        return domain.User{},err
    }
    for rows.Next() {
        err := rows.Scan(&user.Id, &user.Name, &user.LastName, &user.Email, &user.Age, &user.Height, &user.Active, &user.CreatedAt)
        if err != nil {
            log.Fatal(err)
            return domain.User{},err
        }
	}

    return user, nil
}

func (r *repository) StoreUser(name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error) {
	
	var user = domain.User{
		Name: name,
		LastName: lastname,	
		Email: email, 
		Age: age,
		Height: height,
		Active: active,
		CreatedAt: createdat,
	}
	
    stmt, err := r.db.Prepare("INSERT INTO users(name, lastname, email, age, height, active, createdat) VALUES( ?, ?, ?, ?, ?, ?, ? )") // se prepara el SQL
    if err != nil {
		log.Fatal(err)
    }
    defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
    var result sql.Result
    result, err = stmt.Exec(name, lastname, email, age, height, active, createdat) // retorna un sql.Result y un error
    if err != nil {
		return domain.User{}, err
    }
    insertedId, _ := result.LastInsertId() // del sql.Resul devuelto en la ejecución obtenemos el Id insertado
    user.Id = int(insertedId)
	
	return user, nil
}

func(r *repository) UpdateTotal(id int, name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error) {
	// esta parte la agrego yo porque si el usuario no existe no me duelve error, lo tengo que forzar.
	userExist, err := r.GetUserById(id)
	if err!=nil{
		return domain.User{}, err
	}else if userExist.Id == 0{
		return domain.User{}, fmt.Errorf("user not found")
	}
	//--------------------------------------------------------------


	var user = domain.User{
		Id:id,
		Name: name,
		LastName: lastname,	
		Email: email, 
		Age: age,
		Height: height,
		Active: active,
		CreatedAt: createdat,
	}

	stmt, err := r.db.Prepare("UPDATE users SET name = ?, lastname = ?, email = ?, age = ?, height = ?, active = ?, createdat = ? WHERE id = ?") // se prepara la sentencia SQL a ejecutar
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close()     // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
    aaa, err := stmt.Exec(user.Name, user.LastName, user.Email, user.Age, user.Height, user.Active, user.CreatedAt, user.Id)
    if err != nil {
        return domain.User{}, err
    }
	fmt.Println(&aaa)
    return user, nil
}








func (r *repository) GetAllUsers() ([]domain.User, error) {
	return nil, nil
}




func(r *repository) UpdatePartial(id int, lastname string, age int) (domain.User, error) {
	return domain.User{}, nil
}


func (r *repository) Delete(id int) error {
    return nil
}