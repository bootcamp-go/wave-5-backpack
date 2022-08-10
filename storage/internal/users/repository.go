package users

import (
	"context"
	"database/sql"
	"goweb/internal/domain"
)

const (
	getbyID     = "SELECT id, nombre, apellido, email, edad, altura, activo, fechaCreacion FROM users WHERE id = ?"
	getByName   = "SELECT id, nombre, apellido, email, edad, altura, activo, fechaCreacion FROM users WHERE nombre = ?"
	storeUser   = "INSERT INTO users (nombre, apellido, email, edad, altura, activo, fechaCreacion) VALUES(?, ? , ? , ?, ?, ?, ?)"
	getAllUsers = "SELECT id, nombre, apellido, email, edad, altura, activo, fechaCreacion FROM users"
	deleteUser  = "DELETE FROM users WHERE id = ?"
	updateUser  = "UPDATE users SET nombre = ?, apellido = ?, email = ?, edad = ?, altura = ?, activo = ?, fechaCreacion = ? WHERE id = ?"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.User, error)
	Store(ctx context.Context, user domain.User) (int, error)
	GetById(ctx context.Context, id int) (domain.User, error)
	GetByName(ctx context.Context, nombre string) (domain.User, error)
	Update(ctx context.Context, user domain.User) (domain.User, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, id int, apellido string, edad int) (domain.User, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.User, error) {
	var users []domain.User

	rows, err := r.db.Query(getAllUsers)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.Id, &user.Nombre, &user.Apellido, &user.Email, &user.Edad, &user.Altura, &user.Activo, &user.FechaCreacion); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *repository) Store(ctx context.Context, user domain.User) (int, error) {
	stmt, err := r.db.Prepare(storeUser)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(user.Nombre, user.Apellido, user.Email, user.Edad, user.Altura, user.Activo, user.FechaCreacion)
	if err != nil {
		return 0, err
	}

	id, _ := res.LastInsertId()
	user.Id = int(id)

	return int(id), nil
}

func (r *repository) GetById(ctx context.Context, id int) (domain.User, error) {
	stmt, err := r.db.Prepare(getbyID)
	if err != nil {
		return domain.User{}, err
	}

	defer stmt.Close()

	var user domain.User
	if err := stmt.QueryRowContext(ctx, id).Scan(&user.Id, &user.Nombre, &user.Apellido, &user.Email, &user.Edad, &user.Altura, &user.Activo, &user.FechaCreacion); err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (r *repository) GetByName(ctx context.Context, nombre string) (domain.User, error) {
	var user domain.User

	rows, err := r.db.Query(getByName, nombre)
	if err != nil {
		return domain.User{}, err
	}

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Nombre, &user.Apellido, &user.Email, &user.Edad, &user.Altura, &user.Activo, &user.FechaCreacion); err != nil {
			return domain.User{}, err
		}
	}

	return user, nil
}

func (r *repository) Update(ctx context.Context, user domain.User) (domain.User, error) {
	stmt, err := r.db.Prepare(updateUser)
	if err != nil {
		return domain.User{}, err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, user.Nombre, user.Apellido, user.Email, user.Edad, user.Altura, user.Activo, user.FechaCreacion, user.Id)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil

}

func (r *repository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.Prepare(deleteUser)
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

func (r *repository) Patch(ctx context.Context, id int, apellido string, edad int) (domain.User, error) {
	return domain.User{}, nil
}
