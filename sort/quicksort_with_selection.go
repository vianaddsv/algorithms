package sort

func QuicksortSelection(values []int, min int) []int {
	return quickSort(values, 0, len(values)-1, min)
}

func quickSort(values []int, first int, last int, min int) []int {
	if len(values) == min {
		return SelectionSort(values)
	}
	if first < last {
		pivot := partion(values, first, last)
		quickSort(values, first, pivot-1, min)
		quickSort(values, pivot+1, last, min)
	}
	return values
}

func partion(values []int, first int, last int) int {
	pivot := values[last]
	i := first

	for j := first; j < last; j++ {
		if values[j] <= pivot {
			values[j], values[i] = values[i], values[j]
			i++
		}
	}
	values[i], values[last] = values[last], values[i]
	return i
}
