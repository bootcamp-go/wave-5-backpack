package users

import (
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

func (r *repositoryDB) GetAll() ([]domain.Users, error) {
	var users []domain.Users
	return users, nil
}

func (r *repositoryDB) LastID() (int, error) {
	return 0, nil
}

func (r *repositoryDB) GetByName(name string) ([]domain.Users, error) {
	var user domain.Users
	var listUser []domain.Users
	rows, err := r.db.Query("SELECT id, name, last_name, email, age, height, active, creation_date FROM users WHERE name = ?", name)
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

func (r *repositoryDB) Store(id, age int, name, lastName, email, creationDate string, height float64, active bool) (domain.Users, error) {
	stmt, err := r.db.Prepare("INSERT INTO users(name, last_name, email, age, height, active, creation_date) VALUES( ?, ?, ?, ?, ?, ?, ? )")
	if err != nil {
		return domain.Users{}, err
	}
	defer stmt.Close()
	var result sql.Result
	result, err2 := stmt.Exec(name, lastName, email, age, height, active, creationDate)
	if err2 != nil {
		return domain.Users{}, err
	}
	insertedId, _ := result.LastInsertId()
	user := domain.Users{
		Id: int(insertedId), Name: name, LastName: lastName, Email: email, Age: age, Height: height, Active: active, CreationDate: creationDate}
	return user, nil
}

func (r *repositoryDB) Update(id, age int, name, lastName, email, creationDate string, height float64, active bool) (domain.Users, error) {
	return domain.Users{}, fmt.Errorf(UserNotFound, id)
}

func (r *repositoryDB) UpdateLastNameAndAge(id, age int, lastName string) (domain.Users, error) {
	return domain.Users{}, fmt.Errorf(UserNotFound, id)
}

func (r *repositoryDB) Delete(id int) error {
	return fmt.Errorf(UserNotFound, id)
}
