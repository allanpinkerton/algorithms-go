package sorting

import (
	//"math/rand"
	"testing"
)

var arrs = [][]int{
	make([]int, 0),
	[]int{1},
	[]int{3, 5, 2, 1, 4},
	[]int{1, 2, 3, 4, 5},
	[]int{5, 4, 3, 2, 1},
	[]int{1, 1, 1, 1, 1},
	[]int{-3, 1, 8, 10, -9, -4},
	[]int{23, 143, 231, 11, 1, 8, 16, 88, 9, 198, 44, 29, 12,
		90, 86, 57, 55, 21, 71, 60, 27, 2, 128, 721, 333},
}

var sortingFuncs = map[string]func([]int){
	"InsertionSort":        InsertionSort,
	"QuickSortLastElement": QuickSortLastElement,
	"QuickSortRandom":      QuickSortRandom,
	"HeapSort":             HeapSort,
	"MergeSort":            MergeSort,
	"IntroSort":            IntroSort,
	//"CountingSort":         CountingSort,
}

func deepCopy(dest [][]int, src [][]int) {
	for i := range src {
		dest[i] = make([]int, len(src[i]))
		copy(dest[i], src[i])
	}
}

func isEqualArr(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func TestSorting(t *testing.T) {
	for k, v := range sortingFuncs {
		arrsCopy := make([][]int, len(arrs))
		deepCopy(arrsCopy, arrs)

		for i, arr := range arrsCopy {
			v(arr)
			if !isSorted(arr) {
				t.Error(
					"Error running", k,
					"Input:", arrs[i],
					"Got:", arr)
			}
		}
	}
}


func TestBuildMaxHeap(t *testing.T) {
	arrsCopy := make([][]int, len(arrs))
	deepCopy(arrsCopy, arrs)

	var heap = [][]int{
		make([]int, 0),
		[]int{1},
		[]int{5, 4, 2, 1, 3},
		[]int{5, 4, 3, 1, 2},
		[]int{5, 4, 3, 2, 1},
		[]int{1, 1, 1, 1, 1},
		[]int{10, 1, 8, -3, -9, -4},
		[]int{721, 198, 333, 88, 143, 231, 90, 57, 71, 60, 128, 29, 12, 16, 86, 11, 55, 21, 9, 1, 27, 2, 44, 23, 8},
	}

	for i, arr := range arrsCopy {
		buildMaxHeap(arr)
		if !isEqualArr(arr, heap[i]) {
			t.Error(
				"Input:", arrs[i],
				"Expected:", heap[i],
				"Got:", arr)
		}
	}
}
