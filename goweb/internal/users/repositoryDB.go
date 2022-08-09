package users

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/bootcamp-go/wave-5-backpack/internal/domain"
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
	queryByName = "SELECT id, name, lastname, email, age, height, active, doCreation FROM users WHERE name = ?"
	queryAll    = "SELECT id, name, lastname, email, age, height, active, doCreation FROM users"
	queryAllTimeOut    = "SELECT SLEEP(10) FROM DUAL"
	queryStore  = "INSERT INTO users(name, lastname, email, age, height, active, doCreation) VALUES( ?, ?, ?, ?, ?, ?, ? )"
	queryById   = "SELECT id, name, lastname, email, age, height, active, doCreation FROM users us WHERE us.id = ?"
)

func (r *repositoryDB) GetAll(ctx context.Context) ([]domain.User, error) {
	var allUsers []domain.User
	rows, err := r.db.QueryContext(ctx, queryAll)
	if err != nil {
		return []domain.User{}, err
	}
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Lastname, &user.Email, &user.Age, &user.Height, &user.Active, &user.DoCreation); err != nil {
			return []domain.User{}, err
		}
		allUsers = append(allUsers, user)
	}

	return allUsers, nil
}
func (r *repositoryDB) GetById(ctx context.Context, id int) (domain.User, error) {
	var user domain.User
	rows, err := r.db.QueryContext(ctx, queryById, id )
	if err != nil {
		return domain.User{}, err
	}
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Lastname, &user.Email, &user.Age, &user.Height, &user.Active, &user.DoCreation); err != nil {
			return domain.User{}, err
		}
	}
	return user, nil
}

func (r *repositoryDB) DeleteUser(id int) error {

	return nil
}

func (r *repositoryDB) LastId() (int, error) {

	return 0, nil
}

func (r *repositoryDB) GetByName(name string) ([]domain.User, error) {
	var user domain.User
	fmt.Println(name)
	var listUser []domain.User
	rows, err := r.db.Query(queryByName, name)
	if err != nil {
		return []domain.User{}, err
	}
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Lastname, &user.Email, &user.Age, &user.Height, &user.Active, &user.DoCreation); err != nil {
			return []domain.User{}, err
		}
		listUser = append(listUser, user)
	}
	fmt.Println(listUser)
	return listUser, nil
}

func (r *repositoryDB) StoreUser(id int, name, lastname, email string, age int, height float32, active bool, doCreation string) (domain.User, error) {

	stmt, err := r.db.Prepare(queryStore)
	if err != nil {
		return domain.User{}, err
	}
	defer stmt.Close()
	var result sql.Result
	result, err2 := stmt.Exec(name, lastname, email, age, height, active, doCreation)
	if err2 != nil {
		return domain.User{}, err
	}
	insertedId, _ := result.LastInsertId()
	newUser := domain.User{
		ID: int(insertedId), Name: name, Lastname: lastname, Email: email, Age: age, Height: height, Active: active, DoCreation: doCreation}
	return newUser, nil
}

func (r *repositoryDB) UpdateUser(id int, name, lastname, email string, age int, height float32, active bool, doCreation string) (domain.User, error) {

	newUser := domain.User{
		ID:         id,
		Name:       name,
		Lastname:   lastname,
		Email:      email,
		Age:        age,
		Height:     height,
		Active:     active,
		DoCreation: doCreation,
	}

	return newUser, nil

}

func (r *repositoryDB) UpdateLastnameAndAge(id int, lastname string, age int) (*domain.User, error) {

	return nil, nil

}
