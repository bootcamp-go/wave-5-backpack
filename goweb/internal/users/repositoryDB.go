package users

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
)

type repositoryDB struct {
	db *sql.DB
}

func NewRepositoryDB(db *sql.DB) Repository {
	return &repositoryDB{
		db: db,
	}
}

const (
	GetAll       string = "SELECT id, name, last_name, email, age, height, active, creation_date FROM users"
	TimeOut10Sec string = "SELECT SLEEP(10) FROM DUAL"
	GetByName    string = "SELECT id, name, last_name, email, age, height, active, creation_date FROM users WHERE name = ?"
	StoreUser    string = "INSERT INTO users(name, last_name, email, age, height, active, creation_date) VALUES( ?, ?, ?, ?, ?, ?, ? )"
	UpdateUser   string = "UPDATE users SET name = ?, last_name = ?, email = ?, age = ?, height = ?, active = ?, creation_date = ? WHERE id = ?"
)

func (r *repositoryDB) GetAll(ctx context.Context) ([]domain.Users, error) {
	var user domain.Users
	var listUser []domain.Users
	rows, err := r.db.QueryContext(ctx, GetAll)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Name, &user.LastName, &user.Email, &user.Age, &user.Height, &user.Active, &user.CreationDate); err != nil {
			return nil, err
		}
		listUser = append(listUser, user)
	}
	return listUser, nil
}

func (r *repositoryDB) LastID(ctx context.Context) (int, error) {
	return 0, nil
}

func (r *repositoryDB) GetByName(ctx context.Context, name string) ([]domain.Users, error) {
	var user domain.Users
	var listUser []domain.Users
	rows, err := r.db.QueryContext(ctx, GetByName, name)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Name, &user.LastName, &user.Email, &user.Age, &user.Height, &user.Active, &user.CreationDate); err != nil {
			return nil, err
		}
		listUser = append(listUser, user)
	}
	return listUser, nil
}

func (r *repositoryDB) Store(ctx context.Context, id, age int, name, lastName, email, creationDate string, height float64, active bool) (domain.Users, error) {
	stmt, err := r.db.PrepareContext(ctx, StoreUser)
	if err != nil {
		return domain.Users{}, err
	}
	defer stmt.Close()
	var result sql.Result
	result, err2 := stmt.ExecContext(ctx, name, lastName, email, age, height, active, creationDate)
	if err2 != nil {
		return domain.Users{}, err
	}
	insertedId, _ := result.LastInsertId()
	user := domain.Users{
		Id: int(insertedId), Name: name, LastName: lastName, Email: email, Age: age, Height: height, Active: active, CreationDate: creationDate}
	return user, nil
}

func (r *repositoryDB) Update(ctx context.Context, id, age int, name, lastName, email, creationDate string, height float64, active bool) (domain.Users, error) {
	stmt, err := r.db.PrepareContext(ctx, UpdateUser)
	if err != nil {
		return domain.Users{}, err
	}
	defer stmt.Close()
	_, err2 := stmt.ExecContext(ctx, name, lastName, email, age, height, active, creationDate, id)
	if err2 != nil {
		return domain.Users{}, err
	}
	user := domain.Users{
		Id: id, Name: name, LastName: lastName, Email: email, Age: age, Height: height, Active: active, CreationDate: creationDate}
	return user, nil
}

func (r *repositoryDB) UpdateLastNameAndAge(ctx context.Context, id, age int, lastName string) (domain.Users, error) {
	return domain.Users{}, fmt.Errorf(UserNotFound, id)
}

func (r *repositoryDB) Delete(ctx context.Context, id int) error {
	return fmt.Errorf(UserNotFound, id)
}
