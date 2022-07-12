package main

import "fmt"
const (
	pequeño = "pequeño"
	mediano = "mediano"
	grande = "grande"
)

type Producto struct {
	TipProducto string
	Nombre string
	Precio float64
}
type Tienda struct {
	ProductosLs []IProducto 
}

type IProducto interface {
	CalcularCosto() float64
}
type IEcommerce interface {
	Total() float64
	Agregar(prod IProducto) 
}

func nuevoTineda() IEcommerce {
	return &Tienda{}
}
func nuevoProducto(tipProd, nom string, prec float64) IProducto{
	return &Producto{TipProducto: tipProd,Nombre: nom,Precio: prec}
}
func main(){
	nProcto1 := nuevoProducto(pequeño, "sopa",100)
	nProcto2 := nuevoProducto(mediano, "sopa",100)
	nProcto3 := nuevoProducto(grande, "sopa",100)
	nTienda := nuevoTineda()

	nTienda.Agregar(nProcto1)
	nTienda.Agregar(nProcto2)
	nTienda.Agregar(nProcto3)
	nTotal := nTienda.Total();

	fmt.Printf("La tineda tien un valor total de: %.2f\n",nTotal)
}

func (p Producto) CalcularCosto() float64{
	switch p.TipProducto {
	case pequeño:
		return p.Precio 
	case mediano:
		return p.Precio + ((3 * p.Precio)/100)
	case grande:
		return p.Precio + ((6 * p.Precio)/100) + 2500
	}
	return 0
}

func (t *Tienda) Agregar(prod IProducto){
	t.ProductosLs = append(t.ProductosLs, prod)
}
func (t Tienda) Total() float64{
	var total float64
	for _,p := range t.ProductosLs {
		subPrecio := p.CalcularCosto()
		total += subPrecio
	}
	return total
}