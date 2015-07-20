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
}

func deepCopy(dest [][]int, src [][]int) {
	for i := range src {
		dest[i] = make([]int, len(src[i]))
		copy(dest[i], src[i])
	}
}
func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func TestEmptyInsertionSort(t *testing.T) {
	arrsCopy := make([][]int, len(arrs))
	deepCopy(arrsCopy, arrs)

	for i, arr := range arrsCopy {
		InsertionSort(arr)
		if !isSorted(arr) {
			t.Error(
				"Input:", arrs[i],
				"Got:", arr)
		}
	}
}

func TestEmptyQuickSortFirstElement(t *testing.T) {
	arrsCopy := make([][]int, len(arrs))
	deepCopy(arrsCopy, arrs)

	for i, arr := range arrsCopy {
		QuickSortLastElement(arr)
		if !isSorted(arr) {
			t.Error(
				"Input:", arrs[i],
				"Got:", arr)
		}
	}
}
