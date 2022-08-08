package usuarios

import (
	"database/sql"
	"errors"
	"log"

	"github.com/anesquivel/wave-5-backpack/storage/arquitectura_ejercicio/internal/domain"
)

const (
	STRING_HUBO       = "Hubo un error ... "
	ERROR_READING     = "Hubo un error al leer los datos de la BD."
	ERR_WRITING       = "Hubo un error al guardar los datos en la BD."
	ERR_UPDATING_USER = ".. al no encontrar el usuario."
)

type Repository interface {
	GetAll() ([]domain.Usuario, error)
	GetOne(id int) (domain.Usuario, error)
	GetByName(name string) (domain.Usuario, error)
	Store(user domain.Usuario) (domain.Usuario, error)
	Update(id int, user domain.Usuario) (domain.Usuario, error)
	UpdateLastNameAndAge(id, age int, lastname string) (domain.Usuario, error)
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Store(user domain.Usuario) (domain.Usuario, error) {
	stmt, err := r.db.Prepare("INSERT INTO users(names, last_name, email, age, height, is_active, date_created) VALUES(?,?,?,?,?,?,?)")

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
	return []domain.Usuario{}, nil

}

func (r *repository) GetOne(id int) (domain.Usuario, error) {
	var user domain.Usuario

	rows, err := r.db.Query("SELECT * FROM users where id:= ?", id)

	if err != nil {
		return domain.Usuario{}, errors.New(ERROR_READING)
	}

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Names, &user.LastName, &user.Email, &user.Age, &user.Estatura, &user.IsActivo); err != nil {
			return domain.Usuario{}, errors.New("No se encontró el usuario.")
		}
	}
	return user, nil
}
func (r *repository) GetByName(name string) (domain.Usuario, error) {
	var user domain.Usuario

	rows, err := r.db.Query("SELECT * FROM users where names:= ?", name)

	if err != nil {
		return domain.Usuario{}, errors.New(ERROR_READING)
	}

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Names, &user.LastName, &user.Email, &user.Age, &user.Estatura, &user.IsActivo); err != nil {
			return domain.Usuario{}, errors.New("No se encontró el usuario.")
		}
	}
	return user, nil
}

func (r *repository) Delete(id int) error {
	return nil
}
func (r *repository) UpdateLastNameAndAge(id, age int, lastname string) (domain.Usuario, error) {
	return domain.Usuario{}, nil

}
func (r *repository) Update(id int, user domain.Usuario) (domain.Usuario, error) {
	stmt, err := r.db.Prepare("UPDATE products SET names = ?, last_name = ?, email = ?, age = ?, height = ? WHERE id = ?") // se prepara la sentencia SQL a ejecutar
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	userUpdate := domain.Usuario{
		Id:       id,
		Names:    user.Names,
		LastName: user.LastName,
		Email:    user.Email,
	}
	_, err = stmt.Exec(user.Names, user.LastName, user.Email, user.Age, user.Estatura, user.IsActivo)
	if err != nil {
		return domain.Usuario{}, err
	}
	return userUpdate, nil
}
