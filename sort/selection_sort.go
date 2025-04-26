package sort

func SelectionSort(values []int) []int {
	for i := 0; i < len(values)-1; i++ {
		minIndex := i
		for j := i; j < len(values); j++ {
			if values[j] < values[minIndex] {
				minIndex = j
			}
		}

		if values[i] > values[minIndex] {
			aux := values[i]
			values[i] = values[minIndex]
			values[minIndex] = aux
		}

	}

	return values

}
