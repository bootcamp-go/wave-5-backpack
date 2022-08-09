package users

import (
	"database/sql"
	"ejercicioTM/internal/domain"
	"log"
	"time"
)

type Repository interface {
	GetFullData(id int) ([]domain.UserAndWarehouse, error)
	GetByName(nombre string) (domain.Usuarios, error)
	Store(domain.Usuarios) (domain.Usuarios, error)
	Update(domain.Usuarios) (domain.Usuarios, error)
	GetOne(id int) (domain.Usuarios, error)
	GetAll() ([]domain.UserAndWarehouse, error)
	Delete(id int) error
}

const (
	GetUsuarioByName  string = "SELECT id, nombre, apellido, email, edad, altura, activo, fecha FROM usuarios WHERE nombre=?"
	InsertUsuario     string = "INSERT INTO usuarios(nombre, apellido, email, edad, altura, activo, fecha, warehouse_id) VALUES(?,?,?,?,?,?,?,?)"
	GetUsuarioById    string = "SELECT id, nombre, apellido, email, edad, altura, activo, fecha FROM usuarios WHERE id = ?"
	UpdateUsuario     string = "UPDATE usuarios SET nombre = ?, apellido = ?, email = ?, edad = ?, altura = ?, activo = ?, fecha = ? WHERE id = ?"
	GetAllUsuarios    string = "SELECT * FROM usuarios"
	DeleteUsuarioById string = "DELETE FROM usuarios WHERE id = ?"
	GetFullData       string = "SELECT u.id, u.nombre, u.apellido, u.email, u.edad, u.altura, u.activo, u.fecha, w.id, w.name, w.address FROM usuarios AS u INNER JOIN warehouses AS w ON u.warehouse_id = w.id WHERE u.id = ?"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

//Inicialización de base

// Implementación GetByName
func (r *repository) GetByName(nombre string) (domain.Usuarios, error) {
	//Para asignar los valores de la consulta
	var usuario domain.Usuarios
	//Consulta la información del usuario cuyo nombre coincida.
	//TENER EN CUENTA: Es mejor especificar cada campo en vez de hacer uso del *
	rows, err := r.db.Query(GetUsuarioByName, nombre)
	if err != nil {
		return domain.Usuarios{}, err
	}
	//Asignación de valores
	for rows.Next() {
		if err := rows.Scan(&usuario.Id, &usuario.Nombre, &usuario.Apellido, &usuario.Email, &usuario.Edad, &usuario.Altura, &usuario.Activo, &usuario.Fecha); err != nil {
			return domain.Usuarios{}, err
		}
	}
	return usuario, nil
}

// Implementación Store()
func (r *repository) Store(usuario domain.Usuarios) (domain.Usuarios, error) {

	//Preparación de sentencia SQL
	stmt, err := r.db.Prepare(InsertUsuario)
	if err != nil {
		return domain.Usuarios{}, err
	}
	// Cierre de sentencia para evitar consumo de memoria
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(usuario.Nombre, usuario.Apellido, usuario.Email, usuario.Edad, usuario.Altura, true, time.Now().Format(time.RFC3339), 1)
	if err != nil {
		return domain.Usuarios{}, err
	}
	//de sql.Result devuelto en ejecución se obtiene el Id insertado
	insertedId, _ := result.LastInsertId()
	usuario.Id = int(insertedId)
	return usuario, nil
}

//Función para búsqueda por id de usuario
func (r *repository) GetOne(id int) (domain.Usuarios, error) {
	//Para asignación de valores de la consulta
	var usuario domain.Usuarios
	//Consulta la información del usuario cuyo nombre coincida.
	//TENER EN CUENTA: Es mejor especificar cada campo en vez de hacer uso del *
	rows, err := r.db.Query(GetUsuarioById, id)
	if err != nil {
		return domain.Usuarios{}, err
	}
	//Asignación de valores
	for rows.Next() {
		if err := rows.Scan(&usuario.Id, &usuario.Nombre, &usuario.Apellido, &usuario.Email, &usuario.Edad, &usuario.Altura, &usuario.Activo, &usuario.Fecha); err != nil {
			return domain.Usuarios{}, err
		}
	}
	return usuario, nil
}

//Función para actualizar información de un usuario de acuerdo a su id
func (r *repository) Update(usuario domain.Usuarios) (domain.Usuarios, error) {
	//Preparando sentencia
	stmt, err := r.db.Prepare(UpdateUsuario)
	if err != nil {
		log.Fatal(err)
	}
	//Se cierra la sentencia para evitar consumo de memoria
	defer stmt.Close()
	_, err = stmt.Exec(usuario.Nombre, usuario.Apellido, usuario.Email, usuario.Edad, usuario.Altura, true, time.Now().Format(time.RFC3339), usuario.Id)
	if err != nil {
		return domain.Usuarios{}, err
	}
	return usuario, nil
}

//Implementación del GetAll
func (r *repository) GetAll() ([]domain.UserAndWarehouse, error) {
	var userAndWarehouses []domain.UserAndWarehouse
	rows, err := r.db.Query(GetAllUsuarios)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var userAndWarehouse domain.UserAndWarehouse
		err := rows.Scan(
			&userAndWarehouse.Id,
			&userAndWarehouse.Nombre,
			&userAndWarehouse.Apellido,
			&userAndWarehouse.Email,
			&userAndWarehouse.Edad,
			&userAndWarehouse.Altura,
			&userAndWarehouse.Activo,
			&userAndWarehouse.Fecha,
			&userAndWarehouse.Warehouse.Id)
		if err != nil {
			return nil, err
		}
		userAndWarehouses = append(userAndWarehouses, userAndWarehouse)
	}
	return userAndWarehouses, nil
}

//Implementación del Delete
func (r *repository) Delete(id int) error {
	//Sentencia delete
	stmt, err := r.db.Prepare(DeleteUsuarioById)
	if err != nil {
		return err
	}
	//Se cierra para evitar consumo de memoria
	defer stmt.Close()
	//No nos interesa el result por ser delete y no retornar nada
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

//Implementación del GelFullData para usuarios y warehouses
func (r *repository) GetFullData(id int) ([]domain.UserAndWarehouse, error) {
	var userAndWarehouses []domain.UserAndWarehouse
	rows, err := r.db.Query(GetFullData, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var userAndWarehouse domain.UserAndWarehouse
		err := rows.Scan(
			&userAndWarehouse.Id,
			&userAndWarehouse.Nombre,
			&userAndWarehouse.Apellido,
			&userAndWarehouse.Email,
			&userAndWarehouse.Edad,
			&userAndWarehouse.Altura,
			&userAndWarehouse.Activo,
			&userAndWarehouse.Fecha,
			&userAndWarehouse.Warehouse.Id,
			&userAndWarehouse.Warehouse.Name,
			&userAndWarehouse.Warehouse.Address)
		if err != nil {
			return nil, err
		}
		userAndWarehouses = append(userAndWarehouses, userAndWarehouse)
	}
	return userAndWarehouses, nil
}
