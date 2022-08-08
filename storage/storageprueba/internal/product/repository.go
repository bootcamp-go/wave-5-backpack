package product

import (
	"database/sql"
	"log"

	"github.com/bootcamp-go/wave-5-backpack/tree/ospina_christian/storage/storageprueba/internal/domain"
	_ "github.com/go-sql-driver/mysql"
)

type Repository interface {
	Store(domain.Productos) (domain.Productos, error)
	GetOne(id int) (domain.Productos, error)
	Update(id int, nombre, color string, price float64, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Productos, error)
}

type repository struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Store(product domain.Productos) (domain.Productos, error) { // se inicializa la base
	db := r.db
	stmt, err := db.Prepare("INSERT INTO products(nombre, color, precio, stock, codigo, publicado, fechaCreacion) VALUES( ?, ?, ?, ?, ?, ? , ?)") // se prepara la sentencia SQL a ejecutar
	if err != nil {
		return domain.Productos{}, err
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	var result sql.Result
	result, err = stmt.Exec(product.Nombre, product.Color, product.Precio, product.Stock, product.Codigo, product.Publicado, product.FechaCreación) // retorna un sql.Result y un error
	if err != nil {
		return domain.Productos{}, err
	}
	insertedId, _ := result.LastInsertId() // del sql.Resul devuelto en la ejecucion obtenemos el Id insertado
	product.Id = int(insertedId)
	return product, nil
}

func (r *repository) GetOne(id int) (domain.Productos, error) {
	db := r.db
	var product domain.Productos

	rows, err := db.Query("SELECT id, nombre, color, precio, stock, codigo, publicado, fechaCreacion FROM products WHERE id = ?", id)
	if err != nil {
		return domain.Productos{}, err
	}
	for rows.Next() {
		if err := rows.Scan(&product.Id, &product.Nombre, &product.Color, &product.Precio, &product.Stock, &product.Codigo, &product.Publicado, &product.FechaCreación); err != nil {
			return domain.Productos{}, err
		}
	}
	return product, nil
}

func (r *repository) Update(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Productos, error) {
	db := r.db
	stmt, err := db.Prepare("UPDATE products SET nombre = ?, color = ?, precio = ?, stock = ?, codigo = ?, publicado = ?, fechaCreacion = ?  WHERE id = ?") // se prepara la sentencia SQL a ejecutar
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	_, err = stmt.Exec(nombre, color, precio, stock, codigo, publicado, fechaCreacion, id)
	if err != nil {
		return domain.Productos{}, err
	}
	product := domain.Productos{
		Id: id, Nombre: nombre, Color: color, Precio: precio, Stock: stock, Codigo: codigo, Publicado: publicado, FechaCreación: fechaCreacion,
	}
	return product, nil
}
