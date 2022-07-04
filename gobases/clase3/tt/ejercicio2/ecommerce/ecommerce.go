package ecommerce

type Usuario struct {
  Nombre string
  Apellido string
  Correo string
  Productos *[]*Producto
}

func NuevoUsuario(nombre, apellido, correo string) Usuario{
  return Usuario{
    Nombre: nombre,
    Apellido: apellido,
    Correo: correo,
    Productos: &[]*Producto{},
  }
}

type Producto struct {
  Nombre string
  precio float64
  cantidad uint
}

func NuevoProducto(nombre string, precio float64) Producto {
  return Producto{
    Nombre: nombre,
    precio: precio,
  }
}

func AgregarProducto(u *Usuario, producto *Producto, cantidad uint) {
  for _, p:= range *u.Productos {
    // Check si el usuario ya cuenta con el producto
    if p.Nombre == producto.Nombre {

      p.cantidad += cantidad
      return
    }
  }

  producto.cantidad = cantidad

  *u.Productos = append(*u.Productos, producto) // Si no estaba lo agrega a la lista
}

func BorrarProductos(u *Usuario) {
  u.Productos = &[]*Producto{}
}
