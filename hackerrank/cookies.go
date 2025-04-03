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
 * Complete the 'cookies' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY A
 */

func helper(i, j int) int {
	res := i + 2*j
	return res
}

func sorted(arr []int) []int {
	sort.Ints(arr)
	return arr
}

func cookies(k int32, A []int32) int32 {
	// Convert A to an int array
	intA := make([]int, len(A))
	for i, val := range A {
		intA[i] = int(val)
	}

	// Sort the array initially
	intA = sorted(intA)

	operations := 0

	for len(intA) > 1 && intA[0] < int(k) {
		// Combine the two smallest elements
		newElem := helper(intA[0], intA[1])
		intA = append([]int{newElem}, intA[2:]...) // Replace first two elements
		intA = sorted(intA)                        // Re-sort the array
		operations++
	}

	// Check if it's possible to meet the condition
	if intA[0] < int(k) {
		return -1
	}
	return int32(operations)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	ATemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var A []int32

	for i := 0; i < int(n); i++ {
		AItemTemp, err := strconv.ParseInt(ATemp[i], 10, 64)
		checkError(err)
		AItem := int32(AItemTemp)
		A = append(A, AItem)
	}

	result := cookies(k, A)

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
