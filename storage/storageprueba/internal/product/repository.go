package product

import (
	"context"
	"database/sql"
	"log"

	"github.com/bootcamp-go/wave-5-backpack/tree/ospina_christian/storage/storageprueba/internal/domain"
	_ "github.com/go-sql-driver/mysql"
)

type Repository interface {
	Store(domain.Productos) (domain.Productos, error)
	GetOne(ctx context.Context, id int) (domain.Productos, error)
	Update(ctx context.Context, id int, nombre, color string, price float64, stock int, codigo string, publicado bool, fechaCreacion string, idWarehouse int) (domain.Productos, error)
	GetAll(ctx context.Context) ([]domain.Productos, error)
	Delete(ctx context.Context, id int) error
	GetAllProdWare(ctx context.Context, id int) domain.ProductosWarehouse
}

//queries
var (
	storeQueryProducts           string = "INSERT INTO products(nombre, color, precio, stock, codigo, publicado, fechaCreacion, id_warehouse) VALUES( ?, ?, ?, ?, ?, ? , ?, ?)"
	updateQueryProducts          string = "UPDATE products SET nombre = ?, color = ?, precio = ?, stock = ?, codigo = ?, publicado = ?, fechaCreacion = ?, id_warehouse = ?  WHERE id = ?"
	getOneQueryProducts          string = "SELECT id, nombre, color, precio, stock, codigo, publicado, fechaCreacion, id_warehouse FROM products WHERE id = ?"
	getAllQueryProducts          string = "SELECT id, nombre, color, precio, stock, codigo, publicado, fechaCreacion, id_warehouse FROM products"
	deleteQueryProducts          string = "DELETE FROM products WHERE id = ?"
	getAllQueryProductsWarehouse string = "SELECT products.id, products.nombre, products.color, products.precio, products.stock, products.codigo, products.publicado, products.fechaCreacion, products.id_warehouse, warehouses.id,  warehouses.name, warehouses.address FROM products INNER JOIN warehouses ON products.id_warehouse = warehouses.id WHERE products.id = ?;"
)

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
	stmt, err := db.Prepare(storeQueryProducts) // se prepara la sentencia SQL a ejecutar
	if err != nil {
		return domain.Productos{}, err
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	var result sql.Result
	result, err = stmt.Exec(product.Nombre, product.Color, product.Precio, product.Stock, product.Codigo, product.Publicado, product.FechaCreación, product.IdWarehouse) // retorna un sql.Result y un error
	if err != nil {
		return domain.Productos{}, err
	}
	insertedId, _ := result.LastInsertId() // del sql.Resul devuelto en la ejecucion obtenemos el Id insertado
	product.Id = int(insertedId)
	return product, nil
}

func (r *repository) GetOne(ctx context.Context, id int) (domain.Productos, error) {
	db := r.db
	var product domain.Productos

	rows, err := db.QueryContext(ctx, getOneQueryProducts, id)
	if err != nil {
		return domain.Productos{}, err
	}
	for rows.Next() {
		if err := rows.Scan(&product.Id, &product.Nombre, &product.Color, &product.Precio, &product.Stock, &product.Codigo, &product.Publicado, &product.FechaCreación, &product.IdWarehouse); err != nil {
			return domain.Productos{}, err
		}
	}
	return product, nil
}

func (r *repository) Update(ctx context.Context, id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string, idWarehouse int) (domain.Productos, error) {
	db := r.db
	stmt, err := db.PrepareContext(ctx, updateQueryProducts) // se prepara la sentencia SQL a ejecutar
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	_, err = stmt.ExecContext(ctx, nombre, color, precio, stock, codigo, publicado, fechaCreacion, idWarehouse, id)
	if err != nil {
		return domain.Productos{}, err
	}
	product := domain.Productos{
		Id: id, Nombre: nombre, Color: color, Precio: precio, Stock: stock, Codigo: codigo, Publicado: publicado, FechaCreación: fechaCreacion, IdWarehouse: idWarehouse,
	}
	return product, nil
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Productos, error) {
	var products []domain.Productos
	db := r.db
	rows, err := db.QueryContext(ctx, getAllQueryProducts)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//Se recorren todas las filas
	for rows.Next() {
		//Por cada fila se obtiene un objeto del tipo producto
		var product domain.Productos
		if err := rows.Scan(&product.Id, &product.Nombre, &product.Color, &product.Precio, &product.Stock, &product.Codigo, &product.Publicado, &product.FechaCreación, &product.IdWarehouse); err != nil {
			log.Fatal(err)
			return nil, err
		}
		//Se anade el objeto al slice
		products = append(products, product)
	}
	return products, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	db := r.db
	stmt, err := db.PrepareContext(ctx, deleteQueryProducts)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, id)

	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetAllProdWare(ctx context.Context, id int) domain.ProductosWarehouse {
	var prodWare domain.ProductosWarehouse
	db := r.db
	rows, err := db.QueryContext(ctx, getAllQueryProductsWarehouse, id)
	if err != nil {
		log.Println(err)
		return prodWare
	}
	for rows.Next() {
		if err := rows.Scan(
			&prodWare.Id, &prodWare.Nombre, &prodWare.Color, &prodWare.Precio, &prodWare.Stock, &prodWare.Codigo, &prodWare.Publicado, &prodWare.FechaCreación, &prodWare.IdWarehouse,
			&prodWare.Warehouses.Id, &prodWare.Warehouses.Nombre, &prodWare.Warehouses.Direccion); err != nil {
			log.Fatal(err)
			return prodWare
		}
	}
	return prodWare
}
