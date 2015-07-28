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

func TestInsertionSort(t *testing.T) {
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

func TestQuickSortLastElement(t *testing.T) {
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

func TestQuickSortRandom(t *testing.T) {
	arrsCopy := make([][]int, len(arrs))
	deepCopy(arrsCopy, arrs)

	for i, arr := range arrsCopy {
		QuickSortRandom(arr)
		if !isSorted(arr) {
			t.Error(
				"Input:", arrs[i],
				"Got:", arr)
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

func TestHeapSort(t *testing.T) {
	arrsCopy := make([][]int, len(arrs))
	deepCopy(arrsCopy, arrs)

	for i, arr := range arrsCopy {
		HeapSort(arr)
		if !isSorted(arr) {
			t.Error(
				"Input:", arrs[i],
				"Got:", arr)
		}
	}
}
