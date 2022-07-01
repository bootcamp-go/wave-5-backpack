package main
import "fmt"


func impuesto(sueldo float64) float64{
		if sueldo >50000 && sueldo < 150000 {
			return sueldo * 0.17
		}else if sueldo >150000 {
			imp := sueldo * 0.17
			return imp + (imp * 0.10)
		}else{
			return 0
		}			
}
func main(){
fmt.Printf("Impuestos: %2.f \n",impuesto(160000))
}