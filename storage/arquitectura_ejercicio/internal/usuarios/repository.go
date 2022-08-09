package usuarios

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/anesquivel/wave-5-backpack/storage/arquitectura_ejercicio/internal/domain"
)

const (
	STRING_HUBO          = "Hubo un error ... "
	ERROR_READING        = "Hubo un error al leer los datos de la BD."
	ERR_WRITING          = "Hubo un error al guardar los datos en la BD."
	ERR_UPDATING_USER    = ".. al no encontrar el usuario."
	QUERY_STORE          = "INSERT INTO users(names, last_name, email, age, height, is_active, date_created) VALUES(?,?,?,?,?,?,?)"
	QUERY_UPDATE         = "UPDATE users SET names = ?, last_name = ?, email = ?, age = ?, height = ? WHERE id = ?"
	QUERY_UPDATE_LASTAGE = "UPDATE users SET last_name = ?, age = ? WHERE id = ?"
	QUERY_GET_ONE        = "SELECT id, names, last_name, email, age, height, is_active FROM users where id = ?"
	QUERY_GET_BYNAME     = "SELECT id, names, last_name, email FROM users WHERE names = ?"
	QUERY_GET_ALL        = "SELECT id, names, last_name, email, age,height, is_active FROM users"
	QUERY_DELETE         = "DELETE FROM users WHERE id = ?"
)

type Repository interface {
	GetAll() ([]domain.Usuario, error)
	GetOne(id int) (domain.Usuario, error)
	GetByName(name string) (domain.Usuario, error)
	Store(user domain.Usuario) (domain.Usuario, error)
	Update(id int, user domain.Usuario) (domain.Usuario, error)
	UpdateLastNameAndAge(ctx context.Context, id, age int, lastname string) (domain.Usuario, error)
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Store(user domain.Usuario) (domain.Usuario, error) {
	stmt, err := r.db.Prepare(QUERY_STORE)

	if err != nil {
		log.Fatal("err:", err.Error())
	}

	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(user.Names, user.LastName, user.Email, user.Age, user.Estatura, user.IsActivo, user.DateCreated)

	if err != nil {
		return domain.Usuario{}, errors.New(ERR_WRITING)
	}

	insertId, _ := result.LastInsertId()
	user.Id = int(insertId)

	return user, nil
}

func (r *repository) GetAll() ([]domain.Usuario, error) {
	var users []domain.Usuario
	rows, err := r.db.Query(QUERY_GET_ALL)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var user domain.Usuario
		err := rows.Scan(&user.Id,
			&user.Names,
			&user.LastName,
			&user.Email,
			&user.Age,
			&user.Estatura,
			&user.IsActivo)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *repository) GetOne(id int) (domain.Usuario, error) {
	var user domain.Usuario

	rows, err := r.db.Query(QUERY_GET_ONE, id)

	if err != nil {
		return domain.Usuario{}, errors.New(ERROR_READING)
	}

	for rows.Next() {
		if err := rows.Scan(&user.Id,
			&user.Names,
			&user.LastName,
			&user.Email,
			&user.Age,
			&user.Estatura,
			&user.IsActivo); err != nil {
			return domain.Usuario{}, errors.New("No se encontró el usuario.")
		}
	}
	return user, nil
}
func (r *repository) GetByName(name string) (domain.Usuario, error) {
	var user domain.Usuario

	rows, err := r.db.Query(QUERY_GET_BYNAME, name)

	if err != nil {
		return domain.Usuario{}, errors.New(ERROR_READING + err.Error())
	}

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Names, &user.LastName, &user.Email)
		if err != nil {
			return domain.Usuario{}, errors.New(err.Error())
		}
	}
	return user, nil
}

func (r *repository) Delete(id int) error {
	stmt, err := r.db.Prepare(QUERY_DELETE)

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}
	return nil
}
func (r *repository) UpdateLastNameAndAge(ctx context.Context, id, age int, lastname string) (domain.Usuario, error) {
	stmt, err := r.db.Prepare(QUERY_UPDATE_LASTAGE) // se prepara la sentencia SQL a ejecutar
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	userUpdate := domain.Usuario{
		Id:       id,
		LastName: lastname,
		Age:      age,
	}
	_, err = stmt.ExecContext(ctx, userUpdate.LastName, userUpdate.Age, id)
	if err != nil {
		log.Fatal(err)
		return domain.Usuario{}, err
	}
	return userUpdate, nil

}
func (r *repository) Update(id int, user domain.Usuario) (domain.Usuario, error) {
	stmt, err := r.db.Prepare(QUERY_UPDATE) // se prepara la sentencia SQL a ejecutar
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	userUpdate := domain.Usuario{
		Id:       id,
		Names:    user.Names,
		LastName: user.LastName,
		Email:    user.Email,
		Estatura: user.Estatura,
		Age:      user.Age,
	}
	_, err = stmt.Exec(user.Names, user.LastName, user.Email, user.Age, user.Estatura, user.IsActivo, id)
	if err != nil {
		return domain.Usuario{}, err
	}
	return userUpdate, nil
}
