package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'maxSubarray' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func largestNegative(arr []int32) int32 {
	var largestNegative int32 = -2147483648

	for _, val := range arr {
		if val < 0 && val > largestNegative {
			largestNegative = val
		}
	}

	return largestNegative
}

func maxSubarray(arr []int32) []int32 {
	// Write your code here
	var result []int32

	var maxi int32 = arr[0]
	var max_local int32 = arr[0]

	for i := 1; i < len(arr); i++ {
		max_local = max(arr[i], max_local+arr[i])
		maxi = max(maxi, max_local)
	}
	result = append(result, maxi)

	var sum int32 = 0
	for _, val := range arr {
		if val > 0 {
			sum += val
		}
	}
	if sum == 0 {
		largest := largestNegative(arr)
		result = append(result, largest)
	} else {
		result = append(result, sum)
	}

	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
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

		result := maxSubarray(arr)

		for i, resultItem := range result {
			fmt.Fprintf(writer, "%d", resultItem)

			if i != len(result)-1 {
				fmt.Fprintf(writer, " ")
			}
		}

		fmt.Fprintf(writer, "\n")
	}

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
