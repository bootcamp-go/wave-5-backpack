package usuarios

import (
	"database/sql"
	"fmt"
	"storage/Clase1/internal/domain"
	"time"
)

type Repository interface {
	//GetAll() ([]domain.Usuario, error)
	GetByName(nombre string) (domain.Usuario, error)
	Store(domain.Usuario) (domain.Usuario, error)
	//LastID() (int, error)
	//Update(id int, nombre, apellido, email string, edad, altura int, activo bool) (domain.Usuario, error)
	//UpdateSurnameAndAge(id int, surname string, age int) (domain.Usuario, error)
	//Delete(id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

/*
func (r *repository) GetAll() ([]domain.Usuario, error) {
	var us []domain.Usuario
	if err := r.db.Read(&us); err != nil {
		return nil, fmt.Errorf(FailReading)
	}
	return us, nil
}
*/
func (r *repository) Store(user domain.Usuario) (domain.Usuario, error) { // se inicializa la base

	stmt, err := r.db.Prepare("INSERT INTO usuarios(nombre, apellido, email, edad, altura, activo, fecha_de_creacion) VALUES( ?, ?, ?, ?, ?, ?, ? )") // se prepara la sentencia SQL a ejecutar
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

	rows, err := r.db.Query("SELECT id, nombre, apellido, email, edad, altura, activo, fecha_de_creacion FROM usuarios WHERE id = ?", id)
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

	rows, err := r.db.Query("SELECT id, nombre, apellido, email, edad, altura, activo, fecha_de_creacion FROM usuarios WHERE nombre = ?", name)
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

/*
func (r *repository) Update(id int, name, productType string, count int, price float64) (domain.Product, error) {
	stmt, err := r.db.Prepare("UPDATE products SET name = ?, type = ?, count = ?, price = ? WHERE id = ?") // se prepara la sentencia SQL a ejecutar
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	product := domain.Product{ID: id, Name: name, Type: productType, Count: count, Price: price}
	_, err = stmt.Exec(name, productType, count, price, id)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (r *repository) Delete(id int) error {
	var us []domain.Usuario
	if err := r.db.Read(&us); err != nil {
		return fmt.Errorf(FailReading)
	}
	deleted := false
	var index int
	for i := range us {
		if us[i].Id == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf(UserNotFound, id)
	}

	us = append(us[:index], us[index+1:]...)

	if err := r.db.Write(us); err != nil {
		return fmt.Errorf(FailWriting, err)
	}

	return nil
}
*/
