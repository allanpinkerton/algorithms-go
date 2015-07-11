package main

import (
	"fmt"
	"sorting"
)

func main() {
	arrEmpty := make([]int, 0)
	arr5 := []int{3, 1, 4, 2, 5}
	InsertionSort(arrEmpty)
	InsertionSort(arr5)
	fmt.Printf("Empty array: %s", arrEmpty)
	fmt.Printf("Array with 5 elements: %s", arr5)
}
