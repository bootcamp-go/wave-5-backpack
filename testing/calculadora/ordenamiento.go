package calculadora

func Ordenar(listNum []int) []int {

	tmp := 0

	for i, numI := range listNum{
		for j, numJ := range listNum{
			if numI < numJ {
				tmp = listNum[j]
                listNum[j] = listNum[i]
                listNum[i] = tmp
				
			}
		}
	}

return listNum
	
}