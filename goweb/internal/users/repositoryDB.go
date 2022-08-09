package users

import (
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

func (r *repositoryDB) GetAll() ([]domain.User, error) {
	var allUsers []domain.User

	return allUsers, nil
}
func (r *repositoryDB) GetById(id int) (domain.User, error) {

	return domain.User{}, fmt.Errorf(UserNotFound, id)
}
func (r *repositoryDB) DeleteUser(id int) error {

	return nil
}

func (r *repositoryDB) LastId() (int, error) {

	return 0, nil
}
func (r *repositoryDB) GetByName(name string) ([]domain.User, error) {
	var user domain.User
	var listUser []domain.User
	rows, err := r.db.Query("SELECT id, name, last_name, email, age, height, active, creation_date FROM users WHERE name = ?", name)
	if err != nil {
		return []domain.User{}, err
	}
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Lastname, &user.Email, &user.Age, &user.Height, &user.Active, &user.DoCreation); err != nil {
			return []domain.User{}, err
		}
		listUser = append(listUser, user)
	}
	return listUser, nil
}
func (r *repositoryDB) StoreUser(id int, name, lastname, email string, age int, height float32, active bool, doCreation string) (domain.User, error) {

	stmt, err := r.db.Prepare("INSERT INTO users(name, last_name, email, age, height, active, creation_date) VALUES( ?, ?, ?, ?, ?, ?, ? )")
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
