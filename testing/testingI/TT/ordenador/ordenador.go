package ordenador

import "sort"

func Ordenar(lista []int) []int {
	sort.Slice(lista, func(i, j int) bool { return lista[i] < lista[j] })
	return lista
}
