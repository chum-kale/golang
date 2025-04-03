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
 * Complete the 'hackerlandRadioTransmitters' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY x
 *  2. INTEGER k
 */

func findMin(arr []int32) int32 {
	if len(arr) == 0 {
		panic("Array is empty")
	}
	min := arr[0]
	for _, val := range arr {
		if val < min {
			min = val
		}
	}
	return min
}

func findMax(arr []int32) int32 {
	if len(arr) == 0 {
		panic("Array is empty")
	}
	max := arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}

func hackerlandRadioTransmitters(x []int32, k int32) int32 {
	// Write your code here
	var count int32 = 0
	n := len(x)
	i := 0

	sort.Slice(x, func(i, j int) bool { return x[i] < x[j] })

	for i < n {
		count++
		// farthest house covered
		loc := x[i] + k
		for i < n && x[i] <= loc {
			i++
		}

		loc = x[i-1] + k
		for i < n && x[i] <= loc {
			i++
		}
	}

	return count
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

	xTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var x []int32

	for i := 0; i < int(n); i++ {
		xItemTemp, err := strconv.ParseInt(xTemp[i], 10, 64)
		checkError(err)
		xItem := int32(xItemTemp)
		x = append(x, xItem)
	}

	result := hackerlandRadioTransmitters(x, k)

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
