package main

import "fmt"

/*Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios.
Para ello requieren que tanto los usuarios como los productos tengan la misma dirección de memoria en el main del programa
como en las funciones.
Se necesitan las estructuras:
Usuario: Nombre, Apellido, Correo, Productos (array de productos).
Producto: Nombre, precio, cantidad.
Se requieren las funciones:
Nuevo producto: recibe nombre y precio, y retorna un producto.
Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
Borrar productos: recibe un usuario, borra los productos del usuario.
*/

type Users struct {
	Name     string
	LastName string
	Email    string
	Products []Product
}

type Product struct {
	NameProduct string
	Price       float64
	Amount      int64
}

func NewProduct(nameProduct *string, price *float64) *Product {
	return &Product{NameProduct: *nameProduct, Price: *price}
}

func (u *Users) AddProduct(product *Product, cantidad *int64) {
	product.Amount = *cantidad
	u.Products = append(u.Products, *product)
}

func (u *Users) Delete() {
	u.Products = []Product{}
}

func main() {
	var (
		nameProduct string  = "Bicycle"
		price       float64 = 3450.50
		amount      int64   = 5
	)
	p := NewProduct(&nameProduct, &price)

	u := &Users{
		Name:     "José Luis",
		LastName: "Riverón",
		Email:    "jriveronrodriguez@gmail.com",
	}

	u.AddProduct(p, &amount)
	fmt.Printf("El usuario registrado es:\n Nombre: %s\n Apellido: %s\n Correo: %s\n", u.Name, u.LastName, u.Email)
	fmt.Println("EL PRODUCTO AGREGADO")
	for _, values := range u.Products {
		fmt.Printf("El Producto agrgado es: Producto: %s\n Cantidad: %d\n Precio: %.2f\n Precio Total: %v\n", values.NameProduct, values.Amount, values.Price, values.Price*float64(values.Amount))
	}
	fmt.Println("DELETE TEST")
	u.Delete()
	fmt.Printf("After Delete: %v\n", len(u.Products))
}
