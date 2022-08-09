package usuarios

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
)

type repositoryBD struct {
	dbBD *sql.DB
}

const (
	GuardarUser   string = "INSERT INTO storage.users(nombre,apellido,email,edad,altura,activo,fechaCreacion) VALUES(?,?,?,?,?,?,?)"
	GetByNameUser string = "SELECT id, nombre, apellido, email, edad, altura, activo, fechaCreacion FROM storage.users WHERE nombre =?"
	GetAllUser    string = "SELECT id, nombre, apellido, email, edad, altura, activo, fechaCreacion FROM storage.users"
	GetAllUserTO  string = "SELECT SLEEP(10) FROM DUAL"
	UpdateUser    string = "UPDATE storage.users SET nombre = ?, apellido = ?, email = ?, edad = ?, altura = ?, activo = ?, fechaCreacion = ? WHERE  id = ?"
)

func NewRepositoryBD(dbb *sql.DB) Repository {
	return &repositoryBD{
		dbBD: dbb,
	}
}
func (r *repositoryBD) GetAll(ctx context.Context) ([]domain.Usuarios, error) {
	var user domain.Usuarios
	var listUser []domain.Usuarios
	rows, err := r.dbBD.QueryContext(ctx, GetAllUser)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Nombre, &user.Apellido, &user.Email, &user.Edad, &user.Altura, &user.Activo, &user.FechaCreacion); err != nil {
			return nil, err
		}
		listUser = append(listUser, user)
	}
	return listUser, nil
}

func (r *repositoryBD) Guardar(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fecha string) (domain.Usuarios, error) {
	stmt, err := r.dbBD.Prepare(GuardarUser)
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
func (r *repositoryBD) Update(ctx context.Context, id int, nombre, apellido, email string, edad int, altura float64, activo bool, fecha string) (domain.Usuarios, error) {
	stmt, err := r.dbBD.Prepare(UpdateUser)
	if err != nil {
		return domain.Usuarios{}, nil
	}
	defer stmt.Close()
	_, err2 := stmt.ExecContext(ctx, id, nombre, apellido, email, edad, altura, activo, fecha)
	if err2 != nil {
		return domain.Usuarios{}, err2
	}
	var userD domain.Usuarios
	userD.Nombre = nombre
	userD.Apellido = apellido
	userD.Email = email
	userD.Edad = edad
	userD.Altura = altura
	userD.Activo = activo
	userD.FechaCreacion = fecha
	userD.Id = id
	return userD, nil
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
	rows, err := r.dbBD.Query(GetByNameUser, name)
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
