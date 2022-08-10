package users

import (
	"database/sql"
	"fmt"
	"time"

	//"fmt"
	"goweb/internal/domain"
	"log"

	"golang.org/x/net/context"
)

//var lastID int
const (
	UserNotFound = "user %d not found"
	FailReading     = "cant read database"
	FailWriting     = "cant write database"
)


type Repository interface {
	// métodos viejos
	UpdatePartial(id int, lastname string, age int) (domain.User, error)
	Delete(id int) error
	
	// métodos nuevos o modificados
	GetAllUsers(ctx context.Context) ([]domain.User, error)
	GetUserById(ctx context.Context, id int) (domain.User, error)
	GetUserByName(ctx context.Context, name string) (domain.User, error)
	StoreUser(ctx context.Context, name, lastname, email string, age int, height float32, active bool) (domain.User, error)
	UpdateTotal(ctx context.Context, id int, name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error)
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
	GetAllUsers = "SELECT id, name, lastname, email, age, height, active, createdat FROM users"
    GetUserByName = "SELECT id, name, lastname, email, age, height, active, createdat FROM users WHERE name = ?"
    GetUserById = "SELECT id, name, lastname, email, age, height, active, createdat FROM users WHERE id = ?"
	StoreUser = "INSERT INTO users(name, lastname, email, age, height, active, createdat) VALUES( ?, ?, ?, ?, ?, ?, ? )"
	UpdateUser = "UPDATE users SET name = ?, lastname = ?, email = ?, age = ?, height = ?, active = ?, createdat = ? WHERE id = ?"
)

func (r *repository) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
    rows, err := r.db.QueryContext(ctx, GetAllUsers)
    if err != nil {
        log.Println(err)
        return nil,err
    }
    for rows.Next() {
		var user domain.User
        err := rows.Scan(&user.Id, &user.Name, &user.LastName, &user.Email, &user.Age, &user.Height, &user.Active, &user.CreatedAt)
        if err != nil {
            log.Fatal(err)
            return nil,err
        }
		users = append(users, user)
	}

    return users, nil
}


func (r *repository) GetUserByName(ctx context.Context, name string) (domain.User, error) {
	
	var user domain.User
    rows, err := r.db.QueryContext(ctx,GetUserByName, name)
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


func (r *repository) GetUserById(ctx context.Context, id int) (domain.User, error) {
	var user domain.User
    rows, err := r.db.QueryContext(ctx,GetUserById, id)
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

func (r *repository) StoreUser(ctx context.Context, name, lastname, email string, age int, height float32, active bool) (domain.User, error) {
	
	var user = domain.User{
		Name: name,
		LastName: lastname,	
		Email: email, 
		Age: age,
		Height: height,
		Active: active,
		CreatedAt: time.Now().Format("2006-01-01 00:00:00"),
	}
	
    stmt, err := r.db.Prepare(StoreUser) // se prepara el SQL
    if err != nil {
		log.Fatal(err)
    }
    defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
    var result sql.Result
    result, err = stmt.ExecContext(ctx, user.Name, user.LastName, user.Email, user.Age, user.Height, user.Active, user.CreatedAt) // retorna un sql.Result y un error
    if err != nil {
		return domain.User{}, err
    }
    insertedId, err := result.LastInsertId() // del sql.Resul devuelto en la ejecución obtenemos el Id insertado
    if err != nil {
		return domain.User{}, err
	}
	user.Id = int(insertedId)
	
	return user, nil
}

func(r *repository) UpdateTotal(ctx context.Context, id int, name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error) {
	// esta parte la agrego yo porque si el usuario no existe no me duelve error, lo tengo que forzar.
	userExist, err := r.GetUserById(ctx, id)
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

	stmt, err := r.db.Prepare(UpdateUser) // se prepara la sentencia SQL a ejecutar
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close()     // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
    aaa, err := stmt.ExecContext(ctx, user.Name, user.LastName, user.Email, user.Age, user.Height, user.Active, user.CreatedAt, user.Id)
    if err != nil {
        return domain.User{}, err
    }
	fmt.Println(&aaa)
    return user, nil
}



func(r *repository) UpdatePartial(id int, lastname string, age int) (domain.User, error) {
	return domain.User{}, nil
}


func (r *repository) Delete(id int) error {
    return nil
}