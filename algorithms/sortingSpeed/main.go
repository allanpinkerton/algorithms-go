package main

import (
	"allanpinkerton.com/algorithms/sorting"
	"bufio"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"runtime"
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

var INSERTIONSORT_MAX = 100000

func main() {

	sortingFunctions := map[string]func([]int){
		"InsertionSort":        sorting.InsertionSort,
		"QuickSortLastElement": sorting.QuickSortLastElement,
		"QuickSortRandom":      sorting.QuickSortRandom,
		"HeapSort":             sorting.HeapSort,
		"MergeSort":            sorting.MergeSort,
		"IntroSort":            sorting.IntroSort,
	}

	maxCPU := runtime.NumCPU()
	maxProcs := runtime.GOMAXPROCS(0)

	// Allow maximum concurrency
	if maxCPU < maxProcs {
		runtime.GOMAXPROCS(maxProcs)
	} else {
		runtime.GOMAXPROCS(maxCPU)
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
		arr[i] = rand.Intn(math.MaxInt32 - 1)
	}
	fmt.Println("Generated " + size + " integers.")

	if sizeInt > INSERTIONSORT_MAX {
		delete(sortingFunctions, "InsertionSort")
		fmt.Println("Removed InsertionSort")
	}

	mainChannel := make(chan string)
	defer close(mainChannel)
	for k, v := range sortingFunctions {
		//fmt.Println(arr)
		go func(name string, v func([]int)) {
			//fmt.Println("Starting ", name)
			newArr := make([]int, len(arr))
			copy(newArr, arr)
			start := time.Now()
			v(newArr)
			result := timeExecution(start, name, len(newArr))
			mainChannel <- result
		}(k, v)
		select {
		case result := <-mainChannel:
			fmt.Printf(result)
		case <-time.After(5 * time.Second):
			fmt.Println("Timeout for: " + k)
		}

	}

	return
}
