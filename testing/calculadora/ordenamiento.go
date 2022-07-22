package calculadora

func Ordenar(listNum []int) []int {

	tmp := 0

	for i := 0; i < len(listNum); i++ {
        for j := 0; j < len(listNum); j++ {
            if listNum[i] < listNum[j] {
                tmp = listNum[j]
                listNum[j] = listNum[i]
                listNum[i] = tmp
            }
        }
    }

return listNum
	
}