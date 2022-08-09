package usuarios

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"storage/Clase2-TM/internal/domain"
	"time"
)

const (
	GetUser       string = "SELECT id, nombre, apellido, email, edad, altura, activo, fecha_de_creacion FROM usuarios WHERE id = ?"
	GetAllUsers   string = "SELECT id, nombre, apellido, email, edad, altura, activo, fecha_de_creacion FROM usuarios"
	GetUserByName string = "SELECT id, nombre, apellido, email, edad, altura, activo, fecha_de_creacion FROM usuarios WHERE nombre = ?"
	InsertUser    string = "INSERT INTO usuarios(nombre, apellido, email, edad, altura, activo, fecha_de_creacion) VALUES( ?, ?, ?, ?, ?, ?, ? )"
	UpdateUser    string = "UPDATE usuarios SET nombre = ?, apellido = ?, email = ?, edad = ?, altura = ?, activo = ? WHERE id = ?"
	DeleteUser    string = "DELETE FROM usuarios WHERE id = ?"
)

type Repository interface {
	GetAll() ([]domain.Usuario, error)
	GetByName(nombre string) (domain.Usuario, error)
	Store(domain.Usuario) (domain.Usuario, error)
	//LastID() (int, error)
	Update(context.Context, domain.Usuario) (domain.Usuario, error)
	//UpdateSurnameAndAge(id int, surname string, age int) (domain.Usuario, error)
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]domain.Usuario, error) {
	var usuarios []domain.Usuario
	rows, err := r.db.Query(GetAllUsers)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// se recorren todas las filas
	for rows.Next() {
		// por cada fila se obtiene un objeto del tipo Product
		var user domain.Usuario
		if err := rows.Scan(&user.Id, &user.Nombre, &user.Apellido, &user.Email, &user.Edad, &user.Altura, &user.Activo, &user.FechaDeCreacion); err != nil {
			log.Fatal(err)
			return nil, err
		}
		//se añade el objeto obtenido al slice products
		usuarios = append(usuarios, user)
	}
	return usuarios, nil

}

func (r *repository) Store(user domain.Usuario) (domain.Usuario, error) { // se inicializa la base

	stmt, err := r.db.Prepare(InsertUser) // se prepara la sentencia SQL a ejecutar
	if err != nil {
		return domain.Usuario{}, err
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	var result sql.Result
	fechaInsercion := time.Now().String()
	fmt.Println(fechaInsercion)
	result, err = stmt.Exec(user.Nombre, user.Apellido, user.Email, user.Edad, user.Altura, user.Activo, fechaInsercion) // retorna un sql.Result y un error
	if err != nil {
		return domain.Usuario{}, err
	}
	insertedId, _ := result.LastInsertId() // del sql.Resul devuelto en la ejecucion obtenemos el Id insertado
	user.Id = int(insertedId)
	user.FechaDeCreacion = fechaInsercion
	return user, nil
}

/*
func (r *repository) LastID() (int, error) {
	var us []domain.Usuario
	if err := r.db.Read(&us); err != nil {
		return 0, fmt.Errorf(FailReading)
	}
	if len(us) == 0 {
		return 0, nil
	}

	return us[len(us)-1].Id, nil
}
*/
func (r *repository) GetOne(id int) (domain.Usuario, error) {
	var user domain.Usuario

	rows, err := r.db.Query(GetUser, id)
	if err != nil {
		return domain.Usuario{}, err
	}
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Nombre, &user.Apellido, &user.Email, &user.Edad, &user.Altura, &user.Activo, &user.FechaDeCreacion); err != nil {
			return domain.Usuario{}, err
		}
	}
	return user, nil
}
func (r *repository) GetByName(name string) (domain.Usuario, error) {
	var user domain.Usuario

	rows, err := r.db.Query(GetUserByName, name)
	if err != nil {
		return domain.Usuario{}, err
	}
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Nombre, &user.Apellido, &user.Email, &user.Edad, &user.Altura, &user.Activo, &user.FechaDeCreacion); err != nil {
			return domain.Usuario{}, err
		}
	}
	return user, nil
}

func (r *repository) Update(ctx context.Context, user domain.Usuario) (domain.Usuario, error) {
	stmt, err := r.db.PrepareContext(ctx, UpdateUser) // se prepara la sentencia SQL a ejecutar
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	_, err = stmt.ExecContext(ctx, user.Nombre, user.Apellido, user.Email, user.Edad, user.Altura, user.Activo, user.Id)
	if err != nil {
		return domain.Usuario{}, err
	}
	return user, nil
}

func (r *repository) Delete(id int) error {
	stmt, err := r.db.Prepare(DeleteUser)
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
