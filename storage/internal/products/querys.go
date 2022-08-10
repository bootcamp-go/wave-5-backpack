package products

const (
	Store                  = "INSERT INTO products(name, type, price, count, code, public, warehouseid) VALUES (?,?,?,?,?,?,?)"
	GetProductByName       = "SELECT id, name, type, price, count, code, public, warehouseid FROM products WHERE name = ?"
	GetAll                 = "SELECT id, name, type, price, count, code, public FROM products"
	UpdateAll              = "UPDATE products SET name = ?, type = ?, price = ?, count = ?, code = ?, public = ? WHERE id = ?"
	GetProductAndWareHouse = "SELECT p.id, p.name, p.type, p.price, p.count, p.code, p.public, w.id, w.name, w.address FROM products p INNER JOIN warehouse w ON w.id = p.warehouseid"
	Delete                 = "DELETE FROM products WHERE id = ?"
)
