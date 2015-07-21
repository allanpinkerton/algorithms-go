package main

import (
	"allanpinkerton.com/algorithms/sorting"
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Prints out the time elapsed since start
func timeExecution(startTime time.Time, functionName string, inputSize int) {
	executionTime := time.Since(startTime)
	fmt.Printf("%s took %dms to sort %d elements\n", functionName, executionTime.Nanoseconds()/1000, inputSize)
	return
}

// Generates file with n random ints named integerArray + n
func generateRandomIntegers(n int, filename string) {

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Int()
	}

	f, _ := os.Create(filename)
	defer f.Close()

	for _, num := range arr {
		f.WriteString(strconv.Itoa(num) + " ")
	}
	f.WriteString("\n")
	f.Sync()
	fmt.Printf("Generated " + filename + " with " + strconv.Itoa(n) + " elements.\n")
}

// Reads a file of ints into an integer slice
func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	sortingFunctions := map[string]interface{}{
		"InsertionSort":        sorting.InsertionSort,
		"QuickSortLastElement": sorting.QuickSortLastElement,
		"QuickSortRandom":      sorting.QuickSortRandom,
	}
	if len(os.Args) != 2 {
		fmt.Printf("No size specified.\n")
		return
	}
	size := os.Args[1]
	sizeInt, err := strconv.Atoi(size)

	filename := "integerArray" + size
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		generateRandomIntegers(sizeInt, filename)
	}
	f, err := os.Open(filename)
	checkError(err)

	defer f.Close()

	arr, err := readInts(f)
	checkError(err)

	newArr := make([]int, len(arr))
	for k, v := range sortingFunctions {
		copy(newArr, arr)
		start := time.Now()
		v.(func([]int))(newArr)
		timeExecution(start, k, len(newArr))
	}
	return
}
