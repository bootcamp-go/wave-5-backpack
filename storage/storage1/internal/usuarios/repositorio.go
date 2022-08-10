package usuarios

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/del_rio/web-server/internal/domain"
	"github.com/del_rio/web-server/pkg/store"
)

type Repository interface {
	GetAll() ([]domain.Usuario, error)
	GetById(id int) (domain.Usuario, error)
	GetByName(name string) domain.Usuario
	Store(usuario domain.Usuario) (domain.Usuario, error)
	LastId() (int, error)
	UpdateUsuario(ctx context.Context, usuario domain.Usuario) (domain.Usuario, error)
	UpdateAtributos(Nombre, Apellido, Email, Fecha_creacion string, Id, Edad, Altura int, Activo *bool) (domain.Usuario, error)
	DeleteUsuario(Id int) error
}

type repository struct {
	db     store.Store
	dbReal *sql.DB
}

const (
	qGetAll        string = "select id,nombre,apellido,email,edad,altura,activo,fecha_creacion from usuario"
	qGetById       string = "select id,nombre,apellido,email,edad,altura,activo,fecha_creacion from usuario where id = ?"
	qGetByName     string = "select id,nombre,apellido,email,edad,altura,activo,fecha_creacion from usuario where nombre = ?"
	qStore         string = "INSERT INTO usuario(nombre,apellido,email,edad,altura,activo,fecha_creacion) VALUES( ?, ?, ?, ?,?,?,? )"
	qUpdateUsuario string = "UPDATE usuario SET nombre=?, apellido=?, email=?, edad=?, altura=?, activo=?, fecha_creacion=? WHERE id=?"
	qDeleteUsuario string = "DELETE FROM usuario WHERE id = ?"
)

func (r *repository) GetAll() ([]domain.Usuario, error) {
	var usuarios []domain.Usuario
	rows, err := r.dbReal.Query(qGetAll)
	if err != nil {
		log.Println(err)
		return usuarios, err
	}
	fmt.Println("prueba")
	println(rows)
	for rows.Next() {
		usuario := domain.Usuario{}
		if err := rows.Scan(&usuario.Id, &usuario.Nombre, &usuario.Apellido, &usuario.Email, &usuario.Edad, &usuario.Altura, &usuario.Activo, &usuario.Fecha_creacion); err != nil {
			log.Println(err.Error())
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

// `id` int(11) NOT NULL,
// `nombre` varchar(60) NOT NULL,
// `apellido` varchar(60) NOT NULL,
// `email` varchar(60) NOT NULL,
// `edad` int(11) NOT NULL,
// `altura` int(11) NOT NULL,
// `activo` tinyint(1) NOT NULL,
// `Fecha_creacion` timestamp NOT NULL

func (r *repository) GetById(id int) (domain.Usuario, error) {
	var usuario domain.Usuario
	rows, err := r.dbReal.Query(qGetById, id)
	if err != nil {
		log.Println(err)
		return usuario, err
	}
	fmt.Println("prueba")
	println(rows)
	for rows.Next() {
		if err := rows.Scan(&usuario.Id, &usuario.Nombre, &usuario.Apellido, &usuario.Email, &usuario.Edad, &usuario.Altura, &usuario.Activo, &usuario.Fecha_creacion); err != nil {
			log.Println(err.Error())
			nullUsuario := domain.Usuario{}
			return nullUsuario, err
		}
	}
	return usuario, nil
}

func (r *repository) GetByName(name string) domain.Usuario {
	var usuario domain.Usuario
	rows, err := r.dbReal.Query(qGetByName, name)
	if err != nil {
		log.Println(err)
		return usuario
	}
	fmt.Println("prueba")
	println(rows)
	for rows.Next() {
		if err := rows.Scan(&usuario.Id, &usuario.Nombre, &usuario.Apellido, &usuario.Email, &usuario.Edad, &usuario.Altura, &usuario.Activo, &usuario.Fecha_creacion); err != nil {
			log.Println(err.Error())
			return usuario
		}
	}
	return usuario
}

func (r *repository) Store(usuario domain.Usuario) (domain.Usuario, error) {
	stmt, err := r.dbReal.Prepare(qStore) // se prepara el SQL
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	var activoInt int8
	if usuario.Activo {
		activoInt = 1
	}
	var result sql.Result
	result, err = stmt.Exec(usuario.Nombre, usuario.Apellido, usuario.Email, usuario.Edad, usuario.Altura, activoInt, usuario.Fecha_creacion) // retorna un sql.Result y un error
	if err != nil {
		return domain.Usuario{}, err
	}
	insertedId, _ := result.LastInsertId() // del sql.Resul devuelto en la ejecución obtenemos el Id insertado
	usuario.Id = int(insertedId)
	return usuario, nil
}
func (r *repository) LastId() (int, error) {
	//aqui podria pasar algo
	var ListUsuarios []domain.Usuario
	if err := r.db.Read(&ListUsuarios); err != nil {
		return 0, err
	}
	if len(ListUsuarios) == 0 {
		return 0, nil
	}
	lastId := ListUsuarios[len(ListUsuarios)-1].Id
	return lastId, nil
}

func NewRepository(db store.Store, dbReal *sql.DB) Repository {
	return &repository{db, dbReal}
}

func (r *repository) UpdateUsuario(ctx context.Context, usuarioCambiado domain.Usuario) (domain.Usuario, error) {
	stmt, err := r.dbReal.Prepare(qUpdateUsuario) // se prepara la sentencia SQL a ejecutar
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	//nombre=?, apellido=?, email=?, edad=?, altura=?, activo=?, fecha_creacion=?
	result, err := stmt.ExecContext(ctx,
		usuarioCambiado.Nombre,
		usuarioCambiado.Apellido,
		usuarioCambiado.Email,
		usuarioCambiado.Edad,
		usuarioCambiado.Altura,
		usuarioCambiado.Activo,
		usuarioCambiado.Fecha_creacion,
		usuarioCambiado.Id,
	)
	fmt.Println("mira este resultado de put")
	fmt.Println(result)
	if err != nil {
		return domain.Usuario{}, err
	}
	return usuarioCambiado, nil
}

func (r *repository) UpdateAtributos(Nombre, Apellido, Email, Fecha_creacion string, Id, Edad, Altura int, Activo *bool) (domain.Usuario, error) {
	var ListUsuarios []domain.Usuario
	if err := r.db.Read(&ListUsuarios); err != nil {
		return domain.Usuario{}, err
	}
	for i := 0; i < len(ListUsuarios); i++ {
		if Id == ListUsuarios[i].Id {
			// Nombre, Apellido, Email, Fecha_creacion string
			r.changeString(&ListUsuarios[i].Nombre, Nombre)
			r.changeString(&ListUsuarios[i].Apellido, Apellido)
			r.changeString(&ListUsuarios[i].Email, Email)
			r.changeString(&ListUsuarios[i].Fecha_creacion, Fecha_creacion)
			// Edad, Altura int
			r.changeInt(&ListUsuarios[i].Edad, Edad)
			r.changeInt(&ListUsuarios[i].Altura, Altura)
			//Activo *bool
			r.changeBool(&ListUsuarios[i].Activo, Activo)
			if err := r.db.Write(ListUsuarios); err != nil {
				return domain.Usuario{}, err
			}
			return ListUsuarios[i], nil
		}
	}
	usuarioNulo := domain.Usuario{}
	return usuarioNulo, fmt.Errorf("ese Id pareciera no existir dentro de nuestra BD")
}

func (r *repository) DeleteUsuario(id int) error {
	stmt, err := r.dbReal.Prepare(qDeleteUsuario) // se prepara la sentencia SQL a ejecutar
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria

	_, err = stmt.Exec(id) // retorna un sql.Result y un error

	if err != nil {
		return err
	}

	return nil
}

// Nombre, Apellido, Email, Fecha_creacion string
func (r *repository) changeString(myString *string, newString string) {
	if newString != "" {
		*myString = newString
	}
}

// Edad, Altura int
func (r *repository) changeInt(myInt *int, newInt int) {
	if newInt != 0 {
		*myInt = newInt
	}
}

//Activo *bool
func (r *repository) changeBool(myBool *bool, newBool *bool) {
	if newBool != nil {
		*myBool = *newBool
	}
}
