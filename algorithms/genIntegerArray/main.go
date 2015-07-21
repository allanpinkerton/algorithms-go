package main

import (
	//"allanpinkerton.com/algorithms/sorting"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("No size specified.\n")
		return
	}
	size, _ := strconv.Atoi(os.Args[1])
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Int()
	}

	filename := "integerArray" + strconv.Itoa(size)
	f, _ := os.Create(filename)
	defer f.Close()

	for _, num := range arr {
		f.WriteString(strconv.Itoa(num) + " ")
	}
	f.WriteString("\n")
	f.Sync()
	fmt.Printf("Generated " + filename + " with " + strconv.Itoa(size) + " elements.\n")
}
