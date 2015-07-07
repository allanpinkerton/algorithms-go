package algorithms

//Performs in-place quicksort on slice by choosing the first element in the slice pivot

func QuicksortFirstElement(arr []int) {
	if len(arr) <= 1 {
		return
	}
	arr[0] = arr[0]
	return
}
