package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'lilysHomework' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

// sum of adjacent elements in an array needs to be minimized
// to achieve this, array needs to be sorted in ascending order
// to find: min no of swaps required to sort the array
func lilysHomework(arr []int32) int32 {
	n := len(arr)

	// Create two copies of the original array
	arr1 := make([]int32, n)
	arr2 := make([]int32, n)
	copy(arr1, arr)
	copy(arr2, arr)

	// Calculate swaps for ascending and descending sort
	return min(countMinSwaps(arr1, true), countMinSwaps(arr2, false))
}

func countMinSwaps(arr []int32, ascending bool) int32 {
	n := len(arr)

	// Create map to track original indices
	indexMap := make(map[int32]int)
	for i, val := range arr {
		indexMap[val] = i
	}

	// Create sorted version of array
	sorted := make([]int32, n)
	copy(sorted, arr)
	sort.Slice(sorted, func(i, j int) bool {
		if ascending {
			return sorted[i] < sorted[j]
		}
		return sorted[i] > sorted[j]
	})

	var swaps int32 = 0
	visited := make([]bool, n)

	for i := 0; i < n; i++ {
		if visited[i] || arr[i] == sorted[i] {
			continue
		}

		cycleSize := 0
		j := i
		for !visited[j] {
			visited[j] = true
			j = indexMap[sorted[j]]
			cycleSize++
		}

		if cycleSize > 0 {
			swaps += int32(cycleSize - 1)
		}
	}

	return swaps
}

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := lilysHomework(arr)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
