package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) (ratio_arr []float64) {
	// Write your code here
	denominator := len(arr)
	if denominator == 0 {
		err := errors.New("Empty Array")
		fmt.Println(err)
	}
	zeros := 0
	positive := 0
	negative := 0
	for _, val := range arr {
		if val > 0 {
			positive += 1
		} else if val < 0 {
			negative += 1
		} else if val == 0 {
			zeros += 1
		} else {
			fmt.Println("no integer")
		}
	}
	ratio_arr = append(ratio_arr, float64(positive)/float64(denominator))
	ratio_arr = append(ratio_arr, float64(negative)/float64(denominator))
	ratio_arr = append(ratio_arr, float64(zeros)/float64(denominator))

	for i, val := range ratio_arr {
		ratio_arr[i], _ = strconv.ParseFloat(fmt.Sprintf("%.6f", val), 64)
		fmt.Println(val)
	}

	return ratio_arr
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

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

	plusMinus(arr)
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
