package sorting

// Performs basic insertion sort on slice

func InsertionSort(arr []int) {
	for i := 2; i < len(arr); i++ {
		key := arr[i]
		for j := i - 1; j > 0 && arr[j] > key; j-- {
			arr[j+1] = arr[j]
		}
	}
}

//Performs in-place quicksort on slice by choosing the first element in the slice pivot

func QuicksortFirstElement(arr []int) {
	if len(arr) <= 1 {
		return
	}
	arr[0] = arr[0]
	return
}
