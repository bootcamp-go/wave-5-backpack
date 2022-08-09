package repositorio

import (
	"context"
	"database/sql"
	"log"

	"C1-TT/internal/domain"
	"C1-TT/pkg/db"
)

const (
	GetUser = "SELECT users.id, users.firstName, users.lastName, users.email, users.age, users.height, users.activo, users.createdAt " +
		"FROM users WHERE id = ?"
	GetAll = "SELECT users.id, users.firstName, users.lastName, users.email, users.age, users.height, users.activo, users.createdAt" +
		"FROM users"
	GetUserByName = "select SELECT users.id, users.firstName, users.lastName, users.email, users.age, users.height, users.activo, users.createdAt " +
		"FROM users WHERE firstName = ?"
	StoreUser  = "INSERT INTO users(firstName, lastName, email, age, height, activo, createdAt) VALUES(?,?,?,?,?,?,?)"
	UpdateUser = "UPDATE users SET firstName = ?, lastName = ?, email = ?, age = ?, height = ?, activo = ? " +
		"WHERE id = ?"
	DeleteUser = "DELETE FROM users WHERE id = ?"
)

/*
type User struct {
	Id           int     `json:"-"`
	FirstName    string  `json:"firstName" binding:"required"`
	LastName     string  `json:"lastName" binding:"required"`
	Email        string  `json:"email" binding:"required"`
	Age          int     `json:"age" binding:"required"`
	Height       float64 `json:"height" binding:"required"`
	Active       bool    `json:"active" binding:"required"`
	CreationDate string  `json:"creationDate" binding:"required"`
}
*/
type Repository interface {
	Store(user domain.User) (domain.User, error)
	GetOne(id int) (domain.User, error)
	GetOneWithContext(ctx context.Context, id int) (domain.User, error)
	GetAll() ([]domain.User, error)
	GetByName(firstName string) (domain.User, error)
	Update(user domain.User) (domain.User, error)
	UpdateWithContext(ctx context.Context, user domain.User) (domain.User, error)
	//GetAll() ([]domain.User, error)
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository() Repository {

	return &repository{
		db: db.StorageDB,
	}
}

func (r *repository) Store(user domain.User) (domain.User, error) {
	// se inicializa la base
	stmt, err := r.db.Prepare(StoreUser) //se prepara el SQL
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Para optimizar el consumo de memoria.
	var result sql.Result
	result, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Age, user.Height, user.Activo, user.CreatedAt)
	if err != nil {
		return domain.User{}, err
	}
	insertedId, _ := result.LastInsertId() // obtenemos el Id insertado del sql.Result
	user.Id = int(insertedId)

	return user, nil
}

func (r *repository) GetOne(id int) (domain.User, error) {
	var user domain.User
	rows, err := r.db.Query(GetUser, id)
	if err != nil {
		return user, err
	}

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Age, &user.Height, &user.Activo, &user.CreatedAt); err != nil {
			return user, err
		}
	}
	return user, nil
}

func (r *repository) GetByName(firstName string) (domain.User, error) {
	var user domain.User

	rows, err := r.db.Query(GetUserByName, firstName)
	if err != nil {
		return user, err
	}

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Age, &user.Height, &user.Activo, &user.CreatedAt); err != nil {
			return user, err
		}
	}
	return user, nil
}

func (r *repository) GetAll() ([]domain.User, error) {
	var users []domain.User

	db := db.StorageDB
	rows, err := db.Query(GetAll)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for rows.Next() {
		//extraemos el user por cada fila
		var user domain.User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Age, &user.Height, &user.Activo, &user.CreatedAt); err != nil {
			log.Fatal(err)
			return nil, err
		}
		//agregamos los objetos al slice de users
		users = append(users, user)
	}

	return users, nil
}

func (r *repository) Update(user domain.User) (domain.User, error) {
	db := db.StorageDB
	stmt, err := db.Prepare(UpdateUser)
	if err != nil {
		return domain.User{}, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Age, user.Height, user.Activo, user.Id)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *repository) Delete(id int) error {
	db := db.StorageDB
	stmt, err := db.Prepare(DeleteUser)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetOneWithContext(ctx context.Context, id int) (domain.User, error) {
	var user domain.User
	db := db.StorageDB
	getQuery := GetUser

	rows, err := db.QueryContext(ctx, getQuery, id)
	if err != nil {
		return user, err
	}
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Age, &user.Height, &user.Activo, &user.CreatedAt); err != nil {
			return user, err
		}
	}
	return user, nil
}

func (r *repository) UpdateWithContext(ctx context.Context, user domain.User) (domain.User, error) {
	db := db.StorageDB

	stmt, err := db.PrepareContext(ctx, UpdateUser)
	if err != nil {
		return domain.User{}, err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, user.FirstName, user.LastName, user.Email, user.Age, user.Height, user.Activo, user.Id)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
