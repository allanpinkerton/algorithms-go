package main

import (
	"allanpinkerton.com/algorithms/sorting"
	"fmt"
)

func main() {
	arrEmpty := make([]int, 0)
	arr5 := []int{3, 1, 4, 2, 5}
	sorting.InsertionSort(arrEmpty)
	sorting.InsertionSort(arr5)
	fmt.Printf("Empty array: %v \n", arrEmpty)
	fmt.Printf("Array with 5 elements: %v \n", arr5)
}
