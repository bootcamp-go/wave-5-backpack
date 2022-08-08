package products

import (
	"database/sql"
	"ejercicioTT/internal/domain"
	"log"
	"time"
)

type Repository interface {
	GetByName(nombre string) (domain.Usuarios, error)
	Store(domain.Usuarios) (domain.Usuarios, error)
	Update(domain.Usuarios) (domain.Usuarios, error)
	// GetAll() ([]domain.Usuarios, error)
	// Delete(id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

//Inicialización de base

// Ejercicio 1
// Implementación GetByName
func (r *repository) GetByName(nombre string) (domain.Usuarios, error) {
	//Para asignar los valores de la consulta
	var usuario domain.Usuarios
	//Consulta la información del usuario cuyo nombre coincida.
	//TENER EN CUENTA: Es mejor especificar cada campo en vez de hacer uso del *
	rows, err := r.db.Query("SELECT id, nombre, apellido, email, edad, altura, activo, fecha FROM usuarios WHERE nombre = ?", nombre)
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

// Ejercicio 2
// Implementación Store()
func (r *repository) Store(usuario domain.Usuarios) (domain.Usuarios, error) {

	//Preparación de sentencia SQL
	stmt, err := r.db.Prepare("INSERT INTO usuarios(nombre, apellido, email, edad, altura, activo, fecha) VALUES(?,?,?,?,?")
	if err != nil {
		return domain.Usuarios{}, err
	}
	// Cierre de sentencia para evitar consumo de memoria
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(usuario.Nombre, usuario.Apellido, usuario.Email, usuario.Edad, usuario.Altura, usuario.Activo, usuario.Fecha)
	if err != nil {
		return domain.Usuarios{}, err
	}
	//de sql.Result devuelto en ejecución se obtiene el Id insertado
	insertedId, _ := result.LastInsertId()
	usuario.Id = int(insertedId)
	return usuario, nil
}

func (r *repository) Update(usuario domain.Usuarios) (domain.Usuarios, error) {
	//Preparando sentencia
	stmt, err := r.db.Prepare("UPDATE products SET nombre = ?, apellido = ?, email = ?, edad = ?, altura = ?, activo = ?, fecha = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	//Se cierra la sentencia para evitar consumo de memoria
	defer stmt.Close()
	_, err = stmt.Exec(usuario.Nombre, usuario.Apellido, usuario.Email, usuario.Edad, usuario.Altura, true, time.Now(), usuario.Id)
	if err != nil {
		return domain.Usuarios{}, err
	}
	return usuario, nil
}
