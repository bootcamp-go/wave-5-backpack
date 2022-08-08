package usuarios

import (
	"database/sql"
	"fmt"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
)

type RepositoryBD interface {
	GetByName(name string) (domain.Usuarios, error)
}

type repositoryBD struct {
	dbBD *sql.DB
}

func NewRepositoryBD(dbb *sql.DB) RepositoryBD {
	return &repositoryBD{
		dbBD: dbb,
	}
}

func (r *repositoryBD) GetByName(name string) (domain.Usuarios, error) {
	fmt.Println(name)
	var user domain.Usuarios

	rows, err := r.dbBD.Query("SELECT id, nombre, apellido, email, edad, altura, activo, fechaCreacion FROM storage.users WHERE nombre =?", name)
	fmt.Println(err)
	if err != nil {
		return domain.Usuarios{}, err
	}
	fmt.Println(user)
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Nombre, &user.Apellido, &user.Email, &user.Edad, &user.Altura, &user.Activo, &user.FechaCreacion); err != nil {
			return domain.Usuarios{}, err
		}
	}
	fmt.Println(user)
	return user, nil
}
