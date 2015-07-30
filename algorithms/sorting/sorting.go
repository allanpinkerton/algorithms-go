package sorting

import (
	"math"
	"math/rand"
)

// General helper functions
func swap(arr []int, a int, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
	return
}

func random(min, max int) int {
	return rand.Intn(max+1-min) + min
}

// Performs basic insertion sort on slice

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for ; j >= 0 && arr[j] > key; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = key
	}
}

/*
Quick Sort using last eleement as pivot
*/

func partition(arr []int, p int, r int, pivot int) (pivotPos int) {
	x := arr[r]
	pivotPos = p - 1
	for j := p; j < r; j++ {
		if arr[j] <= x {
			pivotPos++
			swap(arr, pivotPos, j)
		}
	}
	pivotPos++
	swap(arr, pivotPos, r)
	return
}

func quickSortLastElementAux(arr []int, p int, r int) {
	if p < r {
		pivot := r
		q := partition(arr, p, r, pivot)
		quickSortLastElementAux(arr, p, q-1)
		quickSortLastElementAux(arr, q+1, r)
	}
}

func QuickSortLastElement(arr []int) {
	quickSortLastElementAux(arr, 0, len(arr)-1)
}

func quickSortRandomAux(arr []int, p int, r int) {
	if p < r {
		pivot := random(p, r)
		q := partition(arr, p, r, pivot)
		quickSortLastElementAux(arr, p, q-1)
		quickSortLastElementAux(arr, q+1, r)
	}
}

func QuickSortRandom(arr []int) {
	quickSortRandomAux(arr, 0, len(arr)-1)
}

func maxHeapify(arr []int, i int, size int) {
	var largest int
	for {
		l := (i+1)*2 - 1
		r := l + 1
		if l < size && arr[l] > arr[i] {
			largest = l
		} else {
			largest = i
		}
		if r < size && arr[r] > arr[largest] {
			largest = r
		}
		if largest != i {
			swap(arr, i, largest)
			i = largest
		} else {
			break
		}
	}
}

func buildMaxHeap(arr []int) {
	size := len(arr)
	for i := size / 2; i >= 0; i-- {
		maxHeapify(arr, i, size)
	}
}

func HeapSort(arr []int) {
	buildMaxHeap(arr)
	size := len(arr) - 1
	for i := size; i > 0; i-- {
		swap(arr, i, 0)
		maxHeapify(arr, 0, size)
		size--
	}
}

func merge(arr []int, p, q, r int) {
	sizeLeft := q - p + 1
	sizeRight := r - q
	left := make([]int, sizeLeft+1)
	right := make([]int, sizeRight+1)
	for i := 0; i < sizeLeft; i++ {
		left[i] = arr[p+i]
	}
	for i := 0; i < sizeRight; i++ {
		right[i] = arr[q+i+1]
	}
	left[sizeLeft] = math.MaxInt64
	right[sizeRight] = math.MaxInt64
	i := 0
	j := 0
	for k := p; k <= r; k++ {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
	}

}

func mergeSortAux(arr []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		mergeSortAux(arr, p, q)
		mergeSortAux(arr, q+1, r)
		merge(arr, p, q, r)
	}
}

func MergeSort(arr []int) {
	mergeSortAux(arr, 0, len(arr)-1)
	return
}

func introSortAux(arr []int, maxDepth int) {
	n := len(arr)
	if n <= 1 {
		return
	}
	if maxDepth == 0 {
		HeapSort(arr)
	} else {
		pivot := random(0, len(arr))
		q := partition(arr, 0, len(arr)-1, pivot)
		introSortAux(arr[:q], maxDepth-1)
		introSortAux(arr[q:], maxDepth-1)
	}
}

func IntroSort(arr []int) {
	maxDepth := int(math.Log(float64(len(arr)))) * 2
	introSortAux(arr, maxDepth)
}
