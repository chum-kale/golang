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
 * Complete the 'equalStacks' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY h1
 *  2. INTEGER_ARRAY h2
 *  3. INTEGER_ARRAY h3
 */

func array_sum(arr []int32) int32 {
	var sum int32 = 0
	for _, val := range arr {
		sum += val
	}

	return sum
}

func check_empty_array(h1 []int32, h2 []int32, h3 []int32) bool {
	if len(h1) == 0 || len(h2) == 0 || len(h3) == 0 {
		return true
	} else {
		return false
	}
}

func check_equality(h1 []int32, h2 []int32, h3 []int32) bool {
	if array_sum(h1) == array_sum(h2) && array_sum(h2) == array_sum(h3) {
		return true
	}
	return false
}

func equalStacks(h1 []int32, h2 []int32, h3 []int32) int32 {
	// Write your code here
	for len(h1) > 0 && len(h2) > 0 && len(h3) > 0 {
		sum1 := array_sum(h1)
		sum2 := array_sum(h2)
		sum3 := array_sum(h3)

		if sum1 == sum2 && sum2 == sum3 {
			return sum1
		}

		if sum1 >= sum2 && sum1 >= sum3 {
			h1 = h1[1:]
		} else if sum2 >= sum1 && sum2 >= sum3 {
			h2 = h2[1:]
		} else {
			h3 = h3[1:]
		}
	}

	return 0
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	n1Temp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n1 := int32(n1Temp)

	n2Temp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	n2 := int32(n2Temp)

	n3Temp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	n3 := int32(n3Temp)

	h1Temp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var h1 []int32

	for i := 0; i < int(n1); i++ {
		h1ItemTemp, err := strconv.ParseInt(h1Temp[i], 10, 64)
		checkError(err)
		h1Item := int32(h1ItemTemp)
		h1 = append(h1, h1Item)
	}

	h2Temp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var h2 []int32

	for i := 0; i < int(n2); i++ {
		h2ItemTemp, err := strconv.ParseInt(h2Temp[i], 10, 64)
		checkError(err)
		h2Item := int32(h2ItemTemp)
		h2 = append(h2, h2Item)
	}

	h3Temp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var h3 []int32

	for i := 0; i < int(n3); i++ {
		h3ItemTemp, err := strconv.ParseInt(h3Temp[i], 10, 64)
		checkError(err)
		h3Item := int32(h3ItemTemp)
		h3 = append(h3, h3Item)
	}

	result := equalStacks(h1, h2, h3)

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

// other way

func equalStacks(h1 []int32, h2 []int32, h3 []int32) int32 {
	counter := 0
	for !check_empty_array(h1, h2, h3) {
		if check_equality(h1, h2, h3) {
			// Return the height of any stack as they are equal
			return array_sum(h1)
		} else if counter == 0 {
			// Remove from stack 1
			h1 = h1[1:]
			counter = 1
		} else if counter == 1 {
			// Remove from stack 2
			h2 = h2[1:]
			counter = 2
		} else if counter == 2 {
			// Remove from stack 3
			h3 = h3[1:]
			counter = 0
		}
	}
	// If any stack becomes empty, return 0
	return 0
}
