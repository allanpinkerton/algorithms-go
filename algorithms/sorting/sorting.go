package sorting

import (
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
