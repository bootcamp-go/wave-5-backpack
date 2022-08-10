package products

const (
	EXISTS_QUERY      = "SELECT id FROM products WHERE id=?;"
	INSERT_QUERY      = "INSERT INTO products (name, type, count, price, id_warehouse) VALUES (?,?,?,?,?)"
	UPDATE_QUERY      = "UPDATE products SET name=?, type=?, count=?, price=?, id_warehouse=? WHERE id=?;"
	GET_BY_NAME_QUERY = "SELECT id, name, type, count, price FROM products WHERE name = ?;"
	GET_ALL_QUERY     = "SELECT products.id, products.name, products.type, products.count, products.price, warehouses.name, warehouses.adress FROM products INNER JOIN warehouses ON products.id_warehouse = warehouses.id WHERE products.id = ?;"
	GET_BY_ID_QUERY   = "SELECT id, name, type, count, price, id_warehouse FROM products WHERE id = ?;"
	DELETE_QUERY      = "DELETE FROM products WHERE id = ?;"
)
