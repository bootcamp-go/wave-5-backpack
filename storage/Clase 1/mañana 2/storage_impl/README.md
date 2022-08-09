### Ejercicio 1 - Implementar GetByName()
Desarrollar un método en el repositorio que permita hacer búsquedas de un producto por nombre. Para lograrlo se deberá:

-  Diseñar interfaz “Repository” en la que exista un método GetByName() que reciba por parámetro un string y retorna una estructura del tipo Product.

-  Implementar el método de forma que con el string recibido lo use para buscar en la DB por el campo “name”.



```go
func (r *repository) GetByName(name string) (domains.Product, error) {
	stmt, err := r.db.Prepare("SELECT id, name, type, count, price FROM products WHERE name = ?;")
	if err != nil {
		return domains.Product{}, fmt.Errorf("error al preparar la consulta - error %v", err)
	}
	defer stmt.Close()

	var product domains.Product
	err = stmt.QueryRow(name).Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price)
	if err != nil {
		return domains.Product{}, fmt.Errorf("no registros para %s - error %v", name, err)
	}

	return product, nil
}
```

### Ejercicio 2 - Replicar Store()
Tomar el ejemplo visto en la clase y diseñar el método Store():
Puede tomar de ejemplo la definición del método Store visto en clase para incorporarlo en la interfaz.
Implementar el método Store.


```go
func (r *repository) Store(p domains.Product) (int, error) {
	stmt, err := r.db.Prepare("INSERT INTO products (name, type, count, price) VALUES (?,?,?,?)")
	if err != nil {
		return 0, fmt.Errorf("error al preparar la consulta - error %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(p.Name, p.Type, p.Count, p.Price)
	if err != nil {
		return 0, fmt.Errorf("error al ejecutar la consulta - error %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error al obtener último id - error %v", err)
	}

	return int(id), nil
}
```


### Ejercicio 3 - Ejecutar Store()
Diseñar un Test que permita ejecutar Store y comprobar la correcta inserción del registro en la tabla.

```go
func TestStoreProduct_Ok(t *testing.T) {
	new := domains.Product{
		Name:  "producto nuevo",
		Type:  "producto tipo",
		Count: 3,
		Price: 84.4,
	}

	product, err := json.Marshal(new)
	require.Nil(t, err)

	req, rr := createRequest(http.MethodPost, "/api/v1/products/", string(product))
	s.ServeHTTP(rr, req)

	// assert code
	assert.Equal(t, http.StatusCreated, rr.Code)
}
```

### Ejercicio 4 - Ejecutar GetByName()
Diseñar un Test que permita ejecutar GetByName y comprobar que retorne el registro buscado en caso de que exista. 

```go
func TestGetByNameProduct_Ok(t *testing.T) {
	req, rr := createRequest(http.MethodGet, "/api/v1/products/", `{"nombre":"producto nuevo"}`)
	s.ServeHTTP(rr, req)

	// assert code
	assert.Equal(t, http.StatusOK, rr.Code)
}

```