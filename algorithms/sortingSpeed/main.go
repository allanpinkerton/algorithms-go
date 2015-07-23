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
func timeExecution(startTime time.Time, functionName string, inputSize int) string {
	executionTime := time.Since(startTime)
	return fmt.Sprintf("%-20s took %10.4fms to sort %d elements\n", functionName, float64(executionTime.Nanoseconds())/1000000, inputSize)
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
	checkError(err)

	arr := make([]int, sizeInt)
	for i := 0; i < sizeInt; i++ {
		arr[i] = rand.Int()
	}
	fmt.Println("Generated " + size + " integers.")

	mainChannel := make(chan string)
	for k, v := range sortingFunctions {
		newArr := make([]int, len(arr))
		copy(newArr, arr)
		go func(name string, v interface{}) {
			start := time.Now()
			v.(func([]int))(newArr)
			result := timeExecution(start, name, len(newArr))
			mainChannel <- result
		}(k, v)
		fmt.Println("0")
	}
	fmt.Println("1")

	fmt.Println("3")
	for _ = range sortingFunctions {
		fmt.Println("4")
		select {
		case result := <-mainChannel:
			fmt.Printf(result)
		case <-time.After(time.Second):
			fmt.Println("5")
		}
	}

	return
}
