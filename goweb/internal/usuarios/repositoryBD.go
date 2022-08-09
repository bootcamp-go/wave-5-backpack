package usuarios

import (
	"database/sql"
	"fmt"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
)

type repositoryBD struct {
	dbBD *sql.DB
}

func NewRepositoryBD(dbb *sql.DB) Repository {
	return &repositoryBD{
		dbBD: dbb,
	}
}
func (r *repositoryBD) GetAll() ([]domain.Usuarios, error) {
	return nil, nil
}
func (r *repositoryBD) Guardar(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fecha string) (domain.Usuarios, error) {
	stmt, err := r.dbBD.Prepare("INSERT INTO storage.users(nombre,apellido,email,edad,altura,activo,fechaCreacion) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		return domain.Usuarios{}, nil
	}
	defer stmt.Close()
	result, err2 := stmt.Exec(nombre, apellido, email, edad, altura, activo, fecha)
	if err2 != nil {
		return domain.Usuarios{}, err2
	}
	insertId, _ := result.LastInsertId()
	var userD domain.Usuarios
	userD.Nombre = nombre
	userD.Apellido = apellido
	userD.Email = email
	userD.Edad = edad
	userD.Altura = altura
	userD.Activo = activo
	userD.FechaCreacion = fecha
	userD.Id = int(insertId)
	return userD, nil
}
func (r *repositoryBD) LastId() (int, error) {
	return 0, nil
}
func (r *repositoryBD) Update(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fecha string) (domain.Usuarios, error) {
	return domain.Usuarios{}, nil
}
func (r *repositoryBD) Delete(id int) error {
	return nil
}
func (r *repositoryBD) UpdateNameAndLastName(id int, name string, apellido string) (domain.Usuarios, error) {
	return domain.Usuarios{}, nil
}
func (r *repositoryBD) GetById(id int) (domain.Usuarios, error) {
	return domain.Usuarios{}, nil
}

func (r *repositoryBD) GetByName(name string) ([]domain.Usuarios, error) {
	var user domain.Usuarios
	var listUser []domain.Usuarios
	rows, err := r.dbBD.Query("SELECT id, nombre, apellido, email, edad, altura, activo, fechaCreacion FROM storage.users WHERE nombre =?", name)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	fmt.Println(user)
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Nombre, &user.Apellido, &user.Email, &user.Edad, &user.Altura, &user.Activo, &user.FechaCreacion); err != nil {
			return nil, err
		}
		listUser = append(listUser, user)
	}
	fmt.Println(user)
	return listUser, nil
}
